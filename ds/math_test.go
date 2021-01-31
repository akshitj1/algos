package ds

import (
	"testing"
)

type powTestEntry struct {
	Base, Quotient int
	Result         int64
}

func TestPow(t *testing.T) {
	var tests []powTestEntry
	tests = append(tests, powTestEntry{Base: 0, Quotient: 0, Result: 1})
	tests = append(tests, powTestEntry{Base: 0, Quotient: 1, Result: 0})
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

func fibRec(n int) int64 {
	if n == 0 {
		return 0
	}
	prev, next := int64(0), int64(1)
	for i := 2; i <= n; i++ {
		prev, next = next, prev+next
	}
	return next
}

func TestFib(t *testing.T) {
	ns := []int{1, 2, 3, 4, 20}
	t.Log(ns)
	for _, n := range ns {
		exp := fibRec(n)
		got := Fib(n)
		if exp != got {
			t.Errorf("%vth Fibonacci number. expected: %v got: %v", n, exp, got)
		}
	}
}
