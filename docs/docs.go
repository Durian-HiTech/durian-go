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
        "/": {
            "get": {
                "description": "测试 Index 页",
                "tags": [
                    "测试"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"gcp\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_all_covid_cases": {
            "get": {
                "description": "获取所有地区的新冠感染人数，返回列表",
                "tags": [
                    "数据"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有地区的新冠感染人数\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_all_covid_cases_response": {
            "get": {
                "description": "获取所有地区的新冠感染人数，返回列表 [根据时间分组]",
                "tags": [
                    "数据"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有地区的新冠感染人数\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_all_covid_cases_response_province": {
            "post": {
                "description": "获取所有地区的新冠感染人数，返回列表 [根据时间分组] [Province]",
                "tags": [
                    "数据"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "区域名",
                        "name": "province",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有地区的新冠感染人数 [Province]\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_all_covid_cdrv": {
            "get": {
                "description": "获取所有地区的新冠感染/死亡/治愈/疫苗接种人数【信息综合】，返回列表",
                "tags": [
                    "数据"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有地区的新冠感染/死亡/治愈/疫苗接种人数\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_all_covid_cdrv_response": {
            "get": {
                "description": "获取所有地区的新冠感染/死亡/治愈/疫苗接种人数【信息综合】，返回列表 [根据时间分组]",
                "tags": [
                    "数据"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有地区的新冠感染/死亡/治愈/疫苗接种人数\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_all_covid_cdrv_response_province": {
            "post": {
                "description": "获取所有地区的新冠感染/死亡/治愈【信息综合】，返回列表 [根据时间分组] [Province]",
                "tags": [
                    "数据"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "区域名",
                        "name": "province",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有地区的新冠感染/死亡/治愈/疫苗接种人数 [Province]\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_all_covid_deaths": {
            "get": {
                "description": "获取所有地区的新冠死亡人数，返回列表",
                "tags": [
                    "数据"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有地区的新冠死亡人数\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_all_covid_deaths_response": {
            "get": {
                "description": "获取所有地区的新冠死亡人数，返回列表 [根据时间分组]",
                "tags": [
                    "数据"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有地区的新冠死亡人数\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_all_covid_deaths_response_province": {
            "post": {
                "description": "获取所有地区的新冠死亡人数，返回列表 [根据时间分组] [Province]",
                "tags": [
                    "数据"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "区域名",
                        "name": "province",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有地区的新冠死亡人数 [Province]\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_all_covid_recovereds": {
            "get": {
                "description": "获取所有地区的新冠治愈人数，返回列表",
                "tags": [
                    "数据"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有地区的新冠治愈人数\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_all_covid_recovereds_response": {
            "get": {
                "description": "获取所有地区的新冠治愈人数，返回列表 [根据时间分组]",
                "tags": [
                    "数据"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有地区的新冠治愈人数\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_all_covid_recovereds_response_province": {
            "post": {
                "description": "获取所有地区的新冠治愈人数，返回列表 [根据时间分组] [Province]",
                "tags": [
                    "数据"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "区域名",
                        "name": "province",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有地区的新冠治愈人数 [Province]\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_all_covid_vaccines": {
            "get": {
                "description": "获取所有地区的新冠疫苗接种人数，返回列表",
                "tags": [
                    "数据"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有地区的新冠疫苗接种人数\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_all_covid_vaccines_response": {
            "get": {
                "description": "获取所有地区的新冠疫苗接种人数，返回列表 [根据时间分组]",
                "tags": [
                    "数据"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有地区的新冠治愈人数\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_all_high_risk_areas": {
            "get": {
                "description": "获取所有中高风险地区，返回列表",
                "tags": [
                    "数据"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有中高风险地区\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/list_overview": {
            "get": {
                "description": "获取世界或中国的现存确诊、新增确诊、累积确诊、累计及新增新冠感染/死亡/治愈【信息综合】，返回列表 [根据时间分组] [Province]",
                "tags": [
                    "数据"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"nowcases\":{\"nownum\": 123, \"newnum\": 123}等数据}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/data/query_data": {
            "post": {
                "description": "获取在数据库中直接存的 Json File",
                "tags": [
                    "数据"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据文件名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询失败，无所需数据\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/news/detail": {
            "post": {
                "description": "查看单条新闻",
                "tags": [
                    "新闻"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "新闻ID",
                        "name": "news_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"该条新闻的详细信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\":true, \"message\":\"查询失败，新闻ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/news/list_all_news": {
            "get": {
                "description": "获取所有新闻，返回列表",
                "tags": [
                    "新闻"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有新闻\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/notice/create_comment": {
            "post": {
                "description": "创建一条评论",
                "tags": [
                    "防控知识板块"
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
                        "description": "用户类型",
                        "name": "user_type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "问题ID",
                        "name": "question_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "评论内容",
                        "name": "comment_content",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"用户评论成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"success\": false, \"message\": \"用户ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "{\"success\": false, \"message\": \"问题ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "402": {
                        "description": "{\"success\": false, \"message\": \"评论失败\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/notice/create_question": {
            "post": {
                "description": "创建一个问题",
                "tags": [
                    "防控知识板块"
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
                        "description": "提问标题",
                        "name": "question_title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "提问内容",
                        "name": "question_content",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"用户提问成功\", \"detail\": \"提问的全部信息\"}",
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
        "/notice/list_all_comments": {
            "post": {
                "description": "列出某个问题的全部评论",
                "tags": [
                    "防控知识板块"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "问题ID",
                        "name": "question_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"查看成功\", \"data\": \"某问题的所有评论\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\": false, \"message\": \"问题ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/notice/list_all_notice": {
            "get": {
                "description": "获取所有公告，返回列表",
                "tags": [
                    "公告"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有公告\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/notice/list_all_questions": {
            "get": {
                "description": "列出全部问题",
                "tags": [
                    "防控知识板块"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"查看成功\", \"data\": \"全部问题\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/notice/list_all_rumor": {
            "get": {
                "description": "获取所有辟谣，返回列表",
                "tags": [
                    "公告"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"所有辟谣\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/notice/notice_detail": {
            "post": {
                "description": "查看单条公告",
                "tags": [
                    "公告"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "公告ID",
                        "name": "notice_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"该条公告的详细信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\":true, \"message\":\"查询失败，公告ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/notice/question_detail": {
            "post": {
                "description": "列出某个问题的详情",
                "tags": [
                    "防控知识板块"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "问题ID",
                        "name": "question_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"查看成功\", \"data\": \"某问题的所有信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\": false, \"message\": \"问题ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/notice/rumor_detail": {
            "post": {
                "description": "查看单条辟谣",
                "tags": [
                    "公告"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "辟谣ID",
                        "name": "rumor_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"该条辟谣的详细信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"success\":true, \"message\":\"查询失败，辟谣ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sub/del_sub": {
            "post": {
                "description": "删除订阅",
                "tags": [
                    "订阅城市"
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
        "/sub/list_all_subs": {
            "post": {
                "description": "获取订阅列表",
                "tags": [
                    "订阅城市"
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
        "/sub/subscribe": {
            "post": {
                "description": "订阅城市疫情信息",
                "tags": [
                    "订阅城市"
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
                        "description": "{\"success\": false, \"message\": \"已经订阅过这个城市的疫情信息\"}",
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
                    "用户管理"
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
                    "用户管理"
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
                    "用户管理"
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
                        "description": "{\"success\": false, \"message\": \"用户ID不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"success\": false, \"message\": \"数据库操作时的其他错误\"}",
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
                    "用户管理"
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
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
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
