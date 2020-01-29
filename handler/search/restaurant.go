package search

import "github.com/gin-gonic/gin"

import "github.com/asynccnu/food_service/model"

type SearchRestaurantModel struct {
	model.Canteen

	Name         string `json:"name"`
	RestaurantID uint32 `json:"restaurant_id"`
	PictureURL   string `json:"picture_url"`
}

type SearchRestaurantList struct {
	Results []SearchRestaurantModel `json:"results"`
}

//@Tags search
//@Summary 搜索餐厅
//@Description 搜索返回一个list
//@Accept json
//@Produce json
//@Param search_text query string true "搜索信息"
//@Success 200 {object} SearchRestaurantList
//@Router /search/restaurant [get]
func SearchRestaurant(c *gin.Context) {

}
