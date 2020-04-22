package testing

import (
	"fmt"
	"errors"
)

func echo(str string) (resp string, err error) {
	if str == "" {
		err = errors.New("empty request")
		return
	}
	resp = fmt.Sprintf("echo: %s\n", str)
	return
}

func TestingError() {
	for _, req := range []string{"", "hello"} {
		fmt.Printf("request: %s\n", req)
		resp, err := echo(req)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("response: %s\n", resp)
	}
}

