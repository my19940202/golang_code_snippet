package main
import (
    "fmt"
)

func main() {
    // 变量申明
    var num int
    var num2 int = 101
    num = 99
    num += num2
    if num <= 200 {
       fmt.Printf("<= 200")
    } else {
        fmt.Printf("> 200")   
    }
    /*
    } else {
    下面这种格式会报错，上面这种不会
    } 
    else {
    */ 
}
