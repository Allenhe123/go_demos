package main

// package被import才会执行package的init函数
import (
	"log"
	"os"
	_ "sample/matchers"
	"sample/search"
	//让 Go 语言对包做初始化操作，但是并不使用包里的标识符。为了让程序的
	//可读性更强，Go 编译器不允许声明导入某个包却不使用。下划线让编译器接受这类导入，并且
	//调用对应包内的所有代码文件里定义的 init 函数。
	// _ "github.com/goinaction/code/chapter2/sample/matchers"
	// "github.com/goinaction/code/chapter2/sample/search"
)

//程序中每个代码文件里的init函数都会在main函数执行前调用
//将标准库里日志类的输出，从默认的标准错误（stderr），设置为标准输出（stdout）设备
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	// 使用特定的项做搜索
	search.Run("cnn")
}
