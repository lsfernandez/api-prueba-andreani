// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Equipo de integraciones",
            "url": "https://developers.andreani.com",
            "email": "apis@andreani.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/pedidos": {
            "post": {
                "description": "Recibe un json con nombre completo y un color. Valida que el nombre sea valido y consulta si el color esta en la base de datos y si lo está, trae su traduccion. Genera un pedido y lo guarda en una base de datos Mongo.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Recibe un nombre completo y un color. Genera un pedido",
                "parameters": [
                    {
                        "description": "Nombre Completo y Color para dar de alta un pedido",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.PedidoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponsePedido"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/pedidos/{id}": {
            "get": {
                "description": "Recibe el id por parametro, lo consulta en la base de datos y si se encontró muestra todos los datos del pedido. Si no se encontró arroja un error",
                "produces": [
                    "application/json"
                ],
                "summary": "Recibe un id de pedido y responde con los datos del mismo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id de Pedido",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/models.ResponsePedido"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Recibe un id por parametro y un json con nombre completo y/o un color. Valida que el nombre sea valido y consulta si el color esta en la base de datos, si lo está, trae su traduccion. Actualiza el pedido y lo guarda en una base de datos Mongo.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Recibe un id de pedido, nombre completo y/o un color. Modifica los datos del pedido",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id de Pedido",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Nombre Completo y/o Color para modificar un pedido",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.PedidoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponsePedido"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.PedidoInput": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "nombreCompleto": {
                    "type": "string"
                }
            }
        },
        "models.ResponsePedido": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "colorIngles": {
                    "type": "string"
                },
                "fechaPedido": {
                    "type": "string"
                },
                "idPedido": {
                    "type": "string"
                },
                "nombreCompleto": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Api Prueba",
	Description:      "API Restful en Golang para obtener la cotizacion para un envío",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
