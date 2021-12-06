package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SipRegistrationsRouter struct {
}

// InitSipRegistrationsRouter 初始化 SipRegistrations 路由信息
func (s *SipRegistrationsRouter) InitSipRegistrationsRouter(Router *gin.RouterGroup) {
	sipRegistrationsRouter := Router.Group("sipRegistrations").Use(middleware.OperationRecord())
	sipRegistrationsRouterWithoutRecord := Router.Group("sipRegistrations")
	var sipRegistrationsApi = v1.ApiGroupApp.AutoCodeApiGroup.SipRegistrationsApi
	{
		sipRegistrationsRouter.POST("createSipRegistrations", sipRegistrationsApi.CreateSipRegistrations)             // 新建SipRegistrations
		sipRegistrationsRouter.DELETE("deleteSipRegistrations", sipRegistrationsApi.DeleteSipRegistrations)           // 删除SipRegistrations
		sipRegistrationsRouter.DELETE("deleteSipRegistrationsByIds", sipRegistrationsApi.DeleteSipRegistrationsByIds) // 批量删除SipRegistrations
		sipRegistrationsRouter.PUT("updateSipRegistrations", sipRegistrationsApi.UpdateSipRegistrations)              // 更新SipRegistrations
	}
	{
		sipRegistrationsRouterWithoutRecord.GET("findSipRegistrations", sipRegistrationsApi.FindSipRegistrations)       // 根据ID获取SipRegistrations
		sipRegistrationsRouterWithoutRecord.GET("getSipRegistrationsList", sipRegistrationsApi.GetSipRegistrationsList) // 获取SipRegistrations列表
	}
}
