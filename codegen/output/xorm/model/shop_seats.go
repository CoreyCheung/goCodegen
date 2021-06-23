package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:35.004969 +0800 CST m=+0.116982517
type ShopSeats struct {
    // 店铺桌号二维码
    Code      string         `xorm:"'code' VARCHAR(255)" json:"code,omitempty"`
    CreatedAt date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    ID        uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    QrCode    string         `xorm:"'qr_code' VARCHAR(255)" json:"qrCode,omitempty"`
    QrPicture string         `xorm:"'qr_picture' VARCHAR(255)" json:"qrPicture,omitempty"`
    SeatId    uint           `xorm:"'seat_id' not null INT UNSIGNED" json:"seatId"`
    SeatName  string         `xorm:"'seat_name' VARCHAR(255)" json:"seatName,omitempty"`
    ShopId    int            `xorm:"'shop_id' not null INT" json:"shopId"`
    UpdatedAt date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (ShopSeats) TableName() string {
	return "shop_seats"
}

func NewShopSeats() ShopSeats {
    return ShopSeats{}
}
