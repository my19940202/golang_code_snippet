package main
import "fmt"
import "time"
func worker(done chan bool, num int) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done", num)
    done <- true
}
func main() {
    done := make(chan bool, 1)
    for i := 0; i < 30; i++ {
        go worker(done, i)
    }
    <-done
}
