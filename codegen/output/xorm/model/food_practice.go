package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.964177 +0800 CST m=+0.076190853
type FoodPractice struct {
    CreatedAt       date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt       *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    // 菜品id
    FoodId          int            `xorm:"'food_id' not null INT" json:"foodId"`
    // 做法
    FoodPractice    string         `xorm:"'food_practice' not null VARCHAR(255)" json:"foodPractice,omitempty"`
    // 做法id
    ID              int            `xorm:"'id' not null pk autoincr INT" json:"id"`
    // 做法选项
    PracticeOptions string         `xorm:"'practice_options' not null VARCHAR(255)" json:"practiceOptions,omitempty"`
    UpdatedAt       date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (FoodPractice) TableName() string {
	return "food_practice"
}

func NewFoodPractice() FoodPractice {
    return FoodPractice{}
}
