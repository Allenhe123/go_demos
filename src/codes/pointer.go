package main

import "fmt"

func swap(x, y *int) (int, int) {
	temp := *x
	*x = *y
	*y = temp
	return *x, *y
}

func main() {
	var x int = 10
	fmt.Println("address of x is: ", &x)

	//pointer is nil by default
	var ip *int
	var fp *float32
	fmt.Println("dump pointers: ", ip, fp)
	if ip == nil && fp == nil {
		fmt.Println("ip and fp are null pointer")
	} else {
		fmt.Println("ip and fp are not all null pointer")
	}
	//will be error!
	//fmt.Println("dump *pointers: ", *ip, *fp)

	ip = &x
	fmt.Println("dump ip and *ip: ", ip, *ip)

	//ptr array
	const MAX int = 3
	array := []int{10, 100, 1000}
	var ptr [MAX]*int
	for i := 0; i < MAX; i++ {
		ptr[i] = &array[i]
	}

	for i := 0; i < MAX; i++ {
		fmt.Println(ptr[i], *ptr[i])
	}

	//ptr to ptr
	var ptrptr **int
	fmt.Println(ptrptr)
	ptrptr = &ip
	fmt.Println(ptrptr, *ptrptr, **ptrptr)

	//ptr param
	p := 99
	q := 88
	fmt.Println(p, q)
	swap(&p, &q)
	fmt.Println(p, q)
}
