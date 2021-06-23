package model

// Created by CodeMachine on 2020-05-09 22:05:34.96084 +0800 CST m=+0.072854617
// [开发助手] 版本记录
type DeveloperVersions struct {
    // 应用code
    AppName        string `xorm:"'app_name' not null VARCHAR(100)" json:"appName,omitempty"`
    // 应用版本号
    AppVersion     string `xorm:"'app_version' not null VARCHAR(30)" json:"appVersion,omitempty"`
    Ctime          uint   `xorm:"'ctime' not null INT UNSIGNED" json:"ctime"`
    // 删除文件记录
    DeleteFile     string `xorm:"'delete_file' TEXT" json:"deleteFile,omitempty"`
    ID             uint   `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    // 安装包
    InstallPackage string `xorm:"'install_package' VARCHAR(255)" json:"installPackage,omitempty"`
    // 发布状态：1已发布，0待发布
    Status         int8   `xorm:"'status' not null TINYINT UNSIGNED" json:"status"`
    // 类型：1模块，2插件，3模板
    Type           int8   `xorm:"'type' not null TINYINT UNSIGNED" json:"type"`
    // 更新文件记录
    UpdateFile     string `xorm:"'update_file' TEXT" json:"updateFile,omitempty"`
    // 应用更新日志
    UpdateLog      string `xorm:"'update_log' TEXT" json:"updateLog,omitempty"`
    // 更新sql
    UpdateSql      string `xorm:"'update_sql' TEXT" json:"updateSql,omitempty"`
    // 升级包
    UpgradePackage string `xorm:"'upgrade_package' VARCHAR(255)" json:"upgradePackage,omitempty"`
}

func (DeveloperVersions) TableName() string {
	return "developer_versions"
}

func NewDeveloperVersions() DeveloperVersions {
    return DeveloperVersions{}
}
