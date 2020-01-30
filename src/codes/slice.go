package main

import "fmt"

//内置类型切片("动态数组"),切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大

//可以声明一个未指定大小的数组来定义切片：
//var numbers []int

//使用make()函数来创建切片:
//var slice1 []type = make([]type, len)
//slice1 := make([]type, len)
//make([]T, length, capacity)
//这里 len 是数组的长度并且也是切片的初始长度。

func printslice(x []int) {
	fmt.Println(len(x), cap(x))
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
	fmt.Println("-------------------")
}

func main() {
	//可以声明一个未指定大小的数组来定义切片：
	var numbers []int
	fmt.Println(len(numbers), cap(numbers))
	if numbers == nil {
		fmt.Println("this slice is nil")
	}

	var numbers1 []int = make([]int, 3, 5)
	printslice(numbers1)

	//直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
	s := []int{1, 2, 3}
	printslice(s)

	var array [5]int = [5]int{6, 7, 8, 9, 10}

	ss := array[:]
	printslice(ss)

	ss1 := array[1:3] //will include array[1], array[2]
	printslice(ss1)

	ss2 := array[:3]
	printslice(ss2)

	ss3 := array[1:]
	printslice(ss3)

	dd := ss[1:3]
	printslice(dd)

	//如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
	//拷贝切片的 copy 方法和向切片追加新元素的 append 方法
	numbers = append(numbers, 0)
	printslice(numbers)

	numbers = append(numbers, 2, 3, 4)
	printslice(numbers)

	double := make([]int, len(numbers), cap(numbers)*2)
	copy(double, numbers)
	printslice(double)

}
