package guru

import (
	"errors"
)

func ReverseInt(x int) (int, error) {
	if x < 0 {
		return x, errors.New("Integer parameter is negative!")
	} else if x < 10 {
		return x, errors.New("Integer parameter doesn't have at least 2 digits!")
	} else {
		reverse := 0
		for x > 0 {
			rightDigit := x % 10
			reverse = (reverse * 10) + rightDigit
			x /= 10
		}
		return reverse, nil
	}
}
