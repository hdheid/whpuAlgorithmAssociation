import service from '@/utils/request'

// @Tags TEST
// @Summary 创建TEST
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TEST true "创建TEST"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TEST_S/createTEST [post]
export const createTEST = (data) => {
  return service({
    url: '/TEST_S/createTEST',
    method: 'post',
    data
  })
}

// @Tags TEST
// @Summary 删除TEST
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TEST true "删除TEST"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TEST_S/deleteTEST [delete]
export const deleteTEST = (data) => {
  return service({
    url: '/TEST_S/deleteTEST',
    method: 'delete',
    data
  })
}

// @Tags TEST
// @Summary 删除TEST
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TEST"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TEST_S/deleteTEST [delete]
export const deleteTESTByIds = (data) => {
  return service({
    url: '/TEST_S/deleteTESTByIds',
    method: 'delete',
    data
  })
}

// @Tags TEST
// @Summary 更新TEST
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TEST true "更新TEST"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TEST_S/updateTEST [put]
export const updateTEST = (data) => {
  return service({
    url: '/TEST_S/updateTEST',
    method: 'put',
    data
  })
}

// @Tags TEST
// @Summary 用id查询TEST
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TEST true "用id查询TEST"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TEST_S/findTEST [get]
export const findTEST = (params) => {
  return service({
    url: '/TEST_S/findTEST',
    method: 'get',
    params
  })
}

// @Tags TEST
// @Summary 分页获取TEST列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TEST列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TEST_S/getTESTList [get]
export const getTESTList = (params) => {
  return service({
    url: '/TEST_S/getTESTList',
    method: 'get',
    params
  })
}
