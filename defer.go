package main
 
import "fmt"
 
func deferIt() {
    defer func() {
        fmt.Println(1)
    }()
    defer func() {
        fmt.Println(2)
    }()
    defer func() {
        fmt.Println(3)
    }()
    fmt.Println(4)
}
func deferIt2() {
    for i := 1; i < 5; i++ {
        defer fmt.Print(i)
    }
}

func deferIt3() {
    f := func(i int) int {
        fmt.Println(i)
        return i * 10
    }
    for i := 1; i < 5; i++ {
        defer fmt.Println(f(i))
    }
}
func deferIt4() {
    for i := 1; i < 5; i++ {
        // 正确的用法是：把要使用的外部变量作为参数传入到匿名函数中。
        defer func(n int) {
            fmt.Print(n)
        }(i)
        // defer fmt.Print(i)
        // 4321
        // defer func() {
        //     fmt.Print(i)
        // }()
        //5555
    }
}

func fibonacci(num int) int {
	if num == 0 {
		return 0
	}
	if num < 2 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

func main() {
    // deferIt()
    // deferIt4()
    // defer fmt.Println("last in the main")
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", fibonacci(i))
        defer func (n int) {
            fmt.Printf("%d ", n) 
        }(fibonacci(i));
	}
}