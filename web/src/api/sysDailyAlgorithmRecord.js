import service from '@/utils/request'

// @Tags DailyAlgorithmRecord
// @Summary 创建DailyAlgorithmRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DailyAlgorithmRecord true "创建DailyAlgorithmRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DAR/createDailyAlgorithmRecord [post]
export const createDailyAlgorithmRecord = (data) => {
  return service({
    url: '/DAR/createDailyAlgorithmRecord',
    method: 'post',
    data
  })
}

// @Tags DailyAlgorithmRecord
// @Summary 删除DailyAlgorithmRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DailyAlgorithmRecord true "删除DailyAlgorithmRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DAR/deleteDailyAlgorithmRecord [delete]
export const deleteDailyAlgorithmRecord = (data) => {
  return service({
    url: '/DAR/deleteDailyAlgorithmRecord',
    method: 'delete',
    data
  })
}

// @Tags DailyAlgorithmRecord
// @Summary 删除DailyAlgorithmRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DailyAlgorithmRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DAR/deleteDailyAlgorithmRecord [delete]
export const deleteDailyAlgorithmRecordByIds = (data) => {
  return service({
    url: '/DAR/deleteDailyAlgorithmRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags DailyAlgorithmRecord
// @Summary 更新DailyAlgorithmRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DailyAlgorithmRecord true "更新DailyAlgorithmRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DAR/updateDailyAlgorithmRecord [put]
export const updateDailyAlgorithmRecord = (data) => {
  return service({
    url: '/DAR/updateDailyAlgorithmRecord',
    method: 'put',
    data
  })
}

// @Tags DailyAlgorithmRecord
// @Summary 用id查询DailyAlgorithmRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DailyAlgorithmRecord true "用id查询DailyAlgorithmRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DAR/findDailyAlgorithmRecord [get]
export const findDailyAlgorithmRecord = (params) => {
  return service({
    url: '/DAR/findDailyAlgorithmRecord',
    method: 'get',
    params
  })
}

// @Tags DailyAlgorithmRecord
// @Summary 用date查询DailyAlgorithmRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DailyAlgorithmRecord true "用date查询DailyAlgorithmRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DAR/findDailyAlgorithmRecord [get]
export const coverDailyAlgorithmRecord = (params) => {
  return service({
    url: '/DAR/coverDailyAlgorithmRecord',
    method: 'get',
    params
  })
}

// @Tags DailyAlgorithmRecord
// @Summary 分页获取DailyAlgorithmRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DailyAlgorithmRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DAR/getDailyAlgorithmRecordList [get]
export const getDailyAlgorithmRecordList = (params) => {
  return service({
    url: '/DAR/getDailyAlgorithmRecordList',
    method: 'get',
    params
  })
}
