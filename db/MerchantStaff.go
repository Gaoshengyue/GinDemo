package db

import (
	"fmt"
	"gorm.io/gorm"
)

type MerchantStaff struct {
	ID         *uint  `gorm:"column:id;primaryKey"`
	Username   string `gorm:"column:username;size:255"`
	ApiToken   string `gorm:"column:api_token;"`
	MerchantId int    `gorm:"column:merchant_id"`
	DepartId   int    `gorm:"column:depart_id"`
	Linkman    string `gorm:"column:linkman"`
}

func (u *MerchantStaff) TableName() string {
	return "merchant_staff" // 指定自定义的表名
}

// CreateUser 创建用户
func (u *MerchantStaff) CreateUser(db *gorm.DB) error {

	// 查询是否存在重复记录
	var count int64
	if err := db.Model(&MerchantStaff{}).Where(MerchantStaff{Username: u.Username, MerchantId: u.MerchantId}).Count(&count).Error; err != nil {
		return err
	}
	// 如果存在重复记录，则不执行插入操作
	if count > 0 {
		return fmt.Errorf("duplicate user")
	}

	if err := db.Create(u).Error; err != nil {
		db.Rollback()
		return err
	}
	// 手动提交事务
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	// 创建用户
	return nil
}
