package model
import "git.wanpinghui.com/WPH/go_common/wph/date"

// Created by CodeMachine on 2020-05-10 15:46:04.688232 +0800 CST m=+0.091548008
type ManageRolesRightsConfigs struct {
    AddActivity          bool           `gorm:"column:add_activity;type:tinyint(1)" json:"addActivity"`
    AddShop              bool           `gorm:"column:add_shop;type:tinyint(1)" json:"addShop"`
    CreatedAt            date.Datetime  `gorm:"column:created_at;type:timestamp" json:"createdAt,omitempty"`
    DataQuery            bool           `gorm:"column:data_query;type:tinyint(1)" json:"dataQuery"`
    DeletedAt            *date.Datetime `gorm:"column:deleted_at;type:timestamp" json:"deletedAt,omitempty"`
    GiftCardManagement   bool           `gorm:"column:gift_card_management;type:tinyint(1)" json:"giftCardManagement"`
    ID                   uint           `gorm:"primary_key;AUTO_INCREMENT;not null;column:id;type:int unsigned" json:"id"`
    ManageDepartment     bool           `gorm:"column:manage_department;type:tinyint(1)" json:"manageDepartment"`
    ManageRole           bool           `gorm:"column:manage_role;type:tinyint(1)" json:"manageRole"`
    ManageUser           bool           `gorm:"column:manage_user;type:tinyint(1)" json:"manageUser"`
    OperationsManagement bool           `gorm:"column:operations_management;type:tinyint(1)" json:"operationsManagement"`
    PicturePlatform      bool           `gorm:"column:picture_platform;type:tinyint(1)" json:"picturePlatform"`
    PictureShop          bool           `gorm:"column:picture_shop;type:tinyint(1)" json:"pictureShop"`
    QueryActivity        bool           `gorm:"column:query_activity;type:tinyint(1)" json:"queryActivity"`
    QueryOrder           bool           `gorm:"column:query_order;type:tinyint(1)" json:"queryOrder"`
    QueryReconciliation  bool           `gorm:"column:query_reconciliation;type:tinyint(1)" json:"queryReconciliation"`
    QuerySettlement      bool           `gorm:"column:query_settlement;type:tinyint(1)" json:"querySettlement"`
    QueryShop            bool           `gorm:"column:query_shop;type:tinyint(1)" json:"queryShop"`
    QueryUser            bool           `gorm:"column:query_user;type:tinyint(1)" json:"queryUser"`
    RoleId               uint           `gorm:"column:role_id;type:int unsigned" json:"roleId"`
    SeparateAccount      bool           `gorm:"column:separate_account;type:tinyint(1)" json:"separateAccount"`
    ShopCoupon           bool           `gorm:"column:shop_coupon;type:tinyint(1)" json:"shopCoupon"`
    ShopManagement       bool           `gorm:"column:shop_management;type:tinyint(1)" json:"shopManagement"`
    UpdatedAt            date.Datetime  `gorm:"column:updated_at;type:timestamp" json:"updatedAt,omitempty"`
}

func (ManageRolesRightsConfigs) TableName() string {
	return "manage_roles_rights_configs"
}
func (this ManageRolesRightsConfigs) Get(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).First(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ManageRolesRightsConfigs) List(selSql,whereSql string,resp interface{}) error {
    if err := AppGormClient.Client.Model(this.TableName).
		Select(selSql).Find(resp,whereSql).Error; err != nil {
		return err
	}
	return nil
}
func (this ManageRolesRightsConfigs) Create() string {
    if AppGormClient.Client.NewRecord(&this) {
		if err := AppGormClient.Client.Create(&this).Error; err != nil {
			return err
		}
	}
	return nil
}
func (this ManageRolesRightsConfigs) Delete() string {
	if err := AppGormClient.Client.Where(whereSql).Delete(ManageRolesRightsConfigs{}).Error; err != nil {
		return err
	}
	return nil
}

func NewManageRolesRightsConfigs() ManageRolesRightsConfigs {
    return ManageRolesRightsConfigs{}
}

