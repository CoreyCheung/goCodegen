package model

// Created by CodeMachine on 2020-05-10 15:46:04.738298 +0800 CST m=+0.141612664
// [系统] 管理用户
type SystemUser struct {
    // 权限
    Auth          string `gorm:"not null;column:auth;type:text;size:65535" json:"auth,omitempty"`
    // 创建时间
    Ctime         uint   `gorm:"not null;column:ctime;type:int unsigned" json:"ctime"`
    // 邮箱
    Email         string `gorm:"not null;column:email;type:varchar(50);size:150" json:"email,omitempty"`
    ID            uint   `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    // 0默认，1框架
    Iframe        int8   `gorm:"not null;column:iframe;type:tinyint unsigned" json:"iframe"`
    // 最后登陆IP
    LastLoginIp   string `gorm:"not null;column:last_login_ip;type:varchar(128);size:384" json:"lastLoginIp,omitempty"`
    // 最后登陆时间
    LastLoginTime uint   `gorm:"not null;column:last_login_time;type:int unsigned" json:"lastLoginTime"`
    Mobile        string `gorm:"not null;column:mobile;type:varchar(11);size:33" json:"mobile,omitempty"`
    // 修改时间
    Mtime         uint   `gorm:"not null;column:mtime;type:int unsigned" json:"mtime"`
    // 昵称
    Nick          string `gorm:"not null;column:nick;type:varchar(50);size:150" json:"nick,omitempty"`
    Password      string `gorm:"not null;column:password;type:varchar(64);size:192" json:"password,omitempty"`
    // 角色ID
    RoleId        uint   `gorm:"not null;column:role_id;type:int unsigned" json:"roleId"`
    // 状态
    Status        int8   `gorm:"not null;column:status;type:tinyint unsigned" json:"status"`
    // 主题
    Theme         string `gorm:"not null;column:theme;type:varchar(50);size:150" json:"theme,omitempty"`
    // 用户名
    Username      string `gorm:"not null;column:username;type:varchar(50);size:150" json:"username,omitempty"`
}

func (SystemUser) TableName() string {
	return "system_user"
}
func (this SystemUser) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemUser) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this SystemUser) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this SystemUser) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(SystemUser{}).Error; err != nil {
		return err
	}
	return nil
}

func NewSystemUser() SystemUser {
    return SystemUser{}
}

