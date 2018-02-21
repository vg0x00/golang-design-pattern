package observer

import "testing"

func TestObserver(t *testing.T) {
	t.Run("msg publish - listen", func(t *testing.T) {
		publisher := Publisher{}
		listener1 := MsgListener{Id: 1}
		listener2 := MsgListener{Id: 2}
		listener3 := MsgListener{Id: 3}
		publisher.Register(&listener1)
		if len(publisher.ListenerList) != 1 {
			t.Errorf("listenr list length should be 1 but got: %d\n",
				len(publisher.ListenerList))
		}
		publisher.Register(&listener2)
		publisher.Register(&listener3)
		publisher.Unregister(&listener2)
		if len(publisher.ListenerList) != 2 {
			t.Errorf("listenr list length should be 2 but got: %d\n",
				len(publisher.ListenerList))
		}

		for _, listener := range publisher.ListenerList {
			tempListener, ok := listener.(*MsgListener)
			if ok && tempListener.Id == 2 {
				t.Errorf("listener with id 2 should not exists in listener list")
			}
		}
		publisher.Publish("hello world")
	})
}
