package main

func checkPowersOfThree(n int) bool {
	for n >= 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3

		if n == 0 {
			return true
		}

	}

	return true
}

func main() {
	println(checkPowersOfThree(12))
}
