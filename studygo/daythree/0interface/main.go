package main

import "fmt"

type cat struct {
	name string
}
func newCat(name string) *cat{
	return &cat{
		name:name,
	}

}
type animal interface {
	speak()
	feed(cat2 cat)
}
func (c cat)feed(){
	fmt.Printf("%s is feeded",c.name)
}
func main() {
	var cat1=cat{name:"miaomiao"}
	cat1.feed()
}
