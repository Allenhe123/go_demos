package main

import "fmt"

//Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了interface的所有这些方法就是实现了这个接口。

//定义了一个接口Phone，接口里面有一个方法call()
type Phone interface {
	call()

	hook(x int) float32
}

//定义结构体Nokiaphone
type Nokiaphone struct {
}

//实现接口方法
func (nokia Nokiaphone) call() {
	fmt.Println("i am nokia phone")
}

func (nokia Nokiaphone) hook(x int) float32 {
	fmt.Println("i am nokia phone, hook method", x)
	return 0.0
}

//定义结构体Iphone
type Iphone struct {
}

//实现接口方法
func (nokia Iphone) call() {
	fmt.Println("i am Iphone phone")
}

func (nokia Iphone) hook(x int) float32 {
	fmt.Println("i am Iphone phone, hook method", x)
	return 0.0
}

func main() {

	//定义了一个Phone类型变量，并分别为之赋值为NokiaPhone和IPhone。然后调用call()方法
	var phone Phone
	phone = new(Nokiaphone)
	phone.call()
	phone.hook(100)

	phone = new(Iphone)
	phone.call()
	phone.hook(99)
}
