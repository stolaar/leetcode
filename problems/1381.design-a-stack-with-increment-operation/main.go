package main

import "fmt"

type CustomStack struct {
	cap  int
	size int
	arr  []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		cap: maxSize,
	}
}

func (stack *CustomStack) Push(x int) {
	if stack.size+1 > stack.cap {
		return
	}
	stack.arr = append(stack.arr, x)
	stack.size += 1
}

func (stack *CustomStack) Pop() int {
	if stack.size == 0 {
		return -1
	}
	last := stack.arr[stack.size-1]
	stack.arr = stack.arr[0 : stack.size-1]

	stack.size -= 1

	return last
}

func (stack *CustomStack) Increment(k int, val int) {
	r := min(k, stack.size)

	for i := 0; i < r; i++ {
		stack.arr[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */

func main() {
	obj := Constructor(3)
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.arr)
	obj.Pop()
	fmt.Println(obj.arr)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	fmt.Println(obj.arr)
	obj.Increment(5, 100)
	fmt.Println(obj.arr)
	obj.Increment(2, 100)
	fmt.Println(obj.arr)
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
}
