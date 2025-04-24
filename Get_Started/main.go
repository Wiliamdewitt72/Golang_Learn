// Generalmente el archivo se debe llamar main, esto es parte del formato handler de Go, este formato es para que la aplicación sea modular.
package main

// -------------------------------------------------------------------------

import (
	"fmt"
	"reflect"
) // Paquete para imprimir en consola. Este paquete es parte de la librería estándar de Go. Se usa para imprimir en consola, al igual que el paquete fmt de C.

/* Así se importan varios paquetes a la vez.
import (
	"libreria"
	"libreria2"
)
*/

// Paquete para trabajar con reflexión, es decir, para obtener información sobre el tipo de dato de una variable en tiempo de ejecución.

// ------------------------------------------------------------------------

// Las constantes se deben declarar fuera de los métodos.
const PI = 3.14 // Cuando el nombre de la constante empieza con mayúscula, es porque es exportable, es decir, puede ser utilizada fuera del paquete donde fue declarada. Es de caracter público.

// ------------------------------------------------------------------------
// Todo en go debe ser ejecutado dentro de una función.
// La función principal del programa se debe llamar main y
func main() { // La llave siempre debe ir en la misma línea.
	fmt.Println("Hello World desde Go!")

	// ------------------------------------------------------------------------
	// variables y constantes

	// declaración por inferencia, aquí se declara directamente el tipo de variable.
	var nombre string = "Jaime"

	// declaración rápida o corta, al estilo matemático. Esto es lo más común.
	nombre2 := "Juan"

	fmt.Println(nombre, nombre2)

	// Cuando trabajamos con Printf podemos usar comodines, los cuales son de la forma: %s, %d, etc.
	// %s es para cadenas de texto, %d es para enteros, %f es para flotantes, %t es para booleanos, %v es para cualquier tipo de dato.
	fmt.Printf("Hola %s, bienvenido a Go\n", nombre)
	// Las formas de imprimir en Go son como las de C, pero con la diferencia que en Go no se necesita el & para imprimir variables.

	// -------------------------------------------------------------------------
	// Tipos de datos
	// String
	var string1 string = "Hola"
	println(string1)
	textoGrande := `Hola mundo, este es un texto muy largo, que no cabe en una sola línea, por lo que se debe dividir en varias líneas. Esto es para que el código sea más legible y fácil de entender.`
	// Esto es un string multilinea, se usa el backtick ` para esto. Para textos grande se usa `.`
	println(textoGrande)

	// Boleanos
	var estado bool = true
	println(estado) // La diferencia entre fmt.Println y println es que la primera es para entornos de producción y la segunda es para entornos de desarrollo. Acá no se pueden usar comodines.

	// Numeros
	var flotante32 float32 = 3.14            // 32 bits
	var flotante64 float64 = 3.14            // 64 bits
	var entero int = 3                       // 32 bits en sistemas de 32 bits y 64 bits en sistemas de 64 bits.
	var entero8 int8 = 4                     // 8 bits de -128 a 127.
	var entero16 int16 = 65                  // 16 bits de -32768 a 32767. [-2^15, 2^15 -1]
	var entero32 int32 = 2147483647          // 32 bits de -2147483648 a 2147483647.	[-2^31, 2^31 -1]
	var entero64 int64 = 9223372036854775807 // 64 bits de -9223372036854775808 a 9223372036854775807.	[-2^63, 2^63 -1]
	fmt.Println(flotante32, flotante64, entero, entero8, entero16, entero32, entero64)
	// El tipo de dato entero es el más usado, ya que es el más rápido y eficiente. El tipo de dato float64 es el más usado para números decimales, ya que es el más preciso. El tipo de dato float32 es el más rápido, pero menos preciso.

	// Positivos
	// Los tipos enteros unit sólo pueden ser positivos. Por lo que no tienen signo y su rango es de 0 a 2^n -1, donde n es el número de bits del tipo de dato.
	var entero_unit8 uint8 = 255                    // 8 bits de 0 a 255. [0, 2^8 -1]
	var entero_unit16 uint16 = 65535                // 16 bits de 0 a 65535. [0, 2^16 -1]
	var entero_unit32 uint32 = 4294967295           // 32 bits de 0 a 4294967295. [0, 2^32 -1]
	var entero_unit64 uint64 = 18446744073709551615 // 64 bits de 0 a 18446744073709551615. [0, 2^64 -1]
	fmt.Println(entero_unit8, entero_unit16, entero_unit32, entero_unit64)

	// -----------------------------------------------------------------------
	// reflect y TypeOf

	var cadena1 string = "Hola"
	fmt.Println(reflect.TypeOf(cadena1)) // Esto imprime el tipo de dato de la variable cadena1.

}
