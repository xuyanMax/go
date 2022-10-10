package main

import (
	"fmt"
	"math"
)
//In Go, a name is exported if it begins with a capital letter. 
//When importing a package, you can refer only to its exported names. 
//Any "unexported" names are not accessible from outside the package.
func main() {
	fmt.Println(math.Pi)
	fmt.Println(math.pi)//error
}
