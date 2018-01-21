package composite

import "fmt"

type Athlete struct {
}

func (a *Athlete) Train() {
	fmt.Println("Traingning...")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func() // func with no params and no returns
}

func swim() {
	fmt.Println("swimming...")
}

type CompositeSwimmerB struct {
	MyAthlete Athlete
	MySwim    *func() // func with no params and no returns
}

type Animal struct {
}

func (a *Animal) Eat() {
	fmt.Println("eating now")
}

type Shark struct {
	Animal
	Swim func()
}

// interface implement

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImplmentor struct {
}

func (s *SwimmerImplmentor) Swim() {
	fmt.Println("swiming now")
}

type CompositeSwimmerC struct {
	Athlete
	Swimmer
}
