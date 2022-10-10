package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first,second:=0,1
	return func() int {
		first, second = second, first+second
		return first
	}
}

func fibonacci_recurs(i int) int {
	if i<=1 {
		return i
	}
	return fibonacci_recurs(i-1) + fibonacci_recurs(i-2)
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	for j := 0; j<10; j++ {
		fmt.Println(fibonacci_recurs(j))
	}
}
