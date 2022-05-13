package models

import (
	"time"
)

type ResponsePedido struct {
	IdPedido       string    `json:"idPedido"`
	NombreCompleto string    `json:"nombreCompleto"`
	Color          string    `json:"color"`
	ColorIngles    string    `json:"colorIngles"`
	FechaPedido    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"fechaPedido"`
}

func ArmarResponse(idConsulta string, nombreCompleto string, color string, colorIngles string, fechaPedido time.Time) ResponsePedido {
	response := ResponsePedido{IdPedido: idConsulta, NombreCompleto: nombreCompleto, Color: color, ColorIngles: colorIngles, FechaPedido: fechaPedido}
	return response
}
