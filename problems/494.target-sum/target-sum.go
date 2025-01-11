package main

func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 1 {
		if target == 0 && nums[0] == 0 {
			return 2
		}
		if nums[0] == target || nums[0]*(-1) == target {
			return 1
		}
		return 0
	}

	rest := nums[1:]
	count := 0

	way1, way2 := findTargetSumWays(rest, target-nums[0]), findTargetSumWays(rest, target+nums[0])

	count += way1
	count += way2

	return count
}

func main() {
	println(findTargetSumWays([]int{1, 0}, 1))
	println(findTargetSumWays([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}
