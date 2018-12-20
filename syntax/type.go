package main

import "fmt"

var boolVar bool
var stringVar string

var intVar   int
var int8Var  int8
var int16Var int16
var int32Var int32
var int64Var int64

var uintVar   uint
var uint8Var  uint8
var uint16Var uint16
var uint32Var uint32
var uint64Var uint64

var uintptrVar uintptr

// alias for uint8
var byteVar byte

// alias for int32
var runeVar rune

var float32Var float32
var float64Var float64

var complex64Var complex64
var complex128Var complex128

func main() {
	fmt.Printf("Zero value of %T is %v\n",boolVar, boolVar)
	fmt.Printf("Zero value of %T is %v\n",stringVar,stringVar)
	fmt.Printf("Zero value of %T is %v\n",intVar,intVar)
	fmt.Printf("Zero value of %T is %v\n",int8Var,int8Var)
	fmt.Printf("Zero value of %T is %v\n",int16Var,int16Var)
	fmt.Printf("Zero value of %T is %v\n",int32Var,int32Var)
	fmt.Printf("Zero value of %T is %v\n",int64Var,int64Var)
	fmt.Printf("Zero value of %T is %v\n",uintVar,uintVar)
	fmt.Printf("Zero value of %T is %v\n",uint8Var,uint8Var)
	fmt.Printf("Zero value of %T is %v\n",uint16Var,uint16Var)
	fmt.Printf("Zero value of %T is %v\n",uint32Var,uint32Var)
	fmt.Printf("Zero value of %T is %v\n",uint64Var,uint64Var)
	fmt.Printf("Zero value of %T is %v\n",uintptrVar,uintptrVar)
	fmt.Printf("Zero value of %T is %v\n",byteVar,byteVar)
	fmt.Printf("Zero value of %T is %v\n",runeVar,runeVar)
	fmt.Printf("Zero value of %T is %v\n",float32Var,float32Var)
	fmt.Printf("Zero value of %T is %v\n",float64Var,float64Var)
	fmt.Printf("Zero value of %T is %v\n",complex64Var,complex64Var)
	fmt.Printf("Zero value of %T is %v\n",complex128Var,complex128Var)
}
