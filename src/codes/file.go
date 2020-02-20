package main

//https://studygolang.com/pkgdoc

import (
	"bufio" //缓存IO
	"fmt"
	"io"
	"io/ioutil" //io 工具包
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileIsExit(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func main() {
	str := "helloworld"
	filename := "a.txt"
	var f *os.File
	var err error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err")
		}
	}()

	if fileIsExit(filename) {
		f, err = os.OpenFile(filename, os.O_APPEND, 0666)
		fmt.Println("file is exist")
	} else {
		f, err = os.Create(filename)
		fmt.Println("file is not exist")
	}
	check(err)
	defer f.Close()

	//使用 io.WriteString 写入文件
	n, err := io.WriteString(f, str)
	check(err)
	fmt.Printf("write %d bytes into a.txt\n", n)

	//使用 ioutil.WriteFile 写入文件
	var d1 []byte = []byte(str)
	err = ioutil.WriteFile(filename, d1, 0666)
	check(err)
	fmt.Printf("write %d bytes into a.txt\n", len(d1))

	//使用 File(Write,WriteString) 写入文件
	ff, err1 := os.Create("b.txt")
	check(err1)
	defer ff.Close()

	nn, err1 := ff.Write(d1)
	check(err1)
	fmt.Printf("write %d byte into b.txt\n", nn)

	nn, err1 = ff.WriteString("writeString")
	fmt.Printf("write %d byte into b.txt\n", nn)
	ff.Sync()

	// 使用 bufio.NewWriter 写入文件
	fff, err := os.Create("d.txt")
	check(err)
	defer fff.Close()

	w := bufio.NewWriter(fff)
	n, err = w.WriteString("bufio.NewWriter")
	check(err)
	fmt.Printf("write %d into d.txt\n", n)
	w.Flush()
}

/*

1)))))).
func ReadFile(filename string) ([]byte, error)
func WriteFile(filename string, data []byte, perm os.FileMode) error

import (
    "fmt"
    "io/ioutil"
)

    b, err := ioutil.ReadFile("test.log")
    if err != nil {
        fmt.Print(err)
    }
    fmt.Println(b)
    str := string(b)
    fmt.Println(str)


    d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile("test.txt", d1, 0644)
    check(err)

2))))).
func Open(name string) (*File, error)  //the associated file descriptor has mode O_RDONLY. If there is an error, it will be of type *PathError.
func OpenFile(name string, flag int, perm FileMode) (*File, error) 需要提供文件路径、打开模式、文件权限

  f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        log.Fatal(err)
    }
	defer f.close()
    if err := f.Close(); err != nil {
        log.Fatal(err)
    }


	   f, err := os.Create("/tmp/dat2")
    check(err)

    defer f.Close()

    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    n3, err := f.WriteString("writes\n")
    fmt.Printf("wrote %d bytes\n", n3)


    f.Sync()


    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    fmt.Printf("wrote %d bytes\n", n4)

    w.Flush()


    fi, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer fi.Close()
    fd, err := ioutil.ReadAll(fi)
    return string(fd)



*/
