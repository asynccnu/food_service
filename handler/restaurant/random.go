package restaurant

import (
	"strconv"

	"github.com/asynccnu/food_service/handler"
	"github.com/asynccnu/food_service/pkg/errno"
	"github.com/asynccnu/food_service/service"
	"github.com/gin-gonic/gin"
)

type Request struct {
	CanteenName string `json:"canteen_name"`
}

//@Tags restaurant
//@Summary 美食首页
//@Description 美食首页的推荐窗口
//@Accept json
//@Produce json
//@Param limit query integer true "每页多少个， limit, 默认为一个"
//@Param canteen_name body string true "食堂名字, 例如东一，学子   注意逗号之间不能有空格，并且为英文逗号"
//@Success 200 {object} service.RecommendRestaurant
//@Router /restaurant/random [get]
func Random(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "1")
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}
	var r Request
	if err := c.ShouldBindJSON(&r); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}
	Results, err := service.RecommendRestaurants(r.CanteenName, limit)
	if err != nil {
		handler.SendError(c, errno.ErrCRUD, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, *Results)
}
