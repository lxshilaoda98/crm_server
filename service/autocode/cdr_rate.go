package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CdrRateService struct {
}

// CreateCdrRate 创建CdrRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrRateService *CdrRateService) CreateCdrRate(cdrRate autocode.CdrRate) (err error) {
	err = global.GVA_DB.Create(&cdrRate).Error
	return err
}

// DeleteCdrRate 删除CdrRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrRateService *CdrRateService) DeleteCdrRate(cdrRate autocode.CdrRate) (err error) {
	err = global.GVA_DB.Delete(&cdrRate).Error
	return err
}

// DeleteCdrRateByIds 批量删除CdrRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrRateService *CdrRateService) DeleteCdrRateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CdrRate{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCdrRate 更新CdrRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrRateService *CdrRateService) UpdateCdrRate(cdrRate autocode.CdrRate) (err error) {
	err = global.GVA_DB.Save(&cdrRate).Error
	return err
}

// GetCdrRate 根据id获取CdrRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrRateService *CdrRateService) GetCdrRate(id uint) (err error, cdrRate autocode.CdrRate) {
	err = global.GVA_DB.Where("id = ?", id).First(&cdrRate).Error
	return
}

// GetCdrRateInfoList 分页获取CdrRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrRateService *CdrRateService) GetCdrRateInfoList(info autoCodeReq.CdrRateSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.CdrRate{})
	var cdrRates []autocode.CdrRate
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("Name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&cdrRates).Error
	return err, cdrRates, total
}
