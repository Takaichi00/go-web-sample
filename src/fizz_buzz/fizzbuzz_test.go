package fizzbuzz_test

import (
	"fmt"
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
		tt := tt                               // prepare local val (for parallel)
		name := fmt.Sprintf("number:%v", tt.n) // test name

		// execute sub test
		t.Run(name, func(t *testing.T) {
			t.Parallel() // execute parallel
			got := fizzbuzz.Convert(tt.n)
			if got != tt.want {
				t.Errorf(`Convert(%v) = %q but want %q`, tt.n, got, tt.want)
			}
		})

		//got := fizzbuzz.Convert(tt.n)
		//if got != tt.want {
		//	t.Errorf(`FizzBuzz(1) is %q`, got)
		//}
	}

	//got := fizzbuzz.Convert(1)
	//if got != "1" {
	//	t.Errorf(`FizzBuzz(1) is %q`, got)
	//}
	//
}
