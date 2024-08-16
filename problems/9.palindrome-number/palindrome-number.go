package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}
	r, t := 0, x

	for t > 0 {
		rm := t % 10
		r = (r * 10) + rm
		t /= 10
	}

	return r == x
}

func main() {
	println(isPalindrome(121))
	println(isPalindrome(10))
}
