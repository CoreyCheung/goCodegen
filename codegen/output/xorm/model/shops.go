package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:35.007215 +0800 CST m=+0.119228298
type Shops struct {
    Address                string         `xorm:"'address' VARCHAR(256)" json:"address,omitempty"`
    AllocationId           string         `xorm:"'allocation_id' VARCHAR(128)" json:"allocationId,omitempty"`
    AppId                  string         `xorm:"'app_id' VARCHAR(10)" json:"appId,omitempty"`
    AvgScore               float64        `xorm:"'avg_score' DOUBLE" json:"avgScore"`
    BadCommentCount        int            `xorm:"'bad_comment_count' INT" json:"badCommentCount"`
    // 分公司ID
    BranchNo               string         `xorm:"'branch_no' VARCHAR(128)" json:"branchNo,omitempty"`
    // 店铺二维码
    Code                   string         `xorm:"'code' VARCHAR(255)" json:"code,omitempty"`
    CommentCount           int            `xorm:"'comment_count' INT" json:"commentCount"`
    CreatedAt              date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    // 商户id
    CustomerId             string         `xorm:"'customer_id' VARCHAR(100)" json:"customerId,omitempty"`
    DeletedAt              *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    Description            string         `xorm:"'description' VARCHAR(1024)" json:"description,omitempty"`
    // 菜品实物图
    DescriptionPicture     string         `xorm:"'description_picture' VARCHAR(255)" json:"descriptionPicture,omitempty"`
    DevelopKey             string         `xorm:"'develop_key' VARCHAR(128)" json:"developKey,omitempty"`
    DistrictSqe            string         `xorm:"'district_sqe' VARCHAR(50)" json:"districtSqe,omitempty"`
    GoodCommentCount       int            `xorm:"'good_comment_count' INT" json:"goodCommentCount"`
    ID                     uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    Latitude               float64        `xorm:"'latitude' DECIMAL(20,7)" json:"latitude"`
    Longitude              float64        `xorm:"'longitude' DECIMAL(20,7)" json:"longitude"`
    // 店铺海报
    MainImage              string         `xorm:"'main_image' VARCHAR(255)" json:"mainImage,omitempty"`
    MessagePush            int            `xorm:"'message_push' INT" json:"messagePush"`
    MiddleCommentCount     int            `xorm:"'middle_comment_count' INT" json:"middleCommentCount"`
    // 店铺联系电话
    Mobile                 string         `xorm:"'mobile' VARCHAR(20)" json:"mobile,omitempty"`
    OpeningTimeDescription string         `xorm:"'opening_time_description' VARCHAR(255)" json:"openingTimeDescription,omitempty"`
    OrderDisplay           int            `xorm:"'order_display' INT" json:"orderDisplay"`
    // 门前照片,logo
    Picture                string         `xorm:"'picture' VARCHAR(1024)" json:"picture,omitempty"`
    // 打印机配置信息
    PrintConfig            string         `xorm:"'print_config' not null VARCHAR(600)" json:"printConfig,omitempty"`
    ProductId              int            `xorm:"'product_id' INT" json:"productId"`
    // 分账状态：1=关闭;2=开启;3=审核中
    ProfitStatus           bool           `xorm:"'profit_status' not null TINYINT(1)" json:"profitStatus"`
    ShopName               string         `xorm:"'shop_name' VARCHAR(50)" json:"shopName,omitempty"`
    // 0-正常  1-审核中    2-已注销
    ShopStatus             int8           `xorm:"'shop_status' TINYINT" json:"shopStatus"`
    // &#39;客户状态 0-正常 1-暂时营业
    Status                 int8           `xorm:"'status' TINYINT" json:"status"`
    Tags                   string         `xorm:"'tags' VARCHAR(128)" json:"tags,omitempty"`
    UpdatedAt              date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (Shops) TableName() string {
	return "shops"
}

func NewShops() Shops {
    return Shops{}
}
