//go:build integration

package integration_test

import (
	"testing"
	"time"

	"github.com/nobbs/uptime-kuma-api/pkg/action"
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
)

func TestLogin(t *testing.T) {
	t.Run("Login with correct credentials", func(t *testing.T) {
		// create new client and wait for connection
		c, err := newConnectedClient()
		if err != nil {
			t.Fatalf("Failed to create new client: %s", err)
		}
		defer c.Close()

		// login
		if _, err = action.Login(c, username, password, ""); err != nil {
			t.Fatalf("Failed to login: %s", err)
		}
	})

	t.Run("Login with wrong credentials", func(t *testing.T) {
		// create new client and wait for connection
		c, err := newConnectedClient()
		if err != nil {
			t.Fatalf("Failed to create new client: %s", err)
		}
		defer c.Close()

		// login with wrong password
		_, err = action.Login(c, username, "wrongpassword", "")
		if err == nil {
			t.Fatalf("Login with wrong password should fail")
		}
	})
}

func TestLoginByToken(t *testing.T) {
	var token string

	t.Run("Login to get token", func(t *testing.T) {
		// create new client and wait for connection
		c, err := newConnectedClient()
		if err != nil {
			t.Fatalf("Failed to create new client: %s", err)
		}
		defer c.Close()

		// login to get valid token
		if token, err = action.Login(c, username, password, ""); err != nil {
			t.Fatalf("Failed to login: %s", err)
		}
	})

	t.Run("Login with valid token", func(t *testing.T) {
		// create new client and wait for connection
		c, err := newConnectedClient()
		if err != nil {
			t.Fatalf("Failed to create new client: %s", err)
		}
		defer c.Close()

		// login by token
		err = action.LoginByToken(c, token)
		if err != nil {
			t.Fatalf("Failed to login by token: %s", err)
		}
	})

	t.Run("Login with invalid token", func(t *testing.T) {
		// create new client and wait for connection
		c, err := newConnectedClient()
		if err != nil {
			t.Fatalf("Failed to create new client: %s", err)
		}
		defer c.Close()

		// login by token with wrong token
		err = action.LoginByToken(c, "wrong")
		if err == nil {
			t.Fatalf("Login by token should fail")
		}
	})
}

func TestAutoLogin(t *testing.T) {
	const awaitTimeout = time.Duration(1) * time.Second

	t.Run("Auto login disabled", func(t *testing.T) {
		// create new client and wait for connection
		c, err := newConnectedClient()
		if err != nil {
			t.Fatalf("Failed to create new client: %s", err)
		}
		defer c.Close()

		// wait for auto login event to not happen
		if err = c.Await(handler.AutoLoginEvent, awaitTimeout); err == nil {
			t.Fatalf("Auto login event should not happen")
		}
	})
}
