package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.704908 +0800 CST m=+0.108224291
type SettlementDetails struct {
    Amount    float64        `gorm:"not null;column:amount;type:double" json:"amount"`
    CreatedAt date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    ID        uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    OrderId   string         `gorm:"not null;column:order_id;type:varchar(255);size:1020" json:"orderId,omitempty"`
    UpdatedAt date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (SettlementDetails) TableName() string {
	return "settlement_details"
}
func (this SettlementDetails) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SettlementDetails) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SettlementDetails) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this SettlementDetails) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(SettlementDetails{}).Error; err != nil {
		return err
	}
	return nil
}

func NewSettlementDetails() SettlementDetails {
    return SettlementDetails{}
}

