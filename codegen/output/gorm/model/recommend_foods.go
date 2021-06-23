package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.699247 +0800 CST m=+0.102562950
type RecommendFoods struct {
    CreatedAt        date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt        *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    Description      string         `gorm:"not null;column:description;type:varchar(512);size:2048" json:"description,omitempty"`
    FoodId           int            `gorm:"column:food_id;type:int" json:"foodId"`
    ID               uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    RecommendPicture string         `gorm:"column:recommend_picture;type:varchar(255);size:1020" json:"recommendPicture,omitempty"`
    ShopId           int            `gorm:"column:shop_id;type:int" json:"shopId"`
    UpdatedAt        date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (RecommendFoods) TableName() string {
	return "recommend_foods"
}
func (this RecommendFoods) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this RecommendFoods) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this RecommendFoods) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this RecommendFoods) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(RecommendFoods{}).Error; err != nil {
		return err
	}
	return nil
}

func NewRecommendFoods() RecommendFoods {
    return RecommendFoods{}
}

