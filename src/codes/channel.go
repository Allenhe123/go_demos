package main

import "fmt"
import "time"

func main() {
	var i = 3

	go func(a int) {
		fmt.Println(a)
		fmt.Println(1)
	}(i)

	fmt.Println(2)
	//因为程序会优先执行主线程，主线程执行完成后，程序会立即退出，没有多余的时间去执行子线程。如果在程序的最后让主线程休眠1秒钟，那程序就会有足够的时间去执行子线程。
	time.Sleep(1 * time.Second)

	//channelTest()
	//channelTest1()
	//channelTest2()

	/*
	   channel是没有缓冲的，所以当生产者给channel赋值后，生产者这个线程会阻塞，直到消费者线程将channel中的数据取出。消费者第一次将数据取出后，
	   进行下一次循环时，消费者的线程也会阻塞，因为生产者还没有将数据存入，这时程序会去执行生产者的线程。程序就这样在消费者和生产者两个线程间不断切换，直到循环结束。
	*/
	//ch := make(chan int)
	//缓冲区可以存储10个int类型的整数，在执行生产者线程的时候，线程就不会阻塞，一次性将10个整数存入channel，在读取的时候，也是一次性读取
	ch := make(chan int, 10)
	go produce(ch)
	go consumer(ch)

	time.Sleep(1 * time.Second)
}

/*
这段代码执行时会出现一个错误：fatal error: all goroutines are asleep - deadlock!
这个错误的意思是说线程陷入了死锁，程序无法继续往下执行。
我们创建了一个无缓冲的channel，然后给这个channel赋值了，程序就是在赋值完成后陷入了死锁。因为我们的channel是无缓冲的，即同步的，赋值完成后来不及读取channel，
程序就已经阻塞了。这里介绍一个非常重要的概念：channel的机制是先进先出，如果你给channel赋值了，那么必须要读取它的值，不然就会造成阻塞，
当然这个只对无缓冲的channel有效。对于有缓冲的channel，发送方会一直阻塞直到数据被拷贝到缓冲区；如果缓冲区已满，则发送方只能在接收方取走数据后才能从阻塞状态恢复。
*/
func channelTest() {
	//创建无缓冲channel
	chreadandwrite := make(chan int)
	//chreadonly := make(<-chan int)
	//chwriteonly := make(chan<- int)

	chreadandwrite <- 1

	//即使没有下面这个goroutine也会deadlock
	go func() {
		a := <-chreadandwrite
		fmt.Println(a)
	}()
}

//给channel增加缓冲区，然后在程序的最后让主线程休眠一秒
func channelTest1() {
	ch := make(chan int, 1) //buffer size is 1, type is chan int
	ch <- 1

	go func() {
		a := <-ch
		fmt.Println(a)
	}()
}

//把ch<-1这一行代码放到子线程代码的后面
func channelTest2() {
	ch := make(chan int)

	go func() {
		a := <-ch
		fmt.Println(a)
	}()

	ch <- 1
}

func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}
func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}
