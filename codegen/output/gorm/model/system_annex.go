package model

// Created by CodeMachine on 2020-05-10 15:46:04.714397 +0800 CST m=+0.117713268
// [系统] 上传附件
type SystemAnnex struct {
    Ctime  uint    `gorm:"not null;column:ctime;type:int unsigned" json:"ctime"`
    // 关联的数据ID
    DataId uint    `gorm:"not null;column:data_id;type:int unsigned" json:"dataId"`
    // 上传文件
    File   string  `gorm:"not null;column:file;type:varchar(255);size:765" json:"file,omitempty"`
    // 文件分组
    Group  string  `gorm:"not null;column:group;type:varchar(100);size:300" json:"group,omitempty"`
    // 文件hash值
    Hash   string  `gorm:"not null;column:hash;type:varchar(64);size:192" json:"hash,omitempty"`
    ID     uint    `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    // 附件大小KB
    Size   float64 `gorm:"not null;column:size;type:decimal(12,2) unsigned" json:"size"`
    // 使用状态(0未使用，1已使用)
    Status int8    `gorm:"not null;column:status;type:tinyint unsigned" json:"status"`
    // 类型
    Type   string  `gorm:"not null;column:type;type:varchar(20);size:60" json:"type,omitempty"`
}

func (SystemAnnex) TableName() string {
	return "system_annex"
}
func (this SystemAnnex) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemAnnex) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemAnnex) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this SystemAnnex) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(SystemAnnex{}).Error; err != nil {
		return err
	}
	return nil
}

func NewSystemAnnex() SystemAnnex {
    return SystemAnnex{}
}

