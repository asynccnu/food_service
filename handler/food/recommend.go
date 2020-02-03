package food

import (
	"github.com/asynccnu/food_service/model"
	"github.com/gin-gonic/gin"
)

type FoodDetails struct {
	Name string `json:"name"`

	ResaurantName string `json:"resaurant_name"`
	model.Canteen

	Ingredient   string `json:"ingredient"`
	Introduction string `json:"introduction"`
	PictureURL   string `json:"picture_url"`
}

type FoodList struct {
	FoodList *[]FoodDetails `json:"food_list"`
}

//@Tags food
//@Summary 华师必吃
//@Description 返回一些推荐的food
//@Accept json
//@Produce json
//@Param page query integer true "页码"
//@Param limit query integer true "每页最大数"
//@Success 200 {object} FoodList
//@Router /food/recommend [get]
func Recommend(c *gin.Context) {

}
