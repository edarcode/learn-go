# Channel

Permite pasar información de una lado al otro, incluso entre **goroutines**, además, No es necesario indicar que las espere explícitamente.

Creemos un channel

- **chan:** Tipo de dato de una channel
- **string:** Tipo de dato que pasará por el channel
- **1:** Cantidad de datos a pasar por el channel

```
channel := make(chan string, 1)
```

Creemos una **func** que posteriormente ejecutaremos concurrentemente. Se debe especificar que va a recibir un **channel** indicando el tipo del mismo y el tipo de información que enviaran por el **channel**.

- **<-** Indica entrada del channel

```
func getUpperText(text string, channel chan<- string) {
  upperText := strings.ToUpper(text)
  channel <- upperText
}
```

Para consumir:

- **->** Indica salida del channel

```
func main() {
  channel := make(chan string, 1)

  fmt.Println("HOLA") // HOLA
  go getUpperText("lore", channel)
  fmt.Println(<-channel) // LORE
}
```

Tener en cuenta que es un poco menos ineficiente que los **wg**. Aunque deja muy limpio el code para gestionar las **goroutines**.
