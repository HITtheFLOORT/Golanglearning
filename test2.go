package main

import (
	"awesomeProject/util"
	"fmt"
	"os"
	"runtime"
)

func printinfo(){
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
	fmt.Println(util.Gettime())
}
var a = "G"
func n() { print(a) }
func m() {
	a := "O"
	print(a)
}
func main() {
	printinfo()
	n()
	m()
	n()
}