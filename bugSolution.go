```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int, 10) // Buffered channel to prevent deadlock

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
                wg.Wait() // Wait for all goroutines to finish sending
                close(ch) // Close the channel after all senders are done
        }()

        for i := 0; i < 10; i++ {
                val := <-ch
                fmt.Printf("Main goroutine received %d
", val)
        }

        fmt.Println("Main goroutine exiting")
}
```