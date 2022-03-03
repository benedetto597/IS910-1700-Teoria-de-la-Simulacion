/*
	@author benedetto597
	@date 2020-04-21
	@version 1.0.0
	@description Paquete "welcome" con funciones de saludo

*/
package welcome

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
