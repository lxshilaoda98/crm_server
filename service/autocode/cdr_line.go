package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CdrLineService struct {
}

// CreateCdrLine 创建CdrLine记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrLineService *CdrLineService) CreateCdrLine(cdrLine autocode.CdrLine) (err error) {
	err = global.GVA_DB.Create(&cdrLine).Error
	return err
}

// DeleteCdrLine 删除CdrLine记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrLineService *CdrLineService) DeleteCdrLine(cdrLine autocode.CdrLine) (err error) {
	err = global.GVA_DB.Delete(&cdrLine).Error
	return err
}

// DeleteCdrLineByIds 批量删除CdrLine记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrLineService *CdrLineService) DeleteCdrLineByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CdrLine{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCdrLine 更新CdrLine记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrLineService *CdrLineService) UpdateCdrLine(cdrLine autocode.CdrLine) (err error) {
	err = global.GVA_DB.Save(&cdrLine).Error
	return err
}

// GetCdrLine 根据id获取CdrLine记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrLineService *CdrLineService) GetCdrLine(id uint) (err error, cdrLine autocode.CdrLine) {
	err = global.GVA_DB.Where("id = ?", id).First(&cdrLine).Error
	return
}

// GetCdrLineInfoList 分页获取CdrLine记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrLineService *CdrLineService) GetCdrLineInfoList(info autoCodeReq.CdrLineSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.CdrLine{})
	var cdrLines []autocode.CdrLine
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.User != "" {
		db = db.Where("User LIKE ?", "%"+info.User+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&cdrLines).Error
	return err, cdrLines, total
}
