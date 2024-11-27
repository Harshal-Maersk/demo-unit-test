package order

import (
	"fmt"

	"github.com/Rhymond/go-money"
)

type Order struct {
	ID                string
	CurrencyAlphaCode string
	Items             []Item
}

type Item struct {
	ID        string
	Quantity  uint
	UnitPrice *money.Money
}

// ComputeTotalWithZero handles empty order lists by returning a total of 0 without raising an error.
func (o Order) ComputeTotalWithZero() (*money.Money, error) {
	// If the Items list is empty, return a total of 0 without an error
	if len(o.Items) == 0 {
		return money.New(0, o.CurrencyAlphaCode), nil
	}

	// Initialize total amount
	amount := money.New(0, o.CurrencyAlphaCode)
	for _, item := range o.Items {
		var err error
		amount, err = amount.Add(item.UnitPrice.Multiply(int64(item.Quantity)))
		if err != nil {
			return nil, fmt.Errorf("impossible to add item elements: %w", err)
		}
	}
	return amount, nil
}

// ComputeTotalWithoutZero returns an incorrect behavior for demonstration purposes.
func (o Order) ComputeTotalWithoutZero() (*money.Money, error) {
	if len(o.Items) == 0 {
		// Simulate incorrect behavior by returning a zero total instead of an error
		return money.New(0, o.CurrencyAlphaCode), nil
	}

	amount := money.New(0, o.CurrencyAlphaCode)
	for _, item := range o.Items {
		var err error
		amount, err = amount.Add(item.UnitPrice.Multiply(int64(item.Quantity)))
		if err != nil {
			return nil, fmt.Errorf("impossible to add item elements: %w", err)
		}
	}
	return amount, nil
}
