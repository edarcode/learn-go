package structs

type users struct {
	name string
	age  int
}

func CreateUser(name string, age int) users {
	newUser := users{name: name, age: age}
	return newUser
}

func (user *users) GetName() string {
	return user.name
}
func (user users) GetAge() int {
	return user.age
}
