# Slice

Slice es similar a un array, sin embargo podemos mutar su tamaño. Definamos un slice de strings.

```
var persons = []string{"edarcode", "lore", "marian"}
```

o usando la función make

```
make([]string, "edarcode", "lore", "marian")
```

Podemos agregar nuevos elementos a un slice con la función append. Tener en cuenta que append crea un copia en base al primer parámetro enviado y agrega todo lo demás.

```
persons = append(persons, "jose", "pablo")
```

Podemos indicar los valores que debe halar. Imprimamos desde la posición 4 hasta el final del array.

```
fmt.Println(persons[4:])
```

Imprimamos lo contrario, vamos desde el inicio hasta la posición 4 del array.

```
fmt.Println(persons[:4])
```

Visualmente hemos jugado con "pablo", pasó de ser el único impreso al único no impreso. Puede aprovechar esto para crear un nuevo array sin un elemento en concreto.

Combinemos ambas sintaxis. Indiquemos que imprima un tramo concreto del array. Vamos desde la posición 2 a la 3.

```
fmt.Println(persons[2:3]) // Si especifica el fin "3", No será incluido.
```

También podemos crear nuevos slice a partir de 2 previamente hechos con la sintaxis de los "...".

```
package main

import "fmt"

func main() {
var students = []string{"edarcode", "lore", "marian"}
var teachers = []string{"jose", "pablo"}

persons := append(students, teachers...)

fmt.Println(persons)
}

```
