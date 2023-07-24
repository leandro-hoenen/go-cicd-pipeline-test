package math

import "testing"

func TestAdd(t *testing.T) {
	given := Add(3, 5)
	expected := 8

	intNotEqual(given, expected, t)
}

func TestSub(t *testing.T) {
	given := Sub(5, 2)
	expected := 3

	intNotEqual(given, expected, t)

}

func intNotEqual(one int, two int, t *testing.T) {
	if one != two {
		t.Errorf("got %q, wanted %q", one, two)
	}
}
