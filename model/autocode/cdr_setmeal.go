// 自动生成模板CdrSetmeal
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CdrSetmeal 结构体
// 如果含有time.Time 请自行import time包
type CdrSetmeal struct {
	global.GVA_MODEL
	Name         string `json:"Name" form:"Name" gorm:"column:Name;comment:套餐名称;type:varchar(50);"`
	Remarks      string `json:"Remarks" form:"Remarks" gorm:"column:Remarks;comment:备注;type:varchar(100);"`
	RentalUnit   string `json:"RentalUnit" form:"RentalUnit" gorm:"column:RentalUnit;comment:租用单位[月];type:varchar(5);"`
	RentalMoney  string `json:"RentalMoney" form:"RentalMoney" gorm:"column:RentalMoney;comment:租用租金;type:varchar(255);"`
	RentalMinute string `json:"RentalMinute" form:"RentalMinute" gorm:"column:RentalMinute;comment:免费分钟数;type:varchar(255);"`
}

// TableName CdrSetmeal 表名
func (CdrSetmeal) TableName() string {
	return "cdr_setmeal"
}
