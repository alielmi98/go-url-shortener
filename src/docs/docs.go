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
        "/v1/shorten": {
            "post": {
                "description": "Create shortn url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shortn_urls"
                ],
                "summary": "Create shortn url",
                "parameters": [
                    {
                        "description": "CreateShortnUrlRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_dto.CreateShortnUrlRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/shorten/{short_code}": {
            "get": {
                "description": "Redirect to original URL using short code",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shortn_urls"
                ],
                "summary": "Redirect to original URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ShortnUrl Short Code",
                        "name": "short_code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Redirect",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    },
                    "404": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    },
                    "500": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update shortn url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shortn_urls"
                ],
                "summary": "Update shortn url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ShortnUrl Short Code",
                        "name": "short_code",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "UpdateShortnUrlRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_dto.UpdateShortnUrlRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    },
                    "404": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    },
                    "500": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete shortn url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shortn_urls"
                ],
                "summary": "Delete shortn url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ShortnUrl Short Code",
                        "name": "short_code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    },
                    "404": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    },
                    "500": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/shorten/{short_code}/stats": {
            "get": {
                "description": "Get shortn url by short code",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shortn_urls"
                ],
                "summary": "Get shortn url by short code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ShortnUrl Short Code",
                        "name": "short_code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    },
                    "404": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    },
                    "500": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_alielmi98_go-url-shortener_api_dto.CreateShortnUrlRequest": {
            "type": "object",
            "required": [
                "original_url"
            ],
            "properties": {
                "original_url": {
                    "type": "string"
                }
            }
        },
        "github_com_alielmi98_go-url-shortener_api_dto.UpdateShortnUrlRequest": {
            "type": "object",
            "required": [
                "original_url"
            ],
            "properties": {
                "access_count": {
                    "type": "integer"
                },
                "original_url": {
                    "type": "string"
                }
            }
        },
        "github_com_alielmi98_go-url-shortener_api_helper.BaseHttpResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "result": {},
                "resultCode": {
                    "$ref": "#/definitions/github_com_alielmi98_go-url-shortener_api_helper.ResultCode"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "github_com_alielmi98_go-url-shortener_api_helper.ResultCode": {
            "type": "integer",
            "enum": [
                0,
                40001,
                40101,
                40301,
                40401,
                42901,
                42902,
                50001,
                50002,
                50003
            ],
            "x-enum-varnames": [
                "Success",
                "ValidationError",
                "AuthError",
                "ForbiddenError",
                "NotFoundError",
                "LimiterError",
                "OtpLimiterError",
                "CustomRecovery",
                "InternalError",
                "InvalidInputError"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
