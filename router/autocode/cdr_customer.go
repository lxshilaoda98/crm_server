package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CdrCustomerRouter struct {
}

// InitCdrCustomerRouter 初始化 CdrCustomer 路由信息
func (s *CdrCustomerRouter) InitCdrCustomerRouter(Router *gin.RouterGroup) {
	cdrCustomerRouter := Router.Group("cdrCustomer").Use(middleware.OperationRecord())
	cdrCustomerRouterWithoutRecord := Router.Group("cdrCustomer")
	var cdrCustomerApi = v1.ApiGroupApp.AutoCodeApiGroup.CdrCustomerApi
	{
		cdrCustomerRouter.POST("createCdrCustomer", cdrCustomerApi.CreateCdrCustomer)             // 新建CdrCustomer
		cdrCustomerRouter.DELETE("deleteCdrCustomer", cdrCustomerApi.DeleteCdrCustomer)           // 删除CdrCustomer
		cdrCustomerRouter.DELETE("deleteCdrCustomerByIds", cdrCustomerApi.DeleteCdrCustomerByIds) // 批量删除CdrCustomer
		cdrCustomerRouter.PUT("updateCdrCustomer", cdrCustomerApi.UpdateCdrCustomer)              // 更新CdrCustomer
	}
	{
		cdrCustomerRouterWithoutRecord.GET("findCdrCustomer", cdrCustomerApi.FindCdrCustomer)       // 根据ID获取CdrCustomer
		cdrCustomerRouterWithoutRecord.GET("getCdrCustomerList", cdrCustomerApi.GetCdrCustomerList) // 获取CdrCustomer列表
	}
}
