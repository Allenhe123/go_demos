package main

import "fmt"

func main() {
	a := 1
	a, b := 2, 4 //can use this to overwrite a
	/*
	   7. 除非特别指定，否则无法使用 nil 对变量赋值
	   nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。但是如果不特别指定的话，Go 语言不能识别类型，所以会报错。
	*/
	//x := nil  //will be error
	//var x = nil //will be error
	var x interface{} = nil

	/*8. Slice 和 Map 的 nil 值.初始值为 nil 的 Slice 是可以进行“添加”操作的，但是对于 Map 的“添加”操作会导致运行时恐慌。*/
	var s []int
	s = append(s, 1)

	var m map[string]int     //a nil map
	m = make(map[string]int) //must call this line, or will be error
	m["one"] = 1

	fmt.Println(a, b, x, s, m)

	//9.创建 Map 的时候可以指定 Map 的长度，但是在运行时是无法使用 cap() 功能重新指定 Map 的大小，Map 是定长的。
	//Map就没有定长一说，cap是针对数组、指针数组、Slice和Channel
	/*
		mm := make(map[string]int, 99)
		cap(mm) //error
	*/

	//10.字符串无法为 nil
	/*
		var x string = nil //error
		if x == nil {      //error
			x = "default"
		}
	*/
	var xx string //defaults to "" (zero value)
	if xx == "" {
		xx = "default"
	}

	/*
	 11.Go 语言中，传递的数组不是内存地址，而是原数组的拷贝，所以是无法通过传递数组的方法去修改原地址的数据的。
	*/
	x1 := [3]int{1, 2, 3}
	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr) //prints [7 2 3]

	}(x1)
	fmt.Println(xx, x1)

	//如果需要修改原数组的数据，需要使用数组指针（array pointer）。
	func(arr *[3]int) {
		(*arr)[0] = 100
		fmt.Println(*arr)
	}(&x1)
	fmt.Println(x1)

	//或者可以使用 Slice
	sx := []int{1, 2, 3}
	func(arr []int) {
		arr[0] = 99
		fmt.Println(arr)
	}(sx)
	fmt.Println(sx)

	//12 使用 Slice 和 Array 的 range 会导致预料外的结果
	//每次的 range 都会返回两个值，第一个值是在 Slice 和 Array 中的编号，第二个是对应的数据。
	x2 := []string{"a", "b", "c"}
	for _, v := range x2 {
		fmt.Println(v)
	}

}
