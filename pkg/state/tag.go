package state

import (
	"fmt"

	"github.com/nobbs/uptime-kuma-api/pkg/utils"
	"github.com/nobbs/uptime-kuma-api/pkg/xerrors"
)

// Tag represents a tag object.
type Tag struct {
	Id    int     `mapstructure:"id"    validate:"required,gt=0"`
	Name  string  `mapstructure:"name"  validate:"required"`
	Color string  `mapstructure:"color" validate:"required,hexcolor"`
	Value *string `json:"-"             mapstructure:"value"`
}

// Validate validates the tag.
func (t *Tag) Validate() error {
	return utils.ValidateStruct(t)
}

// Tag returns the tag with the given id.
func (s *State) Tag(tagId int) (*Tag, error) {
	if s == nil {
		return nil, xerrors.ErrStateNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.tags == nil {
		return nil, xerrors.ErrNotSetYet
	}

	tag, ok := s.tags[tagId]
	if !ok {
		return nil, fmt.Errorf("tag with id %d not found", tagId)
	}

	return tag, nil
}

// Tags returns the tags received from Uptime Kuma.
func (s *State) Tags() ([]Tag, error) {
	if s == nil {
		return nil, xerrors.ErrStateNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.tags == nil {
		return nil, xerrors.ErrNotSetYet
	}

	// Convert map to slice.
	tags := make([]Tag, 0, len(s.tags))
	for i := range s.tags {
		tags = append(tags, *s.tags[i])
	}

	return tags, nil
}

// SetTags clears the current tags and sets the given tags.
func (s *State) SetTags(tags []Tag) error {
	if s == nil {
		return xerrors.ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// Convert slice to map.
	s.tags = make(map[int]*Tag)
	for i := range tags {
		s.tags[tags[i].Id] = &tags[i]
	}

	return nil
}

// SetTag sets the tag with the given id.
func (s *State) SetTag(tag *Tag) error {
	if s == nil {
		return xerrors.ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.tags == nil {
		s.tags = make(map[int]*Tag)
	}

	if tag == nil {
		return fmt.Errorf("tag is nil")
	}

	s.tags[tag.Id] = tag

	return nil
}

// DeleteTag deletes the tag with the given id.
func (s *State) DeleteTag(tagId int) error {
	if s == nil {
		return xerrors.ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.tags == nil {
		return xerrors.ErrNotSetYet
	}

	delete(s.tags, tagId)

	return nil
}
