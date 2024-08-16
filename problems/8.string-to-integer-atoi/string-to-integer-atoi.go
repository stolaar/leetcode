package main

import (
	"strconv"
	"strings"
)

var (
	maxInt = 2147483647
	minInt = -2147483648
)

func isOutOfRange(result string) bool {
	filtered, negative, s := "", strings.HasPrefix(result, "-"), false

	for idx, r := range result {
		if idx == 0 && negative {
			continue
		}
		n, _ := strconv.Atoi(string(r))

		if n > 0 {
			s = true
		}
		if string(r) == "0" && !s {
			continue
		}

		filtered += string(r)
	}

	if len(filtered) < 10 {
		return false
	}

	if len(filtered) > 10 {
		return true
	}

	parsed, e := strconv.Atoi(filtered)

	if e != nil {
		return true
	}

	if negative {
		return parsed > maxInt+1
	}

	return parsed > maxInt
}

func myAtoi(s string) int {
	result, sc := "", 0

	for _, r := range s {
		rs := string(r)
		_, e := strconv.Atoi(rs)

		if sc > 1 {
			break
		}

		if rs == " " && result == "" && sc == 0 {
			continue
		}

		if rs == "-" && !strings.HasPrefix(result, "-") && result == "" {
			result = "-" + result
			sc += 1
			continue
		}

		if rs == "+" && result == "" && sc == 0 {
			sc += 1
			continue
		}

		if e != nil {
			break
		}

		result += rs
	}

	if result == "" || result == "-" {
		return 0
	}

	if isOutOfRange(result) {
		if strings.HasPrefix(result, "-") {
			return minInt
		}
		return maxInt
	}

	val, e := strconv.Atoi(result)

	if e != nil {
		return 0
	}

	return val
}

func main() {
	println(myAtoi(" -042"))
	println(myAtoi("-91283472332"))
	println(myAtoi("words and 987"))
	println(myAtoi("0-1"))
	println(myAtoi("1337c0d3"))
	println(myAtoi("+-12"))
	println(myAtoi("21474836460"))
	println(myAtoi("  0000000000012345678"))
	println(myAtoi("2147483800"))
	println(myAtoi("-2147483647"))
	println(myAtoi("  +  413"))
	println(myAtoi("1095502006p8"))
	println(myAtoi("-2147483647"))
}
