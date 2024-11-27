package order

import (
	"testing" // Import the testing package for writing test cases

	"github.com/Rhymond/go-money"        // Import the go-money library for handling monetary values
	"github.com/stretchr/testify/assert" // Import testify for easier assertion checks
)

// TestComputeTotal_ValidOrder tests the ComputeTotal method for a valid order scenario.
func TestComputeTotal_ValidOrder(t *testing.T) {
	// Create an order with a valid ID, currency, and a single item.
	o := Order{
		ID:                "45",
		CurrencyAlphaCode: "EUR",
		Items: []Item{
			{
				ID:        "21",
				Quantity:  2,
				UnitPrice: money.New(100, "EUR"),
			},
		},
	}

	// Call the ComputeTotal method to calculate the total amount for the order.
	total, err := o.ComputeTotalWithZero()

	// Assert that no error occurred during the computation.
	assert.NoError(t, err)

	// Assert that the computed total amount is as expected (2 items * 100 = 200).
	assert.Equal(t, int64(200), total.Amount())

	// Assert that the currency of the total matches the order's currency.
	assert.Equal(t, "EUR", total.Currency().Code)
}

// Test for ComputeTotalWithZero - Should return total 0 without an error for empty order
func TestComputeTotalWithZero_EmptyOrder(t *testing.T) {
	o := Order{
		ID:                "46",
		CurrencyAlphaCode: "USD",
		Items:             []Item{}, // No items
	}

	total, err := o.ComputeTotalWithZero()
	assert.NoError(t, err)                    // No error expected
	assert.Equal(t, int64(0), total.Amount()) // Total should be 0
	assert.Equal(t, "USD", total.Currency().Code)
}

// Test for ComputeTotalWithoutZero - Should return an error for empty order
func TestComputeTotalWithoutZero_EmptyOrder(t *testing.T) {
	o := Order{
		ID:                "46",
		CurrencyAlphaCode: "USD",
		Items:             []Item{}, // Empty order
	}

	total, err := o.ComputeTotalWithoutZero()

	// The test expects an error but won't get one
	assert.Error(t, err, "An error was expected but not received")
	assert.Nil(t, total, "Total should be nil for an empty order")
}
