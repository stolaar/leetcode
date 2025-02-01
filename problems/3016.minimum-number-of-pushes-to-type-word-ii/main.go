package main

import (
	"sort"
)

type kv struct {
	Key   rune
	Value int
}

func minimumPushes(word string) int {
	runeMap, kpadMap := make(map[rune]int), make(map[int]int)

	for _, r := range word {
		runeMap[r] += 1
	}
	var runeMapArr []kv

	for key, value := range runeMap {
		runeMapArr = append(runeMapArr, kv{key, value})
	}

	sort.Slice(runeMapArr, func(a, b int) bool {
		return runeMapArr[a].Value > runeMapArr[b].Value
	})

	currPadNum, result := 2, 0

	for _, v := range runeMapArr {
		kpl := kpadMap[currPadNum] + 1

		result += v.Value * kpl

		kpadMap[currPadNum] += 1
		if currPadNum == 9 {
			currPadNum = 2
		} else {
			currPadNum += 1
		}

	}

	return result
}

func main() {
	println(minimumPushes("abcde"))
	println(minimumPushes("xyzxyzxyzxyz"))
	println(minimumPushes("aabbccddeeffgghhiiiiii"))
}
