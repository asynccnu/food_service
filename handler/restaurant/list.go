package restaurant

import (
	"strconv"

	"github.com/asynccnu/food_service/handler"
	"github.com/asynccnu/food_service/pkg/errno"
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
//@Param c query  integer true "按照id来给,食堂信息，哪个食堂，楼层"
//@Success 200 {object} ResaurantListResponse
//@Router /restaurant/list [get]
func List(c *gin.Context) {
	canteenStr := c.DefaultQuery("c", "1")
	canteenID, err := strconv.ParseUint(canteenStr, 10, 64)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	Results, err := service.ListRestaurants(uint8(canteenID), page, limit)
	if err != nil {
		handler.SendError(c, errno.ErrCRUD, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, *Results)
}
