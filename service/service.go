package service

import (
	"github.com/asynccnu/food_service/model"
)

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
