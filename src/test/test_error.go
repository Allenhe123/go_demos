package testing

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
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

func WaitServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Get(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s), retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to repond after %s", url, timeout)
}
