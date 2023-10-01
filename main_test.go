package main

import (
	"io"
	"os"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is prime number!"},
		{"not prime", 8, false, "8 is not prime, because it is divisible by 2"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative", -1, false, "Negative numbers are not prime, by definition!"},
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

func Test_prompt(t *testing.T) {
  oldOut := os.Stdout

  r, w, _ := os.Pipe()

  os.Stdout = w

  prompt()

  _ = w.Close()

  os.Stdout = oldOut

  out, _ := io.ReadAll(r)

  if string(out) != "> " {
    t.Errorf("Incorrect prompt: expected > but got %s", out)
  }
}
