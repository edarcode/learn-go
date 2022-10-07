# Funciones de struct

Podemos asociar funciones a cada struct de tal forma que cada instancia de la struct puedad consumirla. Creemos una struct.

```
package structs

import "fmt"

type Person struct {
	Name string
	Age  int
}
```

Creemos una función asociada a la estructura "Person". La idea es que cada instancia de la misma pueda saludar y además mencione su edad.

Para asociar **struct** a la **func** debe agregar luego de func "(person Person)".

```
func (person Person) Greet() {

}
```

- **Person :** Hace ref a la **struct** que desea asociar con la función.
- **person :** Hace ref a cada instancia que se haga de la **struct**, en este caso "Person".

El resto es carpintería. Una función normal y corriente, si desea acceder a datos de la **struct** use la sintaxis de punto.

```
func (person Person) Greet() string {
	var greet string = fmt.Sprintf("Hola soy %s y tengo %d años", person.Name, person.Age)
	return greet
}
```

Para consumir sería:

```
package main

import (
	"fmt"
	"learn-go/structs"
)

func main() {
	edarcode := structs.Person{Name: "edarcode", Age: 26}
	lore := structs.Person{Name: "lorena", Age: 40}

	fmt.Println(edarcode.Greet()) // Hola soy edarcode y tengo 26 años
	fmt.Println(lore.Greet())     // Hola soy lorena y tengo 40 años
}
```

Tener en cuenta que la función asociada en realidad está asociada a una copia de la struct "Person". Es posible cambiar el comportamiento usando punteros.
