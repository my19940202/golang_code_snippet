package main

import (
    "fmt"
)

func main() {
    fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
    // 变量申明
    var num int
    var num2 int = 200
    num = 99
    num += num2
    fmt.Println(num)

    var v1, v2 = num, num2
    fmt.Println(v1, v2)

    v3, v4 := v1, v2
    fmt.Println(v3, v4)

    var v5 string = "100.1231"
    fmt.Println(v5)

    fmt.Printf("the bigger is %d\n", max(v2, v3))

}

/* 函数 */
func max(num1, num2 int) int {
    var result int
    if (num1 > num2) {
        result = num1
    } else {
        result = num2
    }
    return result 
}
