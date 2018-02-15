package memento

import "testing"

func TestMemento(t *testing.T) {
	origin := origin{}
	state := State{Description: "Init"}
	origin.State = state

	snapshot := origin.NewSnapshot()
	if snapshot.State.Description != "Init" {
		t.Errorf("origin state description should be 'Init' but got %s\n",
			snapshot.State.Description)
	}

	snapshotTracker := SnapshotTracker{}
	currentLen := len(snapshotTracker.List)
	snapshotTracker.Add(snapshot)

	if len(snapshotTracker.List) != currentLen+1 {
		t.Errorf("snapTracker should increate 1, but got: %d\n",
			len(snapshotTracker.List))
	}

	// restore state
	preState, err := snapshotTracker.Snapshot(0)
	if err != nil {
		t.Errorf("snapshotTracker return error when fetch previous snapshot: %s\n",
			err.Error())
	}

	if preState.State.Description != "Init" {
		t.Errorf("snapshot description should be 'Init', but got: %s\n",
			preState.State.Description)
	}

	preState, err = snapshotTracker.Snapshot(-1)
	if err == nil {
		t.Errorf("fetch snap should got error with wrong index path: %s\n",
			err.Error())
	}

}
