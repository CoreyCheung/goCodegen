package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:35.00311 +0800 CST m=+0.115123697
type ShopFoodTypes struct {
    CreatedAt    date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt    *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    Description  string         `xorm:"'description' not null VARCHAR(128)" json:"description,omitempty"`
    FoodTypeName string         `xorm:"'food_type_name' VARCHAR(255)" json:"foodTypeName,omitempty"`
    ID           uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    Rank         int            `xorm:"'rank' INT" json:"rank"`
    ShopId       int            `xorm:"'shop_id' INT" json:"shopId"`
    UpdatedAt    date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (ShopFoodTypes) TableName() string {
	return "shop_food_types"
}

func NewShopFoodTypes() ShopFoodTypes {
    return ShopFoodTypes{}
}
