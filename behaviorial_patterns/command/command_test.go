package command

import (
	"testing"
)

func TestCommand(t *testing.T) {
	command1 := CreateCommand("command 1")
	command2 := CreateCommand("command 2")
	command3 := CreateCommand("command 2")

	commandQueue := CommandQueue{}
	commandQueue.AddCommand(command1)
	commandQueue.AddCommand(command2)
	commandQueue.AddCommand(command3)

	if len(commandQueue.Commands) != 0 {
		t.Errorf("queue should clean after 3 commands added, but got: %d\n",
			len(commandQueue.Commands))
	}

	command4 := CreateCommand("command 4")

	commandQueue.AddCommand(command4)
	if len(commandQueue.Commands) != 1 {
		t.Errorf("len of command queue should be 3, but got: %d\n",
			len(commandQueue.Commands))
	}

}
