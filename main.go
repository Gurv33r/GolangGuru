package main

import (
	"os"
	"strconv"

	"guru.org/hello-world-golang/guru"
)

//main function
func main() {
	allArgs := os.Args[1:]
	n, err := strconv.Atoi(allArgs[0])
	//fmt.Scanf("%d", &N)
	/*b := make([]int, N)
	for iter := 0; iter < len(b); iter++ {
		fmt.Scanf("%d", &b[iter])
	}
	c := b*/
	if err == nil {
		guru.FizzBuzz(n)
	}
}
