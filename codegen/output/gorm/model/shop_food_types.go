package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.705873 +0800 CST m=+0.109189256
type ShopFoodTypes struct {
    CreatedAt    date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt    *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    Description  string         `gorm:"not null;column:description;type:varchar(128);size:512" json:"description,omitempty"`
    FoodTypeName string         `gorm:"column:food_type_name;type:varchar(255);size:1020" json:"foodTypeName,omitempty"`
    ID           uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    Rank         int            `gorm:"column:rank;type:int" json:"rank"`
    ShopId       int            `gorm:"column:shop_id;type:int" json:"shopId"`
    UpdatedAt    date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (ShopFoodTypes) TableName() string {
	return "shop_food_types"
}
func (this ShopFoodTypes) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ShopFoodTypes) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ShopFoodTypes) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this ShopFoodTypes) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(ShopFoodTypes{}).Error; err != nil {
		return err
	}
	return nil
}

func NewShopFoodTypes() ShopFoodTypes {
    return ShopFoodTypes{}
}

