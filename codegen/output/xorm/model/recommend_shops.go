package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.994363 +0800 CST m=+0.106376481
type RecommendShops struct {
    BeginTime        date.Datetime  `xorm:"'begin_time' TIMESTAMP" json:"beginTime,omitempty"`
    CreatedAt        date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt        *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    Description      string         `xorm:"'description' not null VARCHAR(1024)" json:"description,omitempty"`
    EndTime          date.Datetime  `xorm:"'end_time' TIMESTAMP" json:"endTime,omitempty"`
    ID               uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    RecommendPicture string         `xorm:"'recommend_picture' VARCHAR(255)" json:"recommendPicture,omitempty"`
    RecommendType    int            `xorm:"'recommend_type' INT" json:"recommendType"`
    ShopId           int            `xorm:"'shop_id' INT" json:"shopId"`
    UpdatedAt        date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (RecommendShops) TableName() string {
	return "recommend_shops"
}

func NewRecommendShops() RecommendShops {
    return RecommendShops{}
}
