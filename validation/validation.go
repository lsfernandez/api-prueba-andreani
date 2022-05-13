package validation

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/customer-experience/api-prueba/models"
)

func LimpiarEspacios(nombre string) string {
	strArray := strings.Fields(nombre)
	strLimpio := strings.Join(strArray, " ")
	return strLimpio
}

func ValidarNombre(nombre string) (string, bool) {
	nombreLimpio := LimpiarEspacios(nombre)
	match, _ := regexp.MatchString(`^([a-zA-Z]{2,24} ){1,3}([A-Za-z']{2,24}){0,2}$`, nombreLimpio)
	return nombreLimpio, match
}

func ValidarRequestCreacion(validacionNombre bool, validacionColor bool) error {
	var err error

	if !validacionNombre && !validacionColor {
		err = errors.New("ambos campos invalidos")
		return err
	}

	if !validacionNombre {
		err = errors.New("nombre invalido")
		return err
	}

	if !validacionColor {
		err = errors.New("color invalido")
		return err
	}

	return err
}

func ValidarRequestModificacion(input models.PedidoInput, consulta models.ResponsePedido) (models.ResponsePedido, error) {
	var err error

	if input.Color == "" && input.NombreCompleto == "" {
		err = errors.New("no se envio ningun campo")
		return consulta, err
	}

	if input.Color != "" {
		traduccion, validacionColor := models.ConsultaColores(input.Color)
		if !validacionColor {
			err = errors.New("color invalido")
			return consulta, err
		}

		consulta.Color = strings.Title(input.Color)
		consulta.ColorIngles = strings.Title(traduccion)
	}

	if input.NombreCompleto != "" {
		nombreValidado, validacionNombre := ValidarNombre(input.NombreCompleto)
		if !validacionNombre {
			err = errors.New("nombre invalido")
			return consulta, err
		}
		consulta.NombreCompleto = strings.Title(nombreValidado)
	}

	consulta.FechaPedido = time.Now().Local()

	return consulta, err
}
