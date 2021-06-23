package model

// Created by CodeMachine on 2020-05-10 15:46:04.725318 +0800 CST m=+0.128633153
// [系统] 操作日志
type SystemLog struct {
    Count  uint   `gorm:"not null;column:count;type:int unsigned" json:"count"`
    Ctime  uint   `gorm:"not null;column:ctime;type:int unsigned" json:"ctime"`
    ID     uint   `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    Ip     string `gorm:"column:ip;type:varchar(128);size:384" json:"ip,omitempty"`
    Mtime  uint   `gorm:"not null;column:mtime;type:int unsigned" json:"mtime"`
    Param  string `gorm:"column:param;type:text;size:65535" json:"param,omitempty"`
    Remark string `gorm:"column:remark;type:varchar(255);size:765" json:"remark,omitempty"`
    Title  string `gorm:"column:title;type:varchar(100);size:300" json:"title,omitempty"`
    Uid    uint   `gorm:"not null;column:uid;type:int unsigned" json:"uid"`
    Url    string `gorm:"column:url;type:varchar(200);size:600" json:"url,omitempty"`
}

func (SystemLog) TableName() string {
	return "system_log"
}
func (this SystemLog) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemLog) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemLog) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this SystemLog) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(SystemLog{}).Error; err != nil {
		return err
	}
	return nil
}

func NewSystemLog() SystemLog {
    return SystemLog{}
}

