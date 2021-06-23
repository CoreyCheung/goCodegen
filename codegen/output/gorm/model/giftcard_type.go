package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.676758 +0800 CST m=+0.080074794
type GiftcardType struct {
    CardId         string         `gorm:"column:card_id;type:varchar(255);size:1020" json:"cardId,omitempty"`
    CompanyLogo    string         `gorm:"column:company_logo;type:varchar(255);size:1020" json:"companyLogo,omitempty"`
    CompanyName    string         `gorm:"column:company_name;type:varchar(255);size:1020" json:"companyName,omitempty"`
    CreatedAt      date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt      *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    Description    string         `gorm:"column:description;type:varchar(255);size:1020" json:"description,omitempty"`
    // 0不限制
    ExpiryDateType int            `gorm:"column:expiry_date_type;type:int" json:"expiryDateType"`
    GiftCardName   string         `gorm:"column:gift_card_name;type:varchar(255);size:1020" json:"giftCardName,omitempty"`
    ID             uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    IssuedCount    int            `gorm:"column:issued_count;type:int" json:"issuedCount"`
    RegularDate    int            `gorm:"column:regular_date;type:int" json:"regularDate"`
    TotalAmount    float64        `gorm:"column:total_amount;type:double" json:"totalAmount"`
    UpdatedAt      date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
    UseEndTime     date.Datetime  `gorm:"column:use_end_time;type:timestamp" json:"useEndTime,omitempty"`
    UseRangeType   int            `gorm:"column:use_range_type;type:int" json:"useRangeType"`
}

func (GiftcardType) TableName() string {
	return "giftcard_type"
}
func (this GiftcardType) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this GiftcardType) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this GiftcardType) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this GiftcardType) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(GiftcardType{}).Error; err != nil {
		return err
	}
	return nil
}

func NewGiftcardType() GiftcardType {
    return GiftcardType{}
}

