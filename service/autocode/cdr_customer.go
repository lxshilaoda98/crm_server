package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CdrCustomerService struct {
}

// CreateCdrCustomer 创建CdrCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrCustomerService *CdrCustomerService) CreateCdrCustomer(cdrCustomer autocode.CdrCustomer) (err error) {
	err = global.GVA_DB.Create(&cdrCustomer).Error
	return err
}

// DeleteCdrCustomer 删除CdrCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrCustomerService *CdrCustomerService) DeleteCdrCustomer(cdrCustomer autocode.CdrCustomer) (err error) {
	err = global.GVA_DB.Delete(&cdrCustomer).Error
	return err
}

// DeleteCdrCustomerByIds 批量删除CdrCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrCustomerService *CdrCustomerService) DeleteCdrCustomerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CdrCustomer{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCdrCustomer 更新CdrCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrCustomerService *CdrCustomerService) UpdateCdrCustomer(cdrCustomer autocode.CdrCustomer) (err error) {
	err = global.GVA_DB.Save(&cdrCustomer).Error
	return err
}

// GetCdrCustomer 根据id获取CdrCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrCustomerService *CdrCustomerService) GetCdrCustomer(id uint) (err error, cdrCustomer autocode.CdrCustomer) {
	err = global.GVA_DB.Where("id = ?", id).First(&cdrCustomer).Error
	return
}

// GetCdrCustomerInfoList 分页获取CdrCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdrCustomerService *CdrCustomerService) GetCdrCustomerInfoList(info autoCodeReq.CdrCustomerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.CdrCustomer{})
	var cdrCustomers []autocode.CdrCustomer
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.User != "" {
		db = db.Where("User LIKE ?", "%"+info.User+"%")
	}
	if info.Name != "" {
		db = db.Where("Name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&cdrCustomers).Error
	return err, cdrCustomers, total
}
