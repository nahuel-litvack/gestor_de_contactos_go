package main

import (
	"fmt"

	"github.com/nahuel-litvack/repositorio01/utilidades"
)

// A mejorar:
// . Hacer que se puedan agregar espacios y no se puedan agregar numeros al ingresar el nombre de contacto.
// . Mejorar el aspecto visual de la consola.
// . Agregar una GUI (interfaz grafica).
// y asi terminamos el proyecto...

// ***** MENU DE OPCIONES *****

func main() {
	opcion := utilidades.MenuDeInicio() // se ejecuta el menu de inicio

	for opcion == 1 || opcion == 2 || opcion == 3 || opcion == 4 || opcion == 5 {
		// se ejecuta mientra la opcion elegida sea 1: agrega un contacto al diccionario
		for opcion == 1 {
			utilidades.ValidacionOpcionUno()
			resultado := utilidades.AgregarContacto(utilidades.Telefono, utilidades.Nombre, utilidades.Correo)
			fmt.Println(resultado)
			fmt.Print("\n\n")
			opcion = utilidades.MenuDeInicio() // se ejecuta el menu de inicio
		}

		// se ejecuta mientras la opcion elegida sea 2: elimina un contacto del diccionario
		for opcion == 2 {
			// Validacion de dominio que verifica que el diccionario no este vacio
			if utilidades.Dict.Size() == 0 {
				fmt.Println("No hay ningun contacto en la lista.")
				opcion = utilidades.MenuDeInicio()
			} else {
				utilidades.ValidacionOpcionDos()
				resultado := utilidades.EliminarContacto(utilidades.Telefono)
				fmt.Println(resultado)
				fmt.Print("\n\n")
				opcion = utilidades.MenuDeInicio() // se ejecuta el menu de inicio
			}
		}

		// se ejecuta mientra la opcion elegida sea 3: busca un contacto en el diccionario
		for opcion == 3 {
			// Validacion de dominio que verifica que el diccionario no este vacio
			if utilidades.Dict.Size() == 0 {
				fmt.Println("No hay ningun contacto en la lista.")
				opcion = utilidades.MenuDeInicio()
			} else {
				utilidades.ValidacionOpcionTres()
				resultado := utilidades.BuscarContacto()
				fmt.Println(resultado)
				fmt.Print("\n\n")
				opcion = utilidades.MenuDeInicio() // se ejecuta el menu de inicio
			}
		}

		// se ejecuta mientra la opcion elegida sea 4: devuelve la lista de contactos en el diccionario
		for opcion == 4 {
			// Validacion de dominio que verifica que el diccionario no este vacio
			if utilidades.Dict.Size() == 0 {
				fmt.Println("No hay ningun contacto en la lista.")
			} else {
				fmt.Println("Lista de contactos:")
				for _, key := range utilidades.Dict.Keys() {
					valor := utilidades.Dict.Get(key)
					fmt.Printf("Nombre: %s, Telefono: %d, Correo electronico: %s\n", valor[0], key, valor[1])
				}
			}
			fmt.Print("\n\n")
			opcion = utilidades.MenuDeInicio() // se ejecuta el menu de inicio
		}

		// se ejecuta si la opcion elegida es 5: termina el programa
		if opcion == 5 {
			println("Fin del programa. Hasta luego!")
			fmt.Print("\n")
			break
		}
	}
}
