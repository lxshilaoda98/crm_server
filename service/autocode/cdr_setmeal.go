package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CdrSetmealService struct {
}

// CreateCdrSetmeal 创建CdrSetmeal记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrSetmealService *CdrSetmealService) CreateCdrSetmeal(cdrSetmeal autocode.CdrSetmeal) (err error) {
	err = global.GVA_DB.Create(&cdrSetmeal).Error
	return err
}

// DeleteCdrSetmeal 删除CdrSetmeal记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrSetmealService *CdrSetmealService) DeleteCdrSetmeal(cdrSetmeal autocode.CdrSetmeal) (err error) {
	err = global.GVA_DB.Delete(&cdrSetmeal).Error
	return err
}

// DeleteCdrSetmealByIds 批量删除CdrSetmeal记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrSetmealService *CdrSetmealService) DeleteCdrSetmealByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CdrSetmeal{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCdrSetmeal 更新CdrSetmeal记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrSetmealService *CdrSetmealService) UpdateCdrSetmeal(cdrSetmeal autocode.CdrSetmeal) (err error) {
	err = global.GVA_DB.Save(&cdrSetmeal).Error
	return err
}

// GetCdrSetmeal 根据id获取CdrSetmeal记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrSetmealService *CdrSetmealService) GetCdrSetmeal(id uint) (err error, cdrSetmeal autocode.CdrSetmeal) {
	err = global.GVA_DB.Where("id = ?", id).First(&cdrSetmeal).Error
	return
}

// GetCdrSetmealInfoList 分页获取CdrSetmeal记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrSetmealService *CdrSetmealService) GetCdrSetmealInfoList(info autoCodeReq.CdrSetmealSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.CdrSetmeal{})
	var cdrSetmeals []autocode.CdrSetmeal
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("Name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&cdrSetmeals).Error
	return err, cdrSetmeals, total
}
