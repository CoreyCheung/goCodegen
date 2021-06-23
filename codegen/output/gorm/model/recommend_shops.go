package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.700332 +0800 CST m=+0.103648106
type RecommendShops struct {
    BeginTime        date.Datetime  `gorm:"column:begin_time;type:timestamp" json:"beginTime,omitempty"`
    CreatedAt        date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt        *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    Description      string         `gorm:"not null;column:description;type:varchar(1024);size:4096" json:"description,omitempty"`
    EndTime          date.Datetime  `gorm:"column:end_time;type:timestamp" json:"endTime,omitempty"`
    ID               uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    RecommendPicture string         `gorm:"column:recommend_picture;type:varchar(255);size:1020" json:"recommendPicture,omitempty"`
    RecommendType    int            `gorm:"column:recommend_type;type:int" json:"recommendType"`
    ShopId           int            `gorm:"column:shop_id;type:int" json:"shopId"`
    UpdatedAt        date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (RecommendShops) TableName() string {
	return "recommend_shops"
}
func (this RecommendShops) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this RecommendShops) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this RecommendShops) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this RecommendShops) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(RecommendShops{}).Error; err != nil {
		return err
	}
	return nil
}

func NewRecommendShops() RecommendShops {
    return RecommendShops{}
}

