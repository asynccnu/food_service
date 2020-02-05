package model

import (
	"fmt"
	"strconv"
	"time"
)

// Canteen 食堂地址
type Canteen struct {
	CanteenName string `json:"canteen_name"`
	Storey      uint8  `json:"storey"`
}

// Menu 每个窗口的菜， 只有名字和价格
type Menu struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

// RestaurantDetails 店家详细信息，用于店家详情页
type RestaurantDetails struct {
	Name         string  `json:"name"`
	Introduction string  `json:"introduction"`
	AveragePrice float32 `json:"average_price"`
	PictureURL   string  `json:"picture_url"` //店家图片信息
	Menus        *[]Menu `json:"menus"`
}

// BaseModel 数据库表 model
type BaseModel struct {
	ID        uint32     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

// CanteenModel 食堂model
type CanteenModel struct {
	ID     uint16 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Name   string `gorm:"column:name" json:"canteen_name"`
	Storey uint8  `gorm:"column:storey" json:"storey"`
}

// TableName 用于gorm 表名字
func (c *CanteenModel) TableName() string {
	return "canteen"
}

// RestaurantModel 店家model
type RestaurantModel struct {
	BaseModel
	Name         string  `gorm:"column:name" json:"name"`
	Location     uint8   `gorm:"column:location" json:"location"`
	Introduction string  `gorm:"column:introduction" json:"introduction"`
	AveragePrice float32 `gorm:"column:average_price" json:"average_price"`
	PictureURL   string  `gorm:"column:picture_url" json:"picture_url"`
}

// TableName 用于gorm 表名字
func (r *RestaurantModel) TableName() string {
	return "restaurant"
}

// FoodModel 菜品model
type FoodModel struct {
	BaseModel
	Name         string  `gorm:"column:name" json:"name"`
	RestaurantID uint32  `gorm:"column:restaurant_id" json:"restaurant_id"`
	Introduction string  `gorm:"column:introduction" json:"introduction"`
	Ingredient   string  `gorm:"column:ingredient" json:"ingredient"`
	Price        float32 `gorm:"column:price" json:"price"`
	IsSpecial    bool    `gorm:"column:is_special" json:"is_special"`
}

// TableName 用于gorm 表名字
func (f *FoodModel) TableName() string {
	return "food"
}

//----------------------------------------------------------------//
// 小写函数

//----------------------------------------------------------------//
// 大写函数

// GetCanteen 获取餐厅信息
func GetCanteen(id uint8) *Canteen {
	var C CanteenModel
	DB.Self.Where("id = ?", id).Find(&C)
	return &Canteen{
		CanteenName: C.Name,
		Storey:      C.Storey,
	}
}

// GetRestaurantByID 通过窗口ID获得 食堂model
func GetRestaurantByID(id uint32) (*RestaurantModel, error) {
	var Restaurant RestaurantModel
	d := DB.Self.Where("id = ?", id).Find(&Restaurant)
	if d.Error != nil {
		return nil, d.Error
	}
	return &Restaurant, d.Error
}

/*
// GetRestaurantForFoodByID 通过窗口ID获得 食堂model,只有名字和位置
func GetRestaurantForFoodByID(id uint32) (*RestaurantModel, error) {
	var Restaurant RestaurantModel
	d := DB.Self.Where("id = ?", id).Find(&Restaurant)
	if d.Error != nil {
		return nil, d.Error
	}
	return &Restaurant, d.Error
}
*/

//GetMenusByRestaurantID 通过窗口ID获得 食堂窗口的菜单，菜单信息有名字，价格。
func GetMenusByRestaurantID(id uint32) (*[]Menu, error) {
	var Menus []Menu
	var foods []FoodModel
	d := DB.Self.Raw("select name, price from food where restaurant_id = ?", id).Scan(&foods)
	if d.Error != nil {
		return nil, d.Error
	}
	for _, food := range foods {
		Menus = append(Menus, Menu{Name: food.Name, Price: food.Price})
	}
	return &Menus, d.Error
}

// CRUDForSearchFoods 用于foods数据库查询
func CRUDForSearchFoods(kws string, page, limit uint64) (*[]FoodModel, error) {
	sql := "select restaurant_id, name from food where name like " + kws + " limit " + strconv.Itoa(int((page-1)*limit)) + ", " + strconv.Itoa(int(limit))
	var Foods []FoodModel
	d := DB.Self.Raw(sql).Scan(&Foods)
	if d.Error != nil {
		return nil, d.Error
	}
	return &Foods, nil
}

// CRUDForSearchRestaurants 用于resta 数据库查询
func CRUDForSearchRestaurants(kws string, page, limit uint64) (*[]RestaurantModel, error) {
	sql := "select id, name, picture_url, location from restaurant where name like " + kws + " limit " + strconv.Itoa(int((page-1)*limit)) + ", " + strconv.Itoa(int(limit))
	var Restaurants []RestaurantModel
	d := DB.Self.Raw(sql).Scan(&Restaurants)
	if d.Error != nil {
		return nil, d.Error
	}
	return &Restaurants, nil
}

// CRUDForListRestaurants 用于在线菜单给出一些食堂
func CRUDForListRestaurants(canteenID uint8, page, limit uint64) (*[]RestaurantModel, error) {
	sql := fmt.Sprintf("select name, picture_url, average_price, id from restaurant where location = %d order by hot desc limit %d, %d;", canteenID, (page-1)*limit, limit)
	var Restaurants []RestaurantModel
	d := DB.Self.Raw(sql).Scan(&Restaurants)
	if d.Error != nil {
		return nil, d.Error
	}
	return &Restaurants, nil
}

// CRUDForSpecialFoods 用于特色推荐
func CRUDForSpecialFoods(restaurantID uint32) {

}

// CRUDForRecommendedFoods 用于华师必吃
func CRUDForRecommendedFoods(page, limit uint64) (*[]FoodModel, error) {
	sql := fmt.Sprintf("select name, ingredient, introduction, restaurant_id from food order by hot desc limit %d, %d;", (page-1)*limit, limit)
	var Foods []FoodModel
	d := DB.Self.Raw(sql).Scan(&Foods)
	if d.Error != nil {
		return nil, d.Error
	}
	return &Foods, nil
}

// func CRUDForRandomRestaurant(page)
