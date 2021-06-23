package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.706936 +0800 CST m=+0.110251774
type ShopSeats struct {
    // 店铺桌号二维码
    Code      string         `gorm:"column:code;type:varchar(255);size:1020" json:"code,omitempty"`
    CreatedAt date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    ID        uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    QrCode    string         `gorm:"column:qr_code;type:varchar(255);size:1020" json:"qrCode,omitempty"`
    QrPicture string         `gorm:"column:qr_picture;type:varchar(255);size:1020" json:"qrPicture,omitempty"`
    SeatId    uint           `gorm:"not null;column:seat_id;type:int unsigned" json:"seatId"`
    SeatName  string         `gorm:"column:seat_name;type:varchar(255);size:1020" json:"seatName,omitempty"`
    ShopId    int            `gorm:"not null;column:shop_id;type:int" json:"shopId"`
    UpdatedAt date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (ShopSeats) TableName() string {
	return "shop_seats"
}
func (this ShopSeats) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ShopSeats) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ShopSeats) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this ShopSeats) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(ShopSeats{}).Error; err != nil {
		return err
	}
	return nil
}

func NewShopSeats() ShopSeats {
    return ShopSeats{}
}

