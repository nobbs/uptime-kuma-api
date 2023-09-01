//go:build integration

package integration_test

import (
	"net/url"
	"regexp"
	"testing"
	"time"

	"github.com/nobbs/uptime-kuma-api/integration/testutil"
	"github.com/nobbs/uptime-kuma-api/pkg/action"
	"github.com/nobbs/uptime-kuma-api/pkg/client"
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Parallel()

	const (
		username string = "testuser"
		password string = "testpassword123"
	)

	// create new uptime-kuma server
	server, err := testutil.NewUptimeKumaServerWithUserSetup(username, password)
	assert.NoError(t, err, "Should not return error")
	t.Cleanup(server.Teardown)

	t.Run("Login with correct credentials", func(t *testing.T) {
		// create new client and wait for connection
		c, err := server.NewClient()
		assert.NoError(t, err, "Should not return error")
		defer c.Close()

		// login
		jwt, err := action.Login(c, username, password, "")
		assert.NoError(t, err, "Should not return error")
		assert.NotEmpty(t, jwt, "Should return jwt token")
	})

	t.Run("Login with wrong credentials", func(t *testing.T) {
		// create new client and wait for connection
		c, err := server.NewClient()
		assert.NoError(t, err, "Should not return error")
		defer c.Close()

		// login with wrong password, should fail
		jwt, err := action.Login(c, username, "wrongpassword", "")
		assert.ErrorAs(t, err, &action.ErrLoginFailed{}, "Should return error")
		assert.ErrorContains(t, err, "Incorrect username or password", "Should return error")
		assert.Empty(t, jwt, "Should not return jwt token")
	})
}

func TestLoginByToken(t *testing.T) {
	t.Parallel()

	const (
		username string = "testuser"
		password string = "testpassword123"
	)

	var jwt string

	// create new uptime-kuma server
	server, err := testutil.NewUptimeKumaServerWithUserSetup(username, password)
	assert.NoError(t, err, "Should not return error")
	t.Cleanup(server.Teardown)

	// create new client and wait for connection
	c, err := server.NewClient()
	assert.NoError(t, err, "Should not return error")

	// login to get valid token
	jwt, err = action.Login(c, username, password, "")
	assert.NoError(t, err, "Should not return error")
	assert.NotEmpty(t, jwt, "Should return jwt token")

	c.Close()

	t.Run("Login with valid token", func(t *testing.T) {
		c, err := server.NewClient()
		assert.NoError(t, err, "Should not return error")
		defer c.Close()

		err = action.LoginByToken(c, jwt)
		assert.NoError(t, err, "Should not return error")
	})

	t.Run("Login with invalid token", func(t *testing.T) {
		c, err := server.NewClient()
		assert.NoError(t, err, "Should not return error")
		defer c.Close()

		// login by token with wrong token
		err = action.LoginByToken(c, "wrong")
		assert.ErrorAs(t, err, &action.ErrLoginFailed{}, "Should return error")
		assert.ErrorContains(t, err, "Invalid token", "Should return error")
	})
}

func TestAutoLogin(t *testing.T) {
	t.Parallel()

	const (
		username     string        = "testuser"
		password     string        = "testpassword123"
		awaitTimeout time.Duration = time.Duration(1) * time.Second
	)

	server, err := testutil.NewUptimeKumaServerWithUserSetup(username, password)
	assert.NoError(t, err, "Should not return error")
	t.Cleanup(server.Teardown)

	t.Run("Auto login disabled", func(t *testing.T) {
		c, err := server.NewClient()
		assert.NoError(t, err, "Should not return error")
		defer c.Close()

		// wait for auto login event to not happen
		err = c.Await(handler.AutoLoginEvent, awaitTimeout)
		assert.ErrorIs(t, err, client.ErrTimeout, "Auto login event should not happen")
	})

	t.Run("Auto login enabled", func(t *testing.T) {
		settings := &action.Settings{
			DisableAuth: utils.NewBool(true),
		}

		c, err := server.NewClient()
		assert.NoError(t, err, "Should not return error")

		_, err = action.Login(c, username, password, "")
		assert.NoError(t, err, "Should not return error")

		// enable auto login by setting disableAuth to true
		err = action.SetSettings(c, settings, password)
		assert.NoError(t, err, "Should not return error")
		c.Close()

		c, err = server.NewClient()
		assert.NoError(t, err, "Should not return error")
		defer c.Close()

		err = c.Await(handler.AutoLoginEvent, awaitTimeout)
		assert.NoError(t, err, "Should not return error")
	})
}

func TestChangePassword(t *testing.T) {
	t.Parallel()

	const (
		username          string = "testuser"
		password          string = "testpassword123"
		temporaryPassword string = "temporarypassword321"
	)

	server, err := testutil.NewUptimeKumaServerWithUserSetup(username, password)
	assert.NoError(t, err, "Should not return error")
	t.Cleanup(server.Teardown)

	c, err := server.NewClient()
	assert.NoError(t, err, "Should not return error")

	defer c.Close()

	// change password to temporary password
	_, err = action.Login(c, username, password, "")
	assert.NoError(t, err, "Should not return error")

	err = action.ChangePassword(c, password, temporaryPassword)
	assert.NoError(t, err, "Should not return error")

	err = action.Logout(c)
	assert.NoError(t, err, "Should not return error")

	// try to change password given wrong password, should fail
	_, err = action.Login(c, username, temporaryPassword, "")
	assert.NoError(t, err, "Should not return error")

	err = action.ChangePassword(c, "wrongpassword", password)
	assert.ErrorAs(t, err, &action.ErrActionFailed{}, "Should return error")
	assert.ErrorContains(t, err, "Incorrect current password", "Should return error")

	err = action.Logout(c)
	assert.NoError(t, err, "Should not return error")

	// change password back to original password
	_, err = action.Login(c, username, temporaryPassword, "")
	assert.NoError(t, err, "Should not return error")

	err = action.ChangePassword(c, temporaryPassword, password)
	assert.NoError(t, err, "Should not return error")

	err = action.Logout(c)
	assert.NoError(t, err, "Should not return error")
}

func TestLogout(t *testing.T) {
	t.Parallel()

	const (
		username          string = "testuser"
		password          string = "testpassword123"
		temporaryPassword string = "temporarypassword321"
	)

	// create new uptime-kuma server
	server, err := testutil.NewUptimeKumaServerWithUserSetup(username, password)
	assert.NoError(t, err, "Should not return error")
	t.Cleanup(server.Teardown)

	// create new client and wait for connection
	c, err := server.NewClient()
	assert.NoError(t, err, "Should not return error")

	defer c.Close()

	_, err = action.Login(c, username, password, "")
	assert.NoError(t, err, "Should not return error")

	// try to change password, should work as we are logged in
	err = action.ChangePassword(c, password, temporaryPassword)
	assert.NoError(t, err, "Should not return error")

	// logout
	err = action.Logout(c)
	assert.NoError(t, err, "Should not return error")

	// try to change password, should fail as we are logged out
	err = action.ChangePassword(c, temporaryPassword, password)
	assert.ErrorAs(t, err, &action.ErrActionFailed{}, "Should return error")
	assert.ErrorContains(t, err, "You are not logged in", "Should return error")

	_, err = action.Login(c, username, temporaryPassword, "")
	assert.NoError(t, err, "Should not return error")

	// change password back to original password
	err = action.ChangePassword(c, temporaryPassword, password)
	assert.NoError(t, err, "Should not return error")

	err = action.Logout(c)
	assert.NoError(t, err, "Should not return error")
}

func Test2fa(t *testing.T) {
	t.Parallel()

	const (
		username string = "testuser"
		password string = "testpassword123"
	)

	// create new uptime-kuma server
	server, err := testutil.NewUptimeKumaServerWithUserSetup(username, password)
	assert.NoError(t, err, "Should not return error")
	t.Cleanup(server.Teardown)

	t.Run("Prepare 2fa, but don't enable it", func(t *testing.T) {
		c, err := server.NewClientWithLoginByUsernameAndPassword()
		assert.NoError(t, err, "Should not return error")
		defer c.Close()

		// prepare 2fa, check if URI is valid
		totpUri, err := action.Prepare2FA(c, password)
		assert.NoError(t, err, "Should not return error")
		assert.Regexp(t, regexp.MustCompile("^otpauth://totp/Uptime.*"), totpUri, "Should match regex")

		// check if 2fa is not enabled
		status, err := action.TwoFAStatus(c)
		assert.NoError(t, err, "Should not return error")
		assert.False(t, status, "Should not be enabled")
	})

	t.Run("Prepare 2fa, enable it and disable it", func(t *testing.T) {
		c, err := server.NewClientWithLoginByUsernameAndPassword()
		assert.NoError(t, err, "Should not return error")
		defer c.Close()

		// prepare 2fa, check if URI is valid
		totpUri, err := action.Prepare2FA(c, password)
		assert.NoError(t, err, "Should not return error")
		assert.Regexp(t, regexp.MustCompile("^otpauth://totp/Uptime.*"), totpUri, "Should match regex")

		// enable 2fa
		err = action.Save2FA(c, password)
		assert.NoError(t, err, "Should not return error")

		// check if 2fa is enabled
		status, err := action.TwoFAStatus(c)
		assert.NoError(t, err, "Should not return error")
		assert.True(t, status, "Should be enabled")

		// disable 2fa, check if 2fa is disabled
		err = action.Disable2FA(c, password)
		assert.NoError(t, err, "Should not return error")

		// check if 2fa is disabled
		status, err = action.TwoFAStatus(c)
		assert.NoError(t, err, "Should not return error")
		assert.False(t, status, "Should not be enabled")
	})

	t.Run("Prepare 2fa, enable it and login with 2fa", func(t *testing.T) {
		c, err := server.NewClientWithLoginByUsernameAndPassword()
		assert.NoError(t, err, "Should not return error")
		defer c.Close()

		// prepare 2fa, check if URI is valid
		totpUri, err := action.Prepare2FA(c, password)
		assert.NoError(t, err, "Should not return error")
		assert.Regexp(t, regexp.MustCompile("^otpauth://totp/Uptime.*"), totpUri, "Should match regex")

		// enable 2fa
		err = action.Save2FA(c, password)
		assert.NoError(t, err, "Should not return error")

		// verify 2fa, should fail as it's invalid
		valid, err := action.VerifyToken(c, password, "wrongtoken")
		assert.Error(t, err, "Should return error")
		assert.ErrorAs(t, err, &action.ErrActionFailed{}, "Should return error")
		assert.ErrorContains(t, err, "Invalid Token", "Should return error")
		assert.False(t, valid, "Should not be valid")

		// parse secret from uri
		u, err := url.Parse(totpUri)
		assert.NoError(t, err, "Should not return error")

		// compute valid token
		secret := u.Query().Get("secret")
		token, err := totp.GenerateCode(secret, time.Now())
		assert.NoError(t, err, "Should not return error")

		// verify token, should be valid
		valid, err = action.VerifyToken(c, password, token)
		assert.NoError(t, err, "Should not return error")
		assert.True(t, valid, "Should be valid")

		// logout
		err = action.Logout(c)
		assert.NoError(t, err, "Should not return error")

		// login again with 2fa
		_, err = action.Login(c, username, password, token)
		assert.NoError(t, err, "Should not return error")

		// check if 2fa is enabled
		status, err := action.TwoFAStatus(c)
		assert.NoError(t, err, "Should not return error")
		assert.True(t, status, "Should be enabled")

		// disable 2fa
		err = action.Disable2FA(c, password)
		assert.NoError(t, err, "Should not return error")
	})
}
