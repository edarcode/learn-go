package structs

type Person struct {
	Name string
	Age  int
}

func (person *Person) SetName(newName string) {
	person.Name = newName
}
