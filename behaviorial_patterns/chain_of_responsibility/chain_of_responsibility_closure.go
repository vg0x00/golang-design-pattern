package chainrep

import "io"

type ChainAClosure struct {
	NextChain ChainOfResposibility
	Msg       *string
	Closure   func(*string)
}

func (c *ChainAClosure) Next(msg *string) {
	if c.Closure != nil {
		c.Closure(msg)
	}

	if c.NextChain != nil {
		c.NextChain.Next(msg)
	}
}

type ChainBClosure struct {
	NextChain ChainOfResposibility
	Closure   func(*string)
}

func (c *ChainBClosure) Next(msg *string) error {
	if c.Closure != nil {
		c.Closure(msg)
	}

	if c.NextChain != nil {
		c.NextChain.Next(msg)
	}

	return nil
}

type ChainCClosure struct {
	NextChain ChainOfResposibility
	Writer    io.Writer
	Msg       *string
	Closure   func(*string)
}

func (c *ChainCClosure) Next(msg *string) error {
	if c.Closure != nil {
		c.Closure(msg)
	}

	if c.NextChain != nil {
		c.NextChain.Next(msg)
	}
	return nil
}

func (c *ChainCClosure) Write(body []byte) (int, error) {
	if c.Writer != nil {
		c.Write(body)
	}

	message := string(body)
	c.Msg = &message

	return len(body), nil
}
