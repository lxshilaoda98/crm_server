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

type CdrLineApi struct {
}

var cdrLineService = service.ServiceGroupApp.AutoCodeServiceGroup.CdrLineService

// CreateCdrLine 创建CdrLine
// @Tags CdrLine
// @Summary 创建CdrLine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CdrLine true "创建CdrLine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cdrLine/createCdrLine [post]
func (cdrLineApi *CdrLineApi) CreateCdrLine(c *gin.Context) {
	var cdrLine autocode.CdrLine
	_ = c.ShouldBindJSON(&cdrLine)
	if err := cdrLineService.CreateCdrLine(cdrLine); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCdrLine 删除CdrLine
// @Tags CdrLine
// @Summary 删除CdrLine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CdrLine true "删除CdrLine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cdrLine/deleteCdrLine [delete]
func (cdrLineApi *CdrLineApi) DeleteCdrLine(c *gin.Context) {
	var cdrLine autocode.CdrLine
	_ = c.ShouldBindJSON(&cdrLine)
	if err := cdrLineService.DeleteCdrLine(cdrLine); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCdrLineByIds 批量删除CdrLine
// @Tags CdrLine
// @Summary 批量删除CdrLine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CdrLine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cdrLine/deleteCdrLineByIds [delete]
func (cdrLineApi *CdrLineApi) DeleteCdrLineByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := cdrLineService.DeleteCdrLineByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCdrLine 更新CdrLine
// @Tags CdrLine
// @Summary 更新CdrLine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CdrLine true "更新CdrLine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cdrLine/updateCdrLine [put]
func (cdrLineApi *CdrLineApi) UpdateCdrLine(c *gin.Context) {
	var cdrLine autocode.CdrLine
	_ = c.ShouldBindJSON(&cdrLine)
	if err := cdrLineService.UpdateCdrLine(cdrLine); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCdrLine 用id查询CdrLine
// @Tags CdrLine
// @Summary 用id查询CdrLine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CdrLine true "用id查询CdrLine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cdrLine/findCdrLine [get]
func (cdrLineApi *CdrLineApi) FindCdrLine(c *gin.Context) {
	var cdrLine autocode.CdrLine
	_ = c.ShouldBindQuery(&cdrLine)
	if err, recdrLine := cdrLineService.GetCdrLine(cdrLine.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recdrLine": recdrLine}, c)
	}
}

// GetCdrLineList 分页获取CdrLine列表
// @Tags CdrLine
// @Summary 分页获取CdrLine列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CdrLineSearch true "分页获取CdrLine列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cdrLine/getCdrLineList [get]
func (cdrLineApi *CdrLineApi) GetCdrLineList(c *gin.Context) {
	var pageInfo autocodeReq.CdrLineSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := cdrLineService.GetCdrLineInfoList(pageInfo); err != nil {
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
