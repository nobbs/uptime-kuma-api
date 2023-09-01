package state_test

import (
	"testing"

	"github.com/nobbs/uptime-kuma-api/pkg/state"
	"github.com/stretchr/testify/assert"
)

func TestTag_Validate(t *testing.T) {
	type fields struct {
		Id    int
		Name  string
		Color string
		Value *string
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "valid tag",
			fields:  fields{Id: 1, Name: "tag", Color: "#000000"},
			wantErr: false,
		},
		{
			name:    "invalid tag",
			fields:  fields{Id: -1, Name: "tag", Color: "#00gg00"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &state.Tag{
				Id:    tt.fields.Id,
				Name:  tt.fields.Name,
				Color: tt.fields.Color,
				Value: tt.fields.Value,
			}

			err := tr.Validate()
			assert.Equal(t, tt.wantErr, err != nil, "Should have same error")
		})
	}
}

func TestState_Tag(t *testing.T) {
	type fields struct {
		s     *state.State
		setup func(*state.State)
	}

	type args struct {
		tagId int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *state.Tag
		wantErr bool
	}{
		{
			name: "tag exists",
			fields: fields{
				s: state.NewState(),
				setup: func(s *state.State) {
					err := s.SetTag(&state.Tag{
						Id:    1,
						Name:  "tag",
						Color: "#000000",
					})
					assert.NoError(t, err, "Should not be error")
				},
			},
			args: args{1},
			want: &state.Tag{
				Id:    1,
				Name:  "tag",
				Color: "#000000",
			},
			wantErr: false,
		},
		{
			name: "tag does not exist",
			fields: fields{
				s: state.NewState(),
			},
			args:    args{1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "state is nil",
			fields: fields{
				s: nil,
			},
			args:    args{1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "tag is nil",
			fields: fields{
				s: &state.State{},
			},
			args:    args{1},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.setup != nil {
				tt.fields.setup(tt.fields.s)
			}

			got, err := tt.fields.s.Tag(tt.args.tagId)
			assert.Equal(t, tt.wantErr, err != nil, "Should have same error")
			assert.Equal(t, tt.want, got, "Should have same tag")
		})
	}
}

func TestState_Tags(t *testing.T) {
	type fields struct {
		s     *state.State
		setup func(*state.State)
	}

	tests := []struct {
		name    string
		fields  fields
		want    []state.Tag
		wantErr bool
	}{
		{
			name: "tags exist",
			fields: fields{
				s: state.NewState(),
				setup: func(s *state.State) {
					err := s.SetTag(&state.Tag{
						Id:    1,
						Name:  "tag",
						Color: "#000000",
					})
					assert.NoError(t, err, "Should not be error")
				},
			},
			want: []state.Tag{
				{
					Id:    1,
					Name:  "tag",
					Color: "#000000",
				},
			},
			wantErr: false,
		},
		{
			name: "tags do not exist",
			fields: fields{
				s: state.NewState(),
			},
			want:    []state.Tag{},
			wantErr: false,
		},
		{
			name: "state is nil",
			fields: fields{
				s: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "tags is nil",
			fields: fields{
				s: &state.State{},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.setup != nil {
				tt.fields.setup(tt.fields.s)
			}

			got, err := tt.fields.s.Tags()
			assert.Equal(t, tt.wantErr, err != nil, "Should have same error")
			assert.Equal(t, tt.want, got, "Should have same tags")
		})
	}
}

func TestState_SetTag(t *testing.T) {
	type fields struct {
		s     *state.State
		setup func(*state.State)
	}

	type args struct {
		tag *state.Tag
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "tag is nil",
			fields: fields{
				s: &state.State{},
			},
			args:    args{nil},
			wantErr: true,
		},
		{
			name: "state is nil",
			fields: fields{
				s: nil,
			},
			args:    args{&state.Tag{}},
			wantErr: true,
		},
		{
			name: "tag is valid",
			fields: fields{
				s: state.NewState(),
			},
			args: args{&state.Tag{
				Id:    1,
				Name:  "tag",
				Color: "#000000",
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.setup != nil {
				tt.fields.setup(tt.fields.s)
			}

			err := tt.fields.s.SetTag(tt.args.tag)
			assert.Equal(t, tt.wantErr, err != nil, "Should have same error")
		})
	}
}

func TestState_SetTags(t *testing.T) {
	type fields struct {
		s     *state.State
		setup func(*state.State)
	}

	type args struct {
		tags []state.Tag
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "tags is nil",
			fields: fields{
				s: &state.State{},
			},
			args:    args{nil},
			wantErr: false,
		},
		{
			name: "state is nil",
			fields: fields{
				s: nil,
			},
			args:    args{[]state.Tag{}},
			wantErr: true,
		},
		{
			name: "tags are valid",
			fields: fields{
				s: state.NewState(),
			},
			args: args{[]state.Tag{
				{
					Id:    1,
					Name:  "tag",
					Color: "#000000",
				},
				{
					Id:    2,
					Name:  "tag2",
					Color: "#000000",
				},
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.setup != nil {
				tt.fields.setup(tt.fields.s)
			}

			err := tt.fields.s.SetTags(tt.args.tags)
			assert.Equal(t, tt.wantErr, err != nil, "Should have same error")
		})
	}
}

func TestState_DeleteTag(t *testing.T) {
	type fields struct {
		s     *state.State
		setup func(*state.State)
	}

	type args struct {
		tagId int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "tag does not exist",
			fields: fields{
				s: state.NewState(),
			},
			args:    args{1},
			wantErr: false,
		},
		{
			name: "state is nil",
			fields: fields{
				s: nil,
			},
			args:    args{1},
			wantErr: true,
		},
		{
			name: "tags is nil",
			fields: fields{
				s: &state.State{},
			},
			args:    args{1},
			wantErr: true,
		},
		{
			name: "tag exists",
			fields: fields{
				s: state.NewState(),
				setup: func(s *state.State) {
					err := s.SetTag(&state.Tag{
						Id:    1,
						Name:  "tag",
						Color: "#000000",
					})
					assert.NoError(t, err, "Should not be error")
				},
			},
			args:    args{1},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.setup != nil {
				tt.fields.setup(tt.fields.s)
			}

			err := tt.fields.s.DeleteTag(tt.args.tagId)
			assert.Equal(t, tt.wantErr, err != nil, "Should have same error")
		})
	}
}

func TestState_TagAll(t *testing.T) {
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
