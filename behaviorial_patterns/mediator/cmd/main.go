package main

import (
	"fmt"

	"github.com/vg0x00/test/behaviorial_patterns/mediator"
)

func main() {
	fmt.Printf("%#v\n", mediator.Sum(mediator.One{}, mediator.Two{}))
	fmt.Printf("%#v\n", mediator.Sum(mediator.Two{}, 20)) // should return Five{}
}
