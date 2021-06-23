package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.679793 +0800 CST m=+0.083109651
type Information struct {
    CreatedAt  date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt  *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    ID         uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    Keyword    string         `gorm:"column:keyword;type:varchar(255);size:1020" json:"keyword,omitempty"`
    SubKeyword string         `gorm:"column:sub_keyword;type:varchar(255);size:1020" json:"subKeyword,omitempty"`
    UpdatedAt  date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
    Value      string         `gorm:"column:value;type:varchar(1024);size:4096" json:"value,omitempty"`
}

func (Information) TableName() string {
	return "information"
}
func (this Information) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Information) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Information) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this Information) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(Information{}).Error; err != nil {
		return err
	}
	return nil
}

func NewInformation() Information {
    return Information{}
}

