package guru

import "fmt"

// HelloWorld will print Hello World
func HelloWorld() string {
	return "Hello World"
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
