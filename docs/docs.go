// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-04-13 08:17:30.684107 +0800 CST m=+0.232426146

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "An example of gin",
        "title": "Golang Gin API",
        "termsOfService": "127.0.0.1:6001",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "127.0.0.1:6001"
        },
        "version": "1.0"
    },
    "basePath": "/v1",
    "paths": {
        "/api/v1/tags/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取S",
                "parameters": [
                    {},
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "State",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ModifiedBy",
                        "name": "modified_by",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
