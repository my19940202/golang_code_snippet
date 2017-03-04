
// 解决爬虫并发过程中 input -> worker -> output 流程中
// worker 动态影响 input的问题
package main
import "sync"
import "fmt"
import "time"
var p = fmt.Println
const Threads = 8
const AvgBufSize = 8
const MaxNum = 30
var index int = 0

func worker(url string, chData chan string,size *int) {
    p(url)
    if  index < 30 {
        chData <- fmt.Sprintf("http://pycm.baidu.com:8081/%d.html", index)    
        (*size) ++
    }
    index ++
}

func main() {
    var wg sync.WaitGroup
    // 并发控制 
    ch := make(chan string, Threads)
    // 
    chData := make(chan string, 100)
    size := 1
    chData <- "http://pycm.baidu.com:8081"

    for i := 0; i< Threads; i++ {
        ch <- "test"
    }
    // 轮巡 检测 size是否已经为0
    go func () {
        for {
            time.Sleep(time.Second * 3)
            if size == 0 {
                chData <- "done"
                break
            }
        }
    }()

    for  {
        if size == 0 {
            break
        }
        urlStr := <- chData
        wg.Add(1)
        go func(url string) {
            <- ch
            worker(url, chData, &size)
            size--
            ch <- "test"
            wg.Done()
        } (urlStr)
    }

    wg.Wait()
    p("结束")

}
