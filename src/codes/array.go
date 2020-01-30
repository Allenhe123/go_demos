package main

import (
	"fmt"
	//    "math"
)

//var variable_name [SIZE] variable_type
//var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
//void myFunction(param [10]int) {...}
//void myFunction(param []int) {...}

func average(array []int, size int) float32 {
	var sum int = 0
	var avg float32
	for i := 0; i < size; i++ {
		sum += array[i]
	}
	avg = float32(sum / size)
	return avg
}

func main() {
	var balance [10]float32
	balance[1] = 99.0
	for i := 0; i < len(balance); i++ {
		fmt.Println(balance[i])
	}

	var arr1 = [5]int{1, 2, 3, 4}
	var arr2 = [...]int{1, 2, 3, 4}
	for j, v := range arr1 {
		fmt.Println(j, v)
	}

	for k, u := range arr2 {
		fmt.Println(k, u)
	}

	////////////////////////

	var threadnum [2][3]int
	threadnum[1][1] = 999

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(threadnum[i][j])
		}
	}

	var a = [3][4]int{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{3, 3, 3, 3},
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(a[i][j])
		}
	}

	var arr3 = []int{1, 2, 3, 4, 5, 6, 7, 8}

	var avg float32
	avg = average(arr3, 8) //only can pass arr3, can not pass arr1 and arr2
	fmt.Println("average is: ", avg)

}
