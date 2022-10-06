# fmt

## fmt.Println()

Imprime por consola la cantidad de parámetros enviadas como parámetro. Siempre se agregan espacios entre parámetros y se agrega una nueva línea.

```
fmt.Println("edarcode", 26)
```

## fmt.Printf()

Imprime por consola un string, permite incrustar variables dentro del mismo indicando con un % y el tipo de dato que desea incrustar. El orden importa.

- **%s** indica un string
- **%d** indica un entero
- **%v** indica cualquier tipo
- **\n** indica un salto de linea

```
fmt.Printf("%s tiene %d años\n", "edarcode", 26)
```

Esta función también puede decir el tipo de dato de lo que le enviemos indicando un %T.

```
fmt.Printf("%T\n%T\n", "edarcode", 26)
```

## fmt.Sprintf()

Sirve para guardar como string el parámetro que se envía al fmt.Printf() sin imprimir por consola.

```
var msg string = fmt.Sprintf("%s tiene %d años", "edarcode", 26)

fmt.Println(msg)
```
