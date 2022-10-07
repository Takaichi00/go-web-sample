package fizzbuzz_test

import (
	"testing"
	fizzbuzz "todo/src/fizz_buzz"
)

func TestFizzBuzz(t *testing.T) {

	tests := []struct {
		n    int
		want string
	}{
		{n: 1, want: "1"},
		{n: 2, want: "2"},
		{n: 3, want: "Fizz"},
		{n: 5, want: "Buzz"},
		{n: 15, want: "FizzBuzz"},
	}

	for _, tt := range tests {
		got := fizzbuzz.Convert(tt.n)
		if got != tt.want {
			t.Errorf(`FizzBuzz(1) is %q`, got)
		}
	}

	//got := fizzbuzz.Convert(1)
	//if got != "1" {
	//	t.Errorf(`FizzBuzz(1) is %q`, got)
	//}
	//
}
