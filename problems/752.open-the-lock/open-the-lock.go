package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=752 lang=golang
 *
 * [752] Open the Lock
 */

func getCurrentNeighbours(current string) []string {
	result := []string{}

	for idx, char := range current {
		charInt := int(char - '0')
		next := min(charInt+1, 9)
		prev := charInt - 1

		if prev < 0 {
			prev = 9
		}
		if idx == len(current)-1 {
			result = append(result, current[:idx]+fmt.Sprint(next))
			result = append(result, current[:idx]+fmt.Sprint(prev))
			continue
		}

		if idx == 0 {
			result = append(result, current[0:0]+fmt.Sprint(next)+current[1:])
			result = append(result, current[0:0]+fmt.Sprint(prev)+current[1:])
			continue
		}

		result = append(result, current[0:idx]+fmt.Sprint(next)+current[idx+1:])
		result = append(result, current[0:idx]+fmt.Sprint(prev)+current[idx+1:])

	}

	return result
}

type Code struct {
	Val   string
	Count int
}

func findShortest(deadends map[string]bool, target string, current string) int {
	result := -1
	visited := make(map[string]bool)
	curr := &Code{
		Val:   current,
		Count: 1,
	}

	q := []*Code{curr}

	if target == current {
		return 0
	}

	for len(q) > 0 {
		c := q[0]
		found := false
		if c.Val == target {
			result = c.Count
			break
		}

		if deadends[c.Val] {
			break
		}

		q = q[1:]
		adj := getCurrentNeighbours(c.Val)
		for _, n := range adj {
			if n == target {
				found = true
				result = c.Count
				break
			}
			if deadends[n] {
				continue
			}
			if !visited[n] {
				visited[n] = true
				q = append(q, &Code{Val: n, Count: c.Count + 1})
			}
		}

		if found {
			break
		}
	}

	return result
}

// @lc code=start
func openLock(deadends []string, target string) int {
	deadendsMap := make(map[string]bool)

	for _, dedeadend := range deadends {
		deadendsMap[dedeadend] = true
	}

	return findShortest(deadendsMap, target, "0000")
}

// @lc code=end

func main() {
	result := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	fmt.Println("Result", result)
	result2 := openLock([]string{"8888"}, "0009")
	fmt.Println("Result2", result2)
	result3 := openLock([]string{"0000"}, "8888")
	fmt.Println("Result3", result3)
}
