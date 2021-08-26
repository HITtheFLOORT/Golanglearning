package main

import (
	"fmt"
	"math"
)

func main() {
	var a =math.MaxFloat64
	fmt.Printf("%T\n",a)
	b1:=true
	b2:=false
	fmt.Println(b2==b1)
	fmt.Printf("%v\n",b1)
}