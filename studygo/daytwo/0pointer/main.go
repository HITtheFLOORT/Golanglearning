package main

import "fmt"
func sum(a int,b int)int {
	return a+b
}
func main() {
	n:=18
	p:=&n
	x:=[5]int{1,2,3,5,5}
	p2:=&x
	v:=new(int)
	v=p
	fmt.Printf("%T\n",p)
	fmt.Printf("%p%T\n",p2,p2)
	fmt.Printf("%v%T\n",v,v)
}
