package money

import (
	"testing"
)

func TestAbsolute(t *testing.T) {
	values := map[float64]float64{
		1.99:            1.99,
		42.987:          42.99,
		-12345.9:        12345.90,
		-1234567890.934: 1234567890.93,
		-1:              1.00,
	}
	for k, v := range values {
		got := New(k, USD).Absolute()
		if got != v {
			t.Errorf("Expected '%f', got '%f'", v, got)
		}
	}
}

func TestCurrency(t *testing.T) {
	expected := "$ 123,456.79"
	if got := New(0123456.789, USD).Currency(); got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	expected = "R$ 123.456,74"
	if got := New(0123456.742, BRL).Currency(); got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
	expected = "R$ 1,99"
	if got := New(01.994, BRL).Currency(); got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestFormatted(t *testing.T) {
	expected := "123,456.79"
	if got := New(0123456.789, USD).Formatted(); got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	expected = "123.456,74"
	if got := New(0123456.742, BRL).Formatted(); got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestFormat(t *testing.T) {
	got := New(0123456.789, BRL).Format('.', ',')
	expected := "123.456,79"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestCalculator(t *testing.T) {
	m := New(12345.9, USD)

	// subtract
	got := m.Subtract(42)
	expected := 12303.9
	if got != expected {
		t.Errorf("Expected '%f', got '%f'", expected, got)
	}

	// sum
	got = m.Sum(57.954)
	expected = 12361.85
	if got != expected {
		t.Errorf("Expected '%f', got '%f'", expected, got)
	}
}

func TestCompare(t *testing.T) {
	values := map[int]map[float64]float64{
		-1: {
			2.00: 2.005,
		},
		0: {
			1.99: 1.989,
			2.0:  1.999,
			3.0:  3.004,
		},
		1: {
			2.0:   1.99,
			3.005: 3.0,
		},
	}
	for c, v := range values {
		for v1, v2 := range v {
			got := New(v1, USD).Compare(v2)
			if got != c {
				t.Errorf("Expected '%d', got '%d' ('%f', '%f')", c, got, v1, v2)
			}
		}
	}
}

func TestEquals(t *testing.T) {
	v1 := 1.99
	v2 := 1.9888
	if !New(v1, USD).Equals(v2) {
		t.Errorf("Should be equals ('%f', '%f')", v1, v2)
	}
}

func TestGreaterThan(t *testing.T) {
	v1 := 1.998
	v2 := 1.99
	if !New(v1, USD).GreaterThan(v2) {
		t.Errorf("Should be greater than ('%f', '%f')", v1, v2)
	}
}

func TestGreaterThanOrEqual(t *testing.T) {
	v1 := 2.00
	v2 := 1.99
	if !New(v1, USD).GreaterThanOrEqual(v2) {
		t.Errorf("Should be greater than or equal ('%f', '%f')", v1, v2)
	}

	v1 = 1.9898
	v2 = 1.99
	if !New(v1, USD).GreaterThanOrEqual(v2) {
		t.Errorf("Should be greater than or equal ('%f', '%f')", v1, v2)
	}
}

func TestLessThan(t *testing.T) {
	v1 := 1.98
	v2 := 1.99
	if !New(v1, USD).LessThan(v2) {
		t.Errorf("Should be less than ('%f', '%f')", v1, v2)
	}
}

func TestLessThanOrEqual(t *testing.T) {
	v1 := 123.45
	v2 := 123.456
	if !New(v1, USD).LessThanOrEqual(v2) {
		t.Errorf("Should be less than or equal ('%f', '%f')", v1, v2)
	}

	v1 = 3.98
	v2 = 3.978
	if !New(v1, USD).LessThanOrEqual(v2) {
		t.Errorf("Should be less than or equal ('%f', '%f')", v1, v2)
	}
}
