// Package docs GENERATED BY SWAG; DO NOT EDIT
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
            "name": "systemallica",
            "url": "http://www.andres.reveronmolina.me",
            "email": "andres@reveronmolina.me"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/rides": {
            "post": {
                "description": "create ride",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rides"
                ],
                "summary": "starts a ride.",
                "parameters": [
                    {
                        "description": "Ride request parameters",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.RideRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/rides.Ride"
                        }
                    }
                }
            }
        },
        "/rides/:id/finish": {
            "post": {
                "description": "finish ride",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rides"
                ],
                "summary": "finishes the ride that matches the given ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ride ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rides.Ride"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.RideRequest": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string"
                },
                "vehicle_id": {
                    "type": "string"
                }
            }
        },
        "rides.Ride": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "finished": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "vehicle_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Rides Swagger API",
	Description:      "This is a basic Rides API using Chi and go-rel.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
