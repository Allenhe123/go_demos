package main

import (
	"fmt"
	"math"
)

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}

//传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
//默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

func swapstring(x, y string) (string, string) {
	return y, x
}

//引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
//引用传递指针参数传递到函数内，以下是交换函数 swap() 使用了引用传递：

func swapstring2(x, y *string) (string, string) {
	var temp string
	temp = *x
	*x = *y
	*y = temp
	return *x, *y
}

//Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
//以下实例中，我们创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量

func getSuquence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

//Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数
//接收者可以是自己定义的一个类型，这个类型可以是struct，interface，甚至我们可以重定义基本数据类型。
//接收者是指针和非指针的区别，我们可以看到当接收者为指针式，我们可以通过方法改变该接收者的属性，但是非指针类型缺做不到。
//我们可以把接受者当作一个class，而这些方法就是类的成员函数，当接收者为指针类型是就是c++中的非const成员函数，
//为非指针时就是const成员函数，不能通过此方法改变累的成员变量。

type Circle struct {
	radius float64
}

func (c1 Circle) getArea() float64 {
	return 3.14 * c1.radius * c1.radius
}

func main() {
	ret := max(90, 100)
	fmt.Println("the max value is: ", ret)

	a, b := swapstring("allenhe", "hommyzhang")
	fmt.Println("after swap string are: ", a, b)

	var str1 string = "OOO"
	var str2 string = "PPP"
	fmt.Println("before swap string are: ", str1, str2)
	swapstring2(&str1, &str2)
	fmt.Println("after swap string are: ", str1, str2)

	//Go 语言可以很灵活的创建函数，并作为值使用。
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	root := getSquareRoot(400)
	fmt.Println("sqrt of 400 is: ", root)

	//nextNumber 为一个函数，i 为 0
	nextNumber := getSuquence()
	//调用 nextNumber 函数，i 变量自增 1 并返回
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	// 创建新的函数 nextNumber1，并查看结果
	nextNumber1 := getSuquence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	var cc Circle
	cc.radius = 10.00
	area := cc.getArea()
	fmt.Println("Area is: ", area)

}
