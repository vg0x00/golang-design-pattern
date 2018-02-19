package visitor

import (
	"testing"
)

func TestVisitorA(t *testing.T) {
	helper := TestHelper{}
	messageA := MessageA{Content: "hello", Writer: &helper}
	messageAVisitor := &MessageAVisitor{}
	messageA.Accept(messageAVisitor)
	messageA.Print()

	expectedString := "visitor A hello"
	if helper.Content != expectedString {
		t.Errorf("visitor should change messageA content, but got: %s\n",
			helper.Content)
	}
}
func TestVisitorB(t *testing.T) {
	helper := TestHelper{}
	messageB := MessageB{Content: "hello", Writer: &helper}
	messageBVisitor := &MessageBVisitor{}
	messageB.Accept(messageBVisitor)
	messageB.Print()

	expectedString := "visitor B hello"
	if helper.Content != expectedString {
		t.Errorf("visitor should change messageA content, but got: %s\n",
			helper.Content)
	}
}

type TestHelper struct {
	Content string
}

func (t *TestHelper) Write(body []byte) (int, error) {
	t.Content = string(body)
	return len(t.Content), nil
}
