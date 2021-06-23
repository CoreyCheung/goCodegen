package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.708199 +0800 CST m=+0.111515406
type Shops struct {
    Address                string         `gorm:"column:address;type:varchar(256);size:1024" json:"address,omitempty"`
    AllocationId           string         `gorm:"column:allocation_id;type:varchar(128);size:512" json:"allocationId,omitempty"`
    AppId                  string         `gorm:"column:app_id;type:varchar(10);size:40" json:"appId,omitempty"`
    AvgScore               float64        `gorm:"column:avg_score;type:double" json:"avgScore"`
    BadCommentCount        int            `gorm:"column:bad_comment_count;type:int" json:"badCommentCount"`
    // 分公司ID
    BranchNo               string         `gorm:"column:branch_no;type:varchar(128);size:512" json:"branchNo,omitempty"`
    // 店铺二维码
    Code                   string         `gorm:"column:code;type:varchar(255);size:1020" json:"code,omitempty"`
    CommentCount           int            `gorm:"column:comment_count;type:int" json:"commentCount"`
    CreatedAt              date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    // 商户id
    CustomerId             string         `gorm:"column:customer_id;type:varchar(100);size:400" json:"customerId,omitempty"`
    DeletedAt              *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    Description            string         `gorm:"column:description;type:varchar(1024);size:4096" json:"description,omitempty"`
    // 菜品实物图
    DescriptionPicture     string         `gorm:"column:description_picture;type:varchar(255);size:1020" json:"descriptionPicture,omitempty"`
    DevelopKey             string         `gorm:"column:develop_key;type:varchar(128);size:512" json:"developKey,omitempty"`
    DistrictSqe            string         `gorm:"column:district_sqe;type:varchar(50);size:200" json:"districtSqe,omitempty"`
    GoodCommentCount       int            `gorm:"column:good_comment_count;type:int" json:"goodCommentCount"`
    ID                     uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    Latitude               float64        `gorm:"column:latitude;type:decimal(20,7)" json:"latitude"`
    Longitude              float64        `gorm:"column:longitude;type:decimal(20,7)" json:"longitude"`
    // 店铺海报
    MainImage              string         `gorm:"column:main_image;type:varchar(255);size:1020" json:"mainImage,omitempty"`
    MessagePush            int            `gorm:"column:message_push;type:int" json:"messagePush"`
    MiddleCommentCount     int            `gorm:"column:middle_comment_count;type:int" json:"middleCommentCount"`
    // 店铺联系电话
    Mobile                 string         `gorm:"column:mobile;type:varchar(20);size:80" json:"mobile,omitempty"`
    OpeningTimeDescription string         `gorm:"column:opening_time_description;type:varchar(255);size:1020" json:"openingTimeDescription,omitempty"`
    OrderDisplay           int            `gorm:"column:order_display;type:int" json:"orderDisplay"`
    // 门前照片,logo
    Picture                string         `gorm:"column:picture;type:varchar(1024);size:4096" json:"picture,omitempty"`
    // 打印机配置信息
    PrintConfig            string         `gorm:"not null;column:print_config;type:varchar(600);size:2400" json:"printConfig,omitempty"`
    ProductId              int            `gorm:"column:product_id;type:int" json:"productId"`
    // 分账状态：1=关闭;2=开启;3=审核中
    ProfitStatus           bool           `gorm:"not null;column:profit_status;type:tinyint(1)" json:"profitStatus"`
    ShopName               string         `gorm:"column:shop_name;type:varchar(50);size:200" json:"shopName,omitempty"`
    // 0-正常  1-审核中    2-已注销
    ShopStatus             int8           `gorm:"column:shop_status;type:tinyint" json:"shopStatus"`
    // &#39;客户状态 0-正常 1-暂时营业
    Status                 int8           `gorm:"column:status;type:tinyint" json:"status"`
    Tags                   string         `gorm:"column:tags;type:varchar(128);size:512" json:"tags,omitempty"`
    UpdatedAt              date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (Shops) TableName() string {
	return "shops"
}
func (this Shops) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Shops) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this Shops) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this Shops) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(Shops{}).Error; err != nil {
		return err
	}
	return nil
}

func NewShops() Shops {
    return Shops{}
}

