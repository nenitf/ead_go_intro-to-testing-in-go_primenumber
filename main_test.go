package main

import "testing"

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is prime number!"},
		{"not prime", 8, false, "8 is not prime, because it is divisible by 2"},
	}

	for _, e := range primeTests {
		actual, msg := isPrime(e.testNum)
		if actual != e.expected {
			t.Errorf("isPrime(%d) = %v, want %v", e.testNum, actual, e.expected)
		}

		if msg != e.msg {
			t.Errorf("isPrime(%d) = %v, want %v", e.testNum, msg, e.msg)
		}
	}
}
