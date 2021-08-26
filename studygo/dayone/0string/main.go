package main

import (
	"fmt"
	"regexp"
	"strings"
)
func Chinesejudge(s string)bool{
	redux:=regexp.MustCompile("^[\u4e00-\u9fa5]$")
	if redux.MatchString(s){
		return true
	}else{
		return false
	}
}
func main() {
	str:="D:\\workspace\\文件夹"
	str2:=`D:\\workspace`
	str=str+str2
	strs:=strings.Split(str,"\\")
	strs2:=strings.Split(str2,"\\")
	strings.Contains(str,"\\")
	strings.Index(str,"\\")
	fmt.Println(strs)
	fmt.Println(strs2)
	for index,value := range strs{
		fmt.Println(strs[index]+value)
	}
	//rune 处理ascii码不能标识的字符
	s1:="abc"
	s2:=[]rune(s1)//[]byte 切片
	s2[0]='c'
	fmt.Printf("%T\n",s2[0])// rune int32 byte uint8
	fmt.Println(string(s2))
	for _,value:=range "nihao你好"{
		if Chinesejudge(string(value)){
			fmt.Print(string(value))
		}
	}
}
