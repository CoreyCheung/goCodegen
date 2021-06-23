package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.666007 +0800 CST m=+0.069324510
type Foods struct {
    AvgScore           float64        `gorm:"column:avg_score;type:double" json:"avgScore"`
    BadCommentCount    int            `gorm:"column:bad_comment_count;type:int" json:"badCommentCount"`
    CommentCount       int            `gorm:"column:comment_count;type:int" json:"commentCount"`
    CreatedAt          date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt          *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    Description        string         `gorm:"not null;column:description;type:varchar(512);size:2048" json:"description,omitempty"`
    DiscountPrice      float64        `gorm:"column:discount_price;type:decimal(15,2)" json:"discountPrice"`
    FoodName           string         `gorm:"not null;column:food_name;type:varchar(100);size:400" json:"foodName,omitempty"`
    // 菜品做法
    FoodPractice       string         `gorm:"column:food_practice;type:varchar(100);size:400" json:"foodPractice,omitempty"`
    // 所属菜单分类ID
    FoodTypesId        uint           `gorm:"column:food_types_id;type:int unsigned" json:"foodTypesId"`
    GoodCommentCount   int            `gorm:"column:good_comment_count;type:int" json:"goodCommentCount"`
    ID                 uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    IsDiscount         int8           `gorm:"column:is_discount;type:tinyint" json:"isDiscount"`
    MiddleCommentCount int            `gorm:"column:middle_comment_count;type:int" json:"middleCommentCount"`
    // 菜品数量
    Number             int            `gorm:"not null;column:number;type:int" json:"number"`
    Picture            string         `gorm:"column:picture;type:varchar(256);size:1024" json:"picture,omitempty"`
    // 菜品单价
    Price              float64        `gorm:"column:price;type:decimal(15,2)" json:"price"`
    ProductId          int            `gorm:"column:product_id;type:int" json:"productId"`
    ShopId             int            `gorm:"not null;column:shop_id;type:int" json:"shopId"`
    Status             int8           `gorm:"column:status;type:tinyint" json:"status"`
    UpdatedAt          date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (Foods) TableName() string {
	return "foods"
}
func (this Foods) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Foods) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Foods) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this Foods) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(Foods{}).Error; err != nil {
		return err
	}
	return nil
}

func NewFoods() Foods {
    return Foods{}
}

