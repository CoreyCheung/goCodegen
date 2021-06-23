package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.998877 +0800 CST m=+0.110890315
type ReconciliationStasitics struct {
    ChanAmount      float64        `xorm:"'chan_amount' DECIMAL(15,2)" json:"chanAmount"`
    ChanCnt         int            `xorm:"'chan_cnt' INT" json:"chanCnt"`
    ChanFee         float64        `xorm:"'chan_fee' not null DECIMAL(15,2) UNSIGNED" json:"chanFee"`
    CreatedAt       date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt       *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    ID              uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    LocAmount       float64        `xorm:"'loc_amount' DECIMAL(15,2)" json:"locAmount"`
    LocCnt          int            `xorm:"'loc_cnt' INT" json:"locCnt"`
    LocFee          float64        `xorm:"'loc_fee' DECIMAL(15,2)" json:"locFee"`
    LongSection     int            `xorm:"'long_section' INT" json:"longSection"`
    Result          uint           `xorm:"'result' not null INT UNSIGNED" json:"result"`
    ShortParagraph  int            `xorm:"'short_paragraph' INT" json:"shortParagraph"`
    TransactionDate string         `xorm:"'transaction_date' VARCHAR(255)" json:"transactionDate,omitempty"`
    UpdatedAt       date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (ReconciliationStasitics) TableName() string {
	return "reconciliation_stasitics"
}

func NewReconciliationStasitics() ReconciliationStasitics {
    return ReconciliationStasitics{}
}
