// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/apiV1/suites": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "suite"
                ],
                "summary": "suite api",
                "parameters": [
                    {
                        "type": "boolean",
                        "name": "isLike",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "执行结果",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user service"
                ],
                "summary": "login api",
                "parameters": [
                    {
                        "description": "登录请求",
                        "name": "entity",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ApiLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "执行结果",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/logout": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user service"
                ],
                "summary": "logout",
                "responses": {
                    "200": {
                        "description": "执行结果",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ApiLoginReq": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.JsonResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "错误码((0:成功, 1:失败, \u003e1:错误码))",
                    "type": "integer"
                },
                "data": {
                    "description": "返回数据(业务接口定义具体数据结构)",
                    "type": "object"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
