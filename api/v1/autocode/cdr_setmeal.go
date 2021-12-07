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

type CdrSetmealApi struct {
}

var cdrSetmealService = service.ServiceGroupApp.AutoCodeServiceGroup.CdrSetmealService

// CreateCdrSetmeal 创建CdrSetmeal
// @Tags CdrSetmeal
// @Summary 创建CdrSetmeal
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CdrSetmeal true "创建CdrSetmeal"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cdrSetmeal/createCdrSetmeal [post]
func (cdrSetmealApi *CdrSetmealApi) CreateCdrSetmeal(c *gin.Context) {
	var cdrSetmeal autocode.CdrSetmeal
	_ = c.ShouldBindJSON(&cdrSetmeal)
	if err := cdrSetmealService.CreateCdrSetmeal(cdrSetmeal); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCdrSetmeal 删除CdrSetmeal
// @Tags CdrSetmeal
// @Summary 删除CdrSetmeal
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CdrSetmeal true "删除CdrSetmeal"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cdrSetmeal/deleteCdrSetmeal [delete]
func (cdrSetmealApi *CdrSetmealApi) DeleteCdrSetmeal(c *gin.Context) {
	var cdrSetmeal autocode.CdrSetmeal
	_ = c.ShouldBindJSON(&cdrSetmeal)
	if err := cdrSetmealService.DeleteCdrSetmeal(cdrSetmeal); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCdrSetmealByIds 批量删除CdrSetmeal
// @Tags CdrSetmeal
// @Summary 批量删除CdrSetmeal
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CdrSetmeal"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cdrSetmeal/deleteCdrSetmealByIds [delete]
func (cdrSetmealApi *CdrSetmealApi) DeleteCdrSetmealByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := cdrSetmealService.DeleteCdrSetmealByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCdrSetmeal 更新CdrSetmeal
// @Tags CdrSetmeal
// @Summary 更新CdrSetmeal
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CdrSetmeal true "更新CdrSetmeal"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cdrSetmeal/updateCdrSetmeal [put]
func (cdrSetmealApi *CdrSetmealApi) UpdateCdrSetmeal(c *gin.Context) {
	var cdrSetmeal autocode.CdrSetmeal
	_ = c.ShouldBindJSON(&cdrSetmeal)
	if err := cdrSetmealService.UpdateCdrSetmeal(cdrSetmeal); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCdrSetmeal 用id查询CdrSetmeal
// @Tags CdrSetmeal
// @Summary 用id查询CdrSetmeal
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CdrSetmeal true "用id查询CdrSetmeal"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cdrSetmeal/findCdrSetmeal [get]
func (cdrSetmealApi *CdrSetmealApi) FindCdrSetmeal(c *gin.Context) {
	var cdrSetmeal autocode.CdrSetmeal
	_ = c.ShouldBindQuery(&cdrSetmeal)
	if err, recdrSetmeal := cdrSetmealService.GetCdrSetmeal(cdrSetmeal.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recdrSetmeal": recdrSetmeal}, c)
	}
}

// GetCdrSetmealList 分页获取CdrSetmeal列表
// @Tags CdrSetmeal
// @Summary 分页获取CdrSetmeal列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CdrSetmealSearch true "分页获取CdrSetmeal列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cdrSetmeal/getCdrSetmealList [get]
func (cdrSetmealApi *CdrSetmealApi) GetCdrSetmealList(c *gin.Context) {
	var pageInfo autocodeReq.CdrSetmealSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := cdrSetmealService.GetCdrSetmealInfoList(pageInfo); err != nil {
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
