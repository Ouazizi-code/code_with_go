package main

import (
	"fmt"
)

// SelectionSort sorts the given array using selection sort
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Find the index of the minimum element
		minimum := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minimum] {
				minimum = j
			}
		}
		// Swap the minimum element with the first element in the unsorted part
		arr[i], arr[minimum] = arr[minimum], arr[i]
	}
}

func main() {
	// Example usage
	array := []int{2, 4, 3, 1, 6, 8, 5, 0, 89}
	SelectionSort(array)
	fmt.Println("Sorted array:", array)
}
