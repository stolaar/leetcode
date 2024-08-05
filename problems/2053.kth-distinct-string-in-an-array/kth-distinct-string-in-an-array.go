package main

func kthDistinct(arr []string, k int) string {
	m := make(map[string]int)

	for _, str := range arr {
		m[str] += 1
	}

	count, result := 1, ""

	for _, str := range arr {
		if m[str] == 1 {
			if count == k {
				result = str
				break
			}
			count += 1
		}
	}

	return result
}

