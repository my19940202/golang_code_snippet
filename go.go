package main

import (
    "fmt"
)

func main() {

	go func() {
		fmt.Println("111")
	}()
	go func() {
		fmt.Println("222")
	}()
	go func() {
		fmt.Println("333")
	}()

    fmt.Println("---------")
    // 使用了通道之后下面三个go会严格按照顺序来
    ch1 := make(chan int, 1)
    ch2 := make(chan int, 1)
    ch3 := make(chan int, 3)
    go func() {
        fmt.Println("1")
        ch1 <- 1
    }()
    go func() {
        <-ch1
        fmt.Println("2")
        ch2 <- 2
    }()
    go func() {
        <-ch2
        fmt.Println("3")
        ch3 <- 3
    }()
    <-ch3
}