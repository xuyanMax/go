package main

import "fmt"

/**
 * 
 * An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view 
 * into the elements of an array. In practice, slices are much more common than arrays.
 * 
 * The type []T is a slice with elements of type T.
 * A slice is formed by specifying two indices, a low and high bound, separated by a colon:
 * a[low : high]
 * 
 * Slices are like references to arrays
 * A slice does not store any data, it just describes a section of an underlying array.
 *
 * Changing the elements of a slice modifies the corresponding elements of its underlying array.
 * 
 * 
 * Slice length and capacity
 * A slice has both a length and a capacity.
 * The length of a slice is the number of elements it contains.
 * The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
 **/
func main() {

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	// Slice mutations
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)

	// nil slices, with length=0 and capacity=0
	fmt.Println("nil slices, with length=0 and capacity=0")
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

	//creating a slice with 
	//The make function allocates a zeroed array and returns a slice that refers to that array:
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	// append to slice
	var s []int
	// append works on nil slices
	s = append(s,0)
	s = append(s,1.2.3,4,5,6)
	printSlice("",s)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}