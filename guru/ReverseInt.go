package guru

func ReverseInt(x int) int {
	reverse := 0
	for x > 0 {
		rightDigit := x % 10
		reverse = (reverse * 10) + rightDigit
		x /= 10
	}
	return reverse
}
