package model

import (
	"sync"
	"time"
)

type Canteen struct {
	CanteenName string `json:"canteen_name"`
	Storey      uint8  `json:"storey"`
}

type BaseModel struct {
	Id        uint32     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

type RestaurantModel struct {
	BaseModel
	Name         string  `gorm:"column:name" json:"name"`
	Location     uint8   `gorm:"column:location" json:"location"`
	Introduction string  `gorm:"column:introduction" json:"introduction"`
	AveragePrice float64 `gorm:"column:average_price" json:"average_price"`
}

type FoodModel struct {
	BaseModel
	Name         string  `gorm:"column:name" json:"name"`
	RestaurantID uint32  `gorm:"column:restaurant" json:"restaurant_id"`
	Location     uint8   `gorm:"column:location" json:"location"`
	Introduction string  `gorm:"column:introduction" json:"introduction"`
	Ingredient   string  `gorm:"column:ingredient" json:"ingredient"`
	Price        float32 `gorm:"column:price" json:"price"`
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"sayHello"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}
