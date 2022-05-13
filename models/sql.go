package models

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	sqlConnection := os.Getenv("SQL_CONNECTION")
	var err error
	db, err = gorm.Open(sqlserver.Open(sqlConnection), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error obteniendo base de datos: %v", err)
	} else {
		fmt.Println("Conectado con exito a", db.Name())
	}
}

func ConsultaColores(color string) (string, bool) {
	var coloresMapeo Traductor
	validacionColor := true
	err := db.Raw(`SELECT color, colour FROM [ANDREANI\mpruyas].colores_mapeo WHERE color = ?`, color).Take(&coloresMapeo).Error
	if err != nil {
		validacionColor = false
		coloresMapeo.Colour = ""
	}
	return coloresMapeo.Colour, validacionColor
}
