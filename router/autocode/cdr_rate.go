package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CdrRateRouter struct {
}

// InitCdrRateRouter 初始化 CdrRate 路由信息
func (s *CdrRateRouter) InitCdrRateRouter(Router *gin.RouterGroup) {
	cdrRateRouter := Router.Group("cdrRate").Use(middleware.OperationRecord())
	cdrRateRouterWithoutRecord := Router.Group("cdrRate")
	var cdrRateApi = v1.ApiGroupApp.AutoCodeApiGroup.CdrRateApi
	{
		cdrRateRouter.POST("createCdrRate", cdrRateApi.CreateCdrRate)             // 新建CdrRate
		cdrRateRouter.DELETE("deleteCdrRate", cdrRateApi.DeleteCdrRate)           // 删除CdrRate
		cdrRateRouter.DELETE("deleteCdrRateByIds", cdrRateApi.DeleteCdrRateByIds) // 批量删除CdrRate
		cdrRateRouter.PUT("updateCdrRate", cdrRateApi.UpdateCdrRate)              // 更新CdrRate
	}
	{
		cdrRateRouterWithoutRecord.GET("findCdrRate", cdrRateApi.FindCdrRate)       // 根据ID获取CdrRate
		cdrRateRouterWithoutRecord.GET("getCdrRateList", cdrRateApi.GetCdrRateList) // 获取CdrRate列表
	}
}
