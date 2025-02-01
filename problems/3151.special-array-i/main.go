package main

func isArraySpecial(nums []int) bool {
	for i, j := 0, 1; j < len(nums) && i < j; i, j = i+1, j+1 {
		if (nums[i]%2 == 0 && nums[j]%2 == 0) || (nums[i]%2 != 0 && nums[j]%2 != 0) {
			return false
		}
	}

	return true
}

func main() {
	println(isArraySpecial([]int{1}))
	println(isArraySpecial([]int{2, 1, 4}))
	println(isArraySpecial([]int{4, 3, 1, 6}))
}
