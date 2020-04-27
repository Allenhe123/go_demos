package testing

import (
	"fmt"
	"os"
)

func TestingDir() {
	var dirs = []string{"/tmp/a", "/tmp/b"}
	var rmdirs []func()
	for _, d := range dirs {
		dir := d
		os.MkdirAll(dir, 0755)
		fmt.Println("create dir: ", dir)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
			fmt.Println("remove dir: ", dir)
		})
	}

	for _, f := range rmdirs {
		f()
	}
}
