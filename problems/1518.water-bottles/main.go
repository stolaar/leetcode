package main

func numWaterBottles(numBottles int, numExchange int) int {
	empt, res := numBottles, numBottles

	for empt >= numExchange {
		exchanged := empt / numExchange
		empt = empt % numExchange

		res += exchanged
		empt += exchanged

		if empt == numExchange {
			res += 1
			break
		}
	}

	return res
}

func main() {
	println("result: ", numWaterBottles(9, 3), " - expected: ", 13)
	println("result: ", numWaterBottles(15, 4), " - expected: ", 19)
	println("result: ", numWaterBottles(1, 2), " - expected: ", 1)
	println("result: ", numWaterBottles(2, 2), " - expected: ", 3)
	println("result: ", numWaterBottles(20, 2), " - expected: ", 38)
}
