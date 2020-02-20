package main

import "fmt"

func main() {
	var b int = 15
	var a int
	numbers := [6]int{1, 2, 3, 5}

	for a := 0; a < 10; a++ {
		fmt.Println("the value of a is: ", a)
	}

	fmt.Println("@@@after loop a is: ", a) //a=0
	fmt.Println("@@@first loop end.")

	for a < b {
		a++
		fmt.Println("the value of a is: ", a)
	}

	for i, x := range numbers {
		fmt.Println("index: ", i, "value of a is: ", x)
	}

	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if (i % j) == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Println(i, " is a sushu")
		}
	}

	for {
		fmt.Println("forever")
		break
	}

	var aa int = 10
LOOP:
	for aa < 20 {
		if aa == 15 {
			aa = aa + 1
			goto LOOP
		}
		fmt.Println("The value of aa is ", aa)
		aa++
	}

}
