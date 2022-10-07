# Punteros

Los punteros son variables que apuntar a otras variables y tienen el poder de cambiarlas. Un puntero almacena la ref o dirección en memoría de la variable a la que esté apuntando.

```
package main

import "fmt"

func main() {
	name := "edarcode"
	pointer := &edarcode

	fmt.Println(name) // edarcode
	fmt.Println(pointer)  // 0xc000014250
}
```

Para acceder al contenido de "name" desde el puntero, usaremos la sintaxis de \*

```
fmt.Println(*pointer) // edarcode
```

Cambiemos de valor a "name" desde el puntero. No olvidar el \*

```
*pointer = "lorena"
fmt.Println(*pointer, name) // lorena lorena
```
