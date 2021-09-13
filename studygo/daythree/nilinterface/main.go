package main

import "fmt"

type Student struct {
	Name string
}
func main() {
	Parames:=make([]interface{},3)
	Parames[0]=88
	Parames[1]="abc"
	Parames[2]=Student{
		"Liming",
	}
	for index,val :=range Parames{
		if _,ok:=val.(int);ok{
			fmt.Printf("Patames[%d]type is int\n",index)
		}else
		if _,ok:=val.(string);ok{
			fmt.Printf("Patames[%d]type is string\n",index)
		}else
		if _,ok:=val.(Student);ok{
			fmt.Printf("Patames[%d]type is Student\n",index)
		}else{
			fmt.Printf("Patames[%d]type is others\n",index)
		}
	}
	for index, v := range Parames {
		switch value := v.(type) {
		case int:
			fmt.Printf("Params[%d] 是int类型, 值：%d \n", index, value)
		case string:
			fmt.Printf("Params[%d] 是字符串类型, 值：%s\n", index, value)
		case Student:
			fmt.Printf("Params[%d] 是Person类型, 值：%s\n", index, value)
		default:
			fmt.Printf("list[%d] 未知类型\n", index)
		}
	}
}
