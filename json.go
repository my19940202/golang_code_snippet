package main
import "encoding/json"
import "fmt"
import "os"
type Response1 struct {
    Page   int
    Fruits []string
}
// string `json:"serverName"`  控制 json化的key是小写的
type Response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}
func main() {
    //Marshal 对象转换为JSON
    // bolB, _ := json.Marshal(123894123894)
    // // fmt.Println("type is ", bolB.(type))
    // fmt.Println(bolB, "-------")
    // fmt.Println(string(bolB))
    // intB, _ := json.Marshal(1)
    // fmt.Println(string(intB))
    // fltB, _ := json.Marshal(2.34)
    // fmt.Println(string(fltB))
    // strB, _ := json.Marshal("gopher")
    // fmt.Println(string(strB))
    // slcD := []string{"apple", "peach", "pear"}
    // fmt.Println(len(slcD) ," - - - ")
    // slcB, _ := json.Marshal(slcD)
    // fmt.Println(string(slcB))

    // mapD := map[string]int{"apple": 5, "lettuce": 7}
    // mapB, state := json.Marshal(mapD)
    // fmt.Println(string(mapB), state)
    // res1D := &Response1{
    //     Page:   1,
    //     Fruits: []string{"apple", "peach", "pear"}}
    // res1B, _ := json.Marshal(res1D)
    // fmt.Println(string(res1B))
    // res2D := Response2{
    //     Page:   1,
    //     Fruits: []string{"apple", "peach", "pear"}}
    // res2B, _ := json.Marshal(res2D)
    // fmt.Println(string(res2B))
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
    type TP struct {
        num float64
        strs []string
    }

    var dat TP
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)
    // 这个断言有啥用
    num := dat.num
    fmt.Println(num)
    strs := data.strs
    str1 := strs[0]
    fmt.Println(str1)

    // byte 和string的区别
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := &Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}

