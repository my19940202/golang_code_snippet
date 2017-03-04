package main
import "log"
import "time"
import "sync"
import "fmt"

func worker(id int, latency int) {
    for in := range input {
        if in == "done" {
            break
        }

        time.Sleep(time.Millisecond * time.Duration(latency))
        log.Printf("[%d] %s this is in worker", id, in)

        s := fmt.Sprintf("%d\t%s\n", id, in)
        output <- s
    }

    output <- fmt.Sprintf("[%d]done", id)
}

const Threads = 2
const AvgBufSize = 20
var input = make(chan string, Threads * AvgBufSize)
var output = make(chan string, Threads * AvgBufSize)

func main() {
    var wg sync.WaitGroup
    var wgOut sync.WaitGroup

    // input 塞进去一堆测试信息
    for i := 0; i < 30; i++ {
        input <- fmt.Sprintf("rec[%d]", i)
    }

    for i := 0; i < Threads; i++ {
        input <- "done"
    }

    for i := 0; i < Threads; i++ {
        // 添加goroutine的数量
        wg.Add(1)
        go func(id int) {
            worker(id, id)
            wg.Done()
        } (i)
    }

    wgOut.Add(1)
    go func() {
        for out := range output {
            if out == "done" {
                break
            }

            log.Printf("[out]%s\n", out)
        }
        wgOut.Done()
        // Done:相当于Add(-1)
    }()

    wg.Wait()

    output <- "done"
    wgOut.Wait()
}
