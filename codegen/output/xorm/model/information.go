package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.976928 +0800 CST m=+0.088941817
type Information struct {
    CreatedAt  date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt  *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    ID         uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    Keyword    string         `xorm:"'keyword' VARCHAR(255)" json:"keyword,omitempty"`
    SubKeyword string         `xorm:"'sub_keyword' VARCHAR(255)" json:"subKeyword,omitempty"`
    UpdatedAt  date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
    Value      string         `xorm:"'value' VARCHAR(1024)" json:"value,omitempty"`
}

func (Information) TableName() string {
	return "information"
}

func NewInformation() Information {
    return Information{}
}
