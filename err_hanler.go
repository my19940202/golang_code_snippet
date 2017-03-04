package main
 
import "fmt"
import "os"
 
func main(){

    // defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
    //     fmt.Println("c")
    //     if err:=recover();err!=nil{
    //         fmt.Println(err) // 这里的err其实就是panic传入的内容，55
    //     }
    //     fmt.Println("d")
    // }()
    // f()
    // 单独使用pannic 啥效果
    panic("a problem")
    // panic 的一个基本用法就是在一个函数返回了错误值
    // 但是我们并不知道（或者不想）处理时终止运行。
    // 这里是一个在创建一个新文件时返回异常错误时的panic 用法。
    _, err := os.Create("~/tmp/file")
    if err != nil {
        panic(err)
    }
}
 
func f(){
    fmt.Println("a")
    panic(55)
    fmt.Println("b")
    fmt.Println("f")
}
