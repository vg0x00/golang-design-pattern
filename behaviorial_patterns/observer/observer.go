package observer

import (
	"errors"
	"fmt"
)

type Listener interface {
	Handle(msg string)
}

type Publisher struct {
	ListenerList []Listener
}

func (p *Publisher) Publish(msg string) {
	for _, listenr := range p.ListenerList {
		listenr.Handle(msg)
	}
}

func (p *Publisher) Register(h Listener) {
	p.ListenerList = append(p.ListenerList, h)
}

func (p *Publisher) Unregister(h Listener) error {
	var index int
	failed := true

	for i, handle := range p.ListenerList {
		if handle == h {
			index = i
			failed = false
			break
		}
	}
	// 0 1 2 3 4
	// ^   ^   ^
	// golang slice: [first, last)
	if !failed {
		// NOTE(vg0x00): syntax sugar
		// append(sliceA, sliceB...) -> append each item in sliceB into sliceA
		p.ListenerList = append(p.ListenerList[:index], p.ListenerList[index+1:]...)
		return nil
	} else {
		return errors.New("Listener not found")
	}
}

type MsgListener struct {
	Id int
}

func (m *MsgListener) Handle(msg string) {
	fmt.Printf("id: %d received msg: %s\n", m.Id, msg)
}
