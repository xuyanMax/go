package main
import "fmt"

//var c = make(chan int,2)
// An unbuffered channel will have no room to store any data. 
// So in order for a value to be passed over an unbuffered channel, the sending Go routine will block until the receiving Go routine has received the value
// in the following case, go routine will block c <- 0 waiting for f() - then a is definitely "hello world" not "nothing"
var c = make(chan int)
var a string

func f() {
    a = "hello, world"
    x := <- c
    fmt.Println(x)
}

func main() {
	a = "nothing"
    go f()
    c <- 0
    print(a)
}