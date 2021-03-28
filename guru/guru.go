package guru

import "fmt"

// HelloWorld will print Hello World
func HelloWorld() string {
	return "Hello World"
}

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

func BubbleSort(arr []int) []int {
	fmt.Println("Bubble sort input =", arr)
	for i := 0; i < len(arr)-1; i++ {
		fmt.Println("Comparing", arr[i], "and", arr[i+1])
		if arr[i] > arr[i+1] {
			tmp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = tmp
		}
	}
	return arr
}

func SelectionSort(arr []int) []int {
	fmt.Println("Selection sort input =", arr)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			fmt.Println("Comparing", arr[i], "and", arr[j])
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func InsertionSort(arr []int) []int {
	fmt.Println("Insertion sort input =", arr)
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func FizzBuzz(x int) {
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
		/*if i%3 == 0 || i%5 == 0 {
			if i%3 == 0 {
				fmt.Print("Fizz")
			}
			if i%5 == 0 {
				fmt.Print("Buzz")
			}
		} else {
			fmt.Print(i)
		}
		fmt.Println()*/
	}
}
