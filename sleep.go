package main
import "time"
import "fmt"
func main() {
    // 就是setTimeout 差不多的功能
    time.Sleep(time.Second * 10)
    fmt.Println("sleep 10s 结束")
}
