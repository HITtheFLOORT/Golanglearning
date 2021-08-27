package main

import "fmt"
//切片是引用类型，指向底层数组
func main() {
	var a1 =[3]int {1,2,3}
	a2 :=[...]int{1,2,3}
	fmt.Println(a1,a2)
	a3:=[][]int{{1,2},{1,2}}
	a4:=a3 //值类型 支持== !=
	fmt.Println(a4)
	fmt.Println(a1==a2)
	a5:=a2[:2]
	a6:=a2[1:]
	fmt.Println(len(a5),cap(a5))
	fmt.Println(len(a6),cap(a6))//cap 看切片头 看源数组尾 容量预估的
	fmt.Println(a5,len(a5),cap(a5))
	fmt.Printf("%p\n",a5)
	s1:=make([]int,5,10)
	fmt.Println(len(s1),cap(s1))
	a5=append(a5,5)
	fmt.Printf("%p\n",a5)
	fmt.Println(a5,len(a5),cap(a5))
	fmt.Println(a1)
}
