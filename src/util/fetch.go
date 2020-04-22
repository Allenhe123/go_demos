package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	fetch2()
}

func fetch1() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

func fetch2() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") {
			fmt.Println("good url")
		} else {
			fmt.Println("bad url, will add http:// for it")
			url = "http://" + url
			fmt.Printf("joined url: %s\n", url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v, status: %v\n", err, resp.Status)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "fetch status: %v\n", resp.Status)
		defer resp.Body.Close()

		bytes, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("copy %d\n", bytes)
	}
}
