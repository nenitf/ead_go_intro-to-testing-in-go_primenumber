package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
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
		t.Errorf("Incorrect prompt: expected > but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("ntro text not correct; got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", "Please enter a valid number."},
		{"not a number", "a", "Please enter a valid number."},
		{"negative", "-1", "Negative numbers are not prime, by definition!"},
    {"zero", "0", "0 is not prime, by definition!"},
    {"one", "1", "1 is not prime, by definition!"},
    {"two", "2", "2 is prime number!"},
    {"three", "3", "3 is prime number!"},
    {"four", "4", "4 is not prime, because it is divisible by 2"},
		{"is prime", "7", "7 is prime number!"},
		{"is not prime", "8", "8 is not prime, because it is divisible by 2"},
    {"typed", "three", "Please enter a valid number."},
    {"typed quit", "q", ""},
    {"typed Quit", "Q", ""},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input) // simulates user input
			reader := bufio.NewScanner(input)
			res, _ := checkNumbers(reader)

			if !strings.EqualFold(res, tc.expected) {
				t.Errorf("Incorrect result; got %s", res)
			}
		})
	}
}

func Test_readUserInput(t *testing.T) {
  doneChan := make(chan bool)

  var stdin bytes.Buffer

  stdin.Write([]byte("1\nq\n"))

  go readUserInput(&stdin, doneChan)
  <- doneChan

  close(doneChan)
}
