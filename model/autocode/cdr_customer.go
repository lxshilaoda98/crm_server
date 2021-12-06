// 自动生成模板CdrCustomer
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CdrCustomer 结构体
// 如果含有time.Time 请自行import time包
type CdrCustomer struct {
	global.GVA_MODEL
	User       string `json:"User" form:"User" gorm:"column:User;comment:账号信息[线路号];type:varchar(50);"`
	Name       string `json:"Name" form:"Name" gorm:"column:Name;comment:名称;type:varchar(50);"`
	Balance    string `json:"Balance" form:"Balance" gorm:"column:Balance;comment:余额;type:varchar(50);"`
	Overdraft  string `json:"Overdraft" form:"Overdraft" gorm:"column:Overdraft;comment:透支限额;type:varchar(50);"`
	SetmealOid *int   `json:"SetmealOid" form:"SetmealOid" gorm:"column:SetmealOid;comment:套餐;type:int"`
	PrantD     string `json:"PrantD" form:"PrantD" gorm:"column:PrantD;comment:来区分代理商还是普通用户;type:varchar(10);"`
}

// TableName CdrCustomer 表名
func (CdrCustomer) TableName() string {
	return "cdr_customer"
}
