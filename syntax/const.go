package main

import "fmt"

/*
	1. Instead of "var", use "const" to declare constants.
	2. Constants CANNOT be declared using ":=" syntax
	3.Type of constants will be automatically inferred.
*/
const Pi = 3.14

const (
	goCreater = "Google"
	location = "California"
)

func main() {
	const localConst = "Hello World!"
	fmt.Printf("Pi is %v\n",Pi)
	fmt.Printf("The creater of Golang is %v at %v\n",goCreater, location)
	fmt.Println(localConst)
}
