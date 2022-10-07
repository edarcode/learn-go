# Map

La estructura de datos "map" es más eficiente que un array o slice, ya que usan concurrencia para ejecutar las operaciones.

```
edarcode := map[string]string{
  "name":    "edarcode",
  "age":     "26",
  "country": "colombia",
}
```

o puede usar la func make

```
temperature := make(map[string]int)
temperature["earth"] = 15
temperature["mars"] = -65
```

Vamos a recorrerlo e imprimir su clave y valor

```
for key, value := range temperature {
  fmt.Println(key, value)
}
```

Si intentamos imprimir un key que no existe, tomará el defult que tanga el tipo de dato alojado. En un map con valores enteros veremos que toma 0 como value.

```
fmt.Println(temperature["xd"]) // 0
```

Para saber si existe o no un valor en el map, se return 2 variables cuando accedemos al mismo:

- **value :** el valor que contenga el map.
- **ok :** un bool que indica si existe o no el valor en el map.

```
value, ok := temperature["xd"]
fmt.Println(value, ok) // 0 false
```
