package main

import "fmt"


func singleNumber(nums []int) []int {
  occurences := make(map[int]int)
  result := []int{}

  for _, num := range nums {
    occurences[num]++
  }


  for key, value := range occurences {
    if value == 1 {
      result = append(result, key)

      if len(result) == 2 {
        break
      }
    }
  }

  return result
}

func main() {
  fmt.Println(singleNumber([]int{1,2,1,3,2,5}))
  fmt.Println(singleNumber([]int{-1, 0}))
  fmt.Println(singleNumber([]int{0,1}))
}
