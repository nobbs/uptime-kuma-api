//go:build e2e

package testutil

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/nobbs/uptime-kuma-api/pkg/action"
	"github.com/nobbs/uptime-kuma-api/pkg/client"
	"github.com/ory/dockertest/v3"
)

const (
	uptimeKumaImage = "louislam/uptime-kuma" // uptimeKumaImage is the image name of the Uptime Kuma container
	uptimeKumaTag   = "1.23.0"               // uptimeKumaTag is the image tag of the Uptime Kuma container

	// defaultContainerExpiration is the default expiration time for the uptime-kuma container in
	// seconds. This is used to expire the container after a certain amount of time to prevent the
	// container from running forever in case of an error.
	defaultContainerExpiration = 360
)

// UptimeKumaServer is a struct that represents a running Uptime Kuma server used for running
// integration tests.
type UptimeKumaServer struct {
	pool     *dockertest.Pool
	resource *dockertest.Resource

	host string
	port int

	username *string
	password *string
	jwtToken *string
}

// NewUptimeKumaServer creates a new Uptime Kuma server and waits for it to be ready.
func NewUptimeKumaServer() (*UptimeKumaServer, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("Could not connect to docker: %w", err)
	}

	// ping docker daemon
	err = pool.Client.Ping()
	if err != nil {
		return nil, fmt.Errorf("Could not ping docker daemon: %w", err)
	}

	// pull and run uptime-kuma
	resource, err := pool.Run(uptimeKumaImage, uptimeKumaTag, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not start resource: %w", err)
	}

	// expirt resource after certain amount of time
	if err := resource.Expire(defaultContainerExpiration); err != nil {
		return nil, fmt.Errorf("Could not expire resource: %w", err)
	}

	// get service host
	host := resource.GetBoundIP("3001/tcp")

	// get service port
	port, err := strconv.Atoi(resource.GetPort("3001/tcp"))
	if err != nil {
		return nil, fmt.Errorf("Could not get service port: %w", err)
	}

	// wait for container to be ready
	if err := pool.Retry(func() error {
		requestURL := fmt.Sprintf("http://%s:%d", host, port)

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
	}); err != nil {
		return nil, fmt.Errorf("Could not connect to uptime kuma: %w", err)
	}

	return &UptimeKumaServer{
		pool:     pool,
		resource: resource,
		host:     host,
		port:     port,
	}, nil
}

// NewUptimeKumaServerWithUserSetup creates a new Uptime Kuma server and sets up a user with the
// given username and password.
func NewUptimeKumaServerWithUserSetup(username, password string) (*UptimeKumaServer, error) {
	// create new server
	server, err := NewUptimeKumaServer()
	if err != nil {
		return nil, fmt.Errorf("Could not create new uptime kuma server: %w", err)
	}

	// setup user
	if err := server.SetupUser(username, password); err != nil {
		return nil, fmt.Errorf("Could not setup user: %w", err)
	}

	return server, nil
}

// NewClient creates a new client for the Uptime Kuma server, without logging in.
func (s *UptimeKumaServer) NewClient() (*client.Client, error) {
	// create new client
	c, err := client.NewClient(s.host, s.port, false)
	if err != nil {
		return nil, fmt.Errorf("Failed to create new client: %w", err)
	}

	return c, nil
}

// NewClientWithLoginByUsernameAndPassword creates a new client for the Uptime Kuma server and logs
// in with the given username and password.
func (s *UptimeKumaServer) NewClientWithLoginByUsernameAndPassword() (*client.Client, error) {
	// create new client
	c, err := s.NewClient()
	if err != nil {
		return nil, fmt.Errorf("Failed to create new client: %w", err)
	}

	// login
	jwtToken, err := action.Login(c, *s.username, *s.password, "")
	if err != nil {
		return nil, fmt.Errorf("Failed to login: %w", err)
	}

	// set jwt token in struct
	s.jwtToken = &jwtToken

	return c, nil
}

// NewClientWithLoginByJWTToken creates a new client for the Uptime Kuma server and logs in with the
// JWT token from a previous username and password login.
func (s *UptimeKumaServer) NewClientWithLoginByJWTToken() (*client.Client, error) {
	// create new client
	c, err := s.NewClient()
	if err != nil {
		return nil, fmt.Errorf("Failed to create new client: %w", err)
	}

	// login
	if err := action.LoginByToken(c, *s.jwtToken); err != nil {
		return nil, fmt.Errorf("Failed to login by token: %w", err)
	}

	return c, nil
}

// SetupUser sets up a user with the given username and password as is required by Uptime Kuma on
// first run.
func (s *UptimeKumaServer) SetupUser(username, password string) error {
	// create new client and wait for connection
	c, err := s.NewClient()
	if err != nil {
		return fmt.Errorf("Failed to create new client: %w", err)
	}
	defer c.Close()

	// set username and password in struct
	s.username = &username
	s.password = &password

	// check if we need to setup
	needSetup, err := action.NeedSetup(c)
	if err != nil {
		return fmt.Errorf("Failed to check if we need to setup: %w", err)
	}

	// if we need to setup, do it
	if needSetup {
		err = action.Setup(c, *s.username, *s.password)
		if err != nil {
			return fmt.Errorf("Failed to setup: %w", err)
		}
	}

	return nil
}

// Teardown tears down the Uptime Kuma server, removing the container.
func (s *UptimeKumaServer) Teardown() {
	err := s.pool.Purge(s.resource)
	if err != nil {
		panic(err)
	}
}
