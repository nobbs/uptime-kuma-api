//go:build integration

package integration_test

import (
	"testing"
	"time"

	"github.com/nobbs/uptime-kuma-api/pkg/action"
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
)

func TestMonitors(t *testing.T) {
	t.Run("Trigger get monitor list event, wait for it; should be empty", func(t *testing.T) {
		c, err := newLoggedInClient()
		if err != nil {
			t.Fatalf("Failed to create new client: %s", err)
		}
		defer c.Close()

		if err := action.GetMonitorList(c); err != nil {
			t.Fatalf("Failed to get monitor list: %s", err)
		}

		if err := c.Await(handler.MonitorListEvent, time.Duration(3)*time.Second); err != nil {
			t.Fatalf("Failed to await monitor list event: %s", err)
		}

		monitors, err := c.State().Monitors()
		if err != nil {
			t.Fatalf("Failed to get monitors: %s", err)
		}

		if len(monitors) != 0 {
			t.Fatalf("Expected 0 monitors, got %d", len(monitors))
		}
	})

	t.Run("Add monitor, trigger get monitor list event, wait for it; should be one", func(t *testing.T) {
		monitor := &state.Monitor{
			Type:                "http",
			Name:                "Test Monitor - google.com",
			Url:                 utils.NewString("https://google.com"),
			AcceptedStatuscodes: []string{"200"},

			Interval:      60,
			RetryInterval: 60,
			Maxretries:    3,
			Method:        utils.NewString("GET"),
		}

		c, err := newLoggedInClient()
		if err != nil {
			t.Fatalf("Failed to create new client: %s", err)
		}
		defer c.Close()

		id, err := action.AddMonitor(c, monitor)
		if err != nil {
			t.Fatalf("Failed to add monitor: %s", err)
		}

		if id != 1 {
			t.Fatalf("Expected monitor ID 1, got %d", id)
		}

		got, err := c.State().Monitor(id)
		if err != nil {
			t.Fatalf("Failed to get monitor: %s", err)
		}

		if got.Name != monitor.Name {
			t.Fatalf("Expected monitor name %s, got %s", monitor.Name, got.Name)
		}

		if got.Url == nil || *got.Url != *monitor.Url {
			t.Fatalf("Expected monitor url %s, got %s", *monitor.Url, *got.Url)
		}
	})
}
