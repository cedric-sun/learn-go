package main

import "fmt"

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
