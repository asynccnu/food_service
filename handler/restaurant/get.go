package restaurant

import (
	"strconv"

	"github.com/asynccnu/food_service/handler"
	"github.com/asynccnu/food_service/pkg/errno"
	"github.com/asynccnu/food_service/service"
	"github.com/gin-gonic/gin"
)

//Get router
//@Tags restaurant
//@Summary 店家详情页
//@Description 给店家id，返回店家详情页
//@Accept json
//@Produce json
//@Param id path int true "店家的id，别的api会给出。"
//@Success 200 {object} model.RestaurantDetails
//@Router /restaurant/detail/{id} [get]
func Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	//RD为窗口详细信息
	RestaurantDetails, err := service.GetRestaurantDetailsByID(uint32(id))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrCRUD, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, *RestaurantDetails)
}
