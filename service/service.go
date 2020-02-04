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
	PictureURL     string `json:"picture_url"`

	model.Canteen
}

// SearchRestaurantModel 用于搜索返回
type SearchRestaurantModel struct {
	model.Canteen

	Name         string `json:"name"`
	RestaurantID uint32 `json:"restaurant_id"`
	PictureURL   string `json:"picture_url"`
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
			PictureURL:     food.PictureURL,
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
