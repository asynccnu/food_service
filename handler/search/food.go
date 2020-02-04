package search

import (
	"strconv"

	"github.com/asynccnu/food_service/handler"
	"github.com/asynccnu/food_service/pkg/errno"
	"github.com/asynccnu/food_service/service"
	"github.com/gin-gonic/gin"
)

type SearchFoodResponse struct {
	Results []service.SearchFoodModel `json:"results"`
}

//@Tags search
//@Summary 搜索食物
//@Description 搜索返回一个list
//@Accept json
//@Produce json
//@Param st query string true "搜索信息, search_text"
//@Param page query integer true "第几页， page"
//@Param limit query integer true "每页多少个， limit"
//@Success 200 {object} SearchFoodResponse
//@Router /search/food [get]
func SearchFood(c *gin.Context) {
	searchText := c.DefaultQuery("st", "Search")
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

	Results, err := service.SearchForFoods(searchText, page, limit)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrCRUD, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, *Results)
}
