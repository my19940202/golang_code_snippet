

// 解决爬虫并发过程中 input -> worker -> output 流程中
// worker 动态影响 input的问题
package main
import "sync"
import "fmt"
import "time"
import "strconv"
var p = fmt.Println
const Threads = 8
const AvgBufSize = 8
const MaxNum = 30
var index int = 0

func worker(url map[string]string, chData chan map[string]string,size *int) {
    p(url)
    depth, _ := strconv.Atoi(url["depth"])
    if depth >= 1 {
        newMsg := make(map[string]string)
        depth--
        newMsg["url"] = fmt.Sprintf("http://pycm.baidu.com:8081/%d.html", index)
        newMsg["depth"] = strconv.Itoa(depth)
        chData <- newMsg
        (*size) ++
    }
}

func main() {
    var wg sync.WaitGroup
    // 并发控制 
    ch := make(chan string, Threads)
    // 存储输入的url
    chData := make(chan map[string]string, 100)
    size := 1
    message1 := make(map[string]string)
    message1["url"] = "http://pycm.baidu.com:8081"
    message1["depth"] = "30"
    chData <- message1

    for i := 0; i< Threads; i++ {
        ch <- "test"
    }
    // 轮巡 检测 size是否已经为0  如果不加上这个 下面的for 不可避免的会遇到的争抢资源时
    // 已经抢到了 但是遇到size = 0 但是
    go func () {
        for {
            time.Sleep(time.Second * 1)
            if size == 0 {
                messageDone := make(map[string]string)
                messageDone["msg"] = "done"
                chData <- messageDone
                break
            }
        }
    }()

    for  {
        if size == 0 {
            break
        }
        urlMap := <- chData
        wg.Add(1)
        go func(url map[string]string) {
            <- ch
            worker(url, chData, &size)
            size--
            ch <- "test"
            wg.Done()
        } (urlMap)
    }

    wg.Wait()
    p("结束")

}
