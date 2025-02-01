package main

import "fmt"

func groupThePeople(groupSizes []int) [][]int {
	pmap, result := map[int][]int{}, [][]int{}

	for idx, s := range groupSizes {
		if len(pmap[s]) == 0 {
			pmap[s] = []int{idx}
			continue
		}
		pmap[s] = append(pmap[s], idx)
	}

	for key, value := range pmap {
		current, total := 0, len(value)
		for current < total {
			group := []int{}
			for i := current; i < key+current; i++ {
				group = append(group, value[i])
			}
			current += key
			result = append(result, group)
		}
	}

	return result
}

func main() {
	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
	fmt.Println(groupThePeople([]int{2, 1, 3, 3, 3, 2}))
}
