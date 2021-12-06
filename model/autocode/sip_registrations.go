// 自动生成模板SipRegistrations
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SipRegistrations 结构体
// 如果含有time.Time 请自行import time包
type SipRegistrations struct {
	global.GVA_MODEL
	CallId        string `json:"callId" form:"callId" gorm:"column:call_id;comment:;type:varchar(255);"`
	SipUser       string `json:"sipUser" form:"sipUser" gorm:"column:sip_user;comment:注册用户;type:varchar(255);"`
	SipHost       string `json:"sipHost" form:"sipHost" gorm:"column:sip_host;comment:注册地址;type:varchar(255);"`
	PresenceHosts string `json:"presenceHosts" form:"presenceHosts" gorm:"column:presence_hosts;comment:地址;type:varchar(255);"`
	Contact       string `json:"contact" form:"contact" gorm:"column:contact;comment:联系地址;type:varchar(1024);"`
	Status        string `json:"status" form:"status" gorm:"column:status;comment:状态;type:varchar(255);"`
	PingStatus    string `json:"pingStatus" form:"pingStatus" gorm:"column:ping_status;comment:ping状态;type:varchar(255);"`
	PingCount     *int   `json:"pingCount" form:"pingCount" gorm:"column:ping_count;comment:ping次数;type:int"`
	PingTime      *int   `json:"pingTime" form:"pingTime" gorm:"column:ping_time;comment:ping时间;type:bigint"`
	Expires       *int   `json:"expires" form:"expires" gorm:"column:expires;comment:失效;type:bigint"`
	PingExpires   *int   `json:"pingExpires" form:"pingExpires" gorm:"column:ping_expires;comment:ping失效;type:int"`
	UserAgent     string `json:"userAgent" form:"userAgent" gorm:"column:user_agent;comment:用户代理;type:varchar(255);"`
	ServerUser    string `json:"serverUser" form:"serverUser" gorm:"column:server_user;comment:服务器用户;type:varchar(255);"`
	ServerHost    string `json:"serverHost" form:"serverHost" gorm:"column:server_host;comment:服务器地址;type:varchar(255);"`
	ProfileName   string `json:"profileName" form:"profileName" gorm:"column:profile_name;comment:配置文件名称;type:varchar(255);"`
	Hostname      string `json:"hostname" form:"hostname" gorm:"column:hostname;comment:主机名;type:varchar(255);"`
	NetworkIp     string `json:"networkIp" form:"networkIp" gorm:"column:network_ip;comment:网络ip;type:varchar(255);"`
	NetworkPort   string `json:"networkPort" form:"networkPort" gorm:"column:network_port;comment:网络端口;type:varchar(6);"`
	SipUsername   string `json:"sipUsername" form:"sipUsername" gorm:"column:sip_username;comment:sip用户名;type:varchar(255);"`
	SipRealm      string `json:"sipRealm" form:"sipRealm" gorm:"column:sip_realm;comment:sip域;type:varchar(255);"`
}

// TableName SipRegistrations 表名
func (SipRegistrations) TableName() string {
	return "sip_registrations"
}
