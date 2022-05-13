package models

import (
	"strings"
	"time"
)

type PedidoInput struct {
	NombreCompleto string `json:"nombreCompleto"`
	Color          string `json:"color"`
}

type Pedido struct {
	NombreCompleto string    `json:"nombreCompleto"`
	Color          string    `json:"color"`
	ColorIngles    string    `json:"colorIngles"`
	FechaPedido    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"fechaPedido"`
}

type Traductor struct {
	Color  string `json:"color" sql:"column:color"`
	Colour string `json:"colour" sql:"column:colour"`
}

func ArmarPedido(nombreCompleto string, color string, colorIngles string, fechaPedido time.Time) Pedido {
	pedido := Pedido{NombreCompleto: strings.Title(nombreCompleto), Color: strings.Title(color), ColorIngles: strings.Title(colorIngles), FechaPedido: fechaPedido.Local()}
	return pedido
}
