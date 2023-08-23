package TEST

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/TEST"
	TESTReq "github.com/flipped-aurora/gin-vue-admin/server/model/TEST/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TESTApi struct {
}

var TEST_SService = service.ServiceGroupApp.TESTServiceGroup.TESTService

// CreateTEST 创建TEST
// @Tags TEST
// @Summary 创建TEST
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body TEST.TEST true "创建TEST"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TEST_S/createTEST [post]
func (TEST_SApi *TESTApi) CreateTEST(c *gin.Context) {
	var TEST_S TEST.TEST
	err := c.ShouldBindJSON(&TEST_S)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := TEST_SService.CreateTEST(&TEST_S); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTEST 删除TEST
// @Tags TEST
// @Summary 删除TEST
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body TEST.TEST true "删除TEST"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TEST_S/deleteTEST [delete]
func (TEST_SApi *TESTApi) DeleteTEST(c *gin.Context) {
	var TEST_S TEST.TEST
	err := c.ShouldBindJSON(&TEST_S)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := TEST_SService.DeleteTEST(TEST_S); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTESTByIds 批量删除TEST
// @Tags TEST
// @Summary 批量删除TEST
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TEST"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /TEST_S/deleteTESTByIds [delete]
func (TEST_SApi *TESTApi) DeleteTESTByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := TEST_SService.DeleteTESTByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTEST 更新TEST
// @Tags TEST
// @Summary 更新TEST
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body TEST.TEST true "更新TEST"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TEST_S/updateTEST [put]
func (TEST_SApi *TESTApi) UpdateTEST(c *gin.Context) {
	var TEST_S TEST.TEST
	err := c.ShouldBindJSON(&TEST_S)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := TEST_SService.UpdateTEST(TEST_S); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTEST 用id查询TEST
// @Tags TEST
// @Summary 用id查询TEST
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query TEST.TEST true "用id查询TEST"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TEST_S/findTEST [get]
func (TEST_SApi *TESTApi) FindTEST(c *gin.Context) {
	var TEST_S TEST.TEST
	err := c.ShouldBindQuery(&TEST_S)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reTEST_S, err := TEST_SService.GetTEST(TEST_S.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reTEST_S": reTEST_S}, c)
	}
}

// GetTESTList 分页获取TEST列表
// @Tags TEST
// @Summary 分页获取TEST列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query TESTReq.TESTSearch true "分页获取TEST列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TEST_S/getTESTList [get]
func (TEST_SApi *TESTApi) GetTESTList(c *gin.Context) {
	var pageInfo TESTReq.TESTSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := TEST_SService.GetTESTInfoList(pageInfo); err != nil {
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
