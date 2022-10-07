# Struct

Permiten escribir nuestros propios tipos de estructuras, como armar una plantilla y luego partir o instanciar muchas en base a la misma.

```
package main

import "fmt"

func main() {
	edarcode := person{name: "edarcode", age: 26}
	lorena := person{name: "lorena"}
	marian := person{name: "marian", age: 35}

	fmt.Println(edarcode) // {edarcode 26}
	fmt.Println(lorena)   // {lorena 0}
	fmt.Println(marian)   // {marian 35}
}

type person struct {
	name string
	age  int
}
```

Con la notación de punto puede acceder a cada uno de los campos si lo desea.

```
fmt.Println(edarcode.name, edarcode.age)
```
