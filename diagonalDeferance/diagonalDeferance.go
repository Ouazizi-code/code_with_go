package main

import "fmt"

func diagonalDeferance(arr [][]int32) int32 {
	var diag01 int32 = 0;
	var diag02 int32 = 0;

	// using for loop to fill diag with each value this method
	// usfull just with the care matrix
	for idx := range arr {
		diag01 += arr[idx][idx]
		diag02 += arr[idx][len(arr)-idx-1]
	}

	// here we returned the abs value for the deferance
	 var deferance int32 = diag01 - diag02
	if deferance < 0 {
		deferance = deferance * -1
	}else{
		deferance = deferance ;
	}

	return deferance ;
}

func main() {

	arr := [][]int32{
		{1, 3, 9},
		{2, 3, 9},
		{9, 6, 8},
	}

	fmt.Printf("=== diagonale deferance ===\n")
	diagonalDeferance(arr)
	fmt.Printf("\n program executed succefully \n")
}
