package search

import "github.com/gin-gonic/gin"

import "github.com/asynccnu/food_service/model"

type SearchFoodModel struct {
	Name           string `json:"name"`
	RestaurantName string `json:"restaurant_name"`
	PictureURL     string `json:"picture_url"`

	model.Canteen
}

type SearchFoodList struct {
	Results []SearchFoodModel `json:"results"`
}

//@Tags search
//@Summary 搜索食物
//@Description 搜索返回一个list
//@Accept json
//@Produce json
//@Param search_text query string true "搜索信息"
//@Success 200 {object} SearchFoodList
//@Router /search/food [get]
func SearchFood(c *gin.Context) {

}
