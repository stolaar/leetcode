package main

import (
	"fmt"
)

var numberWordsMap = map[int]string{
	1:          "One",
	2:          "Two",
	3:          "Three",
	4:          "Four",
	5:          "Five",
	6:          "Six",
	7:          "Seven",
	8:          "Eight",
	9:          "Nine",
	0:          "Zero",
	10:         "Ten",
	11:         "Eleven",
	12:         "Twelve",
	13:         "Thirteen",
	14:         "Fourteen",
	15:         "Fifteen",
	16:         "Sixteen",
	17:         "Seventeen",
	18:         "Eighteen",
	19:         "Nineteen",
	20:         "Twenty",
	30:         "Thirty",
	40:         "Forty",
	50:         "Fifty",
	60:         "Sixty",
	70:         "Seventy",
	80:         "Eighty",
	90:         "Ninety",
	100:        "Hundred",
	1000:       "Thousand",
	1000000:    "Million",
	1000000000: "Billion",
}

var (
	billion  = 1000000000
	million  = 1000000
	thousand = 1000
	hundred  = 100
)

func formatNumber(num string, res string) string {
	if res != "" {
		return fmt.Sprintf("%s %s", res, num)
	}

	return num
}

func extract(num int, res string) string {
	if num == 0 {
		return res
	}
	if numberWordsMap[num] != "" {
		if num == hundred || num == thousand || num == billion || num == million {
			return formatNumber("One "+numberWordsMap[num], res)
		}
		return formatNumber(numberWordsMap[num], res)
	}

	if num > billion {
		s := num / billion
		res = formatNumber(extract(s, "")+" "+"Billion", res)
		return extract(num%billion, res)
	}
	if num > million {
		s := num / million
		res = formatNumber(extract(s, "")+" "+"Million", res)
		return extract(num%million, res)
	}

	if num > thousand {
		s := num / thousand
		res = formatNumber(extract(s, "")+" "+"Thousand", res)

		return extract(num%thousand, res)
	}

	if num > hundred {
		s := num / hundred

		res = formatNumber(extract(s, "")+" "+"Hundred", res)
		return extract(num%hundred, res)
	}

	if num >= 10 {
		s, r := num/10, num%10
		if r == 0 {
			return formatNumber(numberWordsMap[s], res)
		}

		return formatNumber(fmt.Sprintf("%s %s", numberWordsMap[s*10], numberWordsMap[r]), res)
	}

	return formatNumber(numberWordsMap[num], res)
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	return extract(num, "")
}

func checkResult(num int, expected string) {
	result := numberToWords(num)
	if result == expected {
		println("The result is correct")
		return
	}

	println("---Incorrect---")
	println("Input", num)
	println("Result", result)
	println("Expected", expected)
}

func main() {
	checkResult(0, "Zero")
	checkResult(11, "Eleven")
	checkResult(1, "One")
	checkResult(101, "One Hundred One")
	checkResult(23, "Twenty Three")
	checkResult(100, "One Hundred")
	checkResult(123, "One Hundred Twenty Three")
	checkResult(12345, "Twelve Thousand Three Hundred Forty Five")
	checkResult(1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven")
	checkResult(2147483647, "Two Billion One Hundred Forty Seven Million Four Hundred Eighty Three Thousand Six Hundred Forty Seven")
	checkResult(203, "Two Hundred Three")
	checkResult(200, "Two Hundred")
	checkResult(1000, "One Thousand")
	checkResult(1001100, "One Million One Thousand One Hundred")
	checkResult(10000000, "Ten Million")
}
