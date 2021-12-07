// 自动生成模板CdrLine
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CdrLine 结构体
// 如果含有time.Time 请自行import time包
type CdrLine struct {
	global.GVA_MODEL
	User        string `json:"User" form:"User" gorm:"column:User;comment:线路号码;type:varchar(50);"`
	PassWord    string `json:"PassWord" form:"PassWord" gorm:"column:PassWord;comment:密码;type:varchar(50);"`
	Consumption string `json:"Consumption" form:"Consumption" gorm:"column:Consumption;comment:消费;type:varchar(50);"`
	Costomer    string `json:"Costomer" form:"Costomer" gorm:"column:Costomer;comment:所属账号;type:varchar(50);"`
	Gs          string `json:"Gs" form:"Gs" gorm:"column:Gs;comment:归属;type:varchar(10);"`
	Qh          string `json:"Qh" form:"Qh" gorm:"column:Qh;comment:区号;type:varchar(10);"`
}

// TableName CdrLine 表名
func (CdrLine) TableName() string {
	return "cdr_line"
}
