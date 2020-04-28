package testing

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// 实际使用中，更推荐使用Scanner对数据进行读取，而非直接使用Reader类。
// Scanner可以通过splitFunc将输入数据拆分为多个token，然后依次进行读取。
// 和Reader类似，Scanner需要绑定到某个io.Reader上，通过NewScannner进行创建
func TestingIO1() error {
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("->")
		text, err := in.ReadString('\n')
		if err == io.EOF {
			fmt.Println("manual stop")
			break
		}
		if err != nil {
			return fmt.Errorf("ReadString error: %s", err)
		}
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("allen", text) == 0 {
			fmt.Println("input allen")
		}
	}
	return nil
}

// 读取一个UTF-8编码的Unicode字符
// 如果你想从命令行中简单地读取一个 unicode 字符，我建议你使用 bufio.ReadRune()
func TestingIO2() error {
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("->")
		// ReadRune，读取一个utf-8字符
		r, size, err := in.ReadRune()
		if err == io.EOF {
			fmt.Println("manual stop")
			break
		}
		if err != nil {
			return fmt.Errorf("ReadRune error: %s", err)
		}
		fmt.Printf("read: %d, size: %d\n", r, size)

		switch r {
		case 'A', 'a':
			fmt.Println("Key A pressed")
			break
		case 'B', 'b':
			fmt.Println("key B pressed")
			break
		}
	}
	return nil
}

func TestingIO3() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		fmt.Println(in.Text())
	}
}

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
