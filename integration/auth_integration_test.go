//go:build integration

package integration_test

import (
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/nobbs/uptime-kuma-api/pkg/action"
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
	"github.com/pquerna/otp/totp"
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

	t.Run("Auto login enabled", func(t *testing.T) {
		// create new client and wait for connection
		c, err := newConnectedClient()
		if err != nil {
			t.Fatalf("Failed to create new client: %s", err)
		}

		// login
		if _, err := action.Login(c, username, password, ""); err != nil {
			t.Fatalf("Failed to login: %s", err)
		}

		// enable auto login by setting disableAuth to true
		if err := action.SetSettings(c, &action.Settings{
			DisableAuth: utils.NewBool(true),
		}, password); err != nil {
			t.Fatalf("Failed to set settings: %s", err)
		}
		// close client
		c.Close()

		// create new client and wait for connection
		c, err = newConnectedClient()
		if err != nil {
			t.Fatalf("Failed to create new client: %s", err)
		}
		defer c.Close()

		// wait for auto login event to happen
		if err := c.Await(handler.AutoLoginEvent, awaitTimeout); err != nil {
			t.Fatalf("Auto login event should happen")
		}

		// enable auth again
		if err := action.SetSettings(c, &action.Settings{
			DisableAuth: utils.NewBool(false),
		}, password); err != nil {
			t.Fatalf("Failed to set settings: %s", err)
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

func Test2fa(t *testing.T) {
	t.Run("Prepare 2fa, but don't enable it", func(t *testing.T) {
		c, err := newLoggedInClient()
		if err != nil {
			t.Fatalf("Failed to create new client: %s", err)
		}
		defer c.Close()

		// prepare 2fa
		totpUri, err := action.Prepare2FA(c, password)
		if err != nil {
			t.Fatalf("Failed to prepare 2fa: %s", err)
		}

		// check if uri is valid
		if !strings.HasPrefix(totpUri, "otpauth://totp/Uptime%20Kuma") {
			t.Fatalf("Invalid totp uri: %s", totpUri)
		}

		// check if 2fa is not enabled
		status, err := action.TwoFAStatus(c)
		if err != nil {
			t.Fatalf("Failed to get 2fa status: %s", err)
		}

		if status {
			t.Fatalf("2fa should not be enabled")
		}
	})

	t.Run("Prepare 2fa, enable it and disable it", func(t *testing.T) {
		c, err := newLoggedInClient()
		if err != nil {
			t.Fatalf("Failed to create new client: %s", err)
		}
		defer c.Close()

		// prepare 2fa
		totpUri, err := action.Prepare2FA(c, password)
		if err != nil {
			t.Fatalf("Failed to prepare 2fa: %s", err)
		}

		// check if uri is valid
		if !strings.HasPrefix(totpUri, "otpauth://totp/Uptime%20Kuma") {
			t.Fatalf("Invalid totp uri: %s", totpUri)
		}

		// enable 2fa
		if err := action.Save2FA(c, password); err != nil {
			t.Fatalf("Failed to enable 2fa: %s", err)
		}

		// check if 2fa is enabled
		status, err := action.TwoFAStatus(c)
		if err != nil {
			t.Fatalf("Failed to get 2fa status: %s", err)
		}

		if !status {
			t.Fatalf("2fa should be enabled")
		}

		// disable 2fa
		if err := action.Disable2FA(c, password); err != nil {
			t.Fatalf("Failed to disable 2fa: %s", err)
		}

		// check if 2fa is disabled
		status, err = action.TwoFAStatus(c)
		if err != nil {
			t.Fatalf("Failed to get 2fa status: %s", err)
		}

		if status {
			t.Fatalf("2fa should be disabled")
		}
	})

	t.Run("Prepare 2fa, enable it and login with 2fa", func(t *testing.T) {
		c, err := newLoggedInClient()
		if err != nil {
			t.Fatalf("Failed to create new client: %s", err)
		}
		defer c.Close()

		// prepare 2fa
		totpUri, err := action.Prepare2FA(c, password)
		if err != nil {
			t.Fatalf("Failed to prepare 2fa: %s", err)
		}

		// check if uri is valid
		if !strings.HasPrefix(totpUri, "otpauth://totp/Uptime%20Kuma") {
			t.Fatalf("Invalid totp uri: %s", totpUri)
		}

		// enable 2fa
		if err := action.Save2FA(c, password); err != nil {
			t.Fatalf("Failed to enable 2fa: %s", err)
		}

		// verify 2fa, should fail as it's invalid
		valid, err := action.VerifyToken(c, password, "wrongtoken")
		if err == nil {
			t.Fatalf("Token should be invalid")
		}

		if valid {
			t.Fatalf("Token should not be valid")
		}

		// parse secret from uri
		u, err := url.Parse(totpUri)
		if err != nil {
			t.Fatalf("Failed to parse totp uri: %s", err)
		}

		// compute valid token
		secret := u.Query().Get("secret")
		token, err := totp.GenerateCode(secret, time.Now())
		if err != nil {
			t.Fatalf("Failed to generate token: %s", err)
		}

		// verify token, should be valid
		valid, err = action.VerifyToken(c, password, token)
		if err != nil {
			t.Fatalf("Failed to verify token: %s", err)
		}

		if !valid {
			t.Fatalf("Token should be valid")
		}

		// logout
		if err := action.Logout(c); err != nil {
			t.Fatalf("Failed to logout: %s", err)
		}

		// login again with 2fa
		if _, err := action.Login(c, username, password, token); err != nil {
			t.Fatalf("Failed to login: %s", err)
		}

		// check if 2fa is enabled
		status, err := action.TwoFAStatus(c)
		if err != nil {
			t.Fatalf("Failed to get 2fa status: %s", err)
		}

		if !status {
			t.Fatalf("2fa should be enabled")
		}

		// disable 2fa
		if err := action.Disable2FA(c, password); err != nil {
			t.Fatalf("Failed to disable 2fa: %s", err)
		}
	})
}
