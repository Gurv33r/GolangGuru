package guru

import "strings"

func Toolchanger(tools []string, start int32, target string) int32 {
	i, test := start, ""
	var left, right int32 = 0, 0
	//move left
	for strings.Compare(test, target) == 0 {
		test = tools[i]
		i--
		left++
		if i == 0 {
			i = int32(len(tools)) - 1
		}
	}
	i, test = start, ""
	//move right
	for strings.Compare(test, target) == 0 {
		test = tools[i]
		i++
		right++
		if i == int32(len(tools)) {
			i = 0
		}
	}
	if left > right {
		return right
	} else {
		return left
	}
}
