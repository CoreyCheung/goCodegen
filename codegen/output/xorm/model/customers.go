package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.959208 +0800 CST m=+0.071222481
type Customers struct {
    Address             string         `xorm:"'address' VARCHAR(256)" json:"address,omitempty"`
    // 开户人姓名
    BankAccountName     string         `xorm:"'bank_account_name' VARCHAR(20)" json:"bankAccountName,omitempty"`
    // 银行卡号
    BankCard            string         `xorm:"'bank_card' VARCHAR(20)" json:"bankCard,omitempty"`
    // 营业执照
    BusinessLicense     string         `xorm:"'business_license' VARCHAR(255)" json:"businessLicense,omitempty"`
    CreatedAt           date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    CustomerId          string         `xorm:"'customer_id' VARCHAR(255)" json:"customerId,omitempty"`
    CustomerType        int8           `xorm:"'customer_type' TINYINT" json:"customerType"`
    DeletedAt           *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    ID                  uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    IdCard              string         `xorm:"'id_card' VARCHAR(20)" json:"idCard,omitempty"`
    Idcard              string         `xorm:"'idcard' VARCHAR(20)" json:"idcard,omitempty"`
    // 身份证背面照
    IdentityCardBackend string         `xorm:"'identity_card_backend' VARCHAR(255)" json:"identityCardBackend,omitempty"`
    // 身份证正面照
    IdentityCardFront   string         `xorm:"'identity_card_front' VARCHAR(255)" json:"identityCardFront,omitempty"`
    Mobile              string         `xorm:"'mobile' VARCHAR(20)" json:"mobile,omitempty"`
    Name                string         `xorm:"'name' VARCHAR(100)" json:"name,omitempty"`
    // 开户银行名称
    OpenBank            string         `xorm:"'open_bank' VARCHAR(20)" json:"openBank,omitempty"`
    Password            string         `xorm:"'password' VARCHAR(100)" json:"password,omitempty"`
    RealName            string         `xorm:"'real_name' VARCHAR(20)" json:"realName,omitempty"`
    Sex                 int8           `xorm:"'sex' TINYINT" json:"sex"`
    ShopId              int            `xorm:"'shop_id' INT" json:"shopId"`
    UpdatedAt           date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
    UserName            string         `xorm:"'user_name' not null VARCHAR(255)" json:"userName,omitempty"`
}

func (Customers) TableName() string {
	return "customers"
}

func NewCustomers() Customers {
    return Customers{}
}
