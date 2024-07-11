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
        "/status-check": {
            "get": {
                "description": "Returns a 200 status if server is up",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Status check",
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
        "/worker/insert-many-sorted": {
            "post": {
                "description": "Inserts data to be executed in callback at specific date in scheduler worker for each element.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Worker"
                ],
                "summary": "Schedules callback for all elements.",
                "parameters": [
                    {
                        "description": "Insert tasks",
                        "name": "tasks",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/worker_model.InsertSorted"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    }
                }
            }
        },
        "/worker/insert-sorted": {
            "post": {
                "description": "Inserts data to be executed in callback at specific date in scheduler worker.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Worker"
                ],
                "summary": "Schedules callback at data.",
                "parameters": [
                    {
                        "description": "Insert task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/worker_model.InsertSorted"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "fiber.Map": {
            "type": "object",
            "additionalProperties": true
        },
        "time.Duration": {
            "type": "integer",
            "enum": [
                -9223372036854775808,
                9223372036854775807,
                1,
                1000,
                1000000,
                1000000000,
                60000000000,
                3600000000000
            ],
            "x-enum-varnames": [
                "minDuration",
                "maxDuration",
                "Nanosecond",
                "Microsecond",
                "Millisecond",
                "Second",
                "Minute",
                "Hour"
            ]
        },
        "worker_model.InsertSorted": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "Data passed to callback."
                },
                "execute_at": {
                    "description": "Execution date of callback.",
                    "type": "string"
                },
                "future_offset": {
                    "description": "Difference between the current time and the execution date in seconds. Will be loaded just if execute_at is not set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/time.Duration"
                        }
                    ]
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Http Scheduler",
	Description:      "Schedules data propagations using HTTP requests",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
