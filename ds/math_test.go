package ds

import "testing"

type powTestEntry struct {
	Base, Quotient int
	Result         int64
}

func TestPow(t *testing.T) {
	var tests []powTestEntry
	tests = append(tests, powTestEntry{Base: 0, Quotient: 0, Result: 1})
	tests = append(tests, powTestEntry{Base: 0, Quotient: 1, Result: 1})
	tests = append(tests, powTestEntry{Base: -1, Quotient: 2, Result: 1})
	tests = append(tests, powTestEntry{Base: 1, Quotient: 0, Result: 1})
	tests = append(tests, powTestEntry{Base: 1, Quotient: 1, Result: 1})
	tests = append(tests, powTestEntry{Base: 5, Quotient: 1, Result: 5})
	tests = append(tests, powTestEntry{Base: 5, Quotient: 3, Result: 125})

	for _, test := range tests {
		got := Pow(test.Base, test.Quotient)
		if got != test.Result {
			t.Errorf("For %v^%v, expected: %v, got: %v", test.Base, test.Quotient, test.Result, got)
		}
	}
}
