package main

func divideArray(nums []int) bool {
	if len(nums)%2 != 0 {
		return false
	}

	dict := make(map[int]int, 500)

	for _, num := range nums {
		dict[num] ^= num
	}

	for _, val := range dict {
		if val != 0 {
			return false
		}
	}

	return true
}

func main() {
	println(divideArray([]int{3, 2, 3, 2, 2, 2}))
}
