package restaurant

import (
	"github.com/asynccnu/food_service/model"
	"github.com/gin-gonic/gin"
)

type RestaurantForCanteen struct {
	RestaurantName string `json:"restaurant_name"`
	AveragePrice   string `json:"average_price"`
	PictureURL     string `json:"picture_url"`
	RestaurantID   string `json:"restaurant_id"`
}

type ResaurantList struct {
	Restaurants *[]RestaurantForCanteen `json:"restaurant_for_canteen"`
}

type ListRequest model.Canteen

//@Tags restaurant
//@Summary 在线菜单
//@Description 返回一些推荐的食堂
//@Accept json
//@Produce json
//@Param page query integer true "页码"
//@Param limit query integer true "每页最大数"
//@Param payload body Request true "食堂信息，哪个食堂，楼层"
//@Success 200 {object} ResaurantList
//@Router /restaurant/list [get]
func List(c *gin.Context) {

}
