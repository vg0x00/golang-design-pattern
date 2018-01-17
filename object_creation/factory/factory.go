package factory

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash      = 1
	DebitCard = 2
)

// separate implementation
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return &NewCashPM{}, nil // only change this part to update NewCashPM
	case DebitCard:
		return &DebitCardPM{}, nil
	default:
		return nil, errors.New("not implemented yet")
	}
}

type CashPM struct {
}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f amount was paid using cash", amount)
}

type DebitCardPM struct {
}

func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f amount was paid using debit", amount)
}

// new cash payment method here
type NewCashPM struct {
}

func (nc *NewCashPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f was paid using cash\n", amount)
}
