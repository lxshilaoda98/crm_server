package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CdrLineRouter struct {
}

// InitCdrLineRouter 初始化 CdrLine 路由信息
func (s *CdrLineRouter) InitCdrLineRouter(Router *gin.RouterGroup) {
	cdrLineRouter := Router.Group("cdrLine").Use(middleware.OperationRecord())
	cdrLineRouterWithoutRecord := Router.Group("cdrLine")
	var cdrLineApi = v1.ApiGroupApp.AutoCodeApiGroup.CdrLineApi
	{
		cdrLineRouter.POST("createCdrLine", cdrLineApi.CreateCdrLine)             // 新建CdrLine
		cdrLineRouter.DELETE("deleteCdrLine", cdrLineApi.DeleteCdrLine)           // 删除CdrLine
		cdrLineRouter.DELETE("deleteCdrLineByIds", cdrLineApi.DeleteCdrLineByIds) // 批量删除CdrLine
		cdrLineRouter.PUT("updateCdrLine", cdrLineApi.UpdateCdrLine)              // 更新CdrLine
	}
	{
		cdrLineRouterWithoutRecord.GET("findCdrLine", cdrLineApi.FindCdrLine)       // 根据ID获取CdrLine
		cdrLineRouterWithoutRecord.GET("getCdrLineList", cdrLineApi.GetCdrLineList) // 获取CdrLine列表
	}
}
