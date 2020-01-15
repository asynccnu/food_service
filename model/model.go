package model

import (
	"sync"
	"time"
)

type BaseModel struct {
	Id        uint32     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

type RestaurantModel struct {
	BaseModel
	Name         string `gorm:"column:name" json:"name"`
	Location     uint8  `gorm:"column:location" json:"location"`
	Introduction string `gorm:"column:introduction" json:"introduction"`
	SalesVolumn  uint32 `gorm:"column:sales_volumn" json:"sales_volumn"`
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
