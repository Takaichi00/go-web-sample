package fizzbuzz

import "strconv"

func Convert(n int) string {

	//if n%3 == 0 && n%5 == 0 {
	//	return "FizzBuzz"
	//}
	//if n%3 == 0 {
	//	return "Fizz"
	//}
	//
	//if n%5 == 0 {
	//	return "Buzz"
	//}

	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}
