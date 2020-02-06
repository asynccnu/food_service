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
//@Param c query string true "cantenn_name,按照id来给,食堂信息，哪个食堂, 例如学子"
//@Param s query integer true "stoery,楼层, 例如1代表一楼"
//@Success 200 {object} ResaurantListResponse
//@Router /restaurant/list [get]
func List(c *gin.Context) {
	canteenNameStr := c.DefaultQuery("c", "学子")
	storeyStr := c.DefaultQuery("s", "1")
	storey, err := strconv.ParseUint(storeyStr, 10, 64)
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

	Results, err := service.ListRestaurants(canteenNameStr, page, limit, uint8(storey))
	if err != nil {
		handler.SendError(c, errno.ErrCRUD, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, *Results)
}
