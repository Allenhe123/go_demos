package main

import "fmt"

//递归函数
func factorial(n uint64) uint64 {
	if n > 0 {
		//result := n * factorial(n-1)
		//return result
		return n * factorial(n-1)
	} else {
		return 1
	}
}

//类型转换
//type_name(expression)

func main() {

	fmt.Println(factorial(15))

	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Println(mean)

}
