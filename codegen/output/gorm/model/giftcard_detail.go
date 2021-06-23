package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.672196 +0800 CST m=+0.075512631
type GiftcardDetail struct {
    CardId         string         `gorm:"column:card_id;type:varchar(255);size:1020" json:"cardId,omitempty"`
    CreatedAt      date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt      *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    EndTime        date.Datetime  `gorm:"column:end_time;type:timestamp" json:"endTime,omitempty"`
    GiftCardTypeId uint           `gorm:"column:gift_card_type_id;type:int unsigned" json:"giftCardTypeId"`
    ID             uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    Mobile         string         `gorm:"not null;column:mobile;type:varchar(255);size:1020" json:"mobile,omitempty"`
    ReceiveTime    date.Datetime  `gorm:"column:receive_time;type:timestamp" json:"receiveTime,omitempty"`
    RemainAmount   float64        `gorm:"column:remain_amount;type:double" json:"remainAmount"`
    Status         int            `gorm:"not null;column:status;type:int" json:"status"`
    TotalAmount    float64        `gorm:"column:total_amount;type:double" json:"totalAmount"`
    UpdatedAt      date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (GiftcardDetail) TableName() string {
	return "giftcard_detail"
}
func (this GiftcardDetail) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this GiftcardDetail) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this GiftcardDetail) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this GiftcardDetail) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(GiftcardDetail{}).Error; err != nil {
		return err
	}
	return nil
}

func NewGiftcardDetail() GiftcardDetail {
    return GiftcardDetail{}
}

