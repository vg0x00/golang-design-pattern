package decorator

import (
	"errors"
	"fmt"
)

// core obj: pizza
// decorator obj: onion, meat

// the common interface for core obj and decorator
type IngredientAdder interface {
	AddIngredient() (string, error)
}

// inner core obj for decorator
type CorePizzaDecorator struct {
	Ingredient IngredientAdder
}

func (d *CorePizzaDecorator) AddIngredient() (string, error) {
	// return "", errors.New("not implemented yet")
	return "pizza with the following ingredient:", nil
}

// decorator 1: Meat
type Meat struct {
	Ingredient IngredientAdder
}

func (d *Meat) AddIngredient() (string, error) {
	if d.Ingredient != nil {
		result, err := d.Ingredient.AddIngredient()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s %s", result, "meat"), nil
	}
	return "", errors.New("not implemented yet")
}

// decorator 2: Onion
type Onion struct {
	Ingredient IngredientAdder
}

func (d *Onion) AddIngredient() (string, error) {
	if d.Ingredient != nil {
		result, err := d.Ingredient.AddIngredient()
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s %s", result, "onion"), nil
	}
	return "", errors.New("not implemented yet")
}
