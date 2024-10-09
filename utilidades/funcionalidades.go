package utilidades

import "fmt"

var Dict = NewDictionary[int, []string]()
var opcion int
var Telefono int
var Nombre string
var Correo string

// --------------------------------------------------------------------------------------------------------
// MENU PRINCIPAL DEL PROGRAMA
// --------------------------------------------------------------------------------------------------------

func MenuDeInicio() int {
	// Opciones a elegir por el usuario
	fmt.Print("\n\n")
	fmt.Println("Menu de opciones: ")
	fmt.Println("a. Seleccione 1 para agregar un nuevo contacto.")
	fmt.Println("b. Seleccione 2 para eliminar un contacto.")
	fmt.Println("c. Seleccione 3 para buscar un contacto en la lista.")
	fmt.Println("d. Seleccione 4 para ver la lista completa de contactos.")
	fmt.Print("\n\n")

	// VALIDADOR DE OPCIONES: de dominio que verifica que los valores ingresados sean correctos
	fmt.Print("Seleccione una opcion (1-4): ")
	fmt.Scanln(&opcion)
	for opcion != 1 && opcion != 2 && opcion != 3 && opcion != 4 {
		fmt.Print("Seleccione una opcion valida (1-4): ")
		fmt.Scanln(&opcion)
	}
	fmt.Print("\n\n")
	return opcion
}

// --------------------------------------------------------------------------------------------------------
// VALIDACIONES DE DOMINIO
// --------------------------------------------------------------------------------------------------------

func ValidacionOpcionUno() {

	fmt.Println("Ingrese a continuacion el telefono, nombre y correo electronico del contacto: ")
	fmt.Print("\n")
	fmt.Print("Telefono (Ingrese 10 digitos): ")
	fmt.Scanln(&Telefono)
	for Telefono < 1000000000 || Telefono > 9999999999 {
		fmt.Print("Ingrese un telefono valido (10 digitos): ")
		fmt.Scanln(&Telefono)
	}

	fmt.Print("\n")

	fmt.Print("Ingrese el nombre del contacto: ")
	fmt.Scanln(&Nombre)
	for Nombre == "" {
		fmt.Print("Ingrese el nombre del contacto: ")
		fmt.Scanln(&Nombre)
	}

	fmt.Print("\n")

	fmt.Print("Ingrese el correo del contacto: ")
	fmt.Scanln(&Correo)
	for Correo == "" {
		fmt.Print("Ingrese el correo del contacto: ")
		fmt.Scanln(&Correo)
	}
}

func ValidacionOpcionDos() {
	fmt.Print("Ingrese el telefono del contacto que desea eliminar: ")
	fmt.Scanln(&Telefono)
	for !(Dict.Contains(Telefono)) {
		fmt.Print("El numero ingresado no existe, ingrese el telefono del contacto que desea eliminar: ")
		fmt.Scanln(&Telefono)
	}
}

func ValidacionOpcionTres() {
	fmt.Print("Ingrese el telefono del contacto que desea buscar: ")
	fmt.Scanln(&Telefono)
	for !(Dict.Contains(Telefono)) {
		fmt.Print("El numero ingresado no existe, ingrese el telefono del contacto que desea buscar: ")
		fmt.Scanln(&Telefono)
	}
}

// --------------------------------------------------------------------------------------------------------
// FUNCIONALIDADES DEL GESTOR DE CONTACTOS
// --------------------------------------------------------------------------------------------------------

func AgregarContacto(telefono int, nombre string, correo string) string {
	datos := []string{nombre, correo}
	Dict.Put(telefono, datos)
	fmt.Print("\n")
	return "Contacto agregado exitosamente!"
}

func EliminarContacto(telefono int) string {
	Dict.Remove(telefono)
	return "Contacto eliminado"
}

func BuscarContacto(telefono int) int {
	if Dict.Contains(telefono) {
		return telefono
	}
	return -1
}

// * Funcion a considerar para remover
func ListarContacto() []int {
	return Dict.Keys()
}
