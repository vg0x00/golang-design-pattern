package interpreter

import (
	"testing"
)

func TestInterpreter(t *testing.T) {
	t.Run("testing caculator", func(t *testing.T) {
		str := "3 4 sum 1 sub"
		value, err := Caculate(str)
		if err != nil {
			t.Errorf("caculator inner error: %s\n", err.Error())
		}

		if value != 6 {
			t.Errorf("caculator result should: 6 but got: %d\n",
				value)
		}
	})
}
