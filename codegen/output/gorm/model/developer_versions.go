package model

// Created by CodeMachine on 2020-05-10 15:46:04.659968 +0800 CST m=+0.063285384
// [开发助手] 版本记录
type DeveloperVersions struct {
    // 应用code
    AppName        string `gorm:"not null;column:app_name;type:varchar(100);size:300" json:"appName,omitempty"`
    // 应用版本号
    AppVersion     string `gorm:"not null;column:app_version;type:varchar(30);size:90" json:"appVersion,omitempty"`
    Ctime          uint   `gorm:"not null;column:ctime;type:int unsigned" json:"ctime"`
    // 删除文件记录
    DeleteFile     string `gorm:"column:delete_file;type:text;size:65535" json:"deleteFile,omitempty"`
    ID             uint   `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    // 安装包
    InstallPackage string `gorm:"column:install_package;type:varchar(255);size:765" json:"installPackage,omitempty"`
    // 发布状态：1已发布，0待发布
    Status         int8   `gorm:"not null;column:status;type:tinyint unsigned" json:"status"`
    // 类型：1模块，2插件，3模板
    Type           int8   `gorm:"not null;column:type;type:tinyint unsigned" json:"type"`
    // 更新文件记录
    UpdateFile     string `gorm:"column:update_file;type:text;size:65535" json:"updateFile,omitempty"`
    // 应用更新日志
    UpdateLog      string `gorm:"column:update_log;type:text;size:65535" json:"updateLog,omitempty"`
    // 更新sql
    UpdateSql      string `gorm:"column:update_sql;type:text;size:65535" json:"updateSql,omitempty"`
    // 升级包
    UpgradePackage string `gorm:"column:upgrade_package;type:varchar(255);size:765" json:"upgradePackage,omitempty"`
}

func (DeveloperVersions) TableName() string {
	return "developer_versions"
}
func (this DeveloperVersions) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this DeveloperVersions) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this DeveloperVersions) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this DeveloperVersions) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(DeveloperVersions{}).Error; err != nil {
		return err
	}
	return nil
}

func NewDeveloperVersions() DeveloperVersions {
    return DeveloperVersions{}
}

