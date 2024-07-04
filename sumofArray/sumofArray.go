package main

import "fmt"

// this function calculate the sum of array index and return a thier result

func sum(arr [5]int) int {
	var result int = 0
	var i int

	for i = 0; i < len(arr); i++ {
		result += arr[i]
	}

	return result
}

func main() {

	//var max int ;
	fmt.Print("=====hassan using go====== \n\n")
	//fmt.Print("enter maximum number ");
	//fmt.Scan("%d",&max);
	var i int
	var array [5]int
	for i = 0; i < 5; i++ {
		fmt.Print("enter a number :")
		fmt.Scanf("%d", &array[i])
	}
	var h int = sum(array)
	fmt.Printf(" The resul of addition is : %d\n", h)

	//sum(5,6,9,3,8);
}
