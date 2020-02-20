package main

import "fmt"
import "math"
import "errors"

/*
error类型是一个接口类型，这是它的定义：
type error interface {
    Error() string
}
我们可以在编码中通过实现 error 接口类型来生成错误信息。
函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息：
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // 实现
}

fmt.Println(fmt包在处理error时会调用Error方法)被调用，以输出错误
result, err:= Sqrt(-1)

if err != nil {
   fmt.Println(err)
}

*/

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math: square root of negative number")
	} else {
		return math.Sqrt(x), nil
	}
}

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strformat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
    `

	return fmt.Sprint(strformat, de.dividee)
}

func Divide(vardividee int, vardivider int) (result int, errmsg string) {
	if vardivider == 0 {
		data := DivideError{
			dividee: vardividee,
			divider: vardivider,
		}

		errmsg = data.Error()
		return

	} else {
		return vardividee / vardivider, ""
	}
}

func main() {
	var a float64 = -1
	b, err := sqrt(a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sqrt of ", a, " is ", b)
	}

	result, errormsg := Divide(100, 10)
	if errormsg == "" {
		fmt.Println("100/10 = ", result)
	} else {
		fmt.Println(errormsg)
	}

	result, errormsg = Divide(100, 0)
	if errormsg != "" {
		fmt.Println("errorMsg : ", errormsg)
	} else {
		fmt.Println("100 / 0 is: ", result)
	}
}
