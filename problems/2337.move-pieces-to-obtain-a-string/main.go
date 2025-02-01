package main

func canChange(start string, target string) bool {
	i, j := 0, 0
	ls, rs := 0, 0

	for i < len(start) && j < len(target) {

		if start[i] == 'R' && target[j] == 'L' {
			return false
		}

		if start[i] == target[j] {

			if start[i] == 'L' && rs > 0 {
				return false
			}

			if start[i] == 'R' && ls > 0 {
				return false
			}

			j += 1
			i += 1
			continue
		}

		if target[j] == 'L' {
			ls += 1
			i += 1
			j += 1
			continue
		}

		if start[i] == 'L' {
			if rs > 0 || ls <= 0 {
				return false
			}

			ls -= 1
			i += 1
			j += 1
			continue
		}

		if target[j] == 'R' {
			if ls > 0 || rs <= 0 {
				return false
			}
			rs -= 1
			j += 1
			i += 1
			continue
		}

		if start[i] == 'R' {
			rs += 1
			j += 1
			i += 1
			continue
		}

		return false
	}

	return rs == 0 && ls == 0
}

func main() {
	println(canChange("_L__R__R_", "L______RR"))
	println(canChange("R_L_", "__LR"))
	println(canChange("_L__R__RL", "L_____RLR"))
	println(canChange("LL_", "_L_"))
}
