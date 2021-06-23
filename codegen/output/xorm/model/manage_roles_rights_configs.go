package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-09 22:05:34.981109 +0800 CST m=+0.093122831
type ManageRolesRightsConfigs struct {
    AddActivity          bool           `xorm:"'add_activity' TINYINT(1)" json:"addActivity"`
    AddShop              bool           `xorm:"'add_shop' TINYINT(1)" json:"addShop"`
    CreatedAt            date.Datetime  `xorm:"'created_at' TIMESTAMP" json:"createdAt,omitempty"`
    DataQuery            bool           `xorm:"'data_query' TINYINT(1)" json:"dataQuery"`
    DeletedAt            *date.Datetime `xorm:"'deleted_at' TIMESTAMP" json:"deletedAt,omitempty"`
    GiftCardManagement   bool           `xorm:"'gift_card_management' TINYINT(1)" json:"giftCardManagement"`
    ID                   uint           `xorm:"'id' not null pk autoincr INT UNSIGNED" json:"id"`
    ManageDepartment     bool           `xorm:"'manage_department' TINYINT(1)" json:"manageDepartment"`
    ManageRole           bool           `xorm:"'manage_role' TINYINT(1)" json:"manageRole"`
    ManageUser           bool           `xorm:"'manage_user' TINYINT(1)" json:"manageUser"`
    OperationsManagement bool           `xorm:"'operations_management' TINYINT(1)" json:"operationsManagement"`
    PicturePlatform      bool           `xorm:"'picture_platform' TINYINT(1)" json:"picturePlatform"`
    PictureShop          bool           `xorm:"'picture_shop' TINYINT(1)" json:"pictureShop"`
    QueryActivity        bool           `xorm:"'query_activity' TINYINT(1)" json:"queryActivity"`
    QueryOrder           bool           `xorm:"'query_order' TINYINT(1)" json:"queryOrder"`
    QueryReconciliation  bool           `xorm:"'query_reconciliation' TINYINT(1)" json:"queryReconciliation"`
    QuerySettlement      bool           `xorm:"'query_settlement' TINYINT(1)" json:"querySettlement"`
    QueryShop            bool           `xorm:"'query_shop' TINYINT(1)" json:"queryShop"`
    QueryUser            bool           `xorm:"'query_user' TINYINT(1)" json:"queryUser"`
    RoleId               uint           `xorm:"'role_id' INT UNSIGNED" json:"roleId"`
    SeparateAccount      bool           `xorm:"'separate_account' TINYINT(1)" json:"separateAccount"`
    ShopCoupon           bool           `xorm:"'shop_coupon' TINYINT(1)" json:"shopCoupon"`
    ShopManagement       bool           `xorm:"'shop_management' TINYINT(1)" json:"shopManagement"`
    UpdatedAt            date.Datetime  `xorm:"'updated_at' updated TIMESTAMP" json:"updatedAt,omitempty"`
}

func (ManageRolesRightsConfigs) TableName() string {
	return "manage_roles_rights_configs"
}

func NewManageRolesRightsConfigs() ManageRolesRightsConfigs {
    return ManageRolesRightsConfigs{}
}
