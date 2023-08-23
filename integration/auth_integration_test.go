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
		if _, err := action.Login(c, username, password, ""); err != nil {
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

		// login with wrong password, should fail
		if _, err := action.Login(c, username, "wrongpassword", ""); err == nil {
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
		if err := action.LoginByToken(c, token); err != nil {
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
		if err := action.LoginByToken(c, "wrong"); err == nil {
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
		if err := c.Await(handler.AutoLoginEvent, awaitTimeout); err == nil {
			t.Fatalf("Auto login event should not happen")
		}
	})
}

func TestChangePassword(t *testing.T) {
	const temporaryPassword string = "temporarypassword321"

	// create new client and wait for connection
	c, err := newConnectedClient()
	if err != nil {
		t.Fatalf("Failed to create new client: %s", err)
	}
	defer c.Close()

	t.Run("Change password with correct current password", func(t *testing.T) {
		if _, err := action.Login(c, username, password, ""); err != nil {
			t.Fatalf("Failed to login: %s", err)
		}

		// change password
		if err := action.ChangePassword(c, password, temporaryPassword); err != nil {
			t.Fatalf("Failed to change password: %s", err)
		}

		if err := action.Logout(c); err != nil {
			t.Fatalf("Failed to logout: %s", err)
		}
	})

	t.Run("Change password with wrong current password", func(t *testing.T) {
		if _, err := action.Login(c, username, temporaryPassword, ""); err != nil {
			t.Fatalf("Failed to login: %s", err)
		}

		// change password with wrong current password, should fail
		if err := action.ChangePassword(c, "wrongpassword", password); err == nil {
			t.Fatalf("Change password with wrong current password should fail")
		}

		if err := action.Logout(c); err != nil {
			t.Fatalf("Failed to logout: %s", err)
		}
	})

	t.Run("Change password back to original password", func(t *testing.T) {
		if _, err := action.Login(c, username, temporaryPassword, ""); err != nil {
			t.Fatalf("Failed to login: %s", err)
		}

		if err := action.ChangePassword(c, temporaryPassword, password); err != nil {
			t.Fatalf("Failed to change password back to original password: %s", err)
		}

		if err := action.Logout(c); err != nil {
			t.Fatalf("Failed to logout: %s", err)
		}
	})
}

func TestLogout(t *testing.T) {
	const temporaryPassword string = "temporarypassword321"

	// create new client and wait for connection
	c, err := newConnectedClient()
	if err != nil {
		t.Fatalf("Failed to create new client: %s", err)
	}
	defer c.Close()

	if _, err := action.Login(c, username, password, ""); err != nil {
		t.Fatalf("Failed to login: %s", err)
	}

	// try to change password, should work as we are logged in
	if err := action.ChangePassword(c, password, temporaryPassword); err != nil {
		t.Fatalf("Failed to change password: %s", err)
	}

	// logout
	if err := action.Logout(c); err != nil {
		t.Fatalf("Failed to logout: %s", err)
	}

	// try to change password, should fail as we are logged out
	if err := action.ChangePassword(c, temporaryPassword, password); err == nil {
		t.Fatalf("Change password should fail as we are logged out")
	}

	if _, err := action.Login(c, username, temporaryPassword, ""); err != nil {
		t.Fatalf("Failed to login: %s", err)
	}

	// change password back to original password
	if err := action.ChangePassword(c, temporaryPassword, password); err != nil {
		t.Fatalf("Failed to change password back to original password: %s", err)
	}
}
