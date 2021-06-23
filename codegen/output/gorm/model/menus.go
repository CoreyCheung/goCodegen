package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.692827 +0800 CST m=+0.096143352
type Menus struct {
    CreatedAt      date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt      *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    FirstClassify  string         `gorm:"column:first_classify;type:varchar(255);size:1020" json:"firstClassify,omitempty"`
    FoodName       string         `gorm:"column:food_name;type:varchar(255);size:1020" json:"foodName,omitempty"`
    ID             uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    ImgUrl         string         `gorm:"column:img_url;type:varchar(255);size:1020" json:"imgUrl,omitempty"`
    Ingredients    string         `gorm:"column:ingredients;type:varchar(255);size:1020" json:"ingredients,omitempty"`
    SecondClassify string         `gorm:"column:second_classify;type:varchar(255);size:1020" json:"secondClassify,omitempty"`
    ThirdClassify  string         `gorm:"column:third_classify;type:varchar(255);size:1020" json:"thirdClassify,omitempty"`
    UpdatedAt      date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (Menus) TableName() string {
	return "menus"
}
func (this Menus) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Menus) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Menus) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this Menus) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(Menus{}).Error; err != nil {
		return err
	}
	return nil
}

func NewMenus() Menus {
    return Menus{}
}

