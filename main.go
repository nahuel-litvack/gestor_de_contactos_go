package main

import (
	"fmt"

	"github.com/nahuel-litvack/repositorio01/utilidades"

	"time"
)

// A mejorar:
// . Agregar pruebas unitarias a las funcionalidades y al diccionario.
// . Agregar una GUI (interfaz grafica).
// y asi terminamos el proyecto...

// ***** MENU DE OPCIONES *****

func main() {
	opcion := utilidades.MenuDeInicio() // se ejecuta el menu de inicio

	for opcion == 1 || opcion == 2 || opcion == 3 || opcion == 4 || opcion == 5 {
		// se ejecuta mientra la opcion elegida sea 1: agrega un contacto al diccionario
		for opcion == 1 {
			utilidades.ClearScreen()
			utilidades.ValidacionOpcionUno()
			utilidades.ClearScreen()
			resultado := utilidades.AgregarContacto(utilidades.Telefono, utilidades.Nombre, utilidades.Correo)
			fmt.Println(resultado)
			time.Sleep(2 * time.Second)
			utilidades.ClearScreen()
			opcion = utilidades.MenuDeInicio() // se ejecuta el menu de inicio
		}

		// se ejecuta mientras la opcion elegida sea 2: elimina un contacto del diccionario
		for opcion == 2 {
			// Validacion de dominio que verifica que el diccionario no este vacio
			if utilidades.Dict.Size() == 0 {
				fmt.Println(string(utilidades.ColorYellow) + "No hay ningun contacto en la lista." + string(utilidades.ColorReset))
				time.Sleep(2 * time.Second)
				utilidades.ClearScreen()
				opcion = utilidades.MenuDeInicio()
			} else {
				utilidades.ClearScreen()
				utilidades.ValidacionOpcionDos()
				utilidades.ClearScreen()
				resultado := utilidades.EliminarContacto(utilidades.Telefono)
				fmt.Println(resultado)
				time.Sleep(2 * time.Second)
				utilidades.ClearScreen()
				opcion = utilidades.MenuDeInicio() // se ejecuta el menu de inicio
			}
		}

		// se ejecuta mientra la opcion elegida sea 3: busca un contacto en el diccionario
		for opcion == 3 {
			// Validacion de dominio que verifica que el diccionario no este vacio
			if utilidades.Dict.Size() == 0 {
				fmt.Println(string(utilidades.ColorYellow) + "No hay ningun contacto en la lista." + string(utilidades.ColorReset))
				time.Sleep(2 * time.Second)
				utilidades.ClearScreen()
				opcion = utilidades.MenuDeInicio()
			} else {
				utilidades.ClearScreen()
				utilidades.ValidacionOpcionTres()
				utilidades.ClearScreen()
				resultado := utilidades.BuscarContacto()
				fmt.Print((utilidades.ColorMagenta) + "Contacto buscado: " + string(utilidades.ColorReset))
				fmt.Println(resultado)
				fmt.Print("\n")
				fmt.Println("En 10 segundos se regresa al menu de opciones.")
				time.Sleep(10 * time.Second)
				utilidades.ClearScreen()
				opcion = utilidades.MenuDeInicio() // se ejecuta el menu de inicio
			}
		}

		// se ejecuta mientra la opcion elegida sea 4: devuelve la lista de contactos en el diccionario
		for opcion == 4 {
			// Validacion de dominio que verifica que el diccionario no este vacio
			if utilidades.Dict.Size() == 0 {
				fmt.Println(string(utilidades.ColorYellow) + "No hay ningun contacto en la lista." + string(utilidades.ColorReset))
				time.Sleep(2 * time.Second)
				utilidades.ClearScreen()
				opcion = utilidades.MenuDeInicio()
			} else {
				utilidades.ClearScreen()
				fmt.Println(string(utilidades.ColorMagenta) + "Lista de contactos:" + string(utilidades.ColorReset))
				fmt.Print("\n")
				for _, key := range utilidades.Dict.Keys() {
					valor := utilidades.Dict.Get(key)
					fmt.Printf("Nombre: %s, Telefono: %d, Correo electronico: %s\n", valor[0], key, valor[1])
				}
				fmt.Print("\n\n")
				fmt.Println("En 15 segundos se regresa al menu de opciones.")
				time.Sleep(15 * time.Second)
				utilidades.ClearScreen()
				opcion = utilidades.MenuDeInicio()
			}
		}

		// se ejecuta si la opcion elegida es 5: termina el programa
		if opcion == 5 {
			utilidades.ClearScreen()
			println(string(utilidades.ColorBlue) + "Fin del programa. Hasta luego!" + string(utilidades.ColorReset))
			time.Sleep(3 * time.Second)
			utilidades.ClearScreen()
			break
		}
	}
}
