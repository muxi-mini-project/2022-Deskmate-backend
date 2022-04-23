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
        "termsOfService": "http://swagger.io/terrms/",
        "contact": {
            "name": "qianren",
            "email": "1911401642@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/apply": {
            "get": {
                "description": "点击 同桌申请 时调用",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "apply"
                ],
                "summary": "显示收到的同桌申请",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"success\"}"
                    },
                    "400": {
                        "description": "{\"error_code\":\"00001\", \"message\":\"Fail.\"} or {\"error_code\":\"00002\", \"message\":\"Lack Param Or Param Not Satisfiable.\"}"
                    },
                    "401": {
                        "description": "{\"error_code\":\"10001\", \"message\":\"Token Invalid.\"} 身份认证失败 重新登录"
                    },
                    "500": {
                        "description": "{\"error_code\":\"30001\", \"message\":\"Fail.\"} 失败"
                    }
                }
            },
            "put": {
                "description": "用户确认接受的同桌申请",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "apply"
                ],
                "summary": "确认同桌申请,进入同桌打卡",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "要同意申请的同学学号",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Respondent"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"success\"}"
                    },
                    "401": {
                        "description": "{\"msg\":\"confirm faided\"}"
                    }
                }
            },
            "post": {
                "description": "从名片页面像对方发出申请",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "apply"
                ],
                "summary": "同桌申请",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "申请对象的学号(id)",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Respondent"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"success\", \"对象id\":\"string\"}"
                    },
                    "400": {
                        "description": "{\"error_code\":\"00001\", \"message\":\"Fail.\"} or {\"error_code\":\"00002\", \"message\":\"Lack Param Or Param Not Satisfiable.\"}"
                    },
                    "401": {
                        "description": "{\"error_code\":\"10001\", \"message\":\"Token Invalid.\"} 身份认证失败 重新登录"
                    },
                    "500": {
                        "description": "{\"error_code\":\"30001\", \"message\":\"Fail.\"} 失败"
                    }
                }
            }
        },
        "/apply/refuse": {
            "put": {
                "description": "用户拒绝接受的同桌申请",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "apply"
                ],
                "summary": "拒绝同桌申请",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "要拒绝申请的同学学号",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Respondent"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"success\"}"
                    },
                    "401": {
                        "description": "{\"msg\":\"confirm faided\"}"
                    }
                }
            }
        },
        "/card": {
            "get": {
                "description": "\"获取自己的名片信息\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "我的名片",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Card"
                        }
                    },
                    "401": {
                        "description": "身份验证失败"
                    },
                    "404": {
                        "description": "获取失败"
                    }
                }
            },
            "put": {
                "description": "\"修改名片的昵称，标签，头像，用户宣言\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "修改名片信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "名片信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Cardinfor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "修改成功"
                    },
                    "400": {
                        "description": "修改失败"
                    },
                    "401": {
                        "description": "验证失败"
                    }
                }
            },
            "post": {
                "description": "\"设置名片的昵称，标签，头像，用户宣言\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "设置名片",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "名片信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Cardinfor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "设置成功"
                    },
                    "400": {
                        "description": "设置失败"
                    },
                    "401": {
                        "description": "身份验证失败 重新登录"
                    }
                }
            }
        },
        "/card/avatar": {
            "post": {
                "description": "\"修改名片头像\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "修改头像",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"mgs\":\"success\"}",
                        "schema": {
                            "$ref": "#/definitions/model.Card"
                        }
                    },
                    "400": {
                        "description": "上传失败,请检查token与其他配置参数是否正确"
                    }
                }
            }
        },
        "/card/infor": {
            "post": {
                "description": "\"获取他人的名片信息\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "他人名片",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "要查看用户的学号",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Id"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Card"
                        }
                    },
                    "401": {
                        "description": "身份验证失败"
                    },
                    "404": {
                        "description": "获取失败"
                    }
                }
            }
        },
        "/dailyrecord/card": {
            "get": {
                "description": "点击查看同桌名片",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dailyrecord"
                ],
                "summary": "查看同桌名片",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功"
                    },
                    "400": {
                        "description": "{\"error_code\":\"00001\", \"message\":\"Fail.\"} or {\"error_code\":\"00002\", \"message\":\"Lack Param Or Param Not Satisfiable.\"}"
                    },
                    "401": {
                        "description": "{\"error_code\":\"10001\", \"message\":\"Token Invalid.\"} 身份认证失败 重新登录"
                    },
                    "500": {
                        "description": "{\"error_code\":\"30001\", \"message\":\"Fail.\"} 服务器错误"
                    }
                }
            }
        },
        "/dailyrecord/end": {
            "put": {
                "description": "\"强制中断当前同桌\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dailyrecord"
                ],
                "summary": "强制中断同桌",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "强制中断成功"
                    },
                    "400": {
                        "description": "强制中断关系失败"
                    },
                    "401": {
                        "description": "身份验证失败 重新登录"
                    },
                    "500": {
                        "description": "用户更新状态失败"
                    }
                }
            }
        },
        "/dailyrecord/message": {
            "get": {
                "description": "查看本次同桌消息记录，包含时间戳",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dailyrecord"
                ],
                "summary": "浏览消息记录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功"
                    },
                    "400": {
                        "description": "{\"error_code\":\"00001\", \"message\":\"Fail.\"} or {\"error_code\":\"00002\", \"message\":\"Lack Param Or Param Not Satisfiable.\"}"
                    },
                    "401": {
                        "description": "{\"error_code\":\"10001\", \"message\":\"Token Invalid.\"} 身份认证失败 重新登录"
                    },
                    "500": {
                        "description": "{\"error_code\":\"30001\", \"message\":\"Fail.\"} 失败"
                    }
                }
            }
        },
        "/dailyrecord/send": {
            "post": {
                "description": "\"发送一条消息\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dailyrecord"
                ],
                "summary": "发送消息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "每日打卡的内容",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Record"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "发送成功"
                    },
                    "400": {
                        "description": "发送失败"
                    },
                    "401": {
                        "description": "身份验证失败 重新登录"
                    },
                    "500": {
                        "description": "服务器发生错误"
                    }
                }
            }
        },
        "/dailyrecord/update": {
            "put": {
                "description": "\"每日更新打卡天数\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dailyrecord"
                ],
                "summary": "更新打卡",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "更新成功"
                    },
                    "400": {
                        "description": "更新失败，同桌关系已经解除"
                    },
                    "401": {
                        "description": "身份验证失败 重新登录"
                    },
                    "500": {
                        "description": "失败"
                    }
                }
            }
        },
        "/square": {
            "get": {
                "description": "\"显示名片数据流\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "square"
                ],
                "summary": "同桌广场",
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Card"
                            }
                        }
                    },
                    "404": {
                        "description": "搜索失败"
                    }
                }
            }
        },
        "/square/tag": {
            "post": {
                "description": "\"在同桌广场中搜索标签返回对应名片\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "square"
                ],
                "summary": "搜索标签",
                "parameters": [
                    {
                        "description": "输入要搜索的标签",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Tag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "搜索成功",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Card"
                            }
                        }
                    },
                    "401": {
                        "description": "Lack Param Or Param Not Satisfiable."
                    },
                    "404": {
                        "description": "搜索失败"
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "一站式登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "登录用户信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Log"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\":\"将student_id作为token保留\"}",
                        "schema": {
                            "$ref": "#/definitions/response"
                        }
                    },
                    "400": {
                        "description": "{\"error_code\":\"20001\",\"message\":\"Fail.\"} or {\"error_code\":\"00002\",\"message\":\"Lack Param Or Param Not Satisfiable.\"}"
                    },
                    "401": {
                        "description": "{\"error_code\":\"10001\",\"message\":\"password or account wrong.\"} 身份认证失败 重新登录"
                    }
                }
            }
        },
        "/user/infor": {
            "post": {
                "description": "\"获取用户的基本信息\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户界面",
                "parameters": [
                    {
                        "description": "用户学号",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Id"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "搜索成功",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.User"
                            }
                        }
                    },
                    "401": {
                        "description": "身份验证失败"
                    },
                    "404": {
                        "description": "获取基本信息失败"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Card": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "college": {
                    "description": "学院",
                    "type": "string"
                },
                "declaration": {
                    "type": "string"
                },
                "grade": {
                    "description": "年级",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "infor": {
                    "type": "string"
                },
                "major": {
                    "description": "专业",
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "sha": {
                    "type": "string"
                },
                "status": {
                    "description": "记录是否有同桌",
                    "type": "string"
                },
                "tag1": {
                    "type": "string"
                },
                "tag2": {
                    "type": "string"
                },
                "tag3": {
                    "type": "string"
                },
                "tag4": {
                    "type": "string"
                },
                "tag5": {
                    "type": "string"
                },
                "users_id": {
                    "type": "string"
                }
            }
        },
        "model.Cardinfor": {
            "type": "object",
            "properties": {
                "college": {
                    "type": "string"
                },
                "declaration": {
                    "type": "string"
                },
                "grade": {
                    "type": "string"
                },
                "infor": {
                    "type": "string"
                },
                "major": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "tag1": {
                    "type": "string"
                },
                "tag2": {
                    "type": "string"
                },
                "tag3": {
                    "type": "string"
                },
                "tag4": {
                    "type": "string"
                },
                "tag5": {
                    "type": "string"
                }
            }
        },
        "model.Id": {
            "type": "object",
            "properties": {
                "users_id": {
                    "type": "string"
                }
            }
        },
        "model.Log": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "student_id": {
                    "type": "string"
                }
            }
        },
        "model.Record": {
            "type": "object",
            "properties": {
                "information": {
                    "type": "string"
                }
            }
        },
        "model.Respondent": {
            "type": "object",
            "properties": {
                "respondent_id": {
                    "type": "string"
                }
            }
        },
        "model.Tag": {
            "type": "object",
            "properties": {
                "tag": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "college": {
                    "type": "string"
                },
                "grade": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "student_id": {
                    "type": "string"
                }
            }
        },
        "response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {}
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
	Version:     "1.0.0",
	Host:        "119.3.2.168:4016",
	BasePath:    "/api/v1/",
	Schemes:     []string{"http"},
	Title:       "Deskmat",
	Description: "同桌小程序",
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
	swag.Register("swagger", &s{})
}
