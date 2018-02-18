package interpreter

import "testing"

func TestInterpreterWithInterface(t *testing.T) {
	operationStr := "3 4 sum 2 sub"
	val, err := CaculateWithInterface(operationStr)
	if err != nil {
		t.Errorf("caculator internal error: %s\n", err.Error())
	}

	if val != 5 {
		t.Errorf("result expect: 5 but got: %d\n", val)
	}

	t.Logf("result with interface: %d\n", val)
}
