package search

import (
	"fmt"
	"log"
)

type Result struct {
	Field   string
	Content string
}

// Matcher 定义了要实现的新搜索类型的行为
// 声明了一个名为 Matcher 的接口类型。
// interface 关键字声明了一个接口，这个接口声明了结构类型或者具名类型需要实现的行为。
// 一个接口的行为最终由在这个接口类型中声明的方法决定。
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 函数，为每个数据源单独启动 goroutine 来执行这个函数并发地执行搜索
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResult, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResult {
		results <- result
	}
}

// Display 从每个单独的 goroutine 接收到结果后在终端窗口输出
func Display(results chan *Result) {
	// for range 循环会一直阻塞，直到有结果写入通道。
	// 关闭了通道，for range 循环就会终止，Display 函数也会返回。
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
