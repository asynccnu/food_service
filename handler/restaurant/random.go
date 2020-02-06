package restaurant

import (
	"github.com/gin-gonic/gin"
)

//@Tags restaurant
//@Summary 美食首页
//@Description 美食首页的推荐窗口
//@Accept json
//@Produce json
//@Param canteen_name body string true "食堂名字, 例如东一，学子"
//@Success 200 {object} service.RandomRestaurant
//@Router /restaurant/random [get]
func Random(c *gin.Context) {

}
