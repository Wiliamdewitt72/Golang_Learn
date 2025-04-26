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

	// Está es la goroutine, hace que la función se ejecute en paralelo.
	go func() {
		miCanal <- parametro("jitomate") // Se envía el valor de la función al canal.
	}()

	fmt.Println(<-miCanal) // Si se quiere imprimir el valor del canal imprimimos el puntero de la función. Esto es porque la goroutine se ejecuta en paralelo y no se puede imprimir directamente el valor de la función.

	//  Función recursiva.
	recursiva(0)

	// Manejo de errores.
	i := 0
	if i < 10 {
		fmt.Println("\nLa consistencia matemática se ha conservado.")
		time.Sleep(20 * time.Second) // Espera 2 segundos
		defer falla()
	} else {
		panic("Error: La consistencia matemática se ha roto.")
	}

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

// Gorutinas son funciones que se ejecutan en paralelo. Se declaran con la palabra reservada "go" antes de la función. Se ejecutan en un hilo diferente al hilo principal. Esto permite que el programa no se detenga mientras se ejecuta la función.

// Los canales son una forma de comunicación entre gorutinas. Permiten enviar y recibir datos entre gorutinas. Se declaran con la palabra reservada "chan" y se pueden usar para enviar y recibir datos entre gorutinas.

// ---------------------------------------------------------------
// Recursividad
//  Significa que vamos a llamar a la misma función dentro de la misma función.

func recursiva(valor int) {
	dato := valor + 2
	fmt.Println(dato)
	if dato < 10 {
		recursiva(dato) // Llamamos a la misma función dentro de la misma función. Esto es la recursividad.
	} else {
		fmt.Println("Ahora el valor es mayor a 10")
	}

}

// ---------------------------------------------------------------
//  Defer y panic (manejo de errores)

func falla() {
	fmt.Println("Comprobación del sistema terminada.")
}

// defer se ejecuta al final de la función.
// panic se usa para manejar errores. Se usa para detener la ejecución del programa y mostrar un mensaje de error.
