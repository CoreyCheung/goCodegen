package model

// Created by CodeMachine on 2020-05-10 15:46:04.68148 +0800 CST m=+0.084796804
type Jobs struct {
    Attempts    int8   `gorm:"not null;column:attempts;type:tinyint unsigned" json:"attempts"`
    AvailableAt uint   `gorm:"not null;column:available_at;type:int unsigned" json:"availableAt"`
    CreatedAt   uint   `gorm:"not null;column:created_at;type:int unsigned" json:"createdAt"`
    ID          int    `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int" json:"id"`
    Payload     string `gorm:"not null;column:payload;type:longtext;size:4294967295" json:"payload,omitempty"`
    Queue       string `gorm:"not null;column:queue;type:varchar(255);size:765" json:"queue,omitempty"`
    Reserved    int8   `gorm:"not null;column:reserved;type:tinyint unsigned" json:"reserved"`
    ReservedAt  uint   `gorm:"column:reserved_at;type:int unsigned" json:"reservedAt"`
}

func (Jobs) TableName() string {
	return "jobs"
}
func (this Jobs) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Jobs) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Jobs) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this Jobs) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(Jobs{}).Error; err != nil {
		return err
	}
	return nil
}

func NewJobs() Jobs {
    return Jobs{}
}

