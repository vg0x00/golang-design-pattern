package interpreter

import (
	"errors"
	"strconv"
	"strings"
)

type Interpreter interface {
	Read() int
}

type Value int

func (v *Value) Read() int {
	return int(*v) // type conversion
}

type SumOperation struct {
	left, right Interpreter
}

func (s *SumOperation) Read() int {
	return s.left.Read() + s.right.Read()
}

type SubOperation struct {
	left, right Interpreter
}

func (s *SubOperation) Read() int {
	return s.left.Read() - s.right.Read()
}

func operationFactory(o string, left, right Interpreter) Interpreter {
	switch o {
	case SUM:
		return &SumOperation{left: left, right: right}
	case SUB:
		return &SubOperation{left: left, right: right}
	default:
		return nil
	}
}

type CaculatorStackI []Interpreter

func (c *CaculatorStackI) Pop() Interpreter {
	length := len(*c)
	if length > 0 {
		temp := (*c)[length-1]
		*c = (*c)[:length-1]
		return temp
	} else {
		return nil
	}
}

func (c *CaculatorStackI) Push(i Interpreter) {
	*c = append(*c, i)
}

func CaculateWithInterface(o string) (int, error) {
	stack := CaculatorStackI{}
	symbols := strings.Split(o, " ")
	for _, operatorStr := range symbols {
		if operatorStr == SUM || operatorStr == SUB {
			right := stack.Pop()
			left := stack.Pop()
			operation := operationFactory(operatorStr, left, right)
			result := operation.Read()
			temp := Value(result)
			stack.Push(&temp)
		} else {
			val, err := strconv.Atoi(operatorStr)
			if err != nil {
				panic(err)
			}
			temp := Value(val)
			stack.Push(&temp)
		}
	}

	if len(stack) == 1 {
		return stack.Pop().Read(), nil
	} else {
		return 0, errors.New("caculator stack internal error")
	}
}
