package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

//https://golang.google.cn/tour/concurrency/8

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    if t == nil {
        return
    }
    stack := make([]*tree.Tree, 0)
	//stack = append(stack, t)
    for len(stack) > 0 || t!=nil  {
        for t != nil {
            stack = append(stack, t)
            t = t.Left
        }
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        ch <- node.Value
		fmt.Println(node.Value)
        t = node.Right
    }
    close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 && ok2 {
			if v1 != v2 {
				return false
			}
		} else if ok1 || ok2 {
			return false
		} else {
			break
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))//true
	
    fmt.Println(Same(tree.New(2), tree.New(1)))//false

}
