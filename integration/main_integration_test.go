//go:build integration

package integration_test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/nobbs/uptime-kuma-api/pkg/action"
	"github.com/nobbs/uptime-kuma-api/pkg/client"
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/ory/dockertest/v3"
)

const (
	uptimeKumaImage = "louislam/uptime-kuma" // uptimeKumaImage is the image name of the Uptime Kuma container
	uptimeKumaTag   = "1.23.0"               // uptimeKumaTag is the image tag of the Uptime Kuma container

	username = "admin"    // username to use for setup of Uptime Kuma
	password = "admin123" // password to use for setup of Uptime Kuma

	// defaultContainerExpiration is the default expiration time for the uptime-kuma container in
	// seconds. This is used to expire the container after a certain amount of time to prevent the
	// container from running forever in case of an error.
	defaultContainerExpiration = 360
)

var (
	uptimeKumaHost         string = "localhost"
	uptimeKumaPortExternal int
)

func newConnectedClient() (*client.Client, error) {
	// create new client
	c, err := client.NewClient(uptimeKumaHost, uptimeKumaPortExternal, false)
	if err != nil {
		return nil, fmt.Errorf("Failed to create new client: %s", err)
	}

	// Wait for connection event
	err = c.Await(handler.ConnectEvent, time.Duration(5)*time.Second)
	if err != nil {
		return nil, fmt.Errorf("Failed to await connection event: %s", err)
	}

	return c, nil
}

// setupUptimeKumaServer sets up a new Uptime Kuma server, i.e. it sets the username and password
// for the server. If the server is already setup, it will not be setup again.
func setupUptimeKumaServer() error {
	// create new client and wait for connection
	c, err := newConnectedClient()
	if err != nil {
		return fmt.Errorf("Failed to create new client: %s", err)
	}

	// check if we need to setup
	needSetup, err := action.NeedSetup(c)
	if err != nil {
		return fmt.Errorf("Failed to check if we need to setup: %s", err)
	}

	// if we need to setup, do it
	if needSetup {
		err = action.Setup(c, username, password)
		if err != nil {
			return fmt.Errorf("Failed to setup: %s", err)
		}
	}

	// close client
	c.Close()

	return nil
}

// setupUptimeKumaContainer sets up a new Uptime Kuma container and returns the docker pool, the
// container resource and an error if any occurred.‚
func setupUptimeKumaContainer() (*dockertest.Pool, *dockertest.Resource, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, nil, fmt.Errorf("Could not connect to docker: %s", err)
	}

	// ping docker daemon
	err = pool.Client.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("Could not ping docker daemon: %s", err)
	}

	// pull and run uptime-kuma
	resource, err := pool.Run(uptimeKumaImage, uptimeKumaTag, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not start resource: %s", err)
	}

	// expirt resource after certain amount of time
	if err := resource.Expire(defaultContainerExpiration); err != nil {
		return nil, nil, fmt.Errorf("Could not expire resource: %s", err)
	}

	return pool, resource, nil
}

// readinessProbe checks if the Uptime Kuma server is ready to accept requests for the integration
// tests.
func readinessProbe() error {
	requestURL := fmt.Sprintf("http://%s:%d", uptimeKumaHost, uptimeKumaPortExternal)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// check if status code is 200 / OK
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is not 200")
	}

	return nil
}

func TestMain(m *testing.M) {
	// setup
	pool, resource, err := setupUptimeKumaContainer()
	if err != nil {
		log.Fatalf("Could not setup uptime kuma container: %s", err)
	}

	// set imposter port
	uptimeKumaPortExternalStr := resource.GetPort("3001/tcp")
	if uptimeKumaPortExternal, err = strconv.Atoi(uptimeKumaPortExternalStr); err != nil {
		log.Fatalf("Could not convert imposter port to int: %s", err)
	}

	// wait for container to be ready
	err = pool.Retry(readinessProbe)
	if err != nil {
		log.Fatalf("Could not connect to uptime kuma: %s", err)
	}

	// setup uptime kuma with default user and password
	err = setupUptimeKumaServer()
	if err != nil {
		log.Fatalf("Could not setup uptime kuma: %s", err)
	}

	// run tests
	code := m.Run()

	// teardown
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	// exit
	os.Exit(code)
}
