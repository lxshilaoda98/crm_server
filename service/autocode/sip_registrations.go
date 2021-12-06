package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SipRegistrationsService struct {
}

// CreateSipRegistrations 创建SipRegistrations记录
// Author [piexlmax](https://github.com/piexlmax)
func (sipRegistrationsService *SipRegistrationsService) CreateSipRegistrations(sipRegistrations autocode.SipRegistrations) (err error) {
	err = global.GVA_DB.Create(&sipRegistrations).Error
	return err
}

// DeleteSipRegistrations 删除SipRegistrations记录
// Author [piexlmax](https://github.com/piexlmax)
func (sipRegistrationsService *SipRegistrationsService) DeleteSipRegistrations(sipRegistrations autocode.SipRegistrations) (err error) {
	err = global.GVA_DB.Delete(&sipRegistrations).Error
	return err
}

// DeleteSipRegistrationsByIds 批量删除SipRegistrations记录
// Author [piexlmax](https://github.com/piexlmax)
func (sipRegistrationsService *SipRegistrationsService) DeleteSipRegistrationsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.SipRegistrations{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSipRegistrations 更新SipRegistrations记录
// Author [piexlmax](https://github.com/piexlmax)
func (sipRegistrationsService *SipRegistrationsService) UpdateSipRegistrations(sipRegistrations autocode.SipRegistrations) (err error) {
	err = global.GVA_DB.Save(&sipRegistrations).Error
	return err
}

// GetSipRegistrations 根据id获取SipRegistrations记录
// Author [piexlmax](https://github.com/piexlmax)
func (sipRegistrationsService *SipRegistrationsService) GetSipRegistrations(id uint) (err error, sipRegistrations autocode.SipRegistrations) {
	err = global.GVA_DB.Where("id = ?", id).First(&sipRegistrations).Error
	return
}

// GetSipRegistrationsInfoList 分页获取SipRegistrations记录
// Author [piexlmax](https://github.com/piexlmax)
func (sipRegistrationsService *SipRegistrationsService) GetSipRegistrationsInfoList(info autoCodeReq.SipRegistrationsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.SipRegistrations{})
	var sipRegistrationss []autocode.SipRegistrations
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.SipUser != "" {
		db = db.Where("sip_user LIKE ?", "%"+info.SipUser+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&sipRegistrationss).Error
	return err, sipRegistrationss, total
}
