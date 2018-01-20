package absfactory

import (
	"testing"
)

func TestGetCarAbstractFactory(t *testing.T) {
	carFactory, ok := GetAbstractFactory(CarFactoryType)
	if ok != nil {
		t.Logf("GetAbstractFactory with factory id(CarFactoryType) failed: %s\n", ok.Error())
	}

	// return a interface type, a pointer
	car, err := carFactory.GetVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	luxuryCar, found := car.(Car)
	if !found {
		t.Fatalf("luxury car cast failed\n")
	}

	luxuryCar.GetDoors()

	car, err = carFactory.GetVehicle(NormalCarType)
	if err != nil {
		t.Fatal(err)
	}

	normalCar, found := car.(Car)
	if !found {
		t.Fatalf("normal car cast failed\n")
	}

	normalCar.GetDoors()
}

func TestGetMotoAbstractFactory(t *testing.T) {
	motoFactory, ok := GetAbstractFactory(MotoFactoryType)
	if ok != nil {
		t.Logf("GetAbstractFactory with factory id(MotoFactoryType) failed: %s\n", ok.Error())
	}

	// return a interface type, a pointer
	moto, err := motoFactory.GetVehicle(SportMotoType)
	if err != nil {
		t.Fatal(err)
	}

	sportMoto, found := moto.(Moto)
	if !found {
		t.Fatalf("sport moto cast failed\n")
	}

	t.Log("moto has wheels: ", sportMoto.GetWheels())

	moto, err = motoFactory.GetVehicle(NormalMotoType)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("moto label: ", moto.GetLabel())

	normalMoto, found := moto.(Moto)
	if !found {
		t.Fatalf("normal moto cast failed\n")
	}

	t.Log("moto has wheels: ", normalMoto.GetWheels())
}
