package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:35.057415 +0800 CST m=+0.169426653
type TransactionDetails struct {
    Amount                 float64        `xorm:"'amount' DECIMAL(15,2)" json:"amount"`
    CreatedAt              date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    Currency               string         `xorm:"'currency' VARCHAR(255)" json:"currency,omitempty"`
    DeletedAt              *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    ID                     uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    OrderId                string         `xorm:"'order_id' VARCHAR(255)" json:"orderId,omitempty"`
    OutBizNo               string         `xorm:"'out_biz_no' VARCHAR(255)" json:"outBizNo,omitempty"`
    PayDate                date.Datetime  `xorm:"'pay_date' TIMESTAMP" json:"payDate,omitempty"`
    PayeeAccountId         string         `xorm:"'payee_account_id' VARCHAR(255)" json:"payeeAccountId,omitempty"`
    PayeeAccountType       uint           `xorm:"'payee_account_type' not null INT UNSIGNED" json:"payeeAccountType"`
    PayeeNo                string         `xorm:"'payee_no' VARCHAR(255)" json:"payeeNo,omitempty"`
    PayerAccountId         string         `xorm:"'payer_account_id' VARCHAR(255)" json:"payerAccountId,omitempty"`
    PayerAccountType       uint           `xorm:"'payer_account_type' not null INT UNSIGNED" json:"payerAccountType"`
    PayerNo                string         `xorm:"'payer_no' VARCHAR(255)" json:"payerNo,omitempty"`
    // 支付方式：00=支付宝;01=微信;02=线下
    PaymentType            string         `xorm:"'payment_type' VARCHAR(255)" json:"paymentType,omitempty"`
    RefundDate             date.Datetime  `xorm:"'refund_date' TIMESTAMP" json:"refundDate,omitempty"`
    Status                 int            `xorm:"'status' INT" json:"status"`
    TransactionRespCode    string         `xorm:"'transaction_resp_code' VARCHAR(255)" json:"transactionRespCode,omitempty"`
    TransactionRespMessage string         `xorm:"'transaction_resp_message' VARCHAR(255)" json:"transactionRespMessage,omitempty"`
    TransactionType        string         `xorm:"'transaction_type' VARCHAR(255)" json:"transactionType,omitempty"`
    UpdatedAt              date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (TransactionDetails) TableName() string {
	return "transaction_details"
}

func NewTransactionDetails() TransactionDetails {
    return TransactionDetails{}
}
