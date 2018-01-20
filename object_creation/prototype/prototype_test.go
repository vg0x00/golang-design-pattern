package prototype

import (
	"testing"
	"time"
)

func TestShirtPrototype(t *testing.T) {
	// everytime obj cloned from cache shoud be different
	shirtCLoner := GetShirtCloner()
	if shirtCLoner == nil {
		t.Fatalf("cache cloner should not be nil")
	}

	item1, err := shirtCLoner.GetClone(White) // return interface GetItemInfo
	if err != nil {
		t.Error(err)
	}

	if item1 == whiteShirtPrototype {
		t.Fatalf("new clone obj should not equal to prototype obj")
	}

	shirt1, found := item1.(Shirt)
	if !found {
		t.Errorf("item1 can not cast to Shirt type: %s\n", found)
	}

	shirt1.SKU = "absbsb"

	item2, err := shirtCLoner.GetClone(White) // return interface GetItemInfo
	if err != nil {
		t.Error(err)
	}

	if item2 == whiteShirtPrototype {
		t.Fatalf("new clone obj should not equal to prototype obj")
	}

	shirt2, found := item2.(Shirt)
	if !found {
		t.Errorf("item2 can not cast to Shirt type: %s\n", found)
	}

	shirt2.SKU = "ttttt"

	if shirt1.SKU == shirt2.SKU {
		t.Fatalf("two cloned shirt should not have the same SKU")
	}

	t.Log(shirt1.GetInfo())
	t.Log(shirt2.GetInfo())
	t.Logf("&shirt1: %p, &shirt2: %p\n", &shirt1, &shirt2)
}

func TestCloneBenchmark(t *testing.T) {
	tp := time.Now()
	for i := 0; i < 999999999; i++ {
		_ = Shirt{
			Price: 20.0,
			SKU:   "",
			Color: White,
		}
	}

	t.Logf("create new item cost: %0.3v seconds", time.Since(tp).Seconds())

	tp = time.Now()
	sp := &Shirt{
		Price: 20.0,
		SKU:   "",
		Color: White,
	}

	for i := 0; i < 999999999; i++ {
		_ = *sp
	}

	t.Logf("clone new item cost: %.3v seconds", time.Since(tp).Seconds())

}
