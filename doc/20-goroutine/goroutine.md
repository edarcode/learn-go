# Goroutine

Permite ejecutar código de forma concurrente. Basta con poner **go** previamente. Sin embargo, al ejecutar este código solo verá un **"Hola"** en la consola.

```
package main

import "fmt"

func main() {
  fmt.Println("Hola")
  go fmt.Println("edarcode")
}
```

Sucede que la **func main** por defecto es concurrente y al finalizar,automáticamente muere, y las func concurrentes de la misma estarán en otro hilo, por eso No vemos **"edarcode"**.

Debe realizar varios pasos para manejar la concurrencia:

1. Instancie una variable WaitGroup usando el paquete sync

```
var wg sync.WaitGroup
```

2. Indica que va agregar una **"Goroutine"**

```
wg.Add(1)
```

3. Coloque la **func** que desea ejecutar concurrentemente y dentro, al final de la **func**, ejecute **"done"** o puede usar **defer** si así lo desea:

func anonima autoejecutable

```
go func() {
  defer wg.Done()
  fmt.Println("edarcode")
}()
```

Vista general:

```
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hola")

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("edarcode")
	}()

	wg.Wait()
}
```
