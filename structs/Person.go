package structs

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (person Person) Greet() string {
	var greet string = fmt.Sprintf("Hola soy %s y tengo %d años", person.Name, person.Age)
	return greet
}
