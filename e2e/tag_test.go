//go:build e2e

package e2e_test

import (
	"testing"

	"github.com/nobbs/uptime-kuma-api/e2e/testutil"
	"github.com/nobbs/uptime-kuma-api/pkg/action"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
	"github.com/stretchr/testify/assert"
)

func TestTags(t *testing.T) {
	t.Parallel()

	const (
		username string = "testuser"
		password string = "testpassword123"
	)

	// create new uptime-kuma server
	server, err := testutil.NewUptimeKumaServerWithUserSetup(username, password)
	assert.NoError(t, err, "Should not return error")
	t.Cleanup(server.Teardown)

	wantTags := []state.Tag{
		{
			Name:  "Test tag",
			Color: "#ff0000",
		},
		{
			Name:  "Test tag 2",
			Color: "#00ff00",
		},
	}

	editTag := state.Tag{
		Id:    1,
		Name:  "Test tag modified",
		Color: "#0000ff",
	}

	c, err := server.NewClientWithLoginByUsernameAndPassword()
	assert.NoError(t, err, "Should not return error")

	defer c.Close()

	t.Run("Get tags, should be empty", func(t *testing.T) {
		tags, err := action.GetTags(c)
		assert.NoError(t, err, "Should not return error")
		assert.Empty(t, tags, "Should be empty")
	})

	t.Run("Add one tag, check and delete", func(t *testing.T) {
		want := wantTags[0]

		got, err := action.AddTag(c, want.Name, want.Color)
		assert.NoError(t, err, "Should not return error")
		assert.NotNil(t, got, "Should not be nil")
		assert.Equal(t, want.Name, got.Name, "Should have same name")
		assert.Equal(t, want.Color, got.Color, "Should have same color")

		gotTags, err := action.GetTags(c)
		assert.NoError(t, err, "Should not return error")
		assert.Len(t, gotTags, 1, "Should have one tag")
		assert.Equal(t, want.Name, gotTags[0].Name, "Should have same name")
		assert.Equal(t, want.Color, gotTags[0].Color, "Should have same color")

		err = action.DeleteTag(c, got.Id)
		assert.NoError(t, err, "Should not return error")

		gotTags, err = action.GetTags(c)
		assert.NoError(t, err, "Should not return error")
		assert.Empty(t, gotTags, "Should be empty")
	})

	t.Run("Add two tags, check and delete", func(t *testing.T) {
		for i := range wantTags {
			want := wantTags[i]

			got, err := action.AddTag(c, want.Name, want.Color)
			assert.NoError(t, err, "Should not return error")
			assert.NotNil(t, got, "Should not be nil")
			assert.Equal(t, want.Name, got.Name, "Should have same name")
			assert.Equal(t, want.Color, got.Color, "Should have same color")
		}

		gotTags, err := action.GetTags(c)
		assert.NoError(t, err, "Should not return error")
		assert.Equal(t, len(wantTags), len(gotTags), "Should have same length")

		for i := range gotTags {
			assert.Equal(t, wantTags[i].Name, gotTags[i].Name, "Should have same name")
			assert.Equal(t, wantTags[i].Color, gotTags[i].Color, "Should have same color")
		}

		for _, got := range gotTags {
			err = action.DeleteTag(c, got.Id)
			assert.NoError(t, err, "Should not return error")
		}

		gotTags, err = action.GetTags(c)
		assert.NoError(t, err, "Should not return error")
		assert.Empty(t, gotTags, "Should be empty")
	})

	t.Run("Add and edit tag", func(t *testing.T) {
		want := wantTags[0]

		got, err := action.AddTag(c, want.Name, want.Color)
		assert.NoError(t, err, "Should not return error")

		tag, err := action.EditTag(c, got.Id, editTag.Name, editTag.Color)
		assert.NoError(t, err, "Should not return error")
		assert.NotNil(t, tag, "Should not be nil")
		assert.Equal(t, editTag.Name, tag.Name, "Should have same name")
		assert.Equal(t, editTag.Color, tag.Color, "Should have same color")

		gotTags, err := action.GetTags(c)
		assert.NoError(t, err, "Should not return error")
		assert.Len(t, gotTags, 1, "Should have one tag")
		assert.Equal(t, editTag.Name, gotTags[0].Name, "Should have same name")
		assert.Equal(t, editTag.Color, gotTags[0].Color, "Should have same color")

		err = action.DeleteTag(c, got.Id)
		assert.NoError(t, err, "Should not return error")

		gotTags, err = action.GetTags(c)
		assert.NoError(t, err, "Should not return error")
		assert.Empty(t, gotTags, "Should be empty")
	})
}
