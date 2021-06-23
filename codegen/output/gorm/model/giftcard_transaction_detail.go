package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.675008 +0800 CST m=+0.078324434
type GiftcardTransactionDetail struct {
    Amount       float64        `gorm:"column:amount;type:double" json:"amount"`
    CreatedAt    date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt    *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    Description  string         `gorm:"column:description;type:varchar(255);size:1020" json:"description,omitempty"`
    GiftCardId   int            `gorm:"column:gift_card_id;type:int" json:"giftCardId"`
    GiftCardName string         `gorm:"column:gift_card_name;type:varchar(255);size:1020" json:"giftCardName,omitempty"`
    ID           uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    Mobile       string         `gorm:"column:mobile;type:varchar(255);size:1020" json:"mobile,omitempty"`
    // 充值和退款的操作员
    Operator     string         `gorm:"column:operator;type:varchar(255);size:1020" json:"operator,omitempty"`
    OrderId      string         `gorm:"column:order_id;type:varchar(255);size:1020" json:"orderId,omitempty"`
    ShopId       int            `gorm:"column:shop_id;type:int" json:"shopId"`
    Type         int            `gorm:"column:type;type:int" json:"type"`
    UpdatedAt    date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (GiftcardTransactionDetail) TableName() string {
	return "giftcard_transaction_detail"
}
func (this GiftcardTransactionDetail) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this GiftcardTransactionDetail) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this GiftcardTransactionDetail) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this GiftcardTransactionDetail) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(GiftcardTransactionDetail{}).Error; err != nil {
		return err
	}
	return nil
}

func NewGiftcardTransactionDetail() GiftcardTransactionDetail {
    return GiftcardTransactionDetail{}
}

