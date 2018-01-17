package creational

import "testing"

var mc = ManufacturingComplex{}

func TestBuilderPattern(t *testing.T) {
	carBuilder := &CarBuilder{}
	mc.SetBuilder(carBuilder)
	mc.Construct()

	car := carBuilder.GetProduct()

	if car.Wheels != 4 {
		t.Errorf("wheels on a car should be 4 but got %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("structure of car should be \"Car\" but got %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("seats on a car should be 5 but got %d\n", car.Seats)
	}

	motoBuilder := &MotoBuilder{}
	mc.SetBuilder(motoBuilder)
	mc.Construct()

	moto := motoBuilder.GetProduct()

	if moto.Wheels != 2 {
		t.Errorf("wheels on a moto should be 2 but got %d\n", moto.Wheels)
	}

	if moto.Structure != "Moto" {
		t.Errorf("structure of moto should be \"Moto\" but got %s\n", moto.Structure)
	}

	if moto.Seats != 2 {
		t.Errorf("seats on a moto should be 2 but got %d\n", moto.Seats)
	}

}
