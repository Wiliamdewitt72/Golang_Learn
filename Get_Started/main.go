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

	// declaración rápida o corta, al estilo matemático. Esto es lo más común. Para hacerlo no usar la palabra 'var', sino te arrojará un error.
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

	// -----------------------------------------------------------------------
	// puntero
	// En informática, un puntero es una variable que almacena la dirección de memoria de otra variable. En esencia, un puntero "apunta" a una ubicación específica en la memoria donde se guarda el valor de otra variable

	color := "rojo"
	fmt.Println(color, " esta en: ", &color) // Imprime la dirección de memoria de donde está la variable color. O se accede  la dirección de la memoria.

	// ----------------------------------------------------------------------
	// Condicionales

	/* Valores condicionales
	- == igual
	- != diferente
	- > mayor que
	- < menor que
	- >= mayor o igual que
	- <= menor o igual que
	Logicos
	- && y
	- || o
	- ! no
	*/

	edad := 11

	//  if
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}
	// También se pueden usar prentesis para agrupar condiciones, pero no es necesario.
	if (edad == 15) || (edad == 16) {
		fmt.Println("Eres un adolescente")
	}

	// if anidado
	tono := "rojo" // si la variable no coincide en el if anidado, el compilador nos advierte que no se está usando la variable.

	if tono == "rojo" {
		fmt.Println("El color es rojo")
	} else if tono == "verde" {
		fmt.Println("El color es verde")
	} else if tono == "azul" {
		fmt.Println("El color es azul")
	} else {
		fmt.Println("El color no es rojo, verde o azul")
	}

	// Declarar variables en una condicion
	if year := 18; year >= 18 { // Se puede declarar una variable dentro de un if, pero sólo se puede usar dentro del condicional.
		fmt.Println("Eres mayor de edad, tienes ", year, " años")
	}
	// fmt.Println("El año es ", year) --> Esto no se puede usar, ya que la variable year no está declarada fuera del if.

	// ------------------------------
	// switch case

	crayola := "azul"

	switch crayola {
	case "rojo":
		fmt.Println("El color es rojo")
		// break // El break es opcional, ya que el switch case termina automáticamente al encontrar un caso que coincida. No es necesario ponerlo.
	case "verde":
		fmt.Println("El color es verde")

	case "azul":
		fmt.Println("El color es azul")

	default:
		fmt.Println("El color no es rojo, verde o azul")
	}

	// -------------------------------------------------------------------------
	// bucles o ciclos

	// for
	// El for es el único bucle que existe en Go. No hay while ni do while. Se usa el for para todo.
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++ // i = i + 1
	}
	for i2 := 1; i2 <= 10; i2 += 2 { // Se puede declarar la variable dentro del for, pero sólo se puede usar dentro del for.
		fmt.Println(i2)
	}
	// fmt.Println("Fin del bucle for", i2) -> Esto no va a funcionar porque sólo está declarada dentro del bucle.
	for i3 := 100; i3 > -10; i3 -= 7 { // Así se puede decrementar o aumentar la variable.
		fmt.Println(i3)
	}

	// Continue y break

	for i4 := 0; i4 < 10; i4++ {
		if i4 == 5 {
			continue // Esto hace que se salte la iteración actual y pase a la siguiente.
		}
		if i4 == 8 {
			break // Esto hace que se salga del bucle.
		}
		fmt.Println(i4)
	}

	fmt.Println(3 % 4) // El operador % es el operador módulo, que devuelve el residuo de la división entre dos números. En este caso, 3 dividido por 4 es 0 y el residuo es 3.

	// -------------------------------------------------------------------------
	// Arreglos y slices
	// Es un espacio de memoria que te permite almacenar varios datos. El arreglo tiene una capacidad determinada mientras que el slice es dinámico y puede crecer o decrecer.

	// Arreglo
	var paises [4]string // Se declara un arreglo de 4 elementos de tipo string.
	paises[0] = "Colombia"
	paises[1] = "Perú"
	paises[2] = "Chile"
	paises[3] = "Argentina"

	fmt.Println(paises)                 // Imprime el arreglo completo.
	fmt.Println(paises[0])              // Imprime el primer elemento del arreglo.
	fmt.Println(len(paises))            // Imprime la longitud del arreglo.
	fmt.Println(cap(paises))            // Imprime la capacidad del arreglo. La capacidad es la cantidad de elementos que puede almacenar el arreglo. En este caso, 4.
	fmt.Println(reflect.TypeOf(paises)) // Imprime el tipo de dato del arreglo. En este caso, es un arreglo de 4 elementos de tipo string.

	// Slice
	var paises2 = make([]string, 4) // Se declara un slice de 4 elementos de tipo string. El slice se declara con la función make.
	paises2[0] = "Colombia"
	paises2[1] = "Perú"
	paises2[2] = "Chile"
	paises2[3] = "Argentina"

	fmt.Println(paises2)                 // Imprime el arreglo completo.
	fmt.Println(paises2[0])              // Imprime el primer elemento del arreglo.
	fmt.Println(len(paises2))            // Imprime la longitud del arreglo.
	fmt.Println(cap(paises2))            // Imprime la capacidad del arreglo. La capacidad es la cantidad de elementos que puede almacenar el arreglo. En este caso, 4.
	fmt.Println(reflect.TypeOf(paises2)) // Imprime el tipo de dato del arreglo. En este caso, es un arreglo de 4 elementos de tipo string.

	// Agregar elementos a un slice
	paises2 = append(paises2, "Venezuela") // Se agrega un elemento al slice. La función append agrega un elemento al final del slice. Si el slice no tiene suficiente capacidad, se crea un nuevo slice con más capacidad y se copian los elementos del slice original.
	fmt.Println(paises2)                   // Imprime el arreglo completo.

	// Quitar elementos de un slice
	paises2 = append(paises[:4], paises2[4+1:]...) // Se quita el valor de la posición 4. Se puede quitar el valor de cualquier posición solamente hay que cambiar el índice.
	fmt.Println(paises2)

	// -------------------------------------------------------------------------
	// Mapas. Formato de clave-valor. Como un diccionario, la clave es la palabra y el valor es la definición. Pueden ser de diferentes tipos de datos.
	// Forma de hacer un mapa:
	continentes := make(map[string]int)
	continentes["America"] = 12000000
	continentes["Europa"] = 20000000
	continentes["Asia"] = 30000000
	continentes["Africa"] = 40000000
	continentes["Oceania"] = 50000000
	fmt.Println(continentes)

	// Otra forma de declarar un mapa es la siguiente:
	aeropuertos := map[string]string{
		"Colombia":  "Bogotá",
		"Perú":      "Lima",
		"Chile":     "Santiago",
		"Argentina": "Buenos Aires",
		"Venezuela": "Caracas",
	}
	fmt.Println(aeropuertos) // Imprime el mapa completo.

	// Verificar si existe un valor en el mapa
	ciudad, existe := aeropuertos["Colombia"]
	//  Ciudad(string) es el valor que se obtiene del mapa, existe(bool) es un booleano que indica si el valor existe o no.
	if existe {
		fmt.Println("La ciudad donde está el aeropuerto es ", ciudad)
	} else {
		fmt.Println("No existe aeropuerto en la ciudad")
	}

	// Eliminar un valor de un mapa
	delete(aeropuertos, "Colombia") // Se elimina el valor de la clave "Colombia".
	fmt.Println(aeropuertos)

	// Recorrer un mapa con for
	for id, valor := range aeropuertos {
		fmt.Println("El aeropuerto de ", id, " es ", valor)

	}

	// Un ejemplo de un mapa usado en API's
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Todo bien",
	}
	fmt.Println(respuesta) // Imprime el mapa completo.
}
