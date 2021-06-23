package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.984826 +0800 CST m=+0.096840127
type Menus struct {
    CreatedAt      date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt      *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    FirstClassify  string         `xorm:"'first_classify' VARCHAR(255)" json:"firstClassify,omitempty"`
    FoodName       string         `xorm:"'food_name' VARCHAR(255)" json:"foodName,omitempty"`
    ID             uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    ImgUrl         string         `xorm:"'img_url' VARCHAR(255)" json:"imgUrl,omitempty"`
    Ingredients    string         `xorm:"'ingredients' VARCHAR(255)" json:"ingredients,omitempty"`
    SecondClassify string         `xorm:"'second_classify' VARCHAR(255)" json:"secondClassify,omitempty"`
    ThirdClassify  string         `xorm:"'third_classify' VARCHAR(255)" json:"thirdClassify,omitempty"`
    UpdatedAt      date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (Menus) TableName() string {
	return "menus"
}

func NewMenus() Menus {
    return Menus{}
}
