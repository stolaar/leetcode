package main

import (
	"fmt"
	"strconv"
)

func swap(arr []int, i int, j int) {
  arr[i], arr[j] = arr[j], arr[i]
}

func partition(arr []int, low int, high int, sortedMap map[int]int) int {
  pivot := arr[high]

  i := low-1

  for j := low; j <= high -1;j++ {
    if arr[j] > pivot {
      i++
      swap(arr, i, j)
      sortedMap[arr[i]] = i
      sortedMap[arr[j]] = j
    }
  }

  swap(arr, i+1, high)
  sortedMap[arr[i+1]] = i+1
  sortedMap[arr[high]] = high

  return i+1
}

func quickSort(arr []int, low int, high int, sortedMap map[int]int) {
  if low < high {
    pi := partition(arr, low, high, sortedMap)
    quickSort(arr, low, pi-1, sortedMap)
    quickSort(arr, pi+1,high, sortedMap)
  }
}

func findRelativeRanks(score []int) []string {
  sorted := make([]int, len(score))
  copy(sorted, score)

  sortedMap := make(map[int]int)
  quickSort(sorted, 0, len(sorted)-1, sortedMap)

  result := []string{}


  for _, res := range score {
    if sortedMap[res] == 0 {
      result = append(result, "Gold Medal")
      continue
    }

    if sortedMap[res] == 1 {
      result = append(result, "Silver Medal")
      continue
    }

    if sortedMap[res] == 2 {
      result = append(result, "Bronze Medal")
      continue
    }

    result = append(result, strconv.Itoa(sortedMap[res]+1))
  }

  return result
}

func main() {
  fmt.Println(findRelativeRanks([]int{5,4,3,2,1}))
  fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))
}
