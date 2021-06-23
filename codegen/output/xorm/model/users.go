package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:35.060823 +0800 CST m=+0.172834492
type Users struct {
    AuthId    string         `xorm:"'auth_id' VARCHAR(255)" json:"authId,omitempty"`
    // 头像
    Avatar    string         `xorm:"'avatar' TEXT" json:"avatar,omitempty"`
    Birth     date.Datetime  `xorm:"'birth' TIMESTAMP" json:"birth,omitempty"`
    CreatedAt date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    ID        uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    // 是否关注公众号 0-否 1-是
    IsFollow  int8           `xorm:"'is_follow' TINYINT" json:"isFollow"`
    Mobile    string         `xorm:"'mobile' VARCHAR(20)" json:"mobile,omitempty"`
    // 昵称
    Nickname  string         `xorm:"'nickname' VARCHAR(20)" json:"nickname,omitempty"`
    //  0-未填写 1-男 2-女
    Sex       int8           `xorm:"'sex' TINYINT" json:"sex"`
    UpdatedAt date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
    // 用户ID
    UserId    string         `xorm:"'user_id' VARCHAR(100)" json:"userId,omitempty"`
}

func (Users) TableName() string {
	return "users"
}

func NewUsers() Users {
    return Users{}
}
