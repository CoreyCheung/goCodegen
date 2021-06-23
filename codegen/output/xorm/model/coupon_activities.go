package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.951307 +0800 CST m=+0.063321693
type CouponActivities struct {
    ActivitiesDescription string         `xorm:"'activities_description' VARCHAR(255)" json:"activitiesDescription,omitempty"`
    ActivitiesExplain     string         `xorm:"'activities_explain' VARCHAR(255)" json:"activitiesExplain,omitempty"`
    ActivitiesMainUrl     string         `xorm:"'activities_main_url' VARCHAR(255)" json:"activitiesMainUrl,omitempty"`
    ActivitiesName        string         `xorm:"'activities_name' VARCHAR(255)" json:"activitiesName,omitempty"`
    ActivitiesNo          string         `xorm:"'activities_no' VARCHAR(255)" json:"activitiesNo,omitempty"`
    ActivitiesPicture     string         `xorm:"'activities_picture' VARCHAR(255)" json:"activitiesPicture,omitempty"`
    CreatedAt             date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt             *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    EndTime               date.Datetime  `xorm:"'end_time' TIMESTAMP" json:"endTime,omitempty"`
    ID                    uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    StartTime             date.Datetime  `xorm:"'start_time' TIMESTAMP" json:"startTime,omitempty"`
    UpdatedAt             date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (CouponActivities) TableName() string {
	return "coupon_activities"
}

func NewCouponActivities() CouponActivities {
    return CouponActivities{}
}
