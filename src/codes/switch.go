package main

import "fmt"
import "time"

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is week end")
	default:
		fmt.Println("it is work day")
	}

	t := time.Now()
	hour := t.Hour()
	fmt.Println(t)
	switch {
	case hour <= 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}

}
