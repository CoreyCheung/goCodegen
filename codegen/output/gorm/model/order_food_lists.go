package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.694206 +0800 CST m=+0.097522045
type OrderFoodLists struct {
    CreatedAt    date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt    *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    FoodId       int            `gorm:"column:food_id;type:int" json:"foodId"`
    FoodName     string         `gorm:"column:food_name;type:varchar(100);size:400" json:"foodName,omitempty"`
    FoodPractice string         `gorm:"column:food_practice;type:varchar(255);size:1020" json:"foodPractice,omitempty"`
    ID           uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    IsFinish     int            `gorm:"column:is_finish;type:int" json:"isFinish"`
    Number       int            `gorm:"column:number;type:int" json:"number"`
    OrderId      string         `gorm:"column:order_id;type:varchar(255);size:1020" json:"orderId,omitempty"`
    Price        float64        `gorm:"column:price;type:decimal(15,2)" json:"price"`
    TotalPrice   float64        `gorm:"column:total_price;type:double" json:"totalPrice"`
    UpdatedAt    date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (OrderFoodLists) TableName() string {
	return "order_food_lists"
}
func (this OrderFoodLists) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this OrderFoodLists) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this OrderFoodLists) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this OrderFoodLists) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(OrderFoodLists{}).Error; err != nil {
		return err
	}
	return nil
}

func NewOrderFoodLists() OrderFoodLists {
    return OrderFoodLists{}
}

