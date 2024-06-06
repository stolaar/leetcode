package main

import (
	"fmt"
	"sort"
)

func groupHands(hand []int, groupSize int) bool {
	result := true
	group := []int{hand[0]}

	updated := []int{}

	for i := 1; i < len(hand); i++ {
		if len(group) == groupSize {
			updated = append(updated, hand[i:]...)
			break
		}

		last := group[len(group)-1]

		if hand[i] > last+1 {
			result = false
			updated = []int{}
			break
		}

		if hand[i] == last+1 {
			group = append(group, hand[i])
			continue
		}
		updated = append(updated, hand[i])
	}

	if len(group) != groupSize {
		return false
	}

	if len(updated) > 0 {
		return groupHands(updated, groupSize)
	}

	return result
}

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	sort.Ints(hand)

	result := groupHands(hand, groupSize)

	return result
}

func main() {
	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
	fmt.Println(isNStraightHand([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(isNStraightHand([]int{2, 3, 1}, 3))
	fmt.Println(isNStraightHand([]int{8, 10, 12}, 3))
	fmt.Println(isNStraightHand([]int{1, 1, 2, 2, 3, 3}, 2))
}
