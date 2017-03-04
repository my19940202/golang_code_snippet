package main

import "fmt"
import "net/url"
import "strings"

func main() {
    s := "postgres://user:pass@host.com:5432/path?name=xsb&sex=male&k=v#f"
    fmt.Println(s)
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    fmt.Println(u.Scheme)
    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)
    fmt.Println(u.Host)
    h := strings.Split(u.Host, ":")
    fmt.Println(h[0])
    fmt.Println(h[1])
    fmt.Println("Path: ", u.Path)
    fmt.Println("Fragment: ", u.Fragment)
    fmt.Println("RawQuery: ", u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}

