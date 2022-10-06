# Err

Vamos a tratar de convertir un string a un número y manejar el posible err

```
package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	value, err := strconv.Atoi("edar")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(value)
	}
}
```

Note que la función Atoi return 2 valores: **"err"** es cargado en caso tal exista una excepción o err **(err != nil)**. Por consola saldrá esto:

```
2022/10/06 05:27:50 strconv.Atoi: parsing "edar": invalid syntax
exit status 1
```
