package main

import (
	"fmt"
	"math/bits"
	"strconv"
)

func printBinary(num int) {
	fmt.Println(strconv.FormatInt(int64(num), 2))
}

func getBit(num int, position int) int {
	mask := 1 << position
	isolatedBit := num & mask

	return isolatedBit >> position
}

func minBitFlips(start int, goal int) int {
	startBitLen, goalBitLen, count := bits.Len(uint(start)), bits.Len(uint(goal)), 0
	var mask int

	for i := 0; i < max(startBitLen, goalBitLen); i++ {
		if start == goal {
			return count
		}

		startBitPosition, goalBitPosition := getBit(start, i), getBit(goal, i)

		if startBitPosition == goalBitPosition {
			continue
		}
		mask = (1 << i)

		start = start ^ mask
		count += 1
	}

	return count
}

func main() {
	println(minBitFlips(10, 7))
	println(minBitFlips(3, 4))
}
