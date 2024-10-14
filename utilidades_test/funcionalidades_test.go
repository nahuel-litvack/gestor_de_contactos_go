package utilidades_test

import (
	"strconv"
	"testing"

	"github.com/nahuel-litvack/repositorio01/utilidades"
	"github.com/stretchr/testify/assert"
)

// *Revisar test
func TestAgregarContacto(t *testing.T) {
	dict := utilidades.NewDictionary[int, []string]()

	telefono := 1168707070
	nombre := "Gonzalo"
	correo := "gonzalito123@mail.com"

	utilidades.AgregarContacto(telefono, nombre, correo)

	datos := []string{nombre, correo}
	dict.Put(telefono, datos)

	celularStr := strconv.Itoa(telefono)
	assert.Equal(t, utilidades.BuscarContacto(), ("Nombre: " + nombre + " Telefono: " + celularStr + " Correo: " + correo))
}
