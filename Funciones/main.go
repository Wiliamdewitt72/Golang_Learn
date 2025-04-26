package main

import (
	"fmt"
	"time"
)

func main() {
	miFuncion()
	miFuncionConArgumentos(1, "Juan")

	fmt.Println(miFuncionConRetorno("Pedro"))

	// Esto es super usado en Go, el retorno de múltiples valores.
	nombre, apellido, _, edad := miFuncionConRetornoMultiple()
	// Así puedes obtener los multiples valores de retorno de una función.
	// Puedes usar el guion bajo (_) para ignorar un valor de retorno.

	fmt.Println("Mi nombre completo es:", nombre, apellido, "y mi edad es", edad)

	fmt.Println("La suma de 2 + 3 es:", suma(2, 3)) // Uso de una función anónima. Se escribe la variable y se llama como una función.
	anonima()

	//  A cualquier variable a la que le asignemos una función, se le puede llamar como una función. En este caso, se ejecuta la función con el argumento 2.
	tabla := tablas(2) // Se llama a la función clousure y se guarda en una variable.

	for i := 0; i < 10; i++ {
		fmt.Println(tabla()) // Se llama a la función clousure mediante la variable y se imprime el resultado.
	}

	// Goroutine simple.
	fmt.Println(parametro("mayonesa"))
	time.Sleep(2 * time.Second) // Espera 2 segundos
	fmt.Println(parametro("ketchup"))

	// Canal
	miCanal := make(chan string) // Se crea un canal de tipo string.
	//
	go func() {
		miCanal <- parametro("jitomate") // Se envía el valor de la función al canal.
	}()

	fmt.Println(<-miCanal) // Se recibe el valor del canal y se imprime.
}

// -----------------------------------------------------
// Funciones

// Simple (solo se ejecutan)
func miFuncion() {
	fmt.Println("Hola desde mi funcion")
}

// Funciones con argumentos
func miFuncionConArgumentos(numero int, nombre string) {
	fmt.Printf("Soy el número %v y mi nombre es: %v", numero, nombre)
	// En Println no se pueden usar comodines.
}

// Funciones con retorno
func miFuncionConRetorno(nombre string) string {
	return "Hola " + nombre
}

// Funciones con retorno de múltiples valores
func miFuncionConRetornoMultiple() (string, string, string, int) {
	return "Daniel", "Frias", "González", 22
}

// No siempre se retornan tipos de datos nativos del lenguaje. Puedes personalizar el retorno de la función.
// --------------------------------------------------------------

// Funciones anónimas (funciones sin nombre)
// Se declaran dentro de una variable y para poder referenciarla se usa el nombre de la variable.

var anonima = func() {
	fmt.Println("Hola desde una función anónima")
}

// Funciones anónimas con retorno y paso de argumentos
var suma = func(a, b int) int {
	return a + b
}

// ----------------------------------------------------------------

// Funciones clousure
// Son funciones que retornan una función. "Recuerdan" el entorno en el que fueron creadas y pueden acceder a las variables de ese entorno, incluso después de que la función externa haya terminado.

func tablas(multiplo int) func() int {
	numero := multiplo
	return func() int {
		numero++
		return numero * multiplo
	}
}

// ----------------------------------------------------------------
// Gorutinas y chanels

func parametro(parametro string) string {
	return "Este es el parametro: " + parametro
}
