//go:build integration

package integration_test

import (
	"reflect"
	"testing"

	"github.com/nobbs/uptime-kuma-api/pkg/action"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
)

func TestTags(t *testing.T) {
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

	c, err := newLoggedInClient()
	if err != nil {
		t.Fatalf("Failed to create new client: %s", err)
	}

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

	// modify first tag
	wantTags[0].Name = "Test tag edited"
	wantTags[0].Color = "#0000ff"

	t.Run("Edit first tag", func(t *testing.T) {
		// edit tag
		tag, err := action.EditTag(c, wantTags[0].Id, wantTags[0].Name, wantTags[0].Color)
		if err != nil {
			t.Fatalf("Failed to edit tag: %s", err)
		}

		switch {
		case tag == nil:
			t.Fatalf("Expected tag, got none")
		case !reflect.DeepEqual(tag, &wantTags[0]):
			t.Fatalf("Expected tag to be '%v', got '%v'", &wantTags[0], tag)
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
