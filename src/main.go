package main

import (
	"fmt"
	"log"
	"test"
	// "os"
)

func init() {
	log.SetPrefix("BST : ")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}

func modify(s []int) {
	s = append(s, 100)
	fmt.Println(s)
}

func modify_1(s *[]int) {
	*s = append(*s, 100)
}

func main() {
	// fmt.Println(os.Getenv("PATH"))

	// testing.TestingRange()
	// testing.TestingSwitch()
	// testing.TestingError()
	// testing.TestintHtml2()

	// a := make([]string, 5 ,10)
	// a := []int{1, 2, 3, 4, 5}

	// fmt.Println(a)
	// // modify(a)
	// // fmt.Println(a)

	// modify_1(&a)
	// fmt.Println(a)

	testing.TestingCopy()

	fmt.Println("tesing done")
	log.Print("testing done!")
}
