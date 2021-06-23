package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.988963 +0800 CST m=+0.100977036
type Orders struct {
    // 消费卡抵扣了多少钱
    CardNumber         float64        `xorm:"'card_number' not null DECIMAL(15,2)" json:"cardNumber"`
    ChainDate          date.Datetime  `xorm:"'chain_date' TIMESTAMP" json:"chainDate,omitempty"`
    ChainHash          string         `xorm:"'chain_hash' VARCHAR(255)" json:"chainHash,omitempty"`
    ChainId            string         `xorm:"'chain_id' VARCHAR(255)" json:"chainId,omitempty"`
    // 使用的优惠券id，若未使用则为空
    CouponId           int            `xorm:"'coupon_id' INT" json:"couponId"`
    CreatedAt          date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DeletedAt          *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    // 上链错误次数，超过三次不在上链
    ErrorNum           bool           `xorm:"'error_num' not null TINYINT(1)" json:"errorNum"`
    // 使用的消费卡id，若未使用则为空
    GiftcardId         int            `xorm:"'giftcard_id' INT" json:"giftcardId"`
    ID                 uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    // 是否评论过此订单,0未评论，1评论过
    IsComment          int8           `xorm:"'is_comment' TINYINT" json:"isComment"`
    // 0订单菜品未完成 1已完成
    IsFinished         int8           `xorm:"'is_finished' not null TINYINT" json:"isFinished"`
    // 是否上链：0=否;1=是;-1=不上链
    IsProfit           bool           `xorm:"'is_profit' not null TINYINT(1)" json:"isProfit"`
    // 合并的新订单id
    MergeOrderId       string         `xorm:"'merge_order_id' VARCHAR(255)" json:"mergeOrderId,omitempty"`
    // 用户留言
    Message            string         `xorm:"'message' VARCHAR(255)" json:"message,omitempty"`
    OrderId            string         `xorm:"'order_id' not null VARCHAR(255)" json:"orderId,omitempty"`
    // 订单实付金额
    OrderPayPrice      float64        `xorm:"'order_pay_price' DECIMAL(15,2)" json:"orderPayPrice"`
    OrderTime          date.Datetime  `xorm:"'order_time' TIMESTAMP" json:"orderTime,omitempty"`
    // 订单实际价格
    OrderTotalPrice    float64        `xorm:"'order_total_price' DECIMAL(15,2)" json:"orderTotalPrice"`
    // 优惠券抵扣金额
    PreferentialAmount float64        `xorm:"'preferential_amount' not null DECIMAL(15,2)" json:"preferentialAmount"`
    // 桌号
    SeatId             string         `xorm:"'seat_id' not null VARCHAR(255)" json:"seatId,omitempty"`
    // 桌号名称
    SeatName           string         `xorm:"'seat_name' VARCHAR(255)" json:"seatName,omitempty"`
    // 商铺ID
    ShopId             int            `xorm:"'shop_id' INT" json:"shopId"`
    // 商铺名称
    ShopName           string         `xorm:"'shop_name' VARCHAR(50)" json:"shopName,omitempty"`
    // 0:已下单 1:已取消 2:已付款 3:已退款
    Status             int            `xorm:"'status' INT" json:"status"`
    UpdatedAt          date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
    // 下单用户ID
    UserId             string         `xorm:"'user_id' VARCHAR(255)" json:"userId,omitempty"`
}

func (Orders) TableName() string {
	return "orders"
}

func NewOrders() Orders {
    return Orders{}
}
