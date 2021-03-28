package guru

import "fmt"

func FizzBuzz1(x int) {
	for i := 1; i <= x; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
func FizzBuzz2(x int) {
	for i := 1; i <= x; i++ {
		if i%3 == 0 || i%5 == 0 {
			if i%3 == 0 {
				fmt.Print("Fizz")
			}
			if i%5 == 0 {
				fmt.Print("Buzz")
			}
		} else {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
