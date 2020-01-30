package main

import "fmt"

//Map 是一种无序的键值对的集合
//Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

func main() {
	//可以使用内建函数 make 也可以使用 map 关键字来定义 Map
	//如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
	var mmp1 map[string]string
	mmp1 = make(map[string]string)
	mmp1["a"] = "A"
	mmp1["b"] = "B"
	mmp1["c"] = "C"
	for k, v := range mmp1 {
		fmt.Println(k, v, mmp1[k])
	}

	//查看元素在集合中是否存在
	value, ok := mmp1["a"]
	if ok {
		fmt.Println("value of 'a' is: ", value)
	} else {
		fmt.Println("value of 'a' not exist")
	}

	//delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。
	kv := map[string]int{"a": 9, "b": 8, "c": 7}
	for k, v := range kv {
		fmt.Println(k, v)
	}

	delete(kv, "b")
	for k, v := range kv {
		fmt.Println(k, v)
	}

	mm := map[int]string{1: "a", 2: "b"}
	for key, val := range mm {
		fmt.Println(key, val)
	}

}
