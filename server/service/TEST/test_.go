package TEST

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/TEST"
	TESTReq "github.com/flipped-aurora/gin-vue-admin/server/model/TEST/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TESTService struct {
}

// CreateTEST 创建TEST记录
// Author [piexlmax](https://github.com/piexlmax)
func (TEST_SService *TESTService) CreateTEST(TEST_S *TEST.TEST) (err error) {
	err = global.GVA_DB.Create(TEST_S).Error
	return err
}

// DeleteTEST 删除TEST记录
// Author [piexlmax](https://github.com/piexlmax)
func (TEST_SService *TESTService) DeleteTEST(TEST_S TEST.TEST) (err error) {
	err = global.GVA_DB.Delete(&TEST_S).Error
	return err
}

// DeleteTESTByIds 批量删除TEST记录
// Author [piexlmax](https://github.com/piexlmax)
func (TEST_SService *TESTService) DeleteTESTByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]TEST.TEST{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTEST 更新TEST记录
// Author [piexlmax](https://github.com/piexlmax)
func (TEST_SService *TESTService) UpdateTEST(TEST_S TEST.TEST) (err error) {
	err = global.GVA_DB.Save(&TEST_S).Error
	return err
}

// GetTEST 根据id获取TEST记录
// Author [piexlmax](https://github.com/piexlmax)
func (TEST_SService *TESTService) GetTEST(id uint) (TEST_S TEST.TEST, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&TEST_S).Error
	return
}

// GetTESTInfoList 分页获取TEST记录
// Author [piexlmax](https://github.com/piexlmax)
func (TEST_SService *TESTService) GetTESTInfoList(info TESTReq.TESTSearch) (list []TEST.TEST, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&TEST.TEST{})
	var TEST_Ss []TEST.TEST
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&TEST_Ss).Error
	return TEST_Ss, total, err
}
