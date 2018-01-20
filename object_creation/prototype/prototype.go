package prototype

import (
	"errors"
	"fmt"
)

type ShirtCloner interface {
	GetClone(color ShirtColor) (ItemInfoGetter, error)
}

func GetShirtCloner() ShirtCloner {
	return ShirtCache{}
}

type ShirtCache struct {
}

func (s ShirtCache) GetClone(m ShirtColor) (ItemInfoGetter, error) {
	// return nil, errors.New("item info getter not implemented yet")
	switch m {
	case White:
		return *whiteShirtPrototype, nil
	case Black:
		return *blackShirtPrototype, nil
	case Red:
		return redShirtPrototype, nil
	default:
		return nil, errors.New("No Getter for this color type")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

const (
	White = iota
	Black
	Red
)

func (s Shirt) GetInfo() string {
	return fmt.Sprintf("shirt info: price: %0.2f, SKU: %s, Color id: %d\n", s.Price, s.SKU, s.Color)
}

var whiteShirtPrototype *Shirt = &Shirt{
	Price: 20.0,
	SKU:   "",
	Color: White,
}

var blackShirtPrototype *Shirt = &Shirt{
	Price: 10.0,
	SKU:   "",
	Color: Black,
}
var redShirtPrototype *Shirt = &Shirt{
	Price: 10.0,
	SKU:   "",
	Color: Red,
}
