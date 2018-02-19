package visitor

import "fmt"

type ProductInfoRetriver interface {
	GetPrice() float32
	GetName() string
}

type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) GetName() string {
	return p.Name
}

type Visitorp interface {
	Visit(ProductInfoRetriver)
}

type Visitable interface {
	Accept(r Visitorp)
}

func (p *Product) Accept(r Visitorp) {
	r.Visit(p)
}

type Rice struct {
	Product
}

type Phone struct {
	Product
}

type PriceVisitor struct {
	Sum float32
}

// get total price
func (p *PriceVisitor) Visit(r ProductInfoRetriver) {
	p.Sum = p.Sum + r.GetPrice()
}

type NamePrinterVisitor struct {
	NameList string
}

// get total product name list
func (p *NamePrinterVisitor) Visit(r ProductInfoRetriver) {
	p.NameList = fmt.Sprintf("%s\n%s", r.GetName(), p.NameList)
}
