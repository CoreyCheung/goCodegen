package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.963521 +0800 CST m=+0.075534893
type FoodClassifications struct {
    CreatedAt  date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt  *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    FoodId     int            `xorm:"'food_id' INT" json:"foodId"`
    FoodTypeId int            `xorm:"'food_type_id' INT" json:"foodTypeId"`
    ID         uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    ShopId     int            `xorm:"'shop_id' INT" json:"shopId"`
    UpdatedAt  date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (FoodClassifications) TableName() string {
	return "food_classifications"
}

func NewFoodClassifications() FoodClassifications {
    return FoodClassifications{}
}
