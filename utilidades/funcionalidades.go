package utilidades

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// Variables de entorno
var Dict = NewDictionary[int, []string]()
var opcion int
var Telefono int
var Nombre string
var Correo string

var ColorReset = "\033[0m"
var ColorRed = "\033[31m"
var ColorGreen = "\033[32m"
var ColorYellow = "\033[33m"
var ColorBlue = "\033[34m"
var ColorMagenta = "\033[35m"
var ColorCyan = "\033[36m"

// ***** MENU DE OPCIONES *****

// post: se despliega en consola el menu de opciones y retornara aquella que el usuario elija
func MenuDeInicio() int {
	// Opciones a elegir por el usuario
	fmt.Println(string(ColorMagenta) + "Menu de opciones: " + string(ColorReset))
	fmt.Println(string(ColorMagenta) + "~~~~~~~~~~~~~~~~~~~~~~~~~~~" + string(ColorReset))
	fmt.Println("-> 1 para agregar un nuevo contacto.")
	fmt.Println("-> 2 para eliminar un contacto.")
	fmt.Println("-> 3 para buscar un contacto en la lista.")
	fmt.Println("-> 4 para ver la lista completa de contactos.")
	fmt.Println(string(ColorCyan) + "-> 5 para finalizar el programa." + string(ColorReset))
	fmt.Print("\n\n")

	// VALIDADOR DE OPCIONES: validacion de dominio que verifica que los valores ingresados sean correctos
	fmt.Print("Seleccione una opcion (1 - 5): ")
	fmt.Scanln(&opcion)
	for opcion != 1 && opcion != 2 && opcion != 3 && opcion != 4 && opcion != 5 {
		fmt.Print("Seleccione una opcion valida (1 - 5): ")
		fmt.Scanln(&opcion)
	}
	fmt.Print("\n\n")
	return opcion // valor de retorno
}

// post: se limpia el contenido de la consola para reiniciarla
func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") // Comando para Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear") // Comando para Unix
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// ***** VALIDACIONES DE DOMINIO *****

// post: se verifica que los datos ingresados por el usuario cumplan con los requisitos
func ValidacionOpcionUno() {

	// Expresion regular para validar los caracteres del correo
	reCorreo := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	// Expresión regular para validar el nombre
	reNombre := regexp.MustCompile(`^[A-Za-z\s]+$`)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println(string(ColorMagenta) + "Ingrese a continuacion el telefono, nombre y correo electronico del contacto: " + string(ColorReset))
	fmt.Print("\n")

	// se valida el numero de celular del contacto a ingresar
	fmt.Print("Telefono (Ingrese 10 digitos): ")
	fmt.Scanln(&Telefono)
	for Telefono < 1000000000 || Telefono > 9999999999 {
		fmt.Print("Ingrese un telefono valido (10 digitos): ")
		fmt.Scanln(&Telefono)
	}

	// se valida el nombre del contacto a ingresar
	fmt.Print("Ingrese el nombre y apellido del contacto (solo letras): ")
	Nombre, _ = reader.ReadString('\n')
	Nombre = strings.TrimSpace(Nombre)
	for !reNombre.MatchString(Nombre) || len(Nombre) < 2 {
		fmt.Print("Ingrese el nombre del contacto: ")
		Nombre, _ = reader.ReadString('\n')
		Nombre = strings.TrimSpace(Nombre)
	}

	// se valida el correo electronico del contacto a ingresar
	fmt.Print("Ingrese el correo del contacto: ")
	fmt.Scanln(&Correo)
	// mientras el correo no tenga la estructura solicitada, se volvera a pedirlo
	for !reCorreo.MatchString(Correo) {
		fmt.Print("Ingrese el correo del contacto: ")
		fmt.Scanln(&Correo)
	}
}

// post: el usuario ingresa el telefono del contacto a eliminar de la lista
func ValidacionOpcionDos() {
	fmt.Print("Ingrese el telefono del contacto que desea eliminar: ")
	fmt.Scanln(&Telefono)
	for !Dict.Contains(Telefono) {
		fmt.Print("El telefono ingresado no existe, ingrese el telefono del contacto que desea eliminar: ")
		fmt.Scanln(&Telefono)
	}
}

// post: el usuario ingresa el nombre del contacto a buscar en la lista
func ValidacionOpcionTres() {
	fmt.Print("Ingrese el telefono del contacto que desea buscar: ")
	fmt.Scanln(&Telefono)
	for !Dict.Contains(Telefono) {
		fmt.Print("El telefono ingresado no existe, ingrese el telefono del contacto que desea buscar: ")
		fmt.Scanln(&Telefono)
	}
}

// ***** FUNCIONALIDADES DEL GESTOR DE CONTACOS *****

// pre: se ingresa el numero de telefono, el nombre y el correo del contacto a agregar
// post: retornar un string diciendo que el contacto se agrego a la lista
func AgregarContacto(telefono int, nombre string, correo string) string {
	datos := []string{nombre, correo}
	Dict.Put(telefono, datos)
	return (string(ColorGreen) + "Contacto agregado de forma exitosa!" + string(ColorReset))
}

// pre: se ingresa el numero de telefono
// post: retornar un string diciendo que el contacto se agrego a la lista
func EliminarContacto(telefono int) string {
	Dict.Remove(telefono)
	return (string(ColorRed) + "Contacto eliminado." + string(ColorReset))
}

// post: retornar un string con los datos del contacto buscado por el usuario
func BuscarContacto() string {
	celularStr := strconv.Itoa(Telefono)
	return ("Nombre: " + Nombre + " Telefono: " + celularStr + " Correo: " + Correo)
}
