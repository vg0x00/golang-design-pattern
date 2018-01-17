package creational

type ManufacturingComplex struct {
	builder Builder
}

func (f *ManufacturingComplex) SetBuilder(b Builder) {
	f.builder = b
}

func (f *ManufacturingComplex) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

type Product struct {
	Wheels    int
	Seats     int
	Structure string
}

type Builder interface {
	SetWheels() Builder
	SetStructure() Builder
	SetSeats() Builder
	GetProduct() Product
}

type CarBuilder struct {
	product Product
}

func (b *CarBuilder) SetWheels() Builder {
	b.product.Wheels = 4
	return b
}

func (b *CarBuilder) SetSeats() Builder {
	b.product.Seats = 5
	return b
}

func (b *CarBuilder) SetStructure() Builder {
	b.product.Structure = "Car"
	return b
}

func (b *CarBuilder) GetProduct() Product {
	return b.product
}

type MotoBuilder struct {
	product Product
}

func (b *MotoBuilder) SetWheels() Builder {
	b.product.Wheels = 2
	return b
}

func (b *MotoBuilder) SetSeats() Builder {
	b.product.Seats = 2
	return b
}

func (b *MotoBuilder) SetStructure() Builder {
	b.product.Structure = "Moto"
	return b
}

func (b *MotoBuilder) GetProduct() Product {
	return b.product
}
