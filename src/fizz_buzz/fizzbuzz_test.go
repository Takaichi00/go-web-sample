package fizzbuzz_test

import (
	"testing"
	fizzbuzz "todo/src/fizz_buzz"
)

func TestFizzBuzz(t *testing.T) {
	got := fizzbuzz.Convert(1)
	if got != "1" {
		t.Errorf(`FizzBuzz(1) is %q`, got)
	}

	got = fizzbuzz.Convert(2)
	if got != "2" {
		t.Errorf(`FizzBuzz(1) is %q`, got)
	}

	got = fizzbuzz.Convert(3)
	if got != "Fizz" {
		t.Errorf(`FizzBuzz(1) is %q`, got)
	}

	got = fizzbuzz.Convert(5)
	if got != "Buzz" {
		t.Errorf(`FizzBuzz(1) is %q`, got)
	}
}
