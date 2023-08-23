package state

import "fmt"

// Tag represents a tag object.
type Tag struct {
	Color string `mapstructure:"color"`
	Id    int    `mapstructure:"id"`
	Name  string `mapstructure:"name"`
}

// Tag returns the tag with the given id.
func (s *State) Tag(tagId int) (*Tag, error) {
	if s == nil {
		return nil, ErrStateNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.tags == nil {
		return nil, ErrNotSetYet
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
		return nil, ErrStateNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.tags == nil {
		return nil, ErrNotSetYet
	}

	// Convert map to slice.
	tags := make([]Tag, len(s.tags))
	for _, tag := range s.tags {
		tags = append(tags, *tag)
	}

	return tags, nil
}

// SetTags sets the tags received from Uptime Kuma.
func (s *State) SetTags(tags []Tag) error {
	if s == nil {
		return ErrStateNil
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
		return ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.tags == nil {
		return ErrNotSetYet
	}

	s.tags[tag.Id] = tag

	return nil
}

// DeleteTag deletes the tag with the given id.
func (s *State) DeleteTag(tagId int) error {
	if s == nil {
		return ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.tags == nil {
		return ErrNotSetYet
	}

	delete(s.tags, tagId)

	return nil
}
