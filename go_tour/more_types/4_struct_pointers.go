// package main

// import "fmt"

// type Vertex struct {
// 	X, Y int
// }

// func main() {
// 	v := Vertex{1, 2}
// 	p := &v
// 	p.X = 1e3

//		fmt.Println(v)
//		fmt.Println(p)
//	}
package main

import "fmt"

type Person struct {
	Name     string
	LastName string
}

func (p *Person) Talk() {
	var text string = fmt.Sprintf("Hi, my name is %s and my lastname %s", p.Name, p.LastName)
	fmt.Println(text)
}

type Android struct {
	Person Person
	Model  string
}

type Android2 struct {
	Person
	Model string
}

func main() {
	a := new(Android)
	b := new(Android2)
	var c = Android2{
		Model: "Model",
		Person: Person{
			Name: "Name",
		},
	}

	a.Person.Name = "Timur"
	a.Person.LastName = "Teremchuk"

	b.Name = "Timur"
	b.LastName = "Apripov"

	a.Person.Talk()
	b.Talk()
	c.Talk()
}
