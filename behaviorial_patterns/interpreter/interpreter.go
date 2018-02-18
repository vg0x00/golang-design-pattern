package interpreter

import (
	"errors"
	"strconv"
	"strings"
)

type CaculatorStack []int

func (c *CaculatorStack) Push(s int) {
	*c = append(*c, s)
}

func (c *CaculatorStack) Pop() int {
	length := len(*c)
	if length > 0 {
		temp := (*c)[length-1]
		*c = (*c)[:length-1]
		return temp
	}
	return 0
}

func Caculate(s string) (int, error) {
	stack := CaculatorStack{}
	symbols := strings.Split(s, " ")

	for _, symbol := range symbols {
		if isOperator(symbol) {
			mathFunc := getMathFunc(symbol)
			right := stack.Pop()
			left := stack.Pop()
			tempResult := mathFunc(left, right)
			stack.Push(tempResult)
		} else { // symbol is value, skip error checking
			tempValue, err := strconv.Atoi(symbol)
			if err != nil {
				return 0, err
			}
			stack.Push(tempValue)
		}
	}

	if len(stack) == 1 {
		return stack.Pop(), nil
	} else {
		return 0, errors.New("caculator stack internal error")
	}
}

const (
	SUM = "sum"
	SUB = "sub"
	MUL = "mul"
	DIV = "div"
)

func isOperator(s string) bool {
	if (s == SUM) || (s == SUB) || (s == MUL) || (s == DIV) {
		return true
	}
	return false
}

type MathFunc func(int, int) int

func getMathFunc(s string) MathFunc {
	switch s {
	case SUM:
		return func(left, right int) int {
			return left + right
		}
	case SUB:
		return func(left, right int) int {
			return left - right
		}
	case MUL:
		return func(left, right int) int {
			return left * right
		}
	case DIV:
		return func(left, right int) int {
			return left / right
		}
	default: // this should not happened, cause operation filtered by isOperator
		return nil
	}
}
