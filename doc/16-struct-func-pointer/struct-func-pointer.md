# struct / func / pointer

Para asociar una **func** a la **struct** original basta con agregar un \* precedido del nombre de la misma. Creemos una función set para setear el nombre de cada persona.

```
package structs

type Person struct {
Name string
Age int
}

func (person Person) SetName(newName string) {
person.Name = newName
}
```

Seguramente le ha salido un warnning indicando que esa operación es inefectiva, dado que no estamos usando punteros no va a cambiar ninguna instancia original (\*). Consumamos la **struct**.

```
package main

import (
"fmt"
"learn-go/structs"
)

func main() {
edarcode := structs.Person{Name: "edarcode", Age: 26}
}
```

Ahora, imprima la instancia, cambie el nombre y vuelva a imprimir la misma.

```
fmt.Println(edarcode) // {edarcode 26}
edarcode.SetName("edar")
fmt.Println(edarcode) // {edarcode 26}
```

Ahora, use un puntero e indique que apunte a la **struct** "Person" y asociamos la función a dicho puntero.

```
package structs

type Person struct {
Name string
Age int
}

func (person *Person) SetName(newName string) {
person.Name = newName
}
```

Repetimos el proceso para consumir.

```
package main

import (
"fmt"
"learn-go/structs"
)

func main() {
edarcode := structs.Person{Name: "edarcode", Age: 26}
fmt.Println(edarcode) // {edarcode 26}
edarcode.SetName("edar")
fmt.Println(edarcode) // {edar 26}
}
```

Y listo xD, como vemos, ahora si cambiarán todas las instancias originales que se hagan de la **struct**.
