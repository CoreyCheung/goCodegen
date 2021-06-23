package model

// Created by CodeMachine on 2020-05-10 15:46:04.662549 +0800 CST m=+0.065866033
type Districts struct {
    // 0:不展示 1：展示
    Bshow        int    `gorm:"column:bshow;type:int" json:"bshow"`
    // 城市的名字
    DistrictName string `gorm:"column:district_name;type:varchar(255);size:765" json:"districtName,omitempty"`
    // 层级序列
    DistrictSqe  string `gorm:"column:district_sqe;type:varchar(255);size:765" json:"districtSqe,omitempty"`
    // 地区所处的层级
    Hierarchy    int    `gorm:"column:hierarchy;type:int" json:"hierarchy"`
    // 1:热门地区/城市
    Hot          int    `gorm:"column:hot;type:int" json:"hot"`
    // 主键自增
    ID           uint   `gorm:"primary_key;not null;column:id;type:int unsigned" json:"id"`
    // 父类id
    Pid          int    `gorm:"column:pid;type:int" json:"pid"`
    // 城市的类型，0是国，1是省，2是市，3是区
    Type         int    `gorm:"column:type;type:int" json:"type"`
}

func (Districts) TableName() string {
	return "districts"
}
func (this Districts) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Districts) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Districts) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this Districts) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(Districts{}).Error; err != nil {
		return err
	}
	return nil
}

func NewDistricts() Districts {
    return Districts{}
}

