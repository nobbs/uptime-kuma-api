//go:build e2e

package e2e_test

import (
	"testing"
	"time"

	"github.com/nobbs/uptime-kuma-api/e2e/testutil"
	"github.com/nobbs/uptime-kuma-api/pkg/action"
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestMonitor_Basic(t *testing.T) {
	t.Parallel()

	const (
		username string = "testuser"
		password string = "testpassword123"
	)

	// create new uptime-kuma server
	server, err := testutil.NewUptimeKumaServerWithUserSetup(username, password)
	assert.NoError(t, err, "Should not return error")
	t.Cleanup(server.Teardown)

	c, err := server.NewClientWithLoginByUsernameAndPassword()
	assert.NoError(t, err, "Should not return error")

	defer c.Close()

	t.Run("Trigger get monitor list event, wait for it; should be empty", func(t *testing.T) {
		err := action.GetMonitorList(c)
		assert.NoError(t, err, "Should not return error")

		err = c.Await(handler.MonitorListEvent, time.Duration(3)*time.Second)
		assert.NoError(t, err, "Should not return error")

		monitors, err := c.State().Monitors()
		assert.NotNil(t, monitors, "Should not be nil")
		assert.NoError(t, err, "Should not return error")
		assert.Empty(t, monitors, "Should be empty")
	})

	t.Run("Add monitor, trigger get monitor list event, wait for it; should be one", func(t *testing.T) {
		monitor := &state.Monitor{
			Type:                utils.NewString("http"),
			Name:                utils.NewString("Test Monitor - google.com"),
			Url:                 utils.NewString("https://google.com"),
			AcceptedStatuscodes: []string{"200"},

			Interval:      utils.NewInt(60),
			RetryInterval: utils.NewInt(60),
			Maxretries:    utils.NewInt(0),
			Maxredirects:  utils.NewInt(0),
		}

		id, err := action.AddMonitor(c, monitor)
		assert.NoError(t, err, "Should not return error")
		assert.Equal(t, 1, id, "Should be 1")

		got, err := c.State().Monitor(id)
		assert.NoError(t, err, "Should not return error")
		assert.NotNil(t, got, "Should not be nil")
		assert.Equal(t, *monitor.Name, *got.Name, "Should have same name")
		assert.Equal(t, *monitor.Url, *got.Url, "Should have same url")
	})
}

func TestMonitor_Green(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	const (
		username string = "testuser"
		password string = "testpassword123"
	)

	// create new uptime-kuma server
	server, err := testutil.NewUptimeKumaServerWithUserSetup(username, password)
	assert.NoError(t, err, "Should not return error")
	t.Cleanup(server.Teardown)

	c, err := server.NewClientWithLoginByUsernameAndPassword()
	assert.NoError(t, err, "Should not return error")

	defer c.Close()

	t.Run("Add monitor for 'google.com', wait for heartbeats", func(t *testing.T) {
		const timeout = time.Duration(25) * time.Second

		monitor := &state.Monitor{
			Type:                utils.NewString("http"),
			Maxredirects:        utils.NewInt(0),
			Maxretries:          utils.NewInt(0),
			Interval:            utils.NewInt(20),
			RetryInterval:       utils.NewInt(20),
			Name:                utils.NewString("Test Monitor - google.com"),
			Url:                 utils.NewString("https://google.com"),
			AcceptedStatuscodes: []string{},
		}

		id, err := action.AddMonitor(c, monitor)
		assert.NoError(t, err, "Should not return error")
		assert.Equal(t, 1, id, "Should be 1")

		// wait for heartbeats
		err = c.Await(handler.HeartbeatEvent, timeout)
		assert.NoError(t, err, "Should not return error")

		// get heartbeats from state
		heartbeats, err := c.State().Heartbeats(id)
		assert.NoError(t, err, "Should not return error")
		assert.NotNil(t, heartbeats, "Should not be nil")
		assert.NotEmpty(t, heartbeats, "Should not be empty")
		assert.Len(t, heartbeats, 1, "Should have one heartbeat")

		// manually wait for next heartbeat event
		<-time.After(timeout)

		// get heartbeats from state
		heartbeats, err = c.State().Heartbeats(id)
		assert.NoError(t, err, "Should not return error")
		assert.NotNil(t, heartbeats, "Should not be nil")
		assert.NotEmpty(t, heartbeats, "Should not be empty")
		assert.Len(t, heartbeats, 2, "Should have two heartbeats")

		// delete monitor
		err = action.DeleteMonitor(c, id)
		assert.NoError(t, err, "Should not return error")

		// wait for monitor list event
		err = action.GetMonitorList(c)
		assert.NoError(t, err, "Should not return error")

		monitors, err := c.State().Monitors()
		assert.NoError(t, err, "Should not return error")
		assert.NotNil(t, monitors, "Should not be nil")
		assert.Empty(t, monitors, "Should be empty")
	})
}

func TestMonitor_Tags(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	const (
		username string = "testuser"
		password string = "testpassword123"
	)

	// create new uptime-kuma server
	server, err := testutil.NewUptimeKumaServerWithUserSetup(username, password)
	assert.NoError(t, err, "Should not return error")
	t.Cleanup(server.Teardown)

	c, err := server.NewClientWithLoginByUsernameAndPassword()
	assert.NoError(t, err, "Should not return error")

	defer c.Close()

	// add monitor
	monitor := &state.Monitor{
		Type:                utils.NewString("http"),
		Maxredirects:        utils.NewInt(0),
		Maxretries:          utils.NewInt(0),
		Interval:            utils.NewInt(20),
		RetryInterval:       utils.NewInt(20),
		Name:                utils.NewString("Test Monitor - google.com"),
		Url:                 utils.NewString("https://google.com"),
		AcceptedStatuscodes: []string{},
	}

	id, err := action.AddMonitor(c, monitor)
	assert.NoError(t, err, "Should not return error")
	assert.Equal(t, 1, id, "Should be 1")

	// add wantTags
	wantTags := []state.Tag{
		{
			Name:  "tag1",
			Color: "#ff0000",
		},
		{
			Name:  "tag2",
			Color: "#00ff00",
		},
	}

	var gotTags []state.Tag

	for _, wantTag := range wantTags {
		gotTag, err := action.AddTag(c, wantTag.Name, wantTag.Color)
		assert.NoError(t, err, "Should not return error")
		assert.NotNil(t, gotTag, "Should not be nil")
		assert.Equal(t, wantTag.Name, gotTag.Name, "Should have same name")
		assert.Equal(t, wantTag.Color, gotTag.Color, "Should have same color")

		gotTags = append(gotTags, *gotTag)
	}

	// add tags to monitor
	err = action.AddMonitorTag(c, id, gotTags[0].Id, "")
	assert.NoError(t, err, "Should not return error")

	err = action.AddMonitorTag(c, id, gotTags[1].Id, "val")
	assert.NoError(t, err, "Should not return error")

	// get monitor from uptime-kuma
	monitor, err = action.GetMonitor(c, id)
	assert.NoError(t, err, "Should not return error")
	assert.NotNil(t, monitor, "Should not be nil")
	assert.NotNil(t, monitor.Tags, "Should not be nil")
	assert.Len(t, monitor.Tags, 2, "Should have two tags")

	// edit first tag
	gotTagEdited, err := action.EditTag(c, gotTags[0].Id, "tag1 edited", "#0000ff")
	assert.NoError(t, err, "Should not return error")
	assert.NotNil(t, gotTagEdited, "Should not be nil")

	// check if monitor tag was updated
	monitor, err = action.GetMonitor(c, id)
	assert.NoError(t, err, "Should not return error")
	assert.NotNil(t, monitor, "Should not be nil")
	assert.NotNil(t, monitor.Tags, "Should not be nil")
	assert.Len(t, monitor.Tags, 2, "Should have two tags")
	assert.Equal(t, gotTagEdited.Name, monitor.Tags[0].Name, "Should have same name")
	assert.Equal(t, gotTagEdited.Color, monitor.Tags[0].Color, "Should have same color")

	// delete second tag
	err = action.DeleteTag(c, gotTags[1].Id)
	assert.NoError(t, err, "Should not return error")

	// check if monitor tag was deleted
	monitor, err = action.GetMonitor(c, id)
	assert.NoError(t, err, "Should not return error")
	assert.NotNil(t, monitor, "Should not be nil")
	assert.NotNil(t, monitor.Tags, "Should not be nil")
	assert.Len(t, monitor.Tags, 1, "Should have one tag")

	// delete monitor tags
	for _, tag := range monitor.Tags {
		err = action.DeleteMonitorTag(c, id, tag.Id, *tag.Value)
		assert.NoError(t, err, "Should not return error")
	}
}
