package service

import (
	"github.com/asynccnu/food_service/model"
	"github.com/asynccnu/food_service/util"
)

//----------------------------------------------------------------//
//------------------------------model-----------------------------//

// SearchFoodModel 用于搜索返回
type SearchFoodModel struct {
	Name           string `json:"name"`
	RestaurantName string `json:"restaurant_name"`
	PictureURL     string `json:"picture_url"` //店家图片

	model.Canteen
}

// SearchRestaurantModel 用于搜索返回
type SearchRestaurantModel struct {
	model.Canteen

	Name         string `json:"name"`
	RestaurantID uint32 `json:"restaurant_id"`
	PictureURL   string `json:"picture_url"`
}

// RestaurantForCanteen 用于在线菜单中
type RestaurantForCanteen struct {
	RestaurantName string  `json:"restaurant_name"`
	AveragePrice   float32 `json:"average_price"`
	PictureURL     string  `json:"picture_url"`
	RestaurantID   uint32  `json:"restaurant_id"`
}

// FoodDetailsForRecommend 用于华师必吃
type FoodDetailsForRecommend struct {
	Name string `json:"name"`

	RestaurantName string `json:"resaurant_name"`
	model.Canteen

	Ingredient   string `json:"ingredient"`
	Introduction string `json:"introduction"`
	PictureURL   string `json:"picture_url"` //店家图片
}

// RecommendRestaurant 用于美食首页
type RecommendRestaurant struct {
	RestaurantName string `json:"restaurant_name"`

	model.Canteen

	AveragePrice    float32  `json:"average_price"`
	PictureURL      string   `json:"picture_url"`
	Recommendations []string `json:"recommendation"`
}

//----------------------------------------------------------------//
//-----------------------------小写函数----------------------------////

func getKeyWords(st string) string {
	var kws string

	if len([]rune(st)) > 1 {
		//分词
		kws = util.SegWord(st)
	}
	if kws != "" {
		kws = "'%" + st + "%' or name like " + kws
	} else {
		kws = "'%" + st + "%'"
	}

	return kws
}

//----------------------------------------------------------------//
//-------------------------------大写函数--------------------------//

// SearchForFoods 分词之后关键字用于数据库查询
func SearchForFoods(st string, page, limit uint64) (*[]SearchFoodModel, error) {
	kws := getKeyWords(st)

	foods, err := model.CRUDForSearchFoods(kws, page, limit)
	if err != nil {
		return nil, err
	}

	var Results []SearchFoodModel
	for _, food := range *foods {
		restaurant, err := model.GetRestaurantByID(food.RestaurantID)
		if err != nil {
			return nil, err
		}
		result := SearchFoodModel{
			Name:           food.Name,
			RestaurantName: restaurant.Name,
			PictureURL:     restaurant.PictureURL,
			Canteen:        *model.GetCanteen(restaurant.Location),
		}
		Results = append(Results, result)
	}
	return &Results, nil
}

// SearchForRestaurants 分词之后关键字用于数据库查询
func SearchForRestaurants(st string, page, limit uint64) (*[]SearchRestaurantModel, error) {
	kws := getKeyWords(st)

	restaurants, err := model.CRUDForSearchRestaurants(kws, page, limit)
	if err != nil {
		return nil, err
	}

	var Results []SearchRestaurantModel
	for _, restaurant := range *restaurants {
		result := SearchRestaurantModel{
			Name:         restaurant.Name,
			RestaurantID: restaurant.ID,
			PictureURL:   restaurant.PictureURL,
			Canteen:      *model.GetCanteen(restaurant.Location),
		}
		Results = append(Results, result)
	}
	return &Results, nil
}

// GetRestaurantDetailsByID 用于店家详情页.
func GetRestaurantDetailsByID(id uint32) (*model.RestaurantDetails, error) {
	Restaurant, err := model.GetRestaurantByID(id)
	if err != nil {
		return nil, err
	}

	Menus, err := model.GetMenusByRestaurantID(id)
	if err != nil {
		return nil, err
	}

	RD := model.RestaurantDetails{
		Name:         Restaurant.Name,
		Introduction: Restaurant.Introduction,
		AveragePrice: Restaurant.AveragePrice,
		PictureURL:   Restaurant.PictureURL,
		Menus:        Menus,
	}
	return &RD, err
}

// ListRestaurants 用于在线菜单
func ListRestaurants(canteenName string, page, limit uint64, storey uint8) (*[]RestaurantForCanteen, error) {
	canteenID, err := model.GetCanteenID(canteenName, storey)
	if err != nil {
		return nil, err
	}
	restaurants, err := model.CRUDForListRestaurants(canteenID, page, limit)
	if err != nil {
		return nil, err
	}

	var Result []RestaurantForCanteen
	for _, restaurant := range *restaurants {
		resForCanteen := RestaurantForCanteen{
			RestaurantName: restaurant.Name,
			PictureURL:     restaurant.PictureURL,
			RestaurantID:   restaurant.ID,
			AveragePrice:   restaurant.AveragePrice,
		}
		Result = append(Result, resForCanteen)
	}
	return &Result, nil
}

// RecommendFoods 用于华师必吃
func RecommendFoods(page, limit uint64) (*[]FoodDetailsForRecommend, error) {
	foods, err := model.CRUDForRecommendedFoods(page, limit)
	if err != nil {
		return nil, err
	}
	var Results []FoodDetailsForRecommend
	for _, food := range *foods {
		restaurant, err := model.GetRestaurantByID(food.RestaurantID)
		if err != nil {
			return nil, err
		}
		result := FoodDetailsForRecommend{
			Name:         food.Name,
			Ingredient:   food.Ingredient,
			Introduction: food.Introduction,
			PictureURL:   restaurant.PictureURL,
			//下面是食堂信息
			Canteen:        *model.GetCanteen(restaurant.Location),
			RestaurantName: restaurant.Name,
		}
		Results = append(Results, result)
	}
	return &Results, nil
}

// RecommendRestaurants 用于美食首页返回一个商家
func RecommendRestaurants(canteenName string, page, limit uint64) (*[]RecommendRestaurant, error) {
	restaurants, err := model.CRUDForRecommendedRestaurants(canteenName, page, limit)
	if err != nil {
		return nil, err
	}
	var Results []RecommendRestaurant
	for _, restaurant := range *restaurants {
		recommendations, err := model.GetRecommendationByID(restaurant.ID)
		if err != nil {
			return nil, err
		}
		result := RecommendRestaurant{
			RestaurantName:  restaurant.Name,
			AveragePrice:    restaurant.AveragePrice,
			PictureURL:      restaurant.PictureURL,
			Canteen:         *model.GetCanteen(restaurant.Location),
			Recommendations: *recommendations,
		}
		Results = append(Results, result)
	}
	return &Results, nil
}
