package restaurant

import (
	"github.com/asynccnu/food_service/model"
	"github.com/gin-gonic/gin"
)

type RandomRestaurant struct {
	RestaurantName string `json:"restaurant_name"`

	model.Canteen

	AveragePrice   float64 `json:"average_price"`
	PictureURL     string  `json:"picture_url"`
	Recommendation string  `json:"recommendation"`
}

type RandomRequest struct {
	RestaurantName *[]string `json:"restaurant_name"`
}

//@Tags restaurant
//@Summary 美食首页
//@Description 美食首页的推荐窗口
//@Accept json
//@Produce json
//@Param canteen_name body string true "食堂名字"
//@Success 200 {object} RandomRestaurant
//@Router /restaurant/random [get]
func Random(c *gin.Context) {

}
