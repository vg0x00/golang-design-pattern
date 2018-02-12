package template

import (
	"strings"
	"testing"
)

type UserDefinedStruct struct {
}

func (t *UserDefinedStruct) Message() string {
	return "template"
}

// using anonymous function without abstract object creation
// like HandleFunc in net/http package
type MessageHandleFunc func() string

func (t MessageHandleFunc) Message() string {
	return t()
}

func TestTemplate(t *testing.T) {
	t.Run("Template using", func(t *testing.T) {
		tester := &Template{}
		messager := &UserDefinedStruct{} // user customized code
		result := tester.ExecuteProcess(messager)

		expectedString := "template"
		if !strings.Contains(result, expectedString) {
			t.Errorf("final result shold contain user message: 'template', but got: %s",
				result)
		}
	})

	t.Run("Template using Anonymous function", func(t *testing.T) {
		tester := &AnonymousTemplate{}
		result := tester.ExecuteProcess(func() string {
			return "template"
		})

		expectedString := "template"
		if !strings.Contains(result, expectedString) {
			t.Errorf("final result shold contain user message: 'template', but got: %s",
				result)
		}

	})

	t.Run("type alias", func(t *testing.T) {
		tester := &Template{}
		result := tester.ExecuteProcess(MessageHandleFunc(func() string {
			return "template"
		}))

		expectedString := "template"
		if !strings.Contains(result, expectedString) {
			t.Errorf("final result shold contain user message: 'template', but got: %s",
				result)
		}

	})
}
