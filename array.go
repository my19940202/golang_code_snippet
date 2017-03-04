package main
import (
    "fmt"
)
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
func main() {
    var balance = [5]int{12, 1313, 34, 70, 500}
    var j int
    for j = 0; j < len(balance); j++ {
        fmt.Printf("Element[%d] = %d address = %x\n", j, balance[j], &balance[j])
    }

    // 数组的增删改
    arr1 :=[...] int {1,2,3,4,5}
    arr2 :=[5] int {1,2,3,4,5}
    arr1[3] = 1234123
    // out of range 
    // arr1[5] = 5555
    fmt.Println(arr1, arr2)


    // 切片比较方便
    myArr := make([]int, 2)
    myArr[0] = 123413
    myArr[1] = 123
    myArr  = append(myArr, 9090)
    myArr  = append(myArr, 1000)
    // myArr[2] = 3
    // myArr[5] = 3
    fmt.Println(myArr)   
}
