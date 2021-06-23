package model

// Created by CodeMachine on 2020-05-09 22:05:35.053147 +0800 CST m=+0.165159332
// [系统] 管理用户
type SystemUser struct {
    // 权限
    Auth          string `xorm:"'auth' not null TEXT" json:"auth,omitempty"`
    // 创建时间
    Ctime         uint   `xorm:"'ctime' not null INT UNSIGNED" json:"ctime"`
    // 邮箱
    Email         string `xorm:"'email' not null VARCHAR(50)" json:"email,omitempty"`
    ID            uint   `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    // 0默认，1框架
    Iframe        int8   `xorm:"'iframe' not null TINYINT UNSIGNED" json:"iframe"`
    // 最后登陆IP
    LastLoginIp   string `xorm:"'last_login_ip' not null VARCHAR(128)" json:"lastLoginIp,omitempty"`
    // 最后登陆时间
    LastLoginTime uint   `xorm:"'last_login_time' not null INT UNSIGNED" json:"lastLoginTime"`
    Mobile        string `xorm:"'mobile' not null VARCHAR(11)" json:"mobile,omitempty"`
    // 修改时间
    Mtime         uint   `xorm:"'mtime' not null INT UNSIGNED" json:"mtime"`
    // 昵称
    Nick          string `xorm:"'nick' not null VARCHAR(50)" json:"nick,omitempty"`
    Password      string `xorm:"'password' not null VARCHAR(64)" json:"password,omitempty"`
    // 角色ID
    RoleId        uint   `xorm:"'role_id' not null INT UNSIGNED" json:"roleId"`
    // 状态
    Status        int8   `xorm:"'status' not null TINYINT UNSIGNED" json:"status"`
    // 主题
    Theme         string `xorm:"'theme' not null VARCHAR(50)" json:"theme,omitempty"`
    // 用户名
    Username      string `xorm:"'username' not null VARCHAR(50)" json:"username,omitempty"`
}

func (SystemUser) TableName() string {
	return "system_user"
}

func NewSystemUser() SystemUser {
    return SystemUser{}
}
