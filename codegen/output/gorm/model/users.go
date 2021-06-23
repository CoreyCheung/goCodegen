package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.746827 +0800 CST m=+0.150141652
type Users struct {
    AuthId    string         `gorm:"column:auth_id;type:varchar(255);size:1020" json:"authId,omitempty"`
    // 头像
    Avatar    string         `gorm:"column:avatar;type:text;size:65535" json:"avatar,omitempty"`
    Birth     date.Datetime  `gorm:"column:birth;type:timestamp" json:"birth,omitempty"`
    CreatedAt date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    ID        uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    // 是否关注公众号 0-否 1-是
    IsFollow  int8           `gorm:"column:is_follow;type:tinyint" json:"isFollow"`
    Mobile    string         `gorm:"column:mobile;type:varchar(20);size:80" json:"mobile,omitempty"`
    // 昵称
    Nickname  string         `gorm:"column:nickname;type:varchar(20);size:80" json:"nickname,omitempty"`
    //  0-未填写 1-男 2-女
    Sex       int8           `gorm:"column:sex;type:tinyint" json:"sex"`
    UpdatedAt date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
    // 用户ID
    UserId    string         `gorm:"column:user_id;type:varchar(100);size:400" json:"userId,omitempty"`
}

func (Users) TableName() string {
	return "users"
}
func (this Users) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Users) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Users) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this Users) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(Users{}).Error; err != nil {
		return err
	}
	return nil
}

func NewUsers() Users {
    return Users{}
}

