import service from '@/utils/request'

// @Tags ReviewRecords
// @Summary 创建审核记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ReviewRecords true "创建审核记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /reviewRecord/createReviewRecords [post]
export const createReviewRecords = (data) => {
  return service({
    url: '/reviewRecord/createReviewRecords',
    method: 'post',
    data
  })
}

// @Tags ReviewRecords
// @Summary 删除审核记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ReviewRecords true "删除审核记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /reviewRecord/deleteReviewRecords [delete]
export const deleteReviewRecords = (params) => {
  return service({
    url: '/reviewRecord/deleteReviewRecords',
    method: 'delete',
    params
  })
}

// @Tags ReviewRecords
// @Summary 批量删除审核记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除审核记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /reviewRecord/deleteReviewRecords [delete]
export const deleteReviewRecordsByIds = (params) => {
  return service({
    url: '/reviewRecord/deleteReviewRecordsByIds',
    method: 'delete',
    params
  })
}

// @Tags ReviewRecords
// @Summary 更新审核记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ReviewRecords true "更新审核记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /reviewRecord/updateReviewRecords [put]
export const updateReviewRecords = (data) => {
  return service({
    url: '/reviewRecord/updateReviewRecords',
    method: 'put',
    data
  })
}

// @Tags ReviewRecords
// @Summary 用id查询审核记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ReviewRecords true "用id查询审核记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /reviewRecord/findReviewRecords [get]
export const findReviewRecords = (params) => {
  return service({
    url: '/reviewRecord/findReviewRecords',
    method: 'get',
    params
  })
}

// @Tags ReviewRecords
// @Summary 分页获取审核记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取审核记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /reviewRecord/getReviewRecordsList [get]
export const getReviewRecordsList = (params) => {
  return service({
    url: '/reviewRecord/getReviewRecordsList',
    method: 'get',
    params
  })
}
