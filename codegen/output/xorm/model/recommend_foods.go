package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.992863 +0800 CST m=+0.104876651
type RecommendFoods struct {
    CreatedAt        date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt        *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    Description      string         `xorm:"'description' not null VARCHAR(512)" json:"description,omitempty"`
    FoodId           int            `xorm:"'food_id' INT" json:"foodId"`
    ID               uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    RecommendPicture string         `xorm:"'recommend_picture' VARCHAR(255)" json:"recommendPicture,omitempty"`
    ShopId           int            `xorm:"'shop_id' INT" json:"shopId"`
    UpdatedAt        date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (RecommendFoods) TableName() string {
	return "recommend_foods"
}

func NewRecommendFoods() RecommendFoods {
    return RecommendFoods{}
}
