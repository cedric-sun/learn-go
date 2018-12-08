package main

import "fmt"

type Person struct {
	Name string
	Age int
}

// Person implements the Stringer interface in package 'fmt'
func (p Person) String() string {
	// Sprintf is like String.format() in Java
	return fmt.Sprintf("%v (%v year(s) old)",p.Name, p.Age)
}

func person() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a)
	fmt.Println(z)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d",ip[0], ip[1], ip[2], ip[3])
}

func ip() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func main() {
	person()
	ip()
}
