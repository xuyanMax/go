package main

import "fmt"

//A var declaration can include initializers, one per variable.
//If an initializer is present, the type can be omitted; 
//the variable will take the type of the initializer.

//A var statement can be at package or function level.
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

//Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(k, c, python, java)
}
