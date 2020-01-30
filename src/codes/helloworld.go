package main

import "fmt"
import "unsafe"

var a = "hahaha"
var b string = "ice and fire"
var c bool
var x, y int

var (
	m int
	n bool
)

//全局变量是允许声明但不使用。 同一类型的多个变量可以声明在同一行, 多变量可以在同一行进行赋值
var e, f = 888, "my lady!"
var h, i int = 88, 77

//这种不带声明格式的只能在函数体中出现
//如果在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明，例如：a := 20
//就是不被允许的，编译器会提示错误 no new variables on left side of :=，但是 a = 20 是可以的，
//因为这是给相同的变量赋予一个新的值。
//d := 999     will be error!

func main() {
	fmt.Print("my lord!")
	fmt.Println("hello world!")
	println(100)

	d := 999 //right!

	//如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误
	//var q = 666

	fmt.Println(a, b, c, d)
	fmt.Println(x, y)
	fmt.Println(m, n)
	fmt.Println(e, f)
	fmt.Println(h, i)

	//要交换两个变量的值，则可以简单地使用 a, b = b, a。
	r, s := 33, 44
	fmt.Println(r, s)
	r, s = s, r
	fmt.Println(r, s)

	//空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
	//_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，
	//但有时你并不需要使用从一个函数得到的所有返回值。
	_, b = 5, "iceiceice"
	fmt.Println(b)

	//常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。const identifier [type] = value
	const LENGTH int = 111
	const WIDTH int = 222
	const ca, cb, cc = 12, false, "heihei"
	var area int = LENGTH * WIDTH
	println("aarea is : ", area)

	//常量还可以用作枚举
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	//常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过
	const (
		da = "abc"
		db = len(da)
		dc = unsafe.Sizeof(da)
		//dd = cap(da)
	)
	println(da, db, dc)

	//输出结果为：16
	//字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。
	ha := "hello"
	hsize := unsafe.Sizeof(ha)
	println(hsize)

	//iota，特殊常量，可以认为是一个可以被编译器修改的常量。 每出现一次iota，其所代表的数字会自动增加1。
	//第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
	/*	const (
		    	a = iota
		    	b
		    	c
			)
	*/
	const (
		ia = iota
		ib = iota
		ic = iota
	)
	println(ia, ib, ic)

	const (
		ta = iota //0
		tb        //1
		tc        //2
		td = "ha" //独立值  iota += 1
		te        //"ha"   iota += 1
		tf = 100  //100    iota += 1
		tg        //100    iota += 1
		th = iota //7,恢复计数
		ti        //8
	)
	println(ta, tb, tc, td, te, tf, tg, th, ti)

	//iota表示从0开始自动加1，所以si=1<<0, sj=3<<1（<<
	//表示左移的意思），即：i=1,j=6，这没问题，关键在sk和sl，从输出结果看，sk=3<<2，sl=3<<3。
	const (
		si = 1 << iota
		sj = 3 << iota
		sk
		sl
	)
	println(si, sj, sk, sl) //1, 6, 12, 24

}
