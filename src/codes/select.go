package main

import "fmt"
import "time"

/*
使用go这个关键字，后面再跟上一个函数就可以创建一个线程。后面的这个函数可以是已经写好的函数，也可以是一个匿名函数
*/

//producer函数每隔500毫秒向channel里发送一个整数
func producer(ch chan<- int) { //这里的ch是send only channel
	var i = 1
	for i <= 5 {
		ch <- i
		fmt.Println("producer sent ", i)
		time.Sleep(time.Millisecond * 500)
		i++
	}
}

//consumer函数每隔1秒从channel里读取一个整数
func consumer(ch <-chan int) { //这里的channel是receive only channel
	for val := range ch {
		time.Sleep(time.Millisecond * 1000)
		fmt.Println("consumer received ", val)
	}
}

var timeout = make(chan bool, 1) //buffer size is 1, type is chan bool

func make_timeout() {
	time.Sleep(time.Millisecond * 1000)
	timeout <- true
}

func main() {
	//func1()

	//func2()

	//func3()

	func4()
}

func func1() {
	var ch = make(chan int, 6)
	//var sign chan bool //声明一个信号channel，值为nil
	go producer(ch) //创建一个goroutinue，去执行producer函数
	go consumer(ch) //创建一个goroutinue，去执行consumer函数
	time.Sleep(time.Second * 5)
	//<-sign             //sign == nil，所以该操作会导致主goroutinue阻塞,但是也会导致deadlock
}

func func2() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("received ", i1, " from c1")
	case c2 <- i2:
		fmt.Println("sent ", i2, " to c2")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Println("received ", i3, " from c3")
		} else {
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("no communication")
	}
}

func func3() {
	//当超时时间到的时候，case2 会操作成功。 所以 select 语句则会退出。 而不是一直阻塞在 ch 的读取操作上。 从而实现了对 ch 读取操作的超时设置。
	go make_timeout()
	ch := make(chan int)
	select {
	case <-ch:
		fmt.Println("received data from ch")
	case bval := <-timeout:
		fmt.Println("timeout ", bval)
	}
}

func func4() {
	//check is a channel is full using select's default
	//比如我们有一个服务， 当请求进来的时候我们会生成一个 job 扔进 channel， 由其他协程从 channel 中获取 job 去执行。
	//但是我们希望当 channel 瞒了的时候， 将该 job 抛弃并回复 【服务繁忙，请稍微再试。】 就可以用 select 实现该需求
	ch := make(chan int, 3)
	ch <- 1
	ch <- 3
	ch <- 4
	select {
	case ch <- 2:
		fmt.Println("sent 2 to channel ch, ch is not full")
	default: //select 将不会阻塞， 直接执行 default
		fmt.Println("channel ch is full")
	}
}
