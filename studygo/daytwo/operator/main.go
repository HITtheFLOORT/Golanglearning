package main

import "fmt"

func main() {
	var a1 =[3]int {1,2,3}
	a2 :=[...]int{1,2,3}
	fmt.Println(a1,a2)
	a3:=[][]int{{1,2},{1,2}}
	a4:=a3 //值类型 支持== !=
	fmt.Println(a4)
	fmt.Println(a1==a2)
}
