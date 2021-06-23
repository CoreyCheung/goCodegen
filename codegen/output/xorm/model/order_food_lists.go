package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.986676 +0800 CST m=+0.098689902
type OrderFoodLists struct {
    CreatedAt    date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt    *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    FoodId       int            `xorm:"'food_id' INT" json:"foodId"`
    FoodName     string         `xorm:"'food_name' VARCHAR(100)" json:"foodName,omitempty"`
    FoodPractice string         `xorm:"'food_practice' VARCHAR(255)" json:"foodPractice,omitempty"`
    ID           uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    IsFinish     int            `xorm:"'is_finish' INT" json:"isFinish"`
    Number       int            `xorm:"'number' INT" json:"number"`
    OrderId      string         `xorm:"'order_id' VARCHAR(255)" json:"orderId,omitempty"`
    Price        float64        `xorm:"'price' DECIMAL(15,2)" json:"price"`
    TotalPrice   float64        `xorm:"'total_price' DOUBLE" json:"totalPrice"`
    UpdatedAt    date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (OrderFoodLists) TableName() string {
	return "order_food_lists"
}

func NewOrderFoodLists() OrderFoodLists {
    return OrderFoodLists{}
}
