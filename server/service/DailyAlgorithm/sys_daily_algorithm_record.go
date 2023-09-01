package DailyAlgorithm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/DailyAlgorithm"
	DailyAlgorithmReq "github.com/flipped-aurora/gin-vue-admin/server/model/DailyAlgorithm/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"log"
	"time"
)

type DailyAlgorithmRecordService struct {
}

// CreateDailyAlgorithmRecord 创建DailyAlgorithmRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (DARService *DailyAlgorithmRecordService) CreateDailyAlgorithmRecord(DAR *DailyAlgorithm.DailyAlgorithmRecord) (err error) {
	tx := global.GVA_DB.Begin() // 开启事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 发生错误时回滚事务
		}
	}()
	// 操作
	err = global.GVA_DB.Create(DAR).Error

	if err != nil {
		tx.Rollback() // 操作失败，回滚事务
		log.Fatal(err)
	}
	err = tx.Commit().Error // 提交事务

	return err
}

// DeleteDailyAlgorithmRecord 删除DailyAlgorithmRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (DARService *DailyAlgorithmRecordService) DeleteDailyAlgorithmRecord(DAR DailyAlgorithm.DailyAlgorithmRecord) (err error) {
	err = global.GVA_DB.Delete(&DAR).Error
	return err
}

// DeleteDailyAlgorithmRecordByIds 批量删除DailyAlgorithmRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (DARService *DailyAlgorithmRecordService) DeleteDailyAlgorithmRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]DailyAlgorithm.DailyAlgorithmRecord{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDailyAlgorithmRecord 更新DailyAlgorithmRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (DARService *DailyAlgorithmRecordService) UpdateDailyAlgorithmRecord(DAR DailyAlgorithm.DailyAlgorithmRecord) (err error) {
	err = global.GVA_DB.Save(&DAR).Error
	return err
}

// GetDailyAlgorithmRecord 根据id获取DailyAlgorithmRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (DARService *DailyAlgorithmRecordService) GetDailyAlgorithmRecord(id uint) (DAR DailyAlgorithm.DailyAlgorithmRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&DAR).Error
	return
}

// GetCoverDailyAlgorithmRecord 根据date获取DailyAlgorithmRecord记录
// Author [CFDDFC](https://github.com/cfddd)
func (DARService *DailyAlgorithmRecordService) GetCoverDailyAlgorithmRecord(date time.Time, user_name string) (DAR DailyAlgorithm.DailyAlgorithmRecord, err error) {
	dateString := date.Format("2006-01-02")
	//fmt.Println(dateString)
	//fmt.Println(user_name)
	err = global.GVA_DB.Where("date = ?", dateString).Where("user_name = ?", user_name).First(&DAR).Error
	return
}

// GetDailyAlgorithmRecordInfoList 分页获取DailyAlgorithmRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (DARService *DailyAlgorithmRecordService) GetDailyAlgorithmRecordInfoList(info DailyAlgorithmReq.DailyAlgorithmRecordSearch) (list []DailyAlgorithm.DailyAlgorithmRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&DailyAlgorithm.DailyAlgorithmRecord{})
	var DARs []DailyAlgorithm.DailyAlgorithmRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Date != nil {
		db = db.Where("date = ?", info.Date)
	}
	if info.User_name != "" {
		db = db.Where("user_name = ?", info.User_name)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&DARs).Error
	return DARs, total, err
}
