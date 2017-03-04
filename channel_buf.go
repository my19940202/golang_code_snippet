package main
import "fmt"
func main() {
    messages := make(chan string, 4)
    messages <- "buffered"
    messages <- "channel"
    messages <- "3rd messages"
    messages <- "4th messages"
    fmt.Println(<-messages)
    fmt.Println(<-messages)
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}