// 自动生成模板CdrRate
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CdrRate 结构体
// 如果含有time.Time 请自行import time包
type CdrRate struct {
	global.GVA_MODEL
	Name             string `json:"Name" form:"Name" gorm:"column:Name;comment:费率名称;type:varchar(50);"`
	Remarks          string `json:"Remarks" form:"Remarks" gorm:"column:Remarks;comment:备注;type:varchar(100);"`
	Create           string `json:"Create" form:"Create" gorm:"column:Create;comment:创建人;type:varchar(10);"`
	ChargingTime     *int   `json:"ChargingTime" form:"ChargingTime" gorm:"column:ChargingTime;comment:规则时间;type:int"`
	LocalMoneyS      string `json:"LocalMoneyS" form:"LocalMoneyS" gorm:"column:LocalMoneyS;comment:超规则收费;type:varchar(10);"`
	LocalMoney       string `json:"LocalMoney" form:"LocalMoney" gorm:"column:LocalMoney;comment:本地市话;type:varchar(10);"`
	DomesticMoney    string `json:"DomesticMoney" form:"DomesticMoney" gorm:"column:DomesticMoney;comment:国内长途;type:varchar(10);"`
	InternationMoney string `json:"InternationMoney" form:"InternationMoney" gorm:"column:InternationMoney;comment:国际长途;type:varchar(10);"`
	SetMeal          *int   `json:"SetMeal" form:"SetMeal" gorm:"column:SetMeal;comment:套餐;type:int"`
}

// TableName CdrRate 表名
func (CdrRate) TableName() string {
	return "cdr_rate"
}
