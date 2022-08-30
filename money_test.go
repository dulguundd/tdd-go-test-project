package tdd

import (
	"testing"
)

type Money struct {
	amount   float64
	currency string
}

func TestMultiplicationInDollar(t *testing.T) {
	originalMoney := Money{
		amount:   5,
		currency: "USD",
	}
	actualResult := originalMoney.Times(2)
	expectedResult := Money{
		amount:   10,
		currency: "USD",
	}
	assertEqual(t, expectedResult, actualResult)
}

func TestMultiplicationInEuro(t *testing.T) {
	originalMoney := Money{
		amount:   10,
		currency: "EUR",
	}
	actualResult := originalMoney.Times(2)
	expectedResult := Money{
		amount:   20,
		currency: "EUR",
	}
	assertEqual(t, expectedResult, actualResult)
}

func TestDivision(t *testing.T) {
	originalMoney := Money{amount: 4002, currency: "KRW"}
	actualResult := originalMoney.Divide(4)
	expectedResult := Money{amount: 1000.5, currency: "KRW"}
	assertEqual(t, expectedResult, actualResult)
}

func assertEqual(t *testing.T, expected Money, actual Money) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}

func (m Money) Times(multiplier int) Money {
	return Money{
		amount:   m.amount * float64(multiplier),
		currency: m.currency,
	}
}

func (m Money) Divide(divider int) Money {
	return Money{
		amount:   m.amount / float64(divider),
		currency: m.currency,
	}
}
