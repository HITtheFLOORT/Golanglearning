package main

import "fmt"

func main() {
	var a []int
	for i:=1;i<12;i++{
		a=append(a, i)
		fmt.Println(i)
	}
	for _,value:=range a{
		fmt.Println(value)
	}
}
