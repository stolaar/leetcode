package main

func lemonadeChange(bills []int) bool {
	cash, result := make(map[int]int), true

	for _, bill := range bills {
		if bill == 5 {
			cash[5] += 1
			continue
		}

		change := bill - 5

		if change == 15 {
			if cash[10] >= 1 && cash[5] >= 1 {
				cash[10] -= 1
				cash[5] -= 1
				cash[bill] += 1
				continue
			}
			if cash[5] >= 3 {
				cash[5] -= 3
				cash[bill] += 1
				continue
			}
			result = false
			break
		}

		if change == 5 {
			if cash[5] == 0 {
				result = false
				break
			}
			cash[5] -= 1
		}

		cash[bill] += 1
	}

	return result
}

func main() {
	println(lemonadeChange([]int{5, 5, 5, 10, 20}))
	println(lemonadeChange([]int{5, 5, 10, 10, 20}))
	println(lemonadeChange([]int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5})) // expected true
}
