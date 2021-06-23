package model

// Created by CodeMachine on 2020-05-10 15:46:04.730871 +0800 CST m=+0.134186475
// [系统] 管理菜单语言包
type SystemMenuLang struct {
    ID     uint   `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    // 语言包
    Lang   int8   `gorm:"not null;column:lang;type:tinyint unsigned" json:"lang"`
    MenuId uint   `gorm:"not null;column:menu_id;type:int unsigned" json:"menuId"`
    // 标题
    Title  string `gorm:"not null;column:title;type:varchar(120);size:360" json:"title,omitempty"`
}

func (SystemMenuLang) TableName() string {
	return "system_menu_lang"
}
func (this SystemMenuLang) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemMenuLang) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemMenuLang) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this SystemMenuLang) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(SystemMenuLang{}).Error; err != nil {
		return err
	}
	return nil
}

func NewSystemMenuLang() SystemMenuLang {
    return SystemMenuLang{}
}

