package main

// single-line import
import "fmt"

// import multiple package with parentheses
import (
	"math"
	"runtime"
)

// note that all imported package MUST be used, or there will be compile errors
func main() {
	fmt.Printf("Pi is: %v\n",math.Pi)
	fmt.Printf("GO is running on: %v\n",runtime.GOOS)
}
