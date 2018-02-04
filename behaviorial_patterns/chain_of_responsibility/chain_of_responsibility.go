package chainrep

import (
	"fmt"
	"io"
)

type ChainOfResposibility interface {
	Next(msg *string) error
}
type ChainA struct {
	Msg       *string
	NextChain ChainOfResposibility
}

func (c *ChainA) Next(msg *string) error {
	c.Msg = msg
	fmt.Printf("chain A: %s\n", *c.Msg)
	if c.NextChain != nil {
		c.NextChain.Next(msg)
	}
	return nil
}

type ChainB struct {
	NextChain ChainOfResposibility
}

func (c *ChainB) Next(msg *string) error {
	fmt.Println("chain B")
	if c.NextChain != nil {
		c.NextChain.Next(msg)
	}
	return nil
}

type ChainC struct {
	Msg       *string
	NextChain ChainOfResposibility
	Writer    io.Writer
}

func (c *ChainC) Write(body []byte) (int, error) {
	msg := string(body)
	c.Msg = &msg
	c.Writer.Write(body)
	return len(body), nil
}

func (c *ChainC) Next(msg *string) error {
	fmt.Println("chain C")
	if c.Writer != nil {
		c.Write([]byte(*msg))
	}
	return nil
}
