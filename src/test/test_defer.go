package testing

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func BigSlowOperation() {
	defer trace("BigSlowOperation")()
	time.Sleep(time.Second * 10)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s\n", msg)
	return func() { log.Printf("exit %s(%s)", msg, time.Since(start)) }
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func Testf() {
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func TestStack() {
	defer printStack()
	f(3)
}
