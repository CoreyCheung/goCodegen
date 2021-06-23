package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.673671 +0800 CST m=+0.076987855
type GiftcardRange struct {
    CreatedAt      date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt      *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    GiftCardTypeId uint           `gorm:"column:gift_card_type_id;type:int unsigned" json:"giftCardTypeId"`
    ID             uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    ShopId         int            `gorm:"column:shop_id;type:int" json:"shopId"`
    ShopName       string         `gorm:"column:shop_name;type:varchar(255);size:1020" json:"shopName,omitempty"`
    UpdatedAt      date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (GiftcardRange) TableName() string {
	return "giftcard_range"
}
func (this GiftcardRange) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this GiftcardRange) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this GiftcardRange) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this GiftcardRange) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(GiftcardRange{}).Error; err != nil {
		return err
	}
	return nil
}

func NewGiftcardRange() GiftcardRange {
    return GiftcardRange{}
}

