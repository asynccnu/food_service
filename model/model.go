package model

import (
	"fmt"
	"math/rand"
	"strings"
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
	Location     uint16  `gorm:"column:location" json:"location"`
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

// GetCanteenIDs 通过食堂名字(多个名字) 返回id，但是有多个id
func getCanteenIDs(canteenName string) (string, error) {
	var canteens []CanteenModel
	canteenNameSlice := strings.Split(canteenName, ",")
	//sqlCanteenNameSlice 为加了''的餐厅名字，用于sql语句中
	var sqlCanteenNameSlice []string
	for _, canteen := range canteenNameSlice {
		sql := fmt.Sprintf("'%s'", canteen)
		sqlCanteenNameSlice = append(sqlCanteenNameSlice, sql)
	}
	sqlCanteenName := strings.Join(sqlCanteenNameSlice, ",")
	sql := fmt.Sprintf("select id from canteen where name in (%s)", sqlCanteenName)
	d := DB.Self.Raw(sql).Scan(&canteens)
	if d.Error != nil {
		return "", d.Error
	}
	var Result string
	for _, c := range canteens {
		Result += fmt.Sprintf("%d, ", c.ID)
	}
	Result = strings.TrimSuffix(Result, ", ")
	return Result, nil
}

//----------------------------------------------------------------//
// 大写函数

// GetRecommendationByID 用于特色推荐
func GetRecommendationByID(restaurantID uint32) (*[]string, error) {
	sql := fmt.Sprintf("select name from food where restaurant_id = %d and is_special = 1", restaurantID)
	var foods []FoodModel
	d := DB.Self.Raw(sql).Scan(&foods)
	if d.Error != nil {
		return nil, d.Error
	}
	var Recommendations []string
	for _, food := range foods {
		Recommendations = append(Recommendations, food.Name)
	}
	return &Recommendations, nil
}

// GetCanteen 获取餐厅信息
func GetCanteen(id uint16) *Canteen {
	var C CanteenModel
	DB.Self.Where("id = ?", id).Find(&C)
	return &Canteen{
		CanteenName: C.Name,
		Storey:      C.Storey,
	}
}

// GetCanteenID 通过食堂名字和楼层返回canteenID
func GetCanteenID(canteenName string, storey uint8) (uint16, error) {
	var C CanteenModel
	sql := fmt.Sprintf("select id from canteen where name = '%s' and storey = %d", canteenName, storey)
	d := DB.Self.Raw(sql).Scan(&C)
	if d.Error != nil {
		return 0, d.Error
	}
	return C.ID, nil
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
	//sql := fmt.Sprintf("select restaurant_id, name from food where name like %s order by hot desc limit %d,%d", kws, (page-1)*limit, limit)
	// 替换为全文索引
	sql := fmt.Sprintf("select restaurant_id, name from food where match(name) against('%s') order by hot desc limit %d,%d", kws, (page-1)*limit, limit)
	var Foods []FoodModel
	d := DB.Self.Raw(sql).Scan(&Foods)
	if d.Error != nil {
		return nil, d.Error
	}
	return &Foods, nil
}

// CRUDForSearchRestaurants 用于restaurants数据库查询
func CRUDForSearchRestaurants(kws string, page, limit uint64) (*[]RestaurantModel, error) {
	//sql := fmt.Sprintf("select id, name, picture_url, location from restaurant where match(name) against('%s') order by hot desc limit %d,%d", kws, (page-1)*limit, limit)
	// 替换为全文索引
	sql := fmt.Sprintf("select id, name, picture_url, location from restaurant where match(name) against('%s') order by hot desc limit %d,%d", kws, (page-1)*limit, limit)
	var Restaurants []RestaurantModel
	d := DB.Self.Raw(sql).Scan(&Restaurants)
	if d.Error != nil {
		return nil, d.Error
	}
	return &Restaurants, nil
}

// CRUDForListRestaurants 用于在线菜单给出一些食堂
func CRUDForListRestaurants(canteenID uint16, page, limit uint64) (*[]RestaurantModel, error) {
	sql := fmt.Sprintf("select name, picture_url, average_price, id from restaurant where location = %d order by hot desc limit %d, %d;", canteenID, (page-1)*limit, limit)
	var Restaurants []RestaurantModel
	d := DB.Self.Raw(sql).Scan(&Restaurants)
	if d.Error != nil {
		return nil, d.Error
	}
	return &Restaurants, nil
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

// CRUDForRecommendedRestaurants 用于首页推荐
func CRUDForRecommendedRestaurants(canteenName string, limit uint64) (*[]RestaurantModel, error) {
	canteenIDs, err := getCanteenIDs(canteenName)
	if err != nil {
		return nil, err
	}
	seed := rand.NewSource(time.Now().Unix())
	r := rand.New(seed)
	sql := fmt.Sprintf("select name, id, picture_url, average_price, location from restaurant where location in (%s) order by hot desc limit %d, %d", canteenIDs, r.Intn(30), limit)
	var Restaurants []RestaurantModel
	d := DB.Self.Raw(sql).Scan(&Restaurants)
	if d.Error != nil {
		return nil, d.Error
	}
	return &Restaurants, nil
}
