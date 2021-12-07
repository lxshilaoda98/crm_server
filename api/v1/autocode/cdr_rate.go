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

type CdrRateApi struct {
}

var cdrRateService = service.ServiceGroupApp.AutoCodeServiceGroup.CdrRateService

// CreateCdrRate 创建CdrRate
// @Tags CdrRate
// @Summary 创建CdrRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CdrRate true "创建CdrRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cdrRate/createCdrRate [post]
func (cdrRateApi *CdrRateApi) CreateCdrRate(c *gin.Context) {
	var cdrRate autocode.CdrRate
	_ = c.ShouldBindJSON(&cdrRate)
	if err := cdrRateService.CreateCdrRate(cdrRate); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCdrRate 删除CdrRate
// @Tags CdrRate
// @Summary 删除CdrRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CdrRate true "删除CdrRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cdrRate/deleteCdrRate [delete]
func (cdrRateApi *CdrRateApi) DeleteCdrRate(c *gin.Context) {
	var cdrRate autocode.CdrRate
	_ = c.ShouldBindJSON(&cdrRate)
	if err := cdrRateService.DeleteCdrRate(cdrRate); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCdrRateByIds 批量删除CdrRate
// @Tags CdrRate
// @Summary 批量删除CdrRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CdrRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cdrRate/deleteCdrRateByIds [delete]
func (cdrRateApi *CdrRateApi) DeleteCdrRateByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := cdrRateService.DeleteCdrRateByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCdrRate 更新CdrRate
// @Tags CdrRate
// @Summary 更新CdrRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CdrRate true "更新CdrRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cdrRate/updateCdrRate [put]
func (cdrRateApi *CdrRateApi) UpdateCdrRate(c *gin.Context) {
	var cdrRate autocode.CdrRate
	_ = c.ShouldBindJSON(&cdrRate)
	if err := cdrRateService.UpdateCdrRate(cdrRate); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCdrRate 用id查询CdrRate
// @Tags CdrRate
// @Summary 用id查询CdrRate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CdrRate true "用id查询CdrRate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cdrRate/findCdrRate [get]
func (cdrRateApi *CdrRateApi) FindCdrRate(c *gin.Context) {
	var cdrRate autocode.CdrRate
	_ = c.ShouldBindQuery(&cdrRate)
	if err, recdrRate := cdrRateService.GetCdrRate(cdrRate.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recdrRate": recdrRate}, c)
	}
}

// GetCdrRateList 分页获取CdrRate列表
// @Tags CdrRate
// @Summary 分页获取CdrRate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CdrRateSearch true "分页获取CdrRate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cdrRate/getCdrRateList [get]
func (cdrRateApi *CdrRateApi) GetCdrRateList(c *gin.Context) {
	var pageInfo autocodeReq.CdrRateSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := cdrRateService.GetCdrRateInfoList(pageInfo); err != nil {
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
