package main

import "fmt"

func main() {
	q:=[]int{2,3,5,7,11,13}
	fmt.Println(q)
	p:=q[:]
	fmt.Println(p)
	p[3]=666
	fmt.Println(q)
}
