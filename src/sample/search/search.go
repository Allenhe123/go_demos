package search

//从标准库中导入代码时，只需要给出要导入的包名。编译器查找包的时候，
//总是会到GOROOT和GOPATH环境变量引用的位置去查找。
import (
	"log"  //log 包提供打印日志信息到标准输出（stdout）、标准错误（stderr）或者自定义设备的功能
	"sync" // sync 包提供同步 goroutine 的功能
)

//包级变量
//map是Go语言里的一个引用类型，需要使用make来构造。如果不先构造map并将构造后的值赋值
//给变量，会在试图使用这个map变量时收到出错信息。这是因为map变量默认的零值是nil
var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	//获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("feeds size:", len(feeds))

	// 创建一个无缓冲的通道，接收匹配后的结果
	results := make(chan *Result)

	// 构造一个 waitGroup，以便处理所有的数据源
	var waitGroup sync.WaitGroup

	// 设置需要等待处理  每个数据源的 goroutine 的数
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个 goroutine 来查找结果
	for _, feed := range feeds {
		matcher, exist := matchers[feed.Type]
		if !exist {
			matcher = matchers["default"]
		}

		// 启动一个goroutine 来执行搜索
		// go启动了一个匿名函数作为goroutine。匿名函数是指没有明确声明名字的函数。
		// 匿名函数也可以接受声明时指定的参数。matcher和feed两个变量的值被传入匿名函数。
		// searchTerm，results，waitGroup是通过闭包的形式访问的。因为有了闭包，
		// 函数可以直接访问到那些没有作为参数传入的变量。匿名函数并没有拿到这些变量的副本，
		// 而是直接访问外层函数作用域中声明的这些变量本身。
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个goroutine来监控是否所有的工作都做完了
	// 这个匿名函数没有输入参数，使用闭包访问了WaitGroup和results 变量。
	// WaitGroup的Wait方法。这个方法会导致 goroutine阻塞，直到WaitGroup内部的计数到达0。
	go func() {
		// 等候所有任务完成
		waitGroup.Wait()

		// 用关闭通道的方式，通知Display函数可以退出程序了
		close(results)
	}()

	// 启动函数，显示返回的结果，并且在最后一个结果显示完后返回
	Display(results)
}

// Register 调用时，会注册一个匹配器，提供给后面的程序使用
func Register(feedType string, matcher Matcher) {
	if _, exist := matchers[feedType]; exist {
		log.Fatalln(feedType, "matcher already registered")
	}

	log.Println("register", feedType, "matcher")
	matchers[feedType] = matcher
}
