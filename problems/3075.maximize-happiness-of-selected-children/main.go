package main

import (
	"slices"
)

func maximumHappinessSum(happiness []int, k int) int64 {
  var sum int64 = 0
  slices.SortFunc(happiness, func (a int, b int) int {
    if a > b {
      return -1
    }

    if a == b {
      return 0
    }

    return 1
  })

  for i := 0; i < k;i++ {
    sum += max(int64(happiness[i]-i), 0)
  }


  return sum
}

func main() {
  println(maximumHappinessSum([]int{1, 2, 4, 5, 3}, 2))
  println(maximumHappinessSum([]int{1, 1, 1, 1}, 2))
  println(maximumHappinessSum([]int{12,1,42}, 3))
}
