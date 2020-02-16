package search

import (
	"github.com/asynccnu/food_service/handler"
	"github.com/asynccnu/food_service/pkg/errno"
	"github.com/asynccnu/food_service/service"
	"github.com/gin-gonic/gin"
)

type HotSearchResponse struct {
	Result []string `json:"result"`
}

//@Tags search
//@Summary 热门搜索
//@Description 搜索返回字符串数组
//@Accept json
//@Produce json
//@Success 200 {object} HotSearchResponse
//@Router /search/hot [get]
func HotSearch(c *gin.Context) {
	result, err := service.GetHotSearch()
	if err != nil {
		handler.SendError(c, errno.ErrCRUD, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, result)
	return
}
