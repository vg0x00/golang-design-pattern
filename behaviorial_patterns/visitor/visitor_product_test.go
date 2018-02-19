package visitor

import (
	"fmt"
	"testing"
)

func TestVisitorWithProduct(t *testing.T) {
	products := make([]Visitable, 2)
	products[0] = &Rice{Product: Product{Price: 20, Name: "price"}}
	products[1] = &Phone{Product: Product{Price: 100, Name: "phone"}}
	priceVisitor := &PriceVisitor{}
	nameVisitor := &NamePrinterVisitor{}

	for _, p := range products {
		p.Accept(priceVisitor)
	}

	fmt.Printf("total price: %.2f\n", priceVisitor.Sum)

	for _, p := range products {
		p.Accept(nameVisitor)
	}

	fmt.Printf("name list:\n%s", nameVisitor.NameList)
}
