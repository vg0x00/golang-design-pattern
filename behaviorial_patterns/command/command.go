package command

import "fmt"

type Command interface {
	Execute() error
}

type ConsoleOutput struct {
	Content string
}

func (c ConsoleOutput) Execute() error {
	fmt.Println(c.Content)
	return nil
}

func CreateCommand(content string) Command {
	return &ConsoleOutput{Content: content}
}

type CommandQueue struct {
	Commands []Command
}

func (c *CommandQueue) AddCommand(command Command) {
	if c.Commands == nil {
		c.Commands = make([]Command, 0, 3)
	}

	c.Commands = append(c.Commands, command)

	if len(c.Commands) == 3 {
		for _, c := range c.Commands {
			c.Execute()
		}

		c.Commands = make([]Command, 0, 3)
	}
}
