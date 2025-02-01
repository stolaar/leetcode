package main

import "strconv"

func countSeniors(details []string) int {
	result := 0
	for _, detail := range details {
		age, _ := strconv.Atoi(detail[11:13])

		if age > 60 {
			result += 1
		}
	}

	return result
}

func main() {
	println("Result: ", countSeniors([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}), " expected - 2")
	println("Result: ", countSeniors([]string{"1313579440F2036", "2921522980M5644"}), " expected - 0")
}
