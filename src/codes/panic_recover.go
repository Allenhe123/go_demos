package main

import "fmt"

/*
Golang 有2个内置的函数 panic() 和 recover()，用以报告和捕获运行时发生的程序错误，与 error 不同，
panic-recover 一般用在函数内部。一定要注意不要滥用 panic-recover，可能会导致性能问题，我一般只在未知输入和不可靠请求时使用。
golang 的错误处理流程：当一个函数在执行过程中出现了异常或遇到 panic()，正常语句就会立即终止，然后执行 defer 语句，
再报告异常信息，最后退出 goroutine。如果在 defer 中使用了 recover() 函数,则会捕获错误信息，使该错误信息终止报告。

defer是在return之前执行的
大家都知道golang的defer关键字，它可以在函数返回前执行一些操作，最常用的就是打开一个资源（例如一个文件、数据库连接等）
时就用defer延迟关闭改资源，以免引起内存泄漏。
func openfile() (ok bool) {
    file, err := os.Open("c:\a.txt")
    defer file.Close()
    // doSomething
    return true
}

defer关键字用来标记最后执行的Go语句，一般用在资源释放、关闭连接等操作，会在函数关闭前调用。
多个defer的定义与执行类似于栈的操作：先进后出，最先定义的最后执行。

defer 在声明时不会立即执行，而是在函数 return 后，再按照 FILO （先进后出）的原则依次执行每一个 defer，
一般用于异常处理、释放资源、清理数据、记录日志等.

defer会在外围函数或者方法返回之前但是其返回值（如果有的话）计算之后执行。
这样就有可能在一个被延迟执行的函数内部修改外围函数的命名返回值.

return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。


*/

func deferRet(x, y int) (z int) {
	defer func() {
		z += 10
	}() // at last is () -- it is a anonymous func call

	z = x + y
	return z + 50 // 执行顺序 z = z+50 -> (call defer)z = z+100 -> ret
}

func f() {
	fmt.Println("a")
	panic(55)
	fmt.Println("b")
	fmt.Println("f")
}

//return 5
func funcA() int {
	x := 5
	defer func() {
		x += 1
	}()

	//temp = x      #temp变量表示未显示声明的return变量
	//返回temp的值，在将x赋值给temp后，temp未发生改变，最终返回值为5

	return x
}

//return 6
//返回x的值，先对其复制5，接着函数中改变为6，最终返回值为6
func funcB() (x int) {
	defer func() {
		x += 1
	}()
	//x = 5 然后 x += 1
	return 5
}

//return 5
func funcC() (y int) {
	x := 5
	defer func() {
		x += 1
	}()
	//temp = x = 5, return temp
	return x
}

//return 5
func funcD() (x int) {
	defer func(x int) {
		x += 1
	}(x)
	//值传递
	return 5
}

func main() {
	//defer的执行顺序是逆序的，也就是先进后出的顺序
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	x := deferRet(1, 1)
	fmt.Println(x)

	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) //这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()

	f()
}
