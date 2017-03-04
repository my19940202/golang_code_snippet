








package main
import "time"
import "fmt"
func main() {
    c1 := make(chan string, 1)
    go func() {
        fmt.Println("result 1 running")
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()
    select {
        case res := <-c1:
            fmt.Println(res)
        case <-time.After(time.Second * 3):
            fmt.Println("timeout 3 second")
    }
    c2 := make(chan string, 1)
    go func() {
        fmt.Println("result 2 running")
        time.Sleep(time.Second * 2)
        c2 <- "result 2"
    }()
    select {
        case res := <-c2:
            fmt.Println(res)
        case <-time.After(time.Second * 3):
            fmt.Println("timeout 2")
    }
}