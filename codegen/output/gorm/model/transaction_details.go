package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.742746 +0800 CST m=+0.146060581
type TransactionDetails struct {
    Amount                 float64        `gorm:"column:amount;type:decimal(15,2)" json:"amount"`
    CreatedAt              date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    Currency               string         `gorm:"column:currency;type:varchar(255);size:1020" json:"currency,omitempty"`
    DeletedAt              *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    ID                     uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    OrderId                string         `gorm:"column:order_id;type:varchar(255);size:1020" json:"orderId,omitempty"`
    OutBizNo               string         `gorm:"column:out_biz_no;type:varchar(255);size:1020" json:"outBizNo,omitempty"`
    PayDate                date.Datetime  `gorm:"column:pay_date;type:timestamp" json:"payDate,omitempty"`
    PayeeAccountId         string         `gorm:"column:payee_account_id;type:varchar(255);size:1020" json:"payeeAccountId,omitempty"`
    PayeeAccountType       uint           `gorm:"not null;column:payee_account_type;type:int unsigned" json:"payeeAccountType"`
    PayeeNo                string         `gorm:"column:payee_no;type:varchar(255);size:1020" json:"payeeNo,omitempty"`
    PayerAccountId         string         `gorm:"column:payer_account_id;type:varchar(255);size:1020" json:"payerAccountId,omitempty"`
    PayerAccountType       uint           `gorm:"not null;column:payer_account_type;type:int unsigned" json:"payerAccountType"`
    PayerNo                string         `gorm:"column:payer_no;type:varchar(255);size:1020" json:"payerNo,omitempty"`
    // 支付方式：00=支付宝;01=微信;02=线下
    PaymentType            string         `gorm:"column:payment_type;type:varchar(255);size:1020" json:"paymentType,omitempty"`
    RefundDate             date.Datetime  `gorm:"column:refund_date;type:timestamp" json:"refundDate,omitempty"`
    Status                 int            `gorm:"column:status;type:int" json:"status"`
    TransactionRespCode    string         `gorm:"column:transaction_resp_code;type:varchar(255);size:1020" json:"transactionRespCode,omitempty"`
    TransactionRespMessage string         `gorm:"column:transaction_resp_message;type:varchar(255);size:1020" json:"transactionRespMessage,omitempty"`
    TransactionType        string         `gorm:"column:transaction_type;type:varchar(255);size:1020" json:"transactionType,omitempty"`
    UpdatedAt              date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (TransactionDetails) TableName() string {
	return "transaction_details"
}
func (this TransactionDetails) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this TransactionDetails) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this TransactionDetails) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this TransactionDetails) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(TransactionDetails{}).Error; err != nil {
		return err
	}
	return nil
}

func NewTransactionDetails() TransactionDetails {
    return TransactionDetails{}
}

