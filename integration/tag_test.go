//go:build integration

package integration_test

import (
	"reflect"
	"testing"

	"github.com/nobbs/uptime-kuma-api/pkg/action"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
	"github.com/nobbs/uptime-kuma-api/testutil"
)

func TestTags(t *testing.T) {
	t.Parallel()

	const (
		username string = "testuser"
		password string = "testpassword123"
	)

	// create new uptime-kuma server
	server, err := testutil.NewUptimeKumaServerWithUserSetup(username, password)
	if err != nil {
		panic(err)
	}

	t.Cleanup(server.Teardown)

	wantTags := []state.Tag{
		{
			Id:    1,
			Name:  "Test tag",
			Color: "#ff0000",
		},
		{
			Id:    2,
			Name:  "Test tag 2",
			Color: "#00ff00",
		},
	}
	modifiedTag := state.Tag{
		Id:    1,
		Name:  "Test tag modified",
		Color: "#0000ff",
	}

	c, err := server.NewClientWithLoginByUsernameAndPassword()
	if err != nil {
		t.Fatalf("Failed to create new client: %s", err)
	}
	defer c.Close()

	t.Run("Get tags, should be empty", func(t *testing.T) {
		// get tags
		tags, err := action.GetTags(c)
		if err != nil {
			t.Fatalf("Failed to get tags: %s", err)
		}

		// check if tags is empty
		if len(tags) != 0 {
			t.Fatalf("Expected no tags, got %d", len(tags))
		}
	})

	t.Run("Add first tag", func(t *testing.T) {
		// add tag
		tag, err := action.AddTag(c, wantTags[0].Name, wantTags[0].Color)
		if err != nil {
			t.Fatalf("Failed to add tag: %s", err)
		}

		switch {
		case tag == nil:
			t.Fatalf("Expected tag, got none")
		case !reflect.DeepEqual(tag, &wantTags[0]):
			t.Fatalf("Expected tag to be '%v', got '%v'", &wantTags[0], tag)
		}
	})

	t.Run("Get list of tags, should be 1", func(t *testing.T) {
		// get tags
		tags, err := action.GetTags(c)
		if err != nil {
			t.Fatalf("Failed to get tags: %s", err)
		}

		switch {
		case len(tags) != 1:
			t.Fatalf("Expected 1 tag, got %d", len(tags))
		case !reflect.DeepEqual(tags, wantTags[:1]):
			t.Fatalf("Expected tags to be '%v', got '%v'", wantTags[:1], tags)
		}
	})

	t.Run("Add second tag", func(t *testing.T) {
		// add tag
		tag, err := action.AddTag(c, wantTags[1].Name, wantTags[1].Color)
		if err != nil {
			t.Fatalf("Failed to add tag: %s", err)
		}

		switch {
		case tag == nil:
			t.Fatalf("Expected tag, got none")
		case !reflect.DeepEqual(tag, &wantTags[1]):
			t.Fatalf("Expected tag to be '%v', got '%v'", &wantTags[1], tag)
		}
	})

	t.Run("Get list of tags, should be 2", func(t *testing.T) {
		// get tags
		tags, err := action.GetTags(c)
		if err != nil {
			t.Fatalf("Failed to get tags: %s", err)
		}

		switch {
		case len(tags) != 2:
			t.Fatalf("Expected 2 tags, got %d", len(tags))
		case !reflect.DeepEqual(tags, wantTags):
			t.Fatalf("Expected tags to be '%v', got '%v'", wantTags, tags)
		}
	})

	t.Run("Edit first tag", func(t *testing.T) {
		// edit tag
		tag, err := action.EditTag(c, modifiedTag.Id, modifiedTag.Name, modifiedTag.Color)
		if err != nil {
			t.Fatalf("Failed to edit tag: %s", err)
		}

		switch {
		case tag == nil:
			t.Fatalf("Expected tag, got none")
		case !reflect.DeepEqual(tag, &modifiedTag):
			t.Fatalf("Expected tag to be '%v', got '%v'", &modifiedTag, tag)
		}
	})

	t.Run("Delete first tag", func(t *testing.T) {
		// delete tag
		err := action.DeleteTag(c, wantTags[0].Id)
		if err != nil {
			t.Fatalf("Failed to delete tag: %s", err)
		}
	})

	t.Run("Get list of tags, should be 1", func(t *testing.T) {
		// get tags
		tags, err := action.GetTags(c)
		if err != nil {
			t.Fatalf("Failed to get tags: %s", err)
		}

		switch {
		case len(tags) != 1:
			t.Fatalf("Expected 1 tag, got %d", len(tags))
		case !reflect.DeepEqual(tags, wantTags[1:]):
			t.Fatalf("Expected tags to be '%v', got '%v'", wantTags[1:], tags)
		}
	})
}
