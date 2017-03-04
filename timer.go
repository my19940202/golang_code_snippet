package main
import "time"
import "fmt"
func main() {
    timer1 := time.NewTimer(time.Second * 2)
    <-timer1.C
    fmt.Println("Timer 1 expired")
    timer2 := time.NewTimer(time.Second * 2)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}

// 定时器和sleep之间的区别： 定时器是有用原因之一就是你可以在定时器失效之前，取消这个定时器

