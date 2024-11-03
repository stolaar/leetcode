package main

type listnode struct {
	Val  rune
	Next *listnode
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	var list *listnode
	var head *listnode

	tests := []*listnode{}

	for _, r := range goal {
		node := &listnode{
			Val: r,
		}

		if r == rune(s[0]) {
			tests = append(tests, node)
		}

		if list == nil {
			list = node
			head = list
			continue
		}

		list.Next = node
		list = list.Next

	}

	list.Next = head

	for _, test := range tests {
		check := true
		for i := 0; i < len(s); i++ {
			if s[i] != byte(test.Val) {
				check = false
				break
			}
			test = test.Next
		}

		if check {
			return true
		}
	}
	return false
}

func main() {
	println(rotateString("abcde", "cdeab"))
	println(rotateString("abcde", "abced"))
}
