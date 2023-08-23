package DailyAlgorithm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/DailyAlgorithm"
	DailyAlgorithmReq "github.com/flipped-aurora/gin-vue-admin/server/model/DailyAlgorithm/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DailyAlgorithmRecordApi struct {
}

var DARService = service.ServiceGroupApp.DailyAlgorithmServiceGroup.DailyAlgorithmRecordService

// CreateDailyAlgorithmRecord 创建DailyAlgorithmRecord
// @Tags DailyAlgorithmRecord
// @Summary 创建DailyAlgorithmRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body DailyAlgorithm.DailyAlgorithmRecord true "创建DailyAlgorithmRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DAR/createDailyAlgorithmRecord [post]
func (DARApi *DailyAlgorithmRecordApi) CreateDailyAlgorithmRecord(c *gin.Context) {
	var DAR DailyAlgorithm.DailyAlgorithmRecord
	err := c.ShouldBindJSON(&DAR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Date":      {utils.NotEmpty()},
		"User_name": {utils.NotEmpty()}, //User_name传来的是uuid
		"Link":      {utils.NotEmpty()},
		"Code":      {utils.NotEmpty()},
	}
	if err := utils.Verify(DAR, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := DARService.CreateDailyAlgorithmRecord(&DAR); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)

		//创建record成功后，给排行榜上的打卡次数加一
		//第二次再提交因为已经有记录，只是更新记录，不会再一天内再提交，调用该函数

		err := system.AddDACount(DAR.User_name)
		if err != nil {
			return
		}
	}
}

// DeleteDailyAlgorithmRecord 删除DailyAlgorithmRecord
// @Tags DailyAlgorithmRecord
// @Summary 删除DailyAlgorithmRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body DailyAlgorithm.DailyAlgorithmRecord true "删除DailyAlgorithmRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DAR/deleteDailyAlgorithmRecord [delete]
func (DARApi *DailyAlgorithmRecordApi) DeleteDailyAlgorithmRecord(c *gin.Context) {
	var DAR DailyAlgorithm.DailyAlgorithmRecord
	err := c.ShouldBindJSON(&DAR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := DARService.DeleteDailyAlgorithmRecord(DAR); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDailyAlgorithmRecordByIds 批量删除DailyAlgorithmRecord
// @Tags DailyAlgorithmRecord
// @Summary 批量删除DailyAlgorithmRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DailyAlgorithmRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /DAR/deleteDailyAlgorithmRecordByIds [delete]
func (DARApi *DailyAlgorithmRecordApi) DeleteDailyAlgorithmRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := DARService.DeleteDailyAlgorithmRecordByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDailyAlgorithmRecord 更新DailyAlgorithmRecord
// @Tags DailyAlgorithmRecord
// @Summary 更新DailyAlgorithmRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body DailyAlgorithm.DailyAlgorithmRecord true "更新DailyAlgorithmRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DAR/updateDailyAlgorithmRecord [put]
func (DARApi *DailyAlgorithmRecordApi) UpdateDailyAlgorithmRecord(c *gin.Context) {
	var DAR DailyAlgorithm.DailyAlgorithmRecord
	err := c.ShouldBindJSON(&DAR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Date":      {utils.NotEmpty()},
		"User_name": {utils.NotEmpty()},
		"Link":      {utils.NotEmpty()},
		"Code":      {utils.NotEmpty()},
	}
	if err := utils.Verify(DAR, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := DARService.UpdateDailyAlgorithmRecord(DAR); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDailyAlgorithmRecord 用id查询DailyAlgorithmRecord
// @Tags DailyAlgorithmRecord
// @Summary 用id查询DailyAlgorithmRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query DailyAlgorithm.DailyAlgorithmRecord true "用id查询DailyAlgorithmRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DAR/findDailyAlgorithmRecord [get]
func (DARApi *DailyAlgorithmRecordApi) FindDailyAlgorithmRecord(c *gin.Context) {
	var DAR DailyAlgorithm.DailyAlgorithmRecord
	err := c.ShouldBindQuery(&DAR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reDAR, err := DARService.GetDailyAlgorithmRecord(DAR.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDAR": reDAR}, c)
	}
}

// CoverDailyAlgorithmRecord 用date查询DailyAlgorithmRecord
// @Tags DailyAlgorithmRecord
// @Summary  用date查询DailyAlgorithmRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query DailyAlgorithm.DailyAlgorithmRecord true "用date查询DailyAlgorithmRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DAR/findDailyAlgorithmRecord [get]
func (DARApi *DailyAlgorithmRecordApi) CoverDailyAlgorithmRecord(c *gin.Context) {
	var DAR DailyAlgorithm.DailyAlgorithmRecord
	err := c.ShouldBindQuery(&DAR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if reDAR, err := DARService.GetCoverDailyAlgorithmRecord(*(DAR.Date), DAR.User_name); err != nil {
		//global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.OkWithMessage("打卡已提交", c)
	} else {
		response.OkWithDetailed(gin.H{"reDAR": reDAR}, "打卡已更新", c)
	}
}

// GetDailyAlgorithmRecordList 分页获取DailyAlgorithmRecord列表
// @Tags DailyAlgorithmRecord
// @Summary 分页获取DailyAlgorithmRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query DailyAlgorithmReq.DailyAlgorithmRecordSearch true "分页获取DailyAlgorithmRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DAR/getDailyAlgorithmRecordList [get]
func (DARApi *DailyAlgorithmRecordApi) GetDailyAlgorithmRecordList(c *gin.Context) {
	var pageInfo DailyAlgorithmReq.DailyAlgorithmRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := DARService.GetDailyAlgorithmRecordInfoList(pageInfo); err != nil {
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
