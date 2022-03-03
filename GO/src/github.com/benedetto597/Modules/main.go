/*
	@author benedetto597
	@date 2020-04-21
	@version 1.0.0
	@description Prueba de modulos/paquetes en GO

*/
package main

import (
	"fmt"

	"github.com/benedetto597/hello/welcome" // Importamos el paquete
)

func main() {
	fmt.Println(welcome.Hello("Benedetto"))

}
