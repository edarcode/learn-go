# Primer contacto con go

- Cualquier app en go si o si debe empezar con un package main y una función main. Cree un "fileName.go" y ejecutelo por consola con "go run fileName.go"

```
package main

func main() {

}
```

- Para imprimir por consola podemos usar una función que ya está lista para consumir en el paquete fmt.

```
package main

import "fmt"

func main() {
	greet := "Hola"
	name := "edarcode"

	fmt.Println(greet, name, "👋")
}
```
