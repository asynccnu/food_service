// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-02-06 14:21:00.369948331 +0800 CST m=+0.050982201

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/food/recommend": {
            "get": {
                "description": "返回一些推荐的food",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "food"
                ],
                "summary": "华师必吃",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "每页最大数",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/food.FoodList"
                        }
                    }
                }
            }
        },
        "/restaurant/detail/{id}": {
            "get": {
                "description": "给店家id，返回店家详情页",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "restaurant"
                ],
                "summary": "店家详情页",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "店家的id，别的api会给出。",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.RestaurantDetails"
                        }
                    }
                }
            }
        },
        "/restaurant/list": {
            "get": {
                "description": "返回一些推荐的食堂",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "restaurant"
                ],
                "summary": "在线菜单",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "每页最大数",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "cantenn_name,按照id来给,食堂信息，哪个食堂, 例如学子",
                        "name": "c",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "stoery,楼层, 例如1代表一楼",
                        "name": "s",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/restaurant.ResaurantListResponse"
                        }
                    }
                }
            }
        },
        "/restaurant/recommend": {
            "get": {
                "description": "美食首页的推荐窗口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "restaurant"
                ],
                "summary": "美食首页",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "第几页， page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "每页多少个， limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "食堂名字, 例如东一，学子",
                        "name": "canteen_name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.RecommendRestaurant"
                        }
                    }
                }
            }
        },
        "/search/food": {
            "get": {
                "description": "搜索返回一个list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "搜索食物",
                "parameters": [
                    {
                        "type": "string",
                        "description": "搜索信息, search_text",
                        "name": "st",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "第几页， page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "每页多少个， limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/search.SearchFoodResponse"
                        }
                    }
                }
            }
        },
        "/search/restaurant": {
            "get": {
                "description": "搜索返回一个list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "搜索餐厅",
                "parameters": [
                    {
                        "type": "string",
                        "description": "搜索信息, search_text",
                        "name": "st",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "第几页， page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "每页多少个， limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/search.SearchRestaurantResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "food.FoodList": {
            "type": "object",
            "properties": {
                "food_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.FoodDetailsForRecommend"
                    }
                }
            }
        },
        "model.Menu": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "model.RestaurantDetails": {
            "type": "object",
            "properties": {
                "average_price": {
                    "type": "number"
                },
                "introduction": {
                    "type": "string"
                },
                "menus": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Menu"
                    }
                },
                "name": {
                    "type": "string"
                },
                "picture_url": {
                    "description": "店家图片信息",
                    "type": "string"
                }
            }
        },
        "restaurant.ResaurantListResponse": {
            "type": "object",
            "properties": {
                "restaurants": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.RestaurantForCanteen"
                    }
                }
            }
        },
        "search.SearchFoodResponse": {
            "type": "object",
            "properties": {
                "results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.SearchFoodModel"
                    }
                }
            }
        },
        "search.SearchRestaurantResponse": {
            "type": "object",
            "properties": {
                "results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.SearchRestaurantModel"
                    }
                }
            }
        },
        "service.FoodDetailsForRecommend": {
            "type": "object",
            "properties": {
                "canteen_name": {
                    "type": "string"
                },
                "ingredient": {
                    "type": "string"
                },
                "introduction": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "picture_url": {
                    "description": "店家图片",
                    "type": "string"
                },
                "resaurant_name": {
                    "type": "string"
                },
                "storey": {
                    "type": "integer"
                }
            }
        },
        "service.RecommendRestaurant": {
            "type": "object",
            "properties": {
                "average_price": {
                    "type": "number"
                },
                "canteen_name": {
                    "type": "string"
                },
                "picture_url": {
                    "type": "string"
                },
                "recommendation": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "restaurant_name": {
                    "type": "string"
                },
                "storey": {
                    "type": "integer"
                }
            }
        },
        "service.RestaurantForCanteen": {
            "type": "object",
            "properties": {
                "average_price": {
                    "type": "number"
                },
                "picture_url": {
                    "type": "string"
                },
                "restaurant_id": {
                    "type": "integer"
                },
                "restaurant_name": {
                    "type": "string"
                }
            }
        },
        "service.SearchFoodModel": {
            "type": "object",
            "properties": {
                "canteen_name": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "picture_url": {
                    "description": "店家图片",
                    "type": "string"
                },
                "restaurant_name": {
                    "type": "string"
                },
                "storey": {
                    "type": "integer"
                }
            }
        },
        "service.SearchRestaurantModel": {
            "type": "object",
            "properties": {
                "canteen_name": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "picture_url": {
                    "type": "string"
                },
                "restaurant_id": {
                    "type": "integer"
                },
                "storey": {
                    "type": "integer"
                }
            }
        }
    },
    "tags": [
        {
            "description": "店铺(窗口)相关",
            "name": "restaurant"
        },
        {
            "description": "菜品相关",
            "name": "food"
        },
        {
            "description": "搜索相关",
            "name": "search"
        },
        {
            "description": "食堂相关",
            "name": "canteen"
        }
    ]
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
	Host:        "....",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "food_service",
	Description: "美食服务",
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
