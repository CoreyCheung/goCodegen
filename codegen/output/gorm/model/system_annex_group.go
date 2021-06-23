package model

// Created by CodeMachine on 2020-05-10 15:46:04.716259 +0800 CST m=+0.119574452
// [系统] 附件分组
type SystemAnnexGroup struct {
    // 附件数量
    Count uint    `gorm:"not null;column:count;type:int unsigned" json:"count"`
    ID    uint    `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    // 附件分组
    Name  string  `gorm:"not null;column:name;type:varchar(50);size:150" json:"name,omitempty"`
    // 附件大小kb
    Size  float64 `gorm:"not null;column:size;type:decimal(12,2)" json:"size"`
}

func (SystemAnnexGroup) TableName() string {
	return "system_annex_group"
}
func (this SystemAnnexGroup) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemAnnexGroup) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemAnnexGroup) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this SystemAnnexGroup) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(SystemAnnexGroup{}).Error; err != nil {
		return err
	}
	return nil
}

func NewSystemAnnexGroup() SystemAnnexGroup {
    return SystemAnnexGroup{}
}

