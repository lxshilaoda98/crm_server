package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CdrCustomerApi struct {
}

var cdrCustomerService = service.ServiceGroupApp.AutoCodeServiceGroup.CdrCustomerService

// CreateCdrCustomer 创建CdrCustomer
// @Tags CdrCustomer
// @Summary 创建CdrCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CdrCustomer true "创建CdrCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cdrCustomer/createCdrCustomer [post]
func (cdrCustomerApi *CdrCustomerApi) CreateCdrCustomer(c *gin.Context) {
	var cdrCustomer autocode.CdrCustomer
	_ = c.ShouldBindJSON(&cdrCustomer)
	if err := cdrCustomerService.CreateCdrCustomer(cdrCustomer); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCdrCustomer 删除CdrCustomer
// @Tags CdrCustomer
// @Summary 删除CdrCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CdrCustomer true "删除CdrCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cdrCustomer/deleteCdrCustomer [delete]
func (cdrCustomerApi *CdrCustomerApi) DeleteCdrCustomer(c *gin.Context) {
	var cdrCustomer autocode.CdrCustomer
	_ = c.ShouldBindJSON(&cdrCustomer)
	if err := cdrCustomerService.DeleteCdrCustomer(cdrCustomer); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCdrCustomerByIds 批量删除CdrCustomer
// @Tags CdrCustomer
// @Summary 批量删除CdrCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CdrCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cdrCustomer/deleteCdrCustomerByIds [delete]
func (cdrCustomerApi *CdrCustomerApi) DeleteCdrCustomerByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := cdrCustomerService.DeleteCdrCustomerByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCdrCustomer 更新CdrCustomer
// @Tags CdrCustomer
// @Summary 更新CdrCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CdrCustomer true "更新CdrCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cdrCustomer/updateCdrCustomer [put]
func (cdrCustomerApi *CdrCustomerApi) UpdateCdrCustomer(c *gin.Context) {
	var cdrCustomer autocode.CdrCustomer
	_ = c.ShouldBindJSON(&cdrCustomer)
	if err := cdrCustomerService.UpdateCdrCustomer(cdrCustomer); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCdrCustomer 用id查询CdrCustomer
// @Tags CdrCustomer
// @Summary 用id查询CdrCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CdrCustomer true "用id查询CdrCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cdrCustomer/findCdrCustomer [get]
func (cdrCustomerApi *CdrCustomerApi) FindCdrCustomer(c *gin.Context) {
	var cdrCustomer autocode.CdrCustomer
	_ = c.ShouldBindQuery(&cdrCustomer)
	if err, recdrCustomer := cdrCustomerService.GetCdrCustomer(cdrCustomer.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recdrCustomer": recdrCustomer}, c)
	}
}

// GetCdrCustomerList 分页获取CdrCustomer列表
// @Tags CdrCustomer
// @Summary 分页获取CdrCustomer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CdrCustomerSearch true "分页获取CdrCustomer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cdrCustomer/getCdrCustomerList [get]
func (cdrCustomerApi *CdrCustomerApi) GetCdrCustomerList(c *gin.Context) {
	var pageInfo autocodeReq.CdrCustomerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := cdrCustomerService.GetCdrCustomerInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
