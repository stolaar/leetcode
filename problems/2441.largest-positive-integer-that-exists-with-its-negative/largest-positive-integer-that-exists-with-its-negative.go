package main

func findMaxK(nums []int) int {
  numMap := make(map[int]int)
  max := -1
  for _, i := range nums {
    numMap[i] = 1
  }

  for key, _ := range numMap {
    if key < 0 {
      continue
    }

    negative := key - 2*key
    if numMap[negative] == 1 && key > max {
      max = key
    }
  }

  return max
}

func main() {
  println("Max K", findMaxK([]int{-1,2,-3,3}))
  println("Max K", findMaxK([]int{-1,10,6,7,-7,1}))
}
