package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.695865 +0800 CST m=+0.099180724
type Orders struct {
    // 消费卡抵扣了多少钱
    CardNumber         float64        `gorm:"not null;column:card_number;type:decimal(15,2)" json:"cardNumber"`
    ChainDate          date.Datetime  `gorm:"column:chain_date;type:timestamp" json:"chainDate,omitempty"`
    ChainHash          string         `gorm:"column:chain_hash;type:varchar(255);size:1020" json:"chainHash,omitempty"`
    ChainId            string         `gorm:"column:chain_id;type:varchar(255);size:1020" json:"chainId,omitempty"`
    // 使用的优惠券id，若未使用则为空
    CouponId           int            `gorm:"column:coupon_id;type:int" json:"couponId"`
    CreatedAt          date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DeletedAt          *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    // 上链错误次数，超过三次不在上链
    ErrorNum           bool           `gorm:"not null;column:error_num;type:tinyint(1)" json:"errorNum"`
    // 使用的消费卡id，若未使用则为空
    GiftcardId         int            `gorm:"column:giftcard_id;type:int" json:"giftcardId"`
    ID                 uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    // 是否评论过此订单,0未评论，1评论过
    IsComment          int8           `gorm:"column:is_comment;type:tinyint" json:"isComment"`
    // 0订单菜品未完成 1已完成
    IsFinished         int8           `gorm:"not null;column:is_finished;type:tinyint" json:"isFinished"`
    // 是否上链：0=否;1=是;-1=不上链
    IsProfit           bool           `gorm:"not null;column:is_profit;type:tinyint(1)" json:"isProfit"`
    // 合并的新订单id
    MergeOrderId       string         `gorm:"column:merge_order_id;type:varchar(255);size:1020" json:"mergeOrderId,omitempty"`
    // 用户留言
    Message            string         `gorm:"column:message;type:varchar(255);size:1020" json:"message,omitempty"`
    OrderId            string         `gorm:"not null;column:order_id;type:varchar(255);size:1020" json:"orderId,omitempty"`
    // 订单实付金额
    OrderPayPrice      float64        `gorm:"column:order_pay_price;type:decimal(15,2)" json:"orderPayPrice"`
    OrderTime          date.Datetime  `gorm:"column:order_time;type:timestamp" json:"orderTime,omitempty"`
    // 订单实际价格
    OrderTotalPrice    float64        `gorm:"column:order_total_price;type:decimal(15,2)" json:"orderTotalPrice"`
    // 优惠券抵扣金额
    PreferentialAmount float64        `gorm:"not null;column:preferential_amount;type:decimal(15,2)" json:"preferentialAmount"`
    // 桌号
    SeatId             string         `gorm:"not null;column:seat_id;type:varchar(255);size:1020" json:"seatId,omitempty"`
    // 桌号名称
    SeatName           string         `gorm:"column:seat_name;type:varchar(255);size:1020" json:"seatName,omitempty"`
    // 商铺ID
    ShopId             int            `gorm:"column:shop_id;type:int" json:"shopId"`
    // 商铺名称
    ShopName           string         `gorm:"column:shop_name;type:varchar(50);size:200" json:"shopName,omitempty"`
    // 0:已下单 1:已取消 2:已付款 3:已退款
    Status             int            `gorm:"column:status;type:int" json:"status"`
    UpdatedAt          date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
    // 下单用户ID
    UserId             string         `gorm:"column:user_id;type:varchar(255);size:1020" json:"userId,omitempty"`
}

func (Orders) TableName() string {
	return "orders"
}
func (this Orders) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Orders) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Orders) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this Orders) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(Orders{}).Error; err != nil {
		return err
	}
	return nil
}

func NewOrders() Orders {
    return Orders{}
}

