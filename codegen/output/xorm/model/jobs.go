package model

// Created by CodeMachine on 2020-05-09 22:05:34.977572 +0800 CST m=+0.089585802
type Jobs struct {
    Attempts    int8   `xorm:"'attempts' not null TINYINT UNSIGNED" json:"attempts"`
    AvailableAt uint   `xorm:"'available_at' not null INT UNSIGNED" json:"availableAt"`
    CreatedAt   uint   `xorm:"'created_at' not null INT UNSIGNED" json:"createdAt"`
    ID          int    `xorm:"'id' not null pk autoincr INT" json:"id"`
    Payload     string `xorm:"'payload' not null LONGTEXT" json:"payload,omitempty"`
    Queue       string `xorm:"'queue' not null VARCHAR(255)" json:"queue,omitempty"`
    Reserved    int8   `xorm:"'reserved' not null TINYINT UNSIGNED" json:"reserved"`
    ReservedAt  uint   `xorm:"'reserved_at' INT UNSIGNED" json:"reservedAt"`
}

func (Jobs) TableName() string {
	return "jobs"
}

func NewJobs() Jobs {
    return Jobs{}
}
