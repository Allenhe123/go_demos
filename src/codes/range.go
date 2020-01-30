package main

import "fmt"

// range 关键字用于for循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum is: ", sum)

	for i, num := range nums {
		fmt.Println(i, num)
	}

	//range也可以用在map的键值对上。
	mmp := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range mmp {
		fmt.Println(k, v)
	}

	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
