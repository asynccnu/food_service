package restaurant

import "github.com/gin-gonic/gin"

type Menu struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type RestaurantDetails struct {
	Name         string  `json:"name"`
	Introduction string  `json:"introduction"`
	AveragePrice string  `json:"average_price"`
	PictureURL   string  `json:"picture_url"`
	Menus        *[]Menu `json:"menus"`
}

//@Tags restaurant
//@Summary 店家详情页
//@Description 给店家id，返回店家详情页
//@Accept json
//@Produce json
//@Param id path int true "店家的id，别的api会给出。"
//@Success 200 {object} RestaurantDetails
//@Router /restaurant/detail/{id} [get]
func Get(c *gin.Context) {

}
