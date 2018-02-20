package state

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type State interface {
	Execute(*Context) bool
}

type Context struct {
	SecretNumber int
	TotalLife    int
	TotalTried   int
	Won          bool
	Next         State
}

type InitState struct {
}

func (s *InitState) Execute(c *Context) bool {
	rand.Seed(time.Now().UnixNano())
	c.SecretNumber = rand.Intn(10)
	runningState := RunningState{}
	c.Next = &runningState
	fmt.Println("Enter total life:")
	fmt.Fscanf(os.Stdin, "%d", &c.TotalLife)
	return true
}

type RunningState struct {
}

func (s *RunningState) Execute(c *Context) bool {
	fmt.Printf("you have %d life left\n", c.TotalLife-c.TotalTried)
	var answer int
	fmt.Println("enter your answer:")
	fmt.Fscanf(os.Stdin, "%d", &answer)
	final := EndState{}

	if answer == c.SecretNumber {
		c.Next = &final
		c.Won = true
	}

	c.TotalTried = c.TotalTried + 1
	if c.TotalTried == c.TotalLife {
		c.Next = &final
		c.Won = false
	}

	return true
}

type EndState struct {
}

func (s *EndState) Execute(c *Context) bool {
	if c.Won {
		fmt.Println("you won")
		return false
	} else {
		fmt.Println("you lose")
		return false
	}
}
