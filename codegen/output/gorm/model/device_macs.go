package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.661469 +0800 CST m=+0.064786604
type DeviceMacs struct {
    CreatedAt date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    DepId     int            `gorm:"column:depId;type:int" json:"depId"`
    DeviceMac string         `gorm:"column:deviceMac;type:varchar(255);size:1020" json:"deviceMac,omitempty"`
    ID        uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    ShopId    uint           `gorm:"column:shop_id;type:int unsigned" json:"shopId"`
    UpdatedAt date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (DeviceMacs) TableName() string {
	return "device_macs"
}
func (this DeviceMacs) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this DeviceMacs) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this DeviceMacs) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this DeviceMacs) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(DeviceMacs{}).Error; err != nil {
		return err
	}
	return nil
}

func NewDeviceMacs() DeviceMacs {
    return DeviceMacs{}
}

