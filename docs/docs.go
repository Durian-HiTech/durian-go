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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/portal/del_sub": {
            "post": {
                "description": "删除订阅",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "订阅ID",
                        "name": "subscription_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"删除成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "{\"success\": false, \"message\": \"数据库error, 一些其他错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\": false, \"message\": \"用户ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/list_all_subs": {
            "post": {
                "description": "获取订阅列表",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"user的所有订阅\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\": false, \"message\": \"用户ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/sub": {
            "post": {
                "description": "订阅城市疫情信息",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "城市名字",
                        "name": "city_name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"订阅成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "{\"success\": false, \"message\": \"数据库error, 一些其他错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\": false, \"message\": \"用户ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "post": {
                "description": "查看用户个人信息",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"查看用户信息成功\", \"data\": \"model.User的所有信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\": false, \"message\": \"用户ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "登录",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"登录成功\", \"detail\": user的信息}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"success\": false, \"message\": \"密码错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\": false, \"message\": \"没有该用户\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/modify": {
            "post": {
                "description": "修改用户信息（支持修改用户名和密码）",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "原密码",
                        "name": "password_old",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "新密码",
                        "name": "password_new",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"修改成功\", \"data\": \"model.User的所有信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"success\": false, \"message\": \"原密码输入错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\": false, \"message\": \"用户ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "注册",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户类型（0: 普通用户，1: 认证机构用户）",
                        "name": "user_type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "认证机构名",
                        "name": "affiliation",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"用户创建成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"success\": false, \"message\": \"用户已存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{"http", "https"},
	Title:       "Durian Covid-19 Golang Backend",
	Description: "Durian HiTech",
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
