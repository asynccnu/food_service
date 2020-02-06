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
//@Param page query integer true "第几页， page"
//@Param limit query integer true "每页多少个， limit"
//@Param canteen_name body string true "食堂名字, 例如东一，学子"
//@Success 200 {object} service.RecommendRestaurant
//@Router /restaurant/recommend [get]
func Recommend(c *gin.Context) {
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
	var r Request
	if err := c.ShouldBindJSON(&r); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}
	Results, err := service.RecommendRestaurants(r.CanteenName, page, limit)
	if err != nil {
		handler.SendError(c, errno.ErrCRUD, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, *Results)
}
