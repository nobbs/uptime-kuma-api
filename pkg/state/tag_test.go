package state

import (
	"testing"

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
			tr := &Tag{
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
		s     *State
		setup func(*State)
	}

	type args struct {
		tagId int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Tag
		wantErr bool
	}{
		{
			name:    "nil state",
			fields:  fields{s: nil},
			args:    args{tagId: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "nil tags",
			fields: fields{
				s: NewState(),
				setup: func(s *State) {
					s.tags = nil
				},
			},
			args:    args{tagId: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "tag not found",
			fields: fields{
				s: NewState(),
				setup: func(s *State) {
					s.tags[1] = &Tag{Id: 1, Name: "tag", Color: "#000000"}
				},
			},
			args:    args{tagId: 2},
			want:    nil,
			wantErr: true,
		},
		{
			name: "tag found",
			fields: fields{
				s: NewState(),
				setup: func(s *State) {
					s.tags[1] = &Tag{Id: 1, Name: "tag", Color: "#000000"}
				},
			},
			args:    args{tagId: 1},
			want:    &Tag{Id: 1, Name: "tag", Color: "#000000"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields.s
			if tt.fields.setup != nil {
				tt.fields.setup(s)
			}

			got, err := s.Tag(tt.args.tagId)
			assert.Equal(t, tt.wantErr, err != nil, "Should have same error")
			assert.Equal(t, tt.want, got, "Should have same tag")
		})
	}
}

func TestState_Tags(t *testing.T) {
	type fields struct {
		s     *State
		setup func(*State)
	}

	tests := []struct {
		name    string
		fields  fields
		want    []Tag
		wantErr bool
	}{
		{
			name:    "nil state",
			fields:  fields{s: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name: "nil tags",
			fields: fields{
				s: NewState(),
				setup: func(s *State) {
					s.tags = nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "tags found",
			fields: fields{
				s: NewState(),
				setup: func(s *State) {
					s.tags[1] = &Tag{Id: 1, Name: "tag1", Color: "#ff0000"}
					s.tags[2] = &Tag{Id: 2, Name: "tag2", Color: "#00ff00"}
				},
			},
			want: []Tag{
				{Id: 1, Name: "tag1", Color: "#ff0000"},
				{Id: 2, Name: "tag2", Color: "#00ff00"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields.s
			if tt.fields.setup != nil {
				tt.fields.setup(s)
			}

			got, err := s.Tags()
			assert.Equal(t, tt.wantErr, err != nil, "Should have same error")
			assert.ElementsMatch(t, tt.want, got, "Should have same tags")
		})
	}
}
