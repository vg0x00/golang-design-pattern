package decorator

import (
	"strings"
	"testing"
)

func TestDecoratorPizza(t *testing.T) {
	var pizza CorePizzaDecorator
	t.Run("inner core obj", func(t *testing.T) {
		pizza = CorePizzaDecorator{}
		result, err := pizza.AddIngredient()
		if err != nil {
			t.Fatalf("decorator inner obj interface calling wrong: %s", err.Error())
		}

		expectResult := "pizza with the following ingredient:"
		if result != expectResult {
			t.Fatalf("inner obj AddIngredient should return:\n %s\n but got: \n %s\n",
				expectResult, result)
		}
	})

	t.Run("onion decorator without inner core obj", func(t *testing.T) {
		onion := &Onion{}
		_, err := onion.AddIngredient()
		if err == nil {
			t.Fatal("decorator without inner obj should cause fatal error, but not")
		}
	})

	t.Run("onion decorator withinner core obj", func(t *testing.T) {
		onion := &Onion{&pizza}
		onionResult, err := onion.AddIngredient()
		if err != nil {
			t.Fatal("decorator without inner obj should cause fatal error, but not")
		}
		if !strings.Contains(onionResult, "onion") {
			t.Errorf("after calling onion decorator, result should contains word onion, but not")
		}
		t.Log(onionResult)
	})

	t.Run("meat decorator without inner core obj", func(t *testing.T) {
		meat := &Meat{}
		_, err := meat.AddIngredient()
		if err == nil {
			t.Fatal("decorator without inner obj should get fatal error, but not")
		}
	})

	t.Run("meat decorator with inner core obj", func(t *testing.T) {
		meat := &Meat{&pizza}
		meatResult, err := meat.AddIngredient()
		if err != nil {
			t.Error(err.Error())
		}
		if !strings.Contains(meatResult, "meat") {
			t.Errorf("after calling meat decorator, result should contains word meat, but not")
		}
		t.Log(meatResult)

	})

	t.Run("onion inside meat decorator ", func(t *testing.T) {
		deco := &Meat{&Onion{&pizza}}
		result, err := deco.AddIngredient()
		if err != nil {
			t.Error(err.Error())
		}
		if !strings.Contains(result, "meat") || !strings.Contains(result, "onion") {
			t.Errorf("after calling meat decorator, result should contains word botht meat and onino , but got: %s\n", result)
		}

		t.Log(result)
	})

}
