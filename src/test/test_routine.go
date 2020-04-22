package testing

import (
    "fmt"
)

func TestingRoutine() {
    ch := make(chan struct{})
    for i := 0; i<10; i++ {
        go func(i int) {
            fmt.Println(i)
            ch <- struct {} {}
        } (i)
    }

    for j := 0; j<10; j++ {
        <-ch
    }
}
