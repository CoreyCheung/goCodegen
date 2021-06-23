package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.645813 +0800 CST m=+0.049130504
type CouponActivities struct {
    ActivitiesDescription string         `gorm:"column:activities_description;type:varchar(255);size:1020" json:"activitiesDescription,omitempty"`
    ActivitiesExplain     string         `gorm:"column:activities_explain;type:varchar(255);size:1020" json:"activitiesExplain,omitempty"`
    ActivitiesMainUrl     string         `gorm:"column:activities_main_url;type:varchar(255);size:1020" json:"activitiesMainUrl,omitempty"`
    ActivitiesName        string         `gorm:"column:activities_name;type:varchar(255);size:1020" json:"activitiesName,omitempty"`
    ActivitiesNo          string         `gorm:"column:activities_no;type:varchar(255);size:1020" json:"activitiesNo,omitempty"`
    ActivitiesPicture     string         `gorm:"column:activities_picture;type:varchar(255);size:1020" json:"activitiesPicture,omitempty"`
    CreatedAt             date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt             *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    EndTime               date.Datetime  `gorm:"column:end_time;type:timestamp" json:"endTime,omitempty"`
    ID                    uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    StartTime             date.Datetime  `gorm:"column:start_time;type:timestamp" json:"startTime,omitempty"`
    UpdatedAt             date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (CouponActivities) TableName() string {
	return "coupon_activities"
}
func (this CouponActivities) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this CouponActivities) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this CouponActivities) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this CouponActivities) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(CouponActivities{}).Error; err != nil {
		return err
	}
	return nil
}

func NewCouponActivities() CouponActivities {
    return CouponActivities{}
}

