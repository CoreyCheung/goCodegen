package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.703087 +0800 CST m=+0.106402444
type ReconciliationStasitics struct {
    ChanAmount      float64        `gorm:"column:chan_amount;type:decimal(15,2)" json:"chanAmount"`
    ChanCnt         int            `gorm:"column:chan_cnt;type:int" json:"chanCnt"`
    ChanFee         float64        `gorm:"not null;column:chan_fee;type:decimal(15,2) unsigned" json:"chanFee"`
    CreatedAt       date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt       *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    ID              uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    LocAmount       float64        `gorm:"column:loc_amount;type:decimal(15,2)" json:"locAmount"`
    LocCnt          int            `gorm:"column:loc_cnt;type:int" json:"locCnt"`
    LocFee          float64        `gorm:"column:loc_fee;type:decimal(15,2)" json:"locFee"`
    LongSection     int            `gorm:"column:long_section;type:int" json:"longSection"`
    Result          uint           `gorm:"not null;column:result;type:int unsigned" json:"result"`
    ShortParagraph  int            `gorm:"column:short_paragraph;type:int" json:"shortParagraph"`
    TransactionDate string         `gorm:"column:transaction_date;type:varchar(255);size:1020" json:"transactionDate,omitempty"`
    UpdatedAt       date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (ReconciliationStasitics) TableName() string {
	return "reconciliation_stasitics"
}
func (this ReconciliationStasitics) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ReconciliationStasitics) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ReconciliationStasitics) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this ReconciliationStasitics) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(ReconciliationStasitics{}).Error; err != nil {
		return err
	}
	return nil
}

func NewReconciliationStasitics() ReconciliationStasitics {
    return ReconciliationStasitics{}
}

