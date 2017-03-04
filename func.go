


package main
import (
    "fmt"
    "strings"
)
func worker() (int ,int) {
    return 99,90
}
func main() {
    a,b := worker();
    fmt.Println(a,b)
    var testStr := "fuck you"
    var testStr2 := "fuck you"
    var prefix := "nime"
    fmt.Println(strings.HasPrefix(testStr, prefix))
}
