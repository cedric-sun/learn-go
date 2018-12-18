package main

import "fmt"

var pow = []int{1,2,4,8,16,32,64,128}

func main() {
	for index,value := range pow {
		fmt.Printf("2^%d = %d\n",index,value)
	}

	// ignore index
	for _, value := range pow {
		fmt.Println(value)
	}

	// ignore value
	for index, _ := range pow {
		fmt.Println(index)
	}

	// or ignore value by entirely dropping the latter part
	for index := range pow {
		fmt.Println(index)
	}
}
