package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 0; i < n; i++ {
		ii := i + 1

		db3, db5 := ii%3 == 0, ii%5 == 0

		if db3 && db5 {
			result[i] = "FizzBuzz"
			continue
		}

		if db3 {
			result[i] = "Fizz"
			continue
		}

		if db5 {
			result[i] = "Buzz"
			continue
		}

		result[i] = strconv.Itoa(ii)
	}

	return result
}

func main() {
	fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(5))
	fmt.Println(fizzBuzz(15))
}
