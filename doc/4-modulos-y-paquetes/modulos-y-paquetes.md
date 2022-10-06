## Modulos

Para definir un modulo debe ubicarse en el folder que desea establecer como modulo y tirar por consola "go mod init" o crear un go.mod dentro del folder.

## Paquetes

Cree una carpeta, ej "utils", dentro crear archivos que tengan un nombre de paquete definido.

![package](./assets/package.png)

si desea consumir los recursos del paquete debe nombrar el mismo con mayus ej "AddNumbers". Todos los archivos dentro de utils comparten el mismo scope.
