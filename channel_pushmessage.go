package main

import "fmt"

func main() {
    messages := make(chan string)
    // go func() { messages <- "pingpingping" }()
    // go func() { messages <- "ping" }()
    go func() { messages <- "pingping" }()
    //  messages <- "pingping" 往channel里面 send 信息
    
    msg := <-messages
    fmt.Println(msg)
    // if msg {
    //     fmt.Println(msg)
    // } else {
    //     fmt.Println("no messages")
    // }
}
