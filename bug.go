```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int)

        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        fmt.Printf("Goroutine %d: starting
", i)
                        ch <- i
                        fmt.Printf("Goroutine %d: finishing
", i)
                }(i)
        }

        go func() {
                for i := 0; i < 10; i++ {
                        val := <-ch
                        fmt.Printf("Main goroutine received %d
", val)
                }
        }()

        wg.Wait()
        close(ch)
        fmt.Println("Main goroutine exiting")
}
```