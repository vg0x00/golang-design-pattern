package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	ins := GetInstance()
	if ins == nil {
		t.Error("instance is nil")
	}

	expectedCounter := ins

	ins2 := GetInstance()
	if expectedCounter != ins2 {
		t.Errorf("second time calling GetInstance should return same instance, but get diff one")
	}
}

func TestAddOne(t *testing.T) {
	currentCount := instance.AddOne()
	if currentCount != 1 {
		t.Errorf("after first calling AddOne, count expected to 1, but not got %d\n", currentCount)
	}

	currentCount = instance.AddOne()
	if currentCount != 2 {
		t.Errorf("after second calling AddOne, count expected to 1, but not got %d\n", currentCount)
	}

}
