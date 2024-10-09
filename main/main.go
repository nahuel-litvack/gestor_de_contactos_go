package main

import (
	"fmt"

	"github.com/nahuel-litvack/repositorio01/utilidades"
)

/** A mejorar:
	. Permitir que el flujo siempre termine invocando a la funcion MenuDeInicio() para que siempre el programa este funcionando.
	. Fijarse si es necesario pasar el validador de opciones de la funcion MenuDeIncio() a la funcion main().
	. Personalizar la opcion 4 para que devuelva el nombre del contacto.

**/

func main() {
	opcion := utilidades.MenuDeInicio()

	if opcion == 1 {
		utilidades.ValidacionOpcionUno()
		resultado := utilidades.AgregarContacto(utilidades.Telefono, utilidades.Nombre, utilidades.Correo)
		fmt.Println(resultado)
		fmt.Print("\n\n")
		//utilidades.MenuDeInicio()
	}
	if opcion == 2 {
		utilidades.ValidacionOpcionDos()
		resultado := utilidades.EliminarContacto(utilidades.Telefono)
		fmt.Println(resultado)
		fmt.Print("\n\n")
	}
	if opcion == 3 {
		utilidades.ValidacionOpcionTres()
		resultado := utilidades.BuscarContacto(utilidades.Telefono)
		fmt.Println(resultado)
		fmt.Print("\n\n")
	}
	if opcion == 4 {
		fmt.Println("Lista de contactos:")
		for _, key := range utilidades.Dict.Keys() {
			valor := utilidades.Dict.Get(key)
			fmt.Printf("Telefono: %d, Nombre: %s, Correo: %s\n", key, valor[0], valor[1])
		}
		fmt.Print("\n\n")
	}
}
