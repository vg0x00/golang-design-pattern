package chainrep

import (
	"fmt"
	"os"
	"testing"
)

func TestChain(t *testing.T) {
	msg := "hello"
	writer := os.Stdout
	chainC := ChainC{Writer: writer}
	if chainC.Msg != nil {
		t.Errorf("the final chain Msg should be empty not got: %s\n", chainC.Msg)
	}
	chainB := ChainB{NextChain: &chainC}
	chainA := ChainA{NextChain: &chainB}

	chainA.Next(&msg)
	if *chainA.Msg != "hello" {
		t.Errorf("chainA Msg should be: 'hello' but got: %s\n", chainA.Msg)
	}

	if *chainC.Msg != "hello" {
		t.Errorf("the final chain msg should: 'hello', but got: %s\n", chainC.Msg)
	}
}

// not clean for unit testing
// must implement function for testing
func TestChainWithClosure(t *testing.T) {
	writer := os.Stdout
	var chainC ChainCClosure
	chainC = ChainCClosure{
		Writer: writer,
		Closure: func(msg *string) {
			fmt.Println("chain c")
			chainC.Msg = msg // side effect on  outside of block
		},
	}
	chainB := ChainBClosure{
		NextChain: &chainC,
		Closure: func(msg *string) {
			fmt.Println("chain B")
		},
	}

	chainA := ChainAClosure{
		NextChain: &chainB,
		Closure: func(msg *string) {
			fmt.Println("chain A")
		},
	}

	msg := "hello"
	chainA.Next(&msg)

	if *chainC.Msg != "hello" {
		t.Errorf("chainC msg should be 'hello' but got: %s\n", *chainC.Msg)
	}
}
