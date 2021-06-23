package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.663815 +0800 CST m=+0.067132233
type FoodClassifications struct {
    CreatedAt  date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt  *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    FoodId     int            `gorm:"column:food_id;type:int" json:"foodId"`
    FoodTypeId int            `gorm:"column:food_type_id;type:int" json:"foodTypeId"`
    ID         uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    ShopId     int            `gorm:"column:shop_id;type:int" json:"shopId"`
    UpdatedAt  date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (FoodClassifications) TableName() string {
	return "food_classifications"
}
func (this FoodClassifications) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this FoodClassifications) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this FoodClassifications) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this FoodClassifications) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(FoodClassifications{}).Error; err != nil {
		return err
	}
	return nil
}

func NewFoodClassifications() FoodClassifications {
    return FoodClassifications{}
}

