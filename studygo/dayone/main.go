package main

import "fmt"
var a int=1
var b string
var c[] int
var gkd = "a"
//=是赋值，:=是声明与赋值
//批量声明
var(
	x int
	z string
)
//常量必须赋值
//iota是计数器
const s=123
const(
	x1=1
	x2   //未赋值的常量与上一次相同
	x3=3
)
const(
	i1=iota //const 出现时重置为0
	i2		//每次新增+1 实现了一种枚举
	i3=iota
)
const(
	_=iota
	KB=1<<(10*iota)//2^10
	MB=1<<(10*iota)//2^20
)
const name = iota
//无法直接定义二进制
var ba=077
var shiliu=0xff

func main() {
	fmt.Printf("hello golang%d\n",name)
	fmt.Printf("hello golang%d\n",i3)
	fmt.Printf("hello golang%d\n",MB)
	fmt.Printf("hello golang%b\n",MB)//转为二进制
	fmt.Printf("hello golang%o\n",MB)//转为八进制
	fmt.Printf("hello golang%x\n",MB)//转为十六进制
	fmt.Printf("%T\n",MB)//T查看类型
}
