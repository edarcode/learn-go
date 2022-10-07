# Switch

```
package main

import (
	"fmt"
)

func main() {
	num := -8
	switch {
	case num > 0:
		fmt.Println("==>")
	case num < 0:
		fmt.Println("<==")
	default:
		fmt.Println("0")

	}
}
```
