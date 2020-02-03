package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/asynccnu/food_service/util"
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
	PictureURL   string  `json:"picture_url"`
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
	ID     uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
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
	RestaurantID uint32  `gorm:"column:restaurant" json:"restaurant_id"`
	Introduction string  `gorm:"column:introduction" json:"introduction"`
	Ingredient   string  `gorm:"column:ingredient" json:"ingredient"`
	Price        float32 `gorm:"column:price" json:"price"`
	PictureURL   string  `gorm:"column:picture_url" json:"picture_url"`
}

// TableName 用于gorm 表名字
func (f *FoodModel) TableName() string {
	return "food"
}

//----------------------------------------------------------------//
// 小写函数
func getCanteen(id uint8) *Canteen {
	var C CanteenModel
	DB.Self.Where("id = ?", id).Find(&C)
	return &Canteen{
		CanteenName: C.Name,
		Storey:      C.Storey,
	}
}

//----------------------------------------------------------------//
// 大写函数

// GetRestaurantByID 通过窗口ID获得 食堂model
func GetRestaurantByID(id uint32) (*RestaurantModel, error) {
	var Restaurant RestaurantModel
	d := DB.Self.Where("id = ?", id).Find(&Restaurant)
	if d.Error != nil {
		return nil, d.Error
	}
	return &Restaurant, d.Error
}

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

// SearchFoodModel 用于搜索返回
type SearchFoodModel struct {
	Name           string `json:"name"`
	RestaurantName string `json:"restaurant_name"`
	PictureURL     string `json:"picture_url"`

	Canteen
}

// SearchForFoods 分词之后关键字用于数据库查询
func SearchForFoods(st string, page, limit uint64) (*[]SearchFoodModel, error) {
	var kws string
	if len([]rune(st)) > 1 {
		//分词
		kws = util.SegWord(st)
		if kws == "" {
			err := fmt.Errorf("搜索词语过于简单")
			return nil, err
		}
	} else {
		kws = "'%" + st + "%'"
	}
	sql := "select restaurant_id, name, picture_url from food where name like " + kws + " limit " + strconv.Itoa(int((page-1)*limit)) + ", " + strconv.Itoa(int(limit))
	var foods []FoodModel
	d := DB.Self.Raw(sql).Scan(&foods)
	if d.Error != nil {
		return nil, d.Error
	}
	var Results []SearchFoodModel
	for _, food := range foods {
		restaurant, _ := GetRestaurantByID(food.RestaurantID)
		result := SearchFoodModel{
			Name:           food.Name,
			RestaurantName: restaurant.Name,
			PictureURL:     food.PictureURL,
			Canteen:        *getCanteen(restaurant.Location),
		}
		Results = append(Results, result)
	}
	return &Results, nil
}

// SearchRestaurantModel 用于搜索返回
type SearchRestaurantModel struct {
	Canteen

	Name         string `json:"name"`
	RestaurantID uint32 `json:"restaurant_id"`
	PictureURL   string `json:"picture_url"`
}

// SearchForRestaurants 分词之后关键字用于数据库查询
func SearchForRestaurants(st string, page, limit uint64) (*[]SearchRestaurantModel, error) {
	var kws string
	if len([]rune(st)) > 1 {
		kws = util.SegWord(st)
		if kws == "" {
			err := fmt.Errorf("搜索词语过于简单")
			return nil, err
		}
	} else {
		kws = "'%" + st + "%'"
	}
	sql := "select id, name, picture_url, location from restaurant where name like " + kws + " limit " + strconv.Itoa(int((page-1)*limit)) + ", " + strconv.Itoa(int(limit))
	var restaurants []RestaurantModel
	d := DB.Self.Raw(sql).Scan(&restaurants)
	if d.Error != nil {
		return nil, d.Error
	}
	var Results []SearchRestaurantModel
	for _, restaurant := range restaurants {
		result := SearchRestaurantModel{
			Name:         restaurant.Name,
			RestaurantID: restaurant.ID,
			PictureURL:   restaurant.PictureURL,
			Canteen:      *getCanteen(restaurant.Location),
		}
		Results = append(Results, result)
	}
	return &Results, nil
}
