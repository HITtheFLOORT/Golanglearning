package main

import "fmt"

type cat struct {
	animal
}
type animal struct {
	Name string
}
func newAnimal(name string) *cat{
	return &cat{
		animal{
			Name:name,
		},
	}

}
type animalfunc interface {
	speak()
	feed(cat2 cat)
}
func (c cat)feed() {
	fmt.Printf("%s is feeded\n",c.Name)
}
func main() {

	var cat1=cat{animal{Name:"miaomiao"}}
	var cat2=new(cat)
	cat2.Name="bibei"
	cat1.feed()
	cat2.feed()
}
