package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var s string
	reader:=bufio.NewReader(os.Stdin)
	s,_=reader.ReadString('\n')
	fmt.Println(s)
	f,err:=os.Open("README.md")
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	reader=bufio.NewReader(f)

	for{
		line,err:=reader.ReadString('\n')
		if err==io.EOF{
			fmt.Println("读完了")
			break
		}
		if err!=nil{
			fmt.Println(err.Error())
			break
		}
		fmt.Println(line)
	}
	//for{
	//	var tmp [128]byte
	//	n,err:=f.Read(tmp[:])
	//	if err!=nil{
	//		fmt.Println(err.Error())
	//	}
	//	fmt.Println(n)
	//	fmt.Println(string(tmp[:n]))
	//	if n<128{
	//		return
	//	}
	//}
}
