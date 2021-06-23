package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.664651 +0800 CST m=+0.067967659
type FoodPractice struct {
    CreatedAt       date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt       *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    // 菜品id
    FoodId          int            `gorm:"not null;column:food_id;type:int" json:"foodId"`
    // 做法
    FoodPractice    string         `gorm:"not null;column:food_practice;type:varchar(255);size:1020" json:"foodPractice,omitempty"`
    // 做法id
    ID              int            `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int" json:"id"`
    // 做法选项
    PracticeOptions string         `gorm:"not null;column:practice_options;type:varchar(255);size:1020" json:"practiceOptions,omitempty"`
    UpdatedAt       date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (FoodPractice) TableName() string {
	return "food_practice"
}
func (this FoodPractice) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this FoodPractice) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this FoodPractice) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this FoodPractice) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(FoodPractice{}).Error; err != nil {
		return err
	}
	return nil
}

func NewFoodPractice() FoodPractice {
    return FoodPractice{}
}

