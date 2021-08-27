package main

import "fmt"

func main() {
	m1:=make(map[string]int,10)
	m2:='a'
	fmt.Printf("%T%T\n",m1,m2)
}
