package main  
import (  
    "fmt"
    "sync"
    )
var x  = 0  
func increment(wg *sync.WaitGroup) {
   // Each of these Goroutines run concurrently and the race condition occurs when trying to increment x in line no. 8 
   // as multiple Goroutines try to access the value of x concurrently.

    x = x + 1
    wg.Done() //--1
}
func main() {  
    var w sync.WaitGroup
    for i := 0; i < 1000; i++ {
        w.Add(1)        
        go increment(&w)
    }
    w.Wait() //untils zero
    fmt.Println("final value of x", x)
}