package money

import (
	"fmt"
	"math"
	"strings"
)

const (
	decimalDefault = "."
)

// Money represents an amount ot a specific currency
type Money struct {
	amount   float64
	currency Currency
}

// New creates a new Money object
func New(f float64, c Currency) *Money {
	m := &Money{currency: c}
	m.amount = m.round(f)
	return m
}

// Amount returns the numerical value of the money
func (m Money) Amount() float64 {
	return m.amount
}

// CurrencyCode returns the code of currency
func (m Money) CurrencyCode() string {
	return string(m.currency)
}

// Currency returns the string representation of currency
func (m Money) Currency() string {
	d := currencies[m.currency]
	value := m.Format(d.thousand, d.decimal)
	return fmt.Sprintf("%s %s", d.symbol, value)
}

// Symbol returns the currency symbom
func (m Money) Symbol(c Currency) string {
	return currencies[c].symbol
}

// Format returns the formatted value according rules
func (m Money) Format(thousand, decimal byte) string {
	if thousand == 0 {
		thousand = ','
	}
	if decimal == 0 {
		decimal = '.'
	}

	str := fmt.Sprintf("%.2f", m.amount)
	arr := strings.Split(str, decimalDefault)
	n := arr[0]

	var bytes []byte
	l := len(n) - 1
	for i, j := l, 1; i >= 0; i-- {
		if j > 3 {
			j = 1
			bytes = append(bytes, thousand)
		}
		j++

		bytes = append(bytes, n[i])
	}

	var res string
	for _, e := range bytes {
		res = string(e) + res

	}
	return fmt.Sprintf("%s%s%s", res, string(decimal), arr[1])
}

// Formmated returns the formatted value according currency
func (m Money) Formatted() string {
	d := currencies[m.currency]
	return m.Format(d.thousand, d.decimal)
}

// Absolute returns the absolute value of current amount
func (m *Money) Absolute() float64 {
	return math.Abs(m.amount)
}

func (m Money) Compare(v float64) int {
	v = m.round(v)
	if m.amount < v {
		return -1
	}
	if m.amount > v {
		return 1
	}
	return 0
}

func (m Money) Equals(v float64) bool {
	v = m.round(v)
	return m.Compare(v) == 0
}

func (m Money) GreaterThan(v float64) bool {
	v = m.round(v)
	return m.Compare(v) > 0
}

func (m Money) GreaterThanOrEqual(v float64) bool {
	v = m.round(v)
	return m.Compare(v) >= 0
}

func (m Money) LessThan(v float64) bool {
	v = m.round(v)
	return m.Compare(v) < 0
}

func (m Money) LessThanOrEqual(v float64) bool {
	v = m.round(v)
	return m.Compare(v) <= 0
}

func (m *Money) Subtract(v float64) float64 {
	v = m.round(v)
	m.amount -= v
	return m.amount
}

func (m *Money) Sum(v float64) float64 {
	v = m.round(v)
	m.amount += v
	return m.amount
}

func (m Money) round(v float64) float64 {
	precision := currencies[m.currency].precision
	base := math.Pow10(precision)
	return math.Round(v*base) / base
}
