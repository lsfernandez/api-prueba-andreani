package controllers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/customer-experience/api-prueba/models"
	"github.com/customer-experience/api-prueba/validation"
	"github.com/gin-gonic/gin"
)

// CrearPedido POST godoc
// @Summary Recibe un nombre completo y un color. Genera un pedido
// @Description Recibe un json con nombre completo y un color. Valida que el nombre sea valido y consulta si el color esta en la base de datos y si lo está, trae su traduccion. Genera un pedido y lo guarda en una base de datos Mongo.
// @Accept  json
// @Produce  json
// @Param Request body PedidoInput true "Nombre Completo y Color para dar de alta un pedido"
// @Success 200 {object} models.ResponsePedido
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /pedidos [post]
func CrearPedido(ctx *gin.Context) {
	var input models.PedidoInput
	fechaCreacion := time.Now().Local()

	if errorInput := ctx.ShouldBindJSON(&input); errorInput != nil {
		ctx.JSON(http.StatusBadRequest, "Error de validacion del input")
		log.Println("Error de validacion del input: ", errorInput)
		return
	}

	nombreValidado, validacionNombre := validation.ValidarNombre(input.NombreCompleto)
	traduccion, validacionColor := models.ConsultaColores(input.Color)

	errValidacion := validation.ValidarRequestCreacion(validacionNombre, validacionColor)

	if errValidacion != nil {
		ctx.JSON(http.StatusBadRequest, errValidacion.Error())
		log.Println(errValidacion.Error())
		return
	}

	pedido := models.ArmarPedido(strings.Title(nombreValidado), strings.Title(input.Color), strings.Title(traduccion), fechaCreacion)
	idGenerado, errorGuardar := models.GuardarPedido(pedido)

	if errorGuardar != nil {
		ctx.JSON(http.StatusInternalServerError, "No se pudo guardar el pedido en la base de datos")
		log.Println("Error al guardar pedido: ", errorGuardar)
		return
	}

	response := models.ArmarResponse(idGenerado.Hex(), pedido.NombreCompleto, pedido.Color, pedido.ColorIngles, pedido.FechaPedido)
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

// BuscarPedido GET godoc
// @Summary Recibe un id de pedido y responde con los datos del mismo
// @Description Recibe el id por parametro, lo consulta en la base de datos y si se encontró muestra todos los datos del pedido. Si no se encontró arroja un error
// @Produce json
// @Param id path string true "Id de Pedido"
// @Success 202 {object} models.ResponsePedido
// @Failure 400 {object} string
// @Router /pedidos/{id} [get]
func BuscarPedido(ctx *gin.Context) {
	idConsulta := ctx.Param("id")
	consulta, errorConsulta := models.ConsultarPedido(idConsulta)
	if errorConsulta != nil {
		ctx.JSON(http.StatusBadRequest, "ID de pedido incorrecto")
		log.Println("Error al consultar pedido: ", errorConsulta)
		return
	}
	consulta.FechaPedido = consulta.FechaPedido.Local()
	ctx.JSON(http.StatusOK, gin.H{"data": consulta})
}

// ModificarPedido PATCH godoc
// @Summary Recibe un id de pedido, nombre completo y/o un color. Modifica los datos del pedido
// @Description Recibe un id por parametro y un json con nombre completo y/o un color. Valida que el nombre sea valido y consulta si el color esta en la base de datos, si lo está, trae su traduccion. Actualiza el pedido y lo guarda en una base de datos Mongo.
// @Accept  json
// @Produce  json
// @Param id path string true "Id de Pedido"
// @Param Request body PedidoInput true "Nombre Completo y/o Color para modificar un pedido"
// @Success 200 {object} models.ResponsePedido
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /pedidos/{id} [patch]
func ModificarPedido(ctx *gin.Context) {
	idConsulta := ctx.Param("id")
	var input models.PedidoInput
	consulta, errorConsulta := models.ConsultarPedido(idConsulta)

	if errorConsulta != nil {
		ctx.JSON(http.StatusBadRequest, "ID de pedido incorrecto")
		log.Println("Error al consultar pedido: ", errorConsulta)
		return
	}

	if errorInput := ctx.ShouldBindJSON(&input); errorInput != nil {
		ctx.JSON(http.StatusBadRequest, "Error de validacion del input")
		log.Println("Error de validacion del input: ", errorInput)
		return
	}

	modificacion, errValidacion := validation.ValidarRequestModificacion(input, consulta)

	if errValidacion != nil {
		ctx.JSON(http.StatusBadRequest, errValidacion.Error())
		log.Println(errValidacion.Error())
		return
	}

	errorModificarDB := models.ModificarPedido(idConsulta, modificacion)
	if errorModificarDB != nil {
		ctx.JSON(http.StatusInternalServerError, "No se pudo guardar el pedido en la base de datos")
		log.Println("Error al guardar la modificacion del pedido: ", errorModificarDB)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": modificacion})
}
