package guru

//counts characters from a string
func Count(s string) int {
	sum := 0
	i := 0
	for i < len(s) {
		sum++
		i++
	}
	return sum
}
