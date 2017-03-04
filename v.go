package main

import (
    "mymath"
    "fmt"
)



func main() {
    fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
    // 变量申明
    var name1 string = "fuck"
    var name2 string = "dude"
    v2,v3 := name1, name2
    fmt.Println(v2, v3)
}

