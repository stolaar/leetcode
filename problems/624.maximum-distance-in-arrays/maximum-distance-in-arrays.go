package main

import (
	"math"
)

type Distance struct {
	Val   int
	Index int
}

func maxDistance(arrays [][]int) int {
	m, mx, sm, smx := Distance{math.MaxInt32, -1}, Distance{math.MinInt32, -1}, Distance{math.MaxInt32, -1}, Distance{math.MinInt32, -1}

	for index, arr := range arrays {
		minDistance, maxDistance := arr[0], arr[len(arr)-1]

		if minDistance < m.Val {
			m.Val = minDistance
			m.Index = index
		}
		if maxDistance > mx.Val {
			mx.Val = maxDistance
			mx.Index = index
		}
	}

	for index, arr := range arrays {
		minDistance, maxDistance := arr[0], arr[len(arr)-1]

		if minDistance < sm.Val && index != m.Index {
			sm.Val = minDistance
			sm.Index = index
		}

		if maxDistance > smx.Val && index != mx.Index {
			smx.Val = maxDistance
			smx.Index = index
		}
	}

	if m.Index != mx.Index {
		return int(math.Abs(float64(mx.Val - m.Val)))
	}

	return max(int(math.Abs(float64(mx.Val-sm.Val))), int(math.Abs(float64(smx.Val-m.Val))))
}

func main() {
	println(maxDistance([][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}))
	println(maxDistance([][]int{{1}, {1}}))
	println(maxDistance([][]int{{-1000, 0}, {-300, 8000}}))
	println(maxDistance([][]int{{1, 4}, {0, 5}}))
}
