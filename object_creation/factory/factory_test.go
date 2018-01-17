package factory

import (
	"strings"
	"testing"
)

func TestCreatPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("a payment method 'Cash' not exist")
	}

	msg := payment.Pay(10.2)
	if !strings.Contains(msg, "paid using cash") {
		t.Fatal("cash paymetn method msg is not correct")
	}

	t.Log("Log: ", msg)

	payment, err = GetPaymentMethod(20)
	if err == nil {
		t.Error("payment method of id(20) should get error, but not")
	}
}
