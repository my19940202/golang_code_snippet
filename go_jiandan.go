package main
import "sync"
import "log"
import "fmt"
import "time"

func worker(id int, latency int) {
    for in := range input {
        if in == "done" {
            break
        }

        count ++
        time.Sleep(time.Millisecond * time.Duration(latency))
        log.Printf("[%d] %s this is in worker", id, in)
        // input 到达结束时  push Threads 个done进行 完整结束
        // 爬虫如何界定input 可以结束了
        if count == 32 {
            input <- "10000000"
            input <- "10000001"
            input <- "10000002"
            input <- "10000003"
            input <- "10000004"
            for i := 0; i < Threads; i++ {
                input <- "done"
            }
        }


        s := fmt.Sprintf("%d\t%s\n", id, in)
        output <- s
    }

    // output <- fmt.Sprintf("[%d]done", id)
}

const Threads = 8
const AvgBufSize = 8
const MaxNum = 30
var count int = 0

var input = make(chan string, Threads * AvgBufSize)
var output = make(chan string, Threads * AvgBufSize)

func main() {
    var wgInput sync.WaitGroup
    var wgOut sync.WaitGroup

    // input 塞进去一堆测试信息
    for i := 0; i < 30; i++ {
        input <- fmt.Sprintf("rec[%d]", i)
    }

    // for i := 0; i < Threads; i++ {
    //     input <- "done"
    // }

    for i := 0; i < Threads; i++ {
        // 添加goroutine的数量
        wgInput.Add(1)
        go func(id int) {
            worker(id, id)
            wgInput.Done()
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

    wgInput.Wait()

    output <- "done"
    wgOut.Wait()
}
