package state

import (
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
)

const (
	capacityHeartbeats          = 150 // capacityHeartbeats is the maximum number of heartbeats to store per monitor.
	capacityImportantHeartbeats = 25  // capacityImportantHeartbeats is the maximum number of important heartbeats to store per monitor.
)

// Heartbeat represents a heartbeat object.
type Heartbeat struct {
	DownCount int    `mapstructure:"down_count"`
	Duration  int    `mapstructure:"duration"`
	Id        int    `mapstructure:"id"`
	Important bool   `mapstructure:"important"`
	MonitorId int    `mapstructure:"monitorId"` // HACK: the monitor id is sometimes `monitorId` and sometimes `monitor_id` in the heartbeat event payload.
	Msg       string `mapstructure:"msg"`
	Ping      int    `mapstructure:"ping"`
	Status    bool   `mapstructure:"status"`
	Time      string `mapstructure:"time"`
}

// HeartbeatQueue is the interface for a queue of heartbeats.
type HeartbeatQueue interface {
	Push(beat *Heartbeat)
	Slice() []Heartbeat
	Trim(capacity int)
}

// Interface guard to ensure that *utils.Queue[Heartbeat] implements HeartbeatQueue.
var _ HeartbeatQueue = (*utils.Queue[Heartbeat])(nil)

// Heartbeats returns the heartbeats received from Uptime Kuma for the given monitor id.
func (s *State) Heartbeats(monitorId int) ([]Heartbeat, error) {
	if s == nil {
		return nil, ErrStateNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.heartbeats == nil {
		return nil, ErrNotSetYet
	}

	beats, ok := s.heartbeats[monitorId]
	if !ok {
		return nil, NewErrNotFound("heartbeats", monitorId)
	}

	return beats.Slice(), nil
}

// ImportantHeartbeats returns the important heartbeats received from Uptime Kuma for the given monitor id.
func (s *State) ImportantHeartbeats(monitorId int) ([]Heartbeat, error) {
	if s == nil {
		return nil, ErrStateNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.importantHeartbeats == nil {
		return nil, ErrNotSetYet
	}

	beats, ok := s.importantHeartbeats[monitorId]
	if !ok {
		return nil, NewErrNotFound("important heartbeats", monitorId)
	}

	return beats.Slice(), nil
}

// SetHeartbeats sets the heartbeats received from Uptime Kuma for the given monitor id, optionally
// overwriting existing heartbeats.
func (s *State) SetHeartbeats(monitorId int, beats []Heartbeat, overwrite bool) error {
	if s == nil {
		return ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.heartbeats == nil {
		s.heartbeats = make(map[int]HeartbeatQueue)
	}

	// HACK: the monitor id is sometimes `monitorId` and sometimes `monitor_id` in the heartbeat
	// event payload.
	for i := range beats {
		beats[i].MonitorId = monitorId
	}

	// replace all heartbeats if overwrite is true
	if overwrite {
		s.heartbeats[monitorId] = utils.NewQueueFromSlice(beats)
		return nil
	}

	if _, ok := s.heartbeats[monitorId]; !ok {
		s.heartbeats[monitorId] = utils.NewQueue[Heartbeat]()
	}

	// add heartbeats to monitor id
	for i := range beats {
		s.heartbeats[monitorId].Push(&beats[i])
	}

	// trim queue to capacity
	s.heartbeats[monitorId].Trim(capacityHeartbeats)

	return nil
}

// SetImportantHeartbeats sets the important heartbeats received from Uptime Kuma for the given monitor id, optionally
// overwriting existing heartbeats.
func (s *State) SetImportantHeartbeats(monitorId int, beats []Heartbeat, overwrite bool) error {
	if s == nil {
		return ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.importantHeartbeats == nil {
		s.importantHeartbeats = make(map[int]HeartbeatQueue)
	}

	// HACK: the monitor id is sometimes `monitorId` and sometimes `monitor_id` in the heartbeat
	// event payload.
	for i := range beats {
		beats[i].MonitorId = monitorId
	}

	// replace all heartbeats if overwrite is true
	if overwrite {
		s.importantHeartbeats[monitorId] = utils.NewQueueFromSlice(beats)
		return nil
	}

	if _, ok := s.importantHeartbeats[monitorId]; !ok {
		s.importantHeartbeats[monitorId] = utils.NewQueue[Heartbeat]()
	}

	// add heartbeats to monitor id
	for i := range beats {
		s.importantHeartbeats[monitorId].Push(&beats[i])
	}

	// trim queue to capacity
	s.importantHeartbeats[monitorId].Trim(capacityImportantHeartbeats)

	return nil
}

func (s *State) AppendHeartbeat(beat *Heartbeat) error {
	if s == nil {
		return ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	switch beat.Important {
	case true:
		if s.importantHeartbeats == nil {
			s.importantHeartbeats = make(map[int]HeartbeatQueue)
		}

		if _, ok := s.importantHeartbeats[beat.MonitorId]; !ok {
			s.importantHeartbeats[beat.MonitorId] = utils.NewQueue[Heartbeat]()
		}

		// push heartbeat to queue and trim it to capacity
		s.importantHeartbeats[beat.MonitorId].Push(beat)
		s.importantHeartbeats[beat.MonitorId].Trim(capacityImportantHeartbeats)
	case false:
		if s.heartbeats == nil {
			s.heartbeats = make(map[int]HeartbeatQueue)
		}

		if _, ok := s.heartbeats[beat.MonitorId]; !ok {
			s.heartbeats[beat.MonitorId] = utils.NewQueue[Heartbeat]()
		}

		// push heartbeat to queue and trim it to capacity
		s.heartbeats[beat.MonitorId].Push(beat)
		s.heartbeats[beat.MonitorId].Trim(capacityHeartbeats)
	}

	return nil
}
