package main

import (
	"errors"
	"fmt"
)

// constante
const ProbandoPi float64 = 3.14 //se puede escribir dentro de una función

// No sé como funciona esto, solo coloque el ejemplo para ver el uso del error.
func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("error al realizar la operación: división por cero")
	}
	return a / b, nil // nil significa que no hay error
}



func main() {

	//	https://go.dev/ (Sitio oficial del lenguaje)
	/*
		Esta es la forma de colocar un comentario de varias lineas.
	*/

	//crear una variable
	var golang string = "Go"

/*	//datos enteros
	var numero int = 10
	var entero8 int8 = -128
	var enteroSinSigno uint = 4294967295
	var numeroFlotante float32 = 3.14159
	var numeroFlotante64 float64 = 1.234567890123456789
	var caracter rune = 'a'
	var complejo complex64 = 3 + 4i
	var complejo128 complex128 = 1.234 + 5.678i
	var verdadero bool = true
	var mensaje string = "Hola mundo"
	var puntero *int = nil
*/
	fmt.Println("Hola", golang)
	fmt.Println(ProbandoPi)
	
	fmt.Println(dividir(25,5))
}



/*
	Tipos de datos:
	Numericos:
	    Enteros:
			int: Entero de tamaño de palabra, depende de la arquitectura del sistema.
				Ejemplo: var numero int = 10
			int8, int16, int32, int64: Enteros con tamaños específicos.
				Ejemplo: var entero8 int8 = -128
			uint, uint8, uint16, uint32, uint64: Enteros sin signo con tamaños específicos.
				Ejemplo: var enteroSinSigno uint = 4294967295
			rune: Alias para int32, utilizado para representar caracteres Unicode.
				Ejemplo: var caracter rune = 'a'

    	Flotantes:
	        float32: Número de punto flotante de precisión simple.
				Ejemplo: var numeroFlotante float32 = 3.14159
	        float64: Número de punto flotante de precisión doble.
				Ejemplo: var numeroFlotante64 float64 = 1.234567890123456789

    	Complejos:
       	 	complex64: Número complejo con partes real e imaginaria de tipo float32.
				Ejemplo: var complejo complex64 = 3 + 4i
        	complex128: Número complejo con partes real e imaginaria de tipo float64.
				Ejemplo: var complejo128 complex128 = 1.234 + 5.678i
	Booleanos:
		bool: Puede tener dos valores: true o false.
		Ejemplo: var verdadero bool = true

	Cadenas (string):
		string: Representa una secuencia de caracteres Unicode.
		Ejemplo: var mensaje string = "Hola mundo"

	Especiales:
	    nil: Representa la ausencia de un valor, se puede usar para inicializar variables de referencia.
		Ejemplo: var puntero *int = nil

	Predefinidos:
		error: Tipo de error predefinido que se usa para representar errores.
		Ejemplo: return errors.New("error al realizar la operación")
*/
