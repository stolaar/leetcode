package main

func threeConsecutiveOdds(arr []int) bool {
	streak := 0

	for _, num := range arr {
		if streak >= 3 {
			break
		}
		if num%2 == 0 {
			streak = 0
			continue
		}
		streak++
	}

	return streak >= 3
}
