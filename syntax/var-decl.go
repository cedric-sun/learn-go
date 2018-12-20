package main

import "fmt"

// declaration
var a int

// combined declaration of the same type
var b,c int

// declaration with initializer
var d,e int = 3,5

// type can be omitted when initializer exists
var perl,python,java = true, false ,"string!"

// only variable / function with capital first letter is exported
var ExportValue float64 = 3.1415;

func main() {
	// short declaration, can only exist inside function
	ruby := "ruby is a programming langauge"
	fmt.Println(ruby)
}
