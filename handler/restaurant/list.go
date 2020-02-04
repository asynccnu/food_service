package restaurant

import (
	"github.com/asynccnu/food_service/service"
	"github.com/gin-gonic/gin"
)

type ResaurantListResponse struct {
	Restaurants *[]service.RestaurantForCanteen `json:"restaurants"`
}

//@Tags restaurant
//@Summary 在线菜单
//@Description 返回一些推荐的食堂
//@Accept json
//@Produce json
//@Param page query integer true "页码"
//@Param limit query integer true "每页最大数"
//@Param canteen body integer true "按照id来给,食堂信息，哪个食堂，楼层"
//@Success 200 {object} ResaurantListResponse
//@Router /restaurant/list [get]
func List(c *gin.Context) {

}
