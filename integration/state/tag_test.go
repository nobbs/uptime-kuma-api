//go:build integration

package state_test

import (
	"testing"

	"github.com/nobbs/uptime-kuma-api/pkg/state"
	"github.com/stretchr/testify/assert"
)

func TestState_Tag(t *testing.T) {
	s := state.NewState()

	// get tag by id, should be nil
	tag, err := s.Tag(11)
	assert.Error(t, err, "Should return error")
	assert.ErrorContains(t, err, "not found", "Should return error")
	assert.Nil(t, tag, "Should be nil")

	// get all tags in state, should be empty
	tags, err := s.Tags()
	assert.NoError(t, err, "Should not return error")
	assert.NotNil(t, tags, "Should not be nil")
	assert.Empty(t, tags, "Should be empty")

	// add new tag
	want := &state.Tag{
		Id:    3,
		Name:  "tag",
		Color: "#000000",
	}

	err = s.SetTag(want)
	assert.NoError(t, err, "Should not return error")

	// get tag by id
	got, err := s.Tag(want.Id)
	assert.NoError(t, err, "Should not return error")
	assert.NotNil(t, got, "Should not be nil")
	assert.Equal(t, *want, *got, "Should be equal")

	// get all tags in state
	tags, err = s.Tags()
	assert.NoError(t, err, "Should not return error")
	assert.NotNil(t, tags, "Should not be nil")
	assert.NotEmpty(t, tags, "Should not be empty")
	assert.Equal(t, *want, tags[0], "Should be equal")

	// replace with new tags
	wantMore := []state.Tag{
		{
			Id:    2,
			Name:  "tag2",
			Color: "#f000f0",
		},
		{
			Id:    5,
			Name:  "tag5",
			Color: "#0000ff",
		},
	}

	err = s.SetTags(wantMore)
	assert.NoError(t, err, "Should not return error")

	// get all tags in state
	tags, err = s.Tags()
	assert.NoError(t, err, "Should not return error")
	assert.NotNil(t, tags, "Should not be nil")
	assert.NotEmpty(t, tags, "Should not be empty")
	assert.Len(t, tags, 2, "Should have 2 tags")
	assert.NotContains(t, tags, *want, "Should contain tag")
	assert.Contains(t, tags, wantMore[0], "Should contain tag")
	assert.Contains(t, tags, wantMore[1], "Should contain tag")

	// delete tag
	err = s.DeleteTag(wantMore[0].Id)
	assert.NoError(t, err, "Should not return error")

	tags, err = s.Tags()
	assert.NoError(t, err, "Should not return error")
	assert.NotEmpty(t, tags, "Should not be empty")
	assert.Len(t, tags, 1, "Should have 1 tag")
	assert.NotContains(t, tags, wantMore[0], "Should not contain tag")
	assert.Contains(t, tags, wantMore[1], "Should contain tag")
}
