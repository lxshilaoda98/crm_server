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

type SipRegistrationsApi struct {
}

var sipRegistrationsService = service.ServiceGroupApp.AutoCodeServiceGroup.SipRegistrationsService

// CreateSipRegistrations 创建SipRegistrations
// @Tags SipRegistrations
// @Summary 创建SipRegistrations
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SipRegistrations true "创建SipRegistrations"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sipRegistrations/createSipRegistrations [post]
func (sipRegistrationsApi *SipRegistrationsApi) CreateSipRegistrations(c *gin.Context) {
	var sipRegistrations autocode.SipRegistrations
	_ = c.ShouldBindJSON(&sipRegistrations)
	if err := sipRegistrationsService.CreateSipRegistrations(sipRegistrations); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSipRegistrations 删除SipRegistrations
// @Tags SipRegistrations
// @Summary 删除SipRegistrations
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SipRegistrations true "删除SipRegistrations"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sipRegistrations/deleteSipRegistrations [delete]
func (sipRegistrationsApi *SipRegistrationsApi) DeleteSipRegistrations(c *gin.Context) {
	var sipRegistrations autocode.SipRegistrations
	_ = c.ShouldBindJSON(&sipRegistrations)
	if err := sipRegistrationsService.DeleteSipRegistrations(sipRegistrations); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSipRegistrationsByIds 批量删除SipRegistrations
// @Tags SipRegistrations
// @Summary 批量删除SipRegistrations
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SipRegistrations"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sipRegistrations/deleteSipRegistrationsByIds [delete]
func (sipRegistrationsApi *SipRegistrationsApi) DeleteSipRegistrationsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sipRegistrationsService.DeleteSipRegistrationsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSipRegistrations 更新SipRegistrations
// @Tags SipRegistrations
// @Summary 更新SipRegistrations
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.SipRegistrations true "更新SipRegistrations"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sipRegistrations/updateSipRegistrations [put]
func (sipRegistrationsApi *SipRegistrationsApi) UpdateSipRegistrations(c *gin.Context) {
	var sipRegistrations autocode.SipRegistrations
	_ = c.ShouldBindJSON(&sipRegistrations)
	if err := sipRegistrationsService.UpdateSipRegistrations(sipRegistrations); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSipRegistrations 用id查询SipRegistrations
// @Tags SipRegistrations
// @Summary 用id查询SipRegistrations
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.SipRegistrations true "用id查询SipRegistrations"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sipRegistrations/findSipRegistrations [get]
func (sipRegistrationsApi *SipRegistrationsApi) FindSipRegistrations(c *gin.Context) {
	var sipRegistrations autocode.SipRegistrations
	_ = c.ShouldBindQuery(&sipRegistrations)
	if err, resipRegistrations := sipRegistrationsService.GetSipRegistrations(sipRegistrations.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resipRegistrations": resipRegistrations}, c)
	}
}

// GetSipRegistrationsList 分页获取SipRegistrations列表
// @Tags SipRegistrations
// @Summary 分页获取SipRegistrations列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.SipRegistrationsSearch true "分页获取SipRegistrations列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sipRegistrations/getSipRegistrationsList [get]
func (sipRegistrationsApi *SipRegistrationsApi) GetSipRegistrationsList(c *gin.Context) {
	var pageInfo autocodeReq.SipRegistrationsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sipRegistrationsService.GetSipRegistrationsInfoList(pageInfo); err != nil {
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
