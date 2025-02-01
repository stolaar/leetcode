package main

func currentPos(n int, time int, rtl bool) int {
	if time < n {
		if rtl {
			return n - time
		}
		return time + 1
	}

	return currentPos(n, time-n+1, !rtl)
}

func passThePillow(n int, time int) int {
	return currentPos(n, time, false)
}

func main() {
	println("result: ", passThePillow(4, 5), " expected: ", 2)
	println("result: ", passThePillow(3, 2), " expected: ", 3)
	println("result: ", passThePillow(18, 38), " expected: ", 5)
	println("result: ", passThePillow(2, 5), " expected: ", 2)
	println("result: ", passThePillow(2, 1000), " expected: ", 1)
}
