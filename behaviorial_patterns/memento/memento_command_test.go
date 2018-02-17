package memento

import (
	"fmt"
	"testing"
)

func TestMementoWithCommand(t *testing.T) {
	t.Run("memento pattern with command", func(t *testing.T) {
		m := MementoFacade{}
		m.SaveSettings(Value(10))
		m.SaveSettings(Mute(false))
		// stack style restore
		AssertCommandType(m.StoreSettings(0))
		AssertCommandType(m.StoreSettings(1))
	})
}

func AssertCommandType(c Command) {
	switch cast := c.(type) {
	case Value:
		fmt.Printf("Value: %d\n", cast)
	case Mute:
		fmt.Printf("Mute: %t\n", cast)
	}
}
