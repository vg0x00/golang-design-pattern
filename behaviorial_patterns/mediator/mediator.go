package mediator

import "fmt"

type One struct{}
type Two struct{}
type Three struct{}
type Four struct{}
type Five struct{}
type Six struct{}
type Seven struct{}
type Eight struct{}
type Nine struct{}
type Ten struct{}
type Zero struct{}

func Sum(a, b interface{}) interface{} { // return type like void pointer in C
	switch paramA := a.(type) {
	case One:
		switch b.(type) {
		case One:
			return &Two{}
		case Two:
			return &Three{}
		default:
			return fmt.Errorf("Number type caculation not implemented yet")
		}
	case Two:
		switch b.(type) {
		case One:
			return &Three{}
		case Two:
			return &Four{}
		case int:
			return &Five{} // why? cause it's a crazy example
		default:
			return fmt.Errorf("Number type caculation not implemented yet")
		}

	case int:
		switch paramB := b.(type) {
		case One:
			return &Three{}
		case Two:
			return &Four{}
		case int:
			return paramA + paramB
		default:
			return fmt.Errorf("Number type caculation no implemnted yet")
		}

	default:
		return fmt.Errorf("number not found")
	}

}
