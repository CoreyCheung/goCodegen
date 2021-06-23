package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.965132 +0800 CST m=+0.077146807
type Foods struct {
    AvgScore           float64        `xorm:"'avg_score' DOUBLE" json:"avgScore"`
    BadCommentCount    int            `xorm:"'bad_comment_count' INT" json:"badCommentCount"`
    CommentCount       int            `xorm:"'comment_count' INT" json:"commentCount"`
    CreatedAt          date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt          *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    Description        string         `xorm:"'description' not null VARCHAR(512)" json:"description,omitempty"`
    DiscountPrice      float64        `xorm:"'discount_price' DECIMAL(15,2)" json:"discountPrice"`
    FoodName           string         `xorm:"'food_name' not null VARCHAR(100)" json:"foodName,omitempty"`
    // 菜品做法
    FoodPractice       string         `xorm:"'food_practice' VARCHAR(100)" json:"foodPractice,omitempty"`
    // 所属菜单分类ID
    FoodTypesId        uint           `xorm:"'food_types_id' INT UNSIGNED" json:"foodTypesId"`
    GoodCommentCount   int            `xorm:"'good_comment_count' INT" json:"goodCommentCount"`
    ID                 uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    IsDiscount         int8           `xorm:"'is_discount' TINYINT" json:"isDiscount"`
    MiddleCommentCount int            `xorm:"'middle_comment_count' INT" json:"middleCommentCount"`
    // 菜品数量
    Number             int            `xorm:"'number' not null INT" json:"number"`
    Picture            string         `xorm:"'picture' VARCHAR(256)" json:"picture,omitempty"`
    // 菜品单价
    Price              float64        `xorm:"'price' DECIMAL(15,2)" json:"price"`
    ProductId          int            `xorm:"'product_id' INT" json:"productId"`
    ShopId             int            `xorm:"'shop_id' not null INT" json:"shopId"`
    Status             int8           `xorm:"'status' TINYINT" json:"status"`
    UpdatedAt          date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (Foods) TableName() string {
	return "foods"
}

func NewFoods() Foods {
    return Foods{}
}
