package main

import (
	"strings"
	"fmt"
)

func WordCount(str string) map[string]int {
    m1 := make(map[string]int)
	//var m1 map[string]int
	fmt.Print(m1)
	//fmt.Print(m2)
    for _, v := range strings.Fields(str) {
        if _, ok := m1[v]; ok {
            m1[v] = m1[v] + 1
        } else {
            m1[v] = 1
        }
    }
	return m1
}

func main() {
	WordCount("hellow world again world!")
}
