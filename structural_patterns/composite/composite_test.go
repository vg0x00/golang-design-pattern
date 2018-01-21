package composite

import "testing"

func TestComposite(t *testing.T) {
	swimmer := CompositeSwimmerA{
		// zero-initialization for field Athlete
		MySwim: swim,
	}

	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	// shared swim function
	localSwim := swim

	swimmerB := CompositeSwimmerB{
		MySwim: &localSwim,
	}

	swimmerB.MyAthlete.Train()
	(*swimmerB.MySwim)()

	fish := Shark{
		Swim: localSwim,
	}

	fish.Eat()
	fish.Swim()

	swimmerC := CompositeSwimmerC{
		// Athlete{},
		Swimmer: &SwimmerImplmentor{},
	}

	swimmerC.Train()
	swimmerC.Swim()
}
