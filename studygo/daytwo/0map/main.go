package main

import "fmt"
var a=func(x int)int{
	return 1//函数内部不能声明有名字的函数
}
//闭包 python 的装饰器
func f1(f func()){
	f()
}
func f2(x,y int)  {
	fmt.Println(x+y)
}
func f3(f2 func(n ,m int),n,m int)func()int{
	tmp:=func() int {
		f2(n,m)
		return m+n;
	}
	return tmp
}
//func f3(f2 func(n ,m int),n,m int)func(){
//	tmp:=func() {
//		f2(n,m)
//	}
//	return tmp
//}
/*
闭包是匿名函数与匿名函数所引用环境的组合。匿名函数有动态创建的特性，
该特性使得匿名函数不用通过参数传递的方式，就可以直接引用外部的变量。
这就类似于常规函数直接使用全局变量一样，个人理解为：匿名函数和它引用的变量以及环境，
类似常规函数引用全局变量处于一个包的环境。
 */
func adder()func(a int)int{
	x:=100
	return func(a int) int {
		return a+x
	}
}
func main() {
	m1:=make(map[string]int,10)
	m2:='a'
	fmt.Printf( "%T%T\n",m1,m2)

	//立即函数 只执行一次
	func(a,b int ){
		fmt.Println(a+b)
	}(100,200)

	f3(f2,10,20)

	ret:=adder()
	ret2:=ret(100)
	fmt.Println(ret2)
}
