package food

import (
	"strconv"

	"github.com/asynccnu/food_service/handler"
	"github.com/asynccnu/food_service/pkg/errno"
	"github.com/asynccnu/food_service/service"
	"github.com/gin-gonic/gin"
)

type FoodList struct {
	FoodList *[]service.FoodDetailsForRecommend `json:"food_list"`
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
	Results, err := service.RecommendFoods(page, limit)
	if err != nil {
		handler.SendError(c, errno.ErrCRUD, nil, err.Error())
	}
	handler.SendResponse(c, nil, *Results)
}
