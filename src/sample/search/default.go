package search

import "log"

// defaultMatcher 实现了默认匹配器
// 空结构在创建实例时，不会分配任何内存。这种结构很适合创建没有任何状态的类型。
type defaultMatcher struct {
}

// 程序里所有的 init 方法都会在 main 函数启动前被调用。
// 一旦编译器发现 init 函数，它就会给这个函数优先执行的权限，保证其在 main 函数之前被调用
// init 函数将默认匹配器注册到程序里
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search 实现了默认匹配器的行为
// 如果声明函数的时候带有接收者，则意味着声明了一个方法。这个方法会和指定的接收者的类型绑在一起。
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	log.Println("default matcher search")
	return nil, nil
}
