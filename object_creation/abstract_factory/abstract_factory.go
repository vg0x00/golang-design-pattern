package absfactory

import "errors"

const (
	CarFactoryType = iota
	MotoFactoryType
)

const (
	LuxuryCarType = iota
	NormalCarType
)

const (
	SportMotoType = iota
	NormalMotoType
)

type Car interface {
	GetDoors()
}

type Moto interface {
	GetWheels() int
}

type VehicleFactory interface {
	GetVehicle(vehicleTypeId int) (Vehicle, error)
}

type Vehicle interface {
	// only one method for short example
	GetLabel() string
}

func GetAbstractFactory(factoryTypeId int) (VehicleFactory, error) {
	switch factoryTypeId {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotoFactoryType:
		return new(MotoFactory), nil
	default:
		return nil, errors.New("factory not found")
	}
}

type CarFactory struct {
}

type LuxuryCar struct {
}

func (lc *LuxuryCar) GetLabel() string {
	return "LuxuryCar"
}

func (lc *LuxuryCar) GetDoors() {
}

type NormalCar struct {
}

func (nc *NormalCar) GetLabel() string {
	return "NormalCar"
}

func (nc *NormalCar) GetDoors() {
}

func (cf *CarFactory) GetVehicle(vehicleId int) (Vehicle, error) {
	switch vehicleId {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case NormalCarType:
		return new(NormalCar), nil
	default:
		return nil, errors.New("none vehicle object found")
	}
}

type MotoFactory struct {
}

type SportMoto struct {
}

func (sm SportMoto) GetLabel() string {
	return "SportMoto"
}

func (sm SportMoto) GetWheels() int {
	return 2
}

type NormalMoto struct {
}

func (nm NormalMoto) GetLabel() string {
	return "NormalMoto"
}

func (nm NormalMoto) GetWheels() int {
	return 2
}

func (mf *MotoFactory) GetVehicle(vehicleId int) (Vehicle, error) {
	switch vehicleId {
	case SportMotoType:
		return new(SportMoto), nil
	case NormalMotoType:
		return new(NormalMoto), nil
	default:
		return nil, errors.New("Moto object not found")
	}
}
