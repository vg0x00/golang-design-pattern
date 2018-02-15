package memento

import "errors"

type State struct {
	Description string
}

// this layer between snapshotCreator and real state gives us an abstract layer
// so SnapshotTracker and origin known nothing about the real State
type snapshot struct {
	State State
}

type origin struct { // we want to track states of this process
	State State
}

func (o *origin) NewSnapshot() snapshot {
	return snapshot{State: o.State}
}

func (o *origin) ExtractAndSaveSnapshot(s snapshot) {
	o.State = s.State
}

type SnapshotTracker struct {
	List []snapshot
}

func (s *SnapshotTracker) Add(snapshot snapshot) {
	s.List = append(s.List, snapshot)
}

func (s *SnapshotTracker) Snapshot(index int) (snapshot, error) {
	if index < 0 || index > len(s.List) {
		return snapshot{}, errors.New("not implemented yet")
	}

	return s.List[index], nil
}
