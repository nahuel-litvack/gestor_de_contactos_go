package utilidades

import (
	"fmt"
	"regexp"
	"strconv"
)

// Variables de entorno
var Dict = NewDictionary[int, []string]()
var opcion int
var Telefono int
var Nombre string
var Correo string

// ***** MENU DE OPCIONES *****

// post: se despliega en consola el menu de opciones y retornara aquella que el usuario elija
func MenuDeInicio() int {
	// Opciones a elegir por el usuario
	fmt.Print("\n\n")
	fmt.Println("Menu de opciones: ")
	fmt.Println("a. Seleccione 1 para agregar un nuevo contacto.")
	fmt.Println("b. Seleccione 2 para eliminar un contacto.")
	fmt.Println("c. Seleccione 3 para buscar un contacto en la lista.")
	fmt.Println("d. Seleccione 4 para ver la lista completa de contactos.")
	fmt.Println("e. Seleccione 5 para finalizar el programa.")
	fmt.Print("\n\n")

	// VALIDADOR DE OPCIONES: validacion de dominio que verifica que los valores ingresados sean correctos
	fmt.Print("Seleccione una opcion (1-5): ")
	fmt.Scanln(&opcion)
	for opcion != 1 && opcion != 2 && opcion != 3 && opcion != 4 && opcion != 5 {
		fmt.Print("Seleccione una opcion valida (1-5): ")
		fmt.Scanln(&opcion)
	}
	fmt.Print("\n\n")
	return opcion // valor de retorno
}

// ***** VALIDACIONES DE DOMINIO *****

// post: se verifica que los datos ingresados por el usuario cumplan con los requisitos
func ValidacionOpcionUno() {
	// Expresion regular para validar los caracteres del correo
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

	fmt.Println("Ingrese a continuacion el telefono, nombre y correo electronico del contacto: ")
	fmt.Print("\n")

	// se valida el numero de celular del contacto a ingresar
	fmt.Print("Telefono (Ingrese 10 digitos): ")
	fmt.Scanln(&Telefono)
	for Telefono < 1000000000 || Telefono > 9999999999 {
		fmt.Print("Ingrese un telefono valido (10 digitos): ")
		fmt.Scanln(&Telefono)
	}
	fmt.Print("\n")

	// se valida el nombre del contacto a ingresar
	fmt.Print("Ingrese el nombre del contacto: ")
	fmt.Scanln(&Nombre)
	for len(Nombre) < 2 {
		fmt.Print("Ingrese el nombre del contacto: ")
		fmt.Scanln(&Nombre)
	}
	fmt.Print("\n")

	// se valida el correo electronico del contacto a ingresar
	fmt.Print("Ingrese el correo del contacto: ")
	fmt.Scanln(&Correo)
	// mientras el correo no tenga la estructura solicitada, se volvera a pedirlo
	for !re.MatchString(Correo) {
		fmt.Print("Ingrese el correo del contacto: ")
		fmt.Scanln(&Correo)
	}
}

// post: el usuario ingresa el nombre del contacto a eliminar de la lista
func ValidacionOpcionDos() {
	valor := Dict.Get(Telefono)
	fmt.Print("Ingrese el nombre del contacto que desea eliminar: ")
	fmt.Scanln(&Nombre)
	for valor[0] != Nombre {
		fmt.Print("El nombre ingresado no existe, ingrese el nombre del contacto que desea eliminar: ")
		fmt.Scanln(&Nombre)
	}
}

// post: el usuario ingresa el nombre del contacto a buscar en la lista
func ValidacionOpcionTres() {
	valor := Dict.Get(Telefono)
	fmt.Print("Ingrese el nombre del contacto que desea buscar: ")
	fmt.Scanln(&Nombre)
	for valor[0] != Nombre {
		fmt.Print("El nombre ingresado no existe, ingrese el nombre del contacto que desea buscar: ")
		fmt.Scanln(&Nombre)
	}
}

// ***** FUNCIONALIDADES DEL GESTOR DE CONTACOS *****

// pre: se ingresa el numero de telefono, el nombre y el correo del contacto a agregar
// post: retornar un string diciendo que el contacto se agrego a la lista
func AgregarContacto(telefono int, nombre string, correo string) string {
	datos := []string{nombre, correo}
	Dict.Put(telefono, datos)
	fmt.Print("\n")
	return "Contacto agregado de forma exitosa!"
}

// pre: se ingresa el numero de telefono
// post: retornar un string diciendo que el contacto se agrego a la lista
func EliminarContacto(telefono int) string {
	Dict.Remove(telefono)
	fmt.Print("\n")
	return "Contacto eliminado"
}

// post: retornar un string con los datos del contacto buscado por el usuario
func BuscarContacto() string {
	fmt.Print("\n")
	celularStr := strconv.Itoa(Telefono)
	return ("Nombre: " + Nombre + " Telefono: " + celularStr + " Correo: " + Correo)
}
