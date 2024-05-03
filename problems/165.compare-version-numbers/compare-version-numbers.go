package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	version1Arr := strings.Split(version1, ".")
	version2Arr := strings.Split(version2, ".")

	result := 0
	idx := 0

	for idx < len(version1Arr) || idx < len(version2Arr) {
		ver1HasRevision := len(version1Arr)-1 >= idx
		ver2HasRevision := len(version2Arr)-1 >= idx

		var val1 int
		var val2 int

		if !ver1HasRevision {
			val2, _ = strconv.Atoi(version2Arr[idx])

			if val2 > 0 {
				result = -1
				break
			}
			idx++
			continue
		}

		if !ver2HasRevision {
			val1, _ = strconv.Atoi(version1Arr[idx])

			if val1 > 0 {
				result = 1
				break
			}
			idx++
			continue
		}
		val1, _ = strconv.Atoi(version1Arr[idx])
		val2, _ = strconv.Atoi(version2Arr[idx])

		if val1 > val2 {
			result = 1
			break
		}

		if val1 < val2 {
			result = -1
			break
		}

		idx++
	}

	return result
}

func main() {
	println("1.0 and 1.1", compareVersion("1.0", "1.1"))
	println("1.1 and 1.0", compareVersion("1.1", "1.0"))
	println("1.01 and 1.001", compareVersion("1.01", "1.001"))
	println("1.0 and 1.0.0", compareVersion("1.0", "1.0.0"))

	// Failed 59/84
	println("1.0.1 and 1", compareVersion("1.0.1", "1"))
}
