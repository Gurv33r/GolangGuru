package guru

import (
	"fmt"
	"strings"
)

func Sort(b []int) {
	var alg string
	fmt.Println("Choose a phrase to choose a sorting algorithm")
	fmt.Println("	Bubble Sort = bubble")
	fmt.Println("	Insertion Sort = insert")
	fmt.Println("	Selection Sort = select")
	fmt.Print("Enter Here: ")
	fmt.Scanf("%s\n", alg)
	fmt.Println("You choose", alg)
	if strings.Compare(alg, "bubble") == 0 {
		fmt.Println(BubbleSort(b))
	} else if strings.Compare(alg, "insert") == 0 {
		fmt.Println(InsertionSort(b))
	} else if strings.Compare(alg, "select") == 0 {
		fmt.Println(SelectionSort(b))
	} else {
		fmt.Println("Action not supported!")
	}
}

func BubbleSort(arr []int) []int {
	fmt.Println("Starting Bubble Sort on", arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			fmt.Println("Comparing", arr[j], "and", arr[j+1])
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	fmt.Println("Finished Bubble Sort")
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
	fmt.Println("Finished Selection Sort")
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
	fmt.Println("Finished Insertion Sort")
	return arr
}
