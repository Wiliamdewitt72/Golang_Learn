package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Una estructura es un objeto de datos que tiene diferentes campos. Es como una clase, pero sin métodos. Solamente se usa para agrupar datos relacionados.
	// Estructura de datos en Go. (Forma 1)
	// En general, se usa está forma paqra crear registros.
	estructura := Persona{
		Id:     1,
		Nombre: "Juan",
		Correo: "estudios@gmail.com",
		Edad:   25,
	}
	fmt.Println(estructura)                 // Imprime la estructura completa.
	fmt.Printf("%+v \n", estructura)        // Imprime la estructura completa con los nombres de los campos.
	fmt.Println(reflect.TypeOf(estructura)) // Imprime el tipo de la estructura.
	fmt.Println(estructura.Correo)          // Imprime el dato guardado en la estructura, en el campo Correo.

	// Estructura de datos en Go. (Forma 2)
	// En general, esta forma se usa para modificar registros.
	p := new(Persona)              // Se crea un puntero a la estructura.
	fmt.Println(reflect.TypeOf(p)) // Imprime el tipo de la estructura.

	p.Id = 2
	p.Correo = "correo@gmail.com"
	p.Nombre = "Pedro"
	p.Edad = 30

	fmt.Printf("%+v \n", p)
	//  ------------------------------------------------------------------------

}

// Lo más recomendable es usar el nombre de la estructura en mayúscula, para que sea exportable (es decir, publica) y se pueda usar en otros paquetes. Cuando es minuscula, es privada y solo se puede usar en el mismo paquete.

type Persona struct {
	Id     int
	Nombre string
	Correo string
	Edad   int
}
type Categoria struct {
	Id     int
	Nombre string
	Slug   string
	// También pueden haber campos privados si se ponen en minúscula.
}

type Producto struct {
	Id        int
	Nombre    string
	Slug      string
	Precio    int
	Categoria Categoria // Se puede usar una estructura dentro de otra estructura.
}
