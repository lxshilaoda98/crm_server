package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CdrSetmealRouter struct {
}

// InitCdrSetmealRouter 初始化 CdrSetmeal 路由信息
func (s *CdrSetmealRouter) InitCdrSetmealRouter(Router *gin.RouterGroup) {
	cdrSetmealRouter := Router.Group("cdrSetmeal").Use(middleware.OperationRecord())
	cdrSetmealRouterWithoutRecord := Router.Group("cdrSetmeal")
	var cdrSetmealApi = v1.ApiGroupApp.AutoCodeApiGroup.CdrSetmealApi
	{
		cdrSetmealRouter.POST("createCdrSetmeal", cdrSetmealApi.CreateCdrSetmeal)             // 新建CdrSetmeal
		cdrSetmealRouter.DELETE("deleteCdrSetmeal", cdrSetmealApi.DeleteCdrSetmeal)           // 删除CdrSetmeal
		cdrSetmealRouter.DELETE("deleteCdrSetmealByIds", cdrSetmealApi.DeleteCdrSetmealByIds) // 批量删除CdrSetmeal
		cdrSetmealRouter.PUT("updateCdrSetmeal", cdrSetmealApi.UpdateCdrSetmeal)              // 更新CdrSetmeal
	}
	{
		cdrSetmealRouterWithoutRecord.GET("findCdrSetmeal", cdrSetmealApi.FindCdrSetmeal)       // 根据ID获取CdrSetmeal
		cdrSetmealRouterWithoutRecord.GET("getCdrSetmealList", cdrSetmealApi.GetCdrSetmealList) // 获取CdrSetmeal列表
	}
}
