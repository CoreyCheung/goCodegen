package model

import "git.wanpinghui.com/WPH/go_common/wph"
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:35.055706 +0800 CST m=+0.167717940
type ThirdProductRelApplies struct {
    CertImg        string        `xorm:"'cert_img' VARCHAR(255)" json:"certImg,omitempty"`
    CreatedAt      date.Datetime `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    CustomerId     wph.Long      `xorm:"'customer_id' BIGINT" json:"customerId"`
    ID             uint          `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    Stat           int8          `xorm:"'stat' TINYINT UNSIGNED" json:"stat"`
    ThirdProductId uint          `xorm:"'third_product_id' INT UNSIGNED" json:"thirdProductId"`
    UpdatedAt      date.Datetime `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (ThirdProductRelApplies) TableName() string {
	return "third_product_rel_applies"
}

func NewThirdProductRelApplies() ThirdProductRelApplies {
    return ThirdProductRelApplies{}
}
