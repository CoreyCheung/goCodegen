package model

import "git.wanpinghui.com/WPH/go_common/wph"
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.701891 +0800 CST m=+0.105206883
type ReconciliationDetails struct {
    CreatedAt          date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt          *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    ID                 uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    InconsistentReason string         `gorm:"column:inconsistent_reason;type:varchar(255);size:1020" json:"inconsistentReason,omitempty"`
    OrderId            string         `gorm:"column:order_id;type:varchar(255);size:1020" json:"orderId,omitempty"`
    PaymentType        string         `gorm:"column:payment_type;type:varchar(255);size:1020" json:"paymentType,omitempty"`
    States             wph.Long       `gorm:"column:states;type:bigint" json:"states"`
    UpdatedAt          date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (ReconciliationDetails) TableName() string {
	return "reconciliation_details"
}
func (this ReconciliationDetails) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ReconciliationDetails) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ReconciliationDetails) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this ReconciliationDetails) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(ReconciliationDetails{}).Error; err != nil {
		return err
	}
	return nil
}

func NewReconciliationDetails() ReconciliationDetails {
    return ReconciliationDetails{}
}

