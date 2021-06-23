package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.657697 +0800 CST m=+0.061014713
type Customers struct {
    Address             string         `gorm:"column:address;type:varchar(256);size:1024" json:"address,omitempty"`
    // 开户人姓名
    BankAccountName     string         `gorm:"column:bank_account_name;type:varchar(20);size:80" json:"bankAccountName,omitempty"`
    // 银行卡号
    BankCard            string         `gorm:"column:bank_card;type:varchar(20);size:80" json:"bankCard,omitempty"`
    // 营业执照
    BusinessLicense     string         `gorm:"column:business_license;type:varchar(255);size:1020" json:"businessLicense,omitempty"`
    CreatedAt           date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    CustomerId          string         `gorm:"column:customer_id;type:varchar(255);size:1020" json:"customerId,omitempty"`
    CustomerType        int8           `gorm:"column:customer_type;type:tinyint" json:"customerType"`
    DeletedAt           *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    ID                  uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    IdCard              string         `gorm:"column:id_card;type:varchar(20);size:80" json:"idCard,omitempty"`
    Idcard              string         `gorm:"column:idcard;type:varchar(20);size:80" json:"idcard,omitempty"`
    // 身份证背面照
    IdentityCardBackend string         `gorm:"column:identity_card_backend;type:varchar(255);size:1020" json:"identityCardBackend,omitempty"`
    // 身份证正面照
    IdentityCardFront   string         `gorm:"column:identity_card_front;type:varchar(255);size:1020" json:"identityCardFront,omitempty"`
    Mobile              string         `gorm:"column:mobile;type:varchar(20);size:80" json:"mobile,omitempty"`
    Name                string         `gorm:"column:name;type:varchar(100);size:400" json:"name,omitempty"`
    // 开户银行名称
    OpenBank            string         `gorm:"column:open_bank;type:varchar(20);size:80" json:"openBank,omitempty"`
    Password            string         `gorm:"column:password;type:varchar(100);size:400" json:"password,omitempty"`
    RealName            string         `gorm:"column:real_name;type:varchar(20);size:80" json:"realName,omitempty"`
    Sex                 int8           `gorm:"column:sex;type:tinyint" json:"sex"`
    ShopId              int            `gorm:"column:shop_id;type:int" json:"shopId"`
    UpdatedAt           date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
    UserName            string         `gorm:"not null;column:user_name;type:varchar(255);size:1020" json:"userName,omitempty"`
}

func (Customers) TableName() string {
	return "customers"
}
func (this Customers) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Customers) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Customers) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this Customers) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(Customers{}).Error; err != nil {
		return err
	}
	return nil
}

func NewCustomers() Customers {
    return Customers{}
}

