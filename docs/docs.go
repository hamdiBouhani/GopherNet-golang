// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/burrows": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show Burrow Status",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.IndexResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "results": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Burrow"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rent-burrow/{id}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Rent Burrow",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Burrow"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.IndexResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "results": {}
            }
        },
        "model.Burrow": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "created_date": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "depth": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "occupied": {
                    "type": "boolean"
                },
                "updated_date": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                },
                "wide": {
                    "type": "number"
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
	Title:            "GopherNet API",
	Description:      "GopherNet API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}