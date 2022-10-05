package fizzbuzz

import "strconv"

func Convert(n int) string {
	//return "1"
	if n%3 == 0 {
		return "Fizz"
	}

	if n%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(n)
}
