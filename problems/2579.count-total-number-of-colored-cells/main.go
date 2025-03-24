package main

func coloredCells(n int) int64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 5
	}

	var current int64 = 5
	for i := 3; i <= n; i++ {
		current += int64((i - 1) * 4)
	}

	return current
}

func main() {
	println(coloredCells(5))
	println(coloredCells(100000))
}
