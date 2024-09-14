package main

func removeElement(nums []int, val int) int {
	k, i, j := 0, 0, len(nums)-1
	if j < 0 {
		return 0
	}
	if j == 0 {
		if nums[0] == val {
			return 0
		}
		return 1
	}

	for j >= 0 {
		if nums[j] == val {
			k += 1
			j--
			continue
		}
		break
	}

	if i == j {
		return len(nums) - k
	}

	for j >= i {
		if nums[j] == val {
			k += 1
			j--
			continue
		}

		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			continue
		}

		i++
	}

	return len(nums) - k
}

func main() {
	println(removeElement([]int{3, 2, 2, 3}, 3))
	println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	println(removeElement([]int{2, 2, 2}, 2))
	println(removeElement([]int{3}, 2))
}
