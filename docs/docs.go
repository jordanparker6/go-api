// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/jordanparker6/go-api/LICENSE",
        "contact": {
            "name": "Jordan Parker",
            "url": "https://github.com/jordanparker6/go-api",
            "email": "jordanhparker6@gmail.com"
        },
        "license": {
            "name": "MIT Licence",
            "url": "https://github.com/jordanparker6/go-api/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/openapi.json": {
            "get": {
                "description": "Fetches the OpenAPI json specification for the API.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Core"
                ],
                "summary": "OpenAPI",
                "operationId": "open-api",
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPError"
                        }
                    }
                }
            }
        },
        "/test": {
            "get": {
                "description": "An endpoint to test the API's connectivity.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Core"
                ],
                "summary": "Test API",
                "operationId": "test-api",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPSuccess"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPError"
                        }
                    }
                }
            }
        },
        "/user/create": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Creates a new user record.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "CreateUser",
                "operationId": "create-user",
                "parameters": [
                    {
                        "description": "User ID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CreateUserSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPError"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Gets a user by user ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetUser",
                "operationId": "get-user-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Delete a user record by user ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "DeleteUser",
                "operationId": "delete-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/core.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.CreateUserSchema": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                }
            }
        },
        "core.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "core.HTTPSuccess": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "message": {
                    "type": "string",
                    "example": "request completed"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AccessToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "tags": [
        {
            "description": "Core API enpoints for documentation and health checks.",
            "name": "Core"
        },
        {
            "description": "User API endpoints for CRUD operations on user records.",
            "name": "User"
        }
    ],
    "x-extension-openapi": {
        "example": "value on a json format"
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8000",
	BasePath:    "/api/v1",
	Schemes:     []string{"http", "https"},
	Title:       "Golang Example API",
	Description: "This is boilerplate Golang REST API. Please feel free to contribute to this boilerplate.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
