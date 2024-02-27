import service from '@/utils/request'

// @Tags CommentScore
// @Summary 创建评分
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CommentScore true "创建评分"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /commentScore/createCommentScore [post]
export const createCommentScore = (data) => {
  return service({
    url: '/commentScore/createCommentScore',
    method: 'post',
    data
  })
}

// @Tags CommentScore
// @Summary 删除评分
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CommentScore true "删除评分"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /commentScore/deleteCommentScore [delete]
export const deleteCommentScore = (params) => {
  return service({
    url: '/commentScore/deleteCommentScore',
    method: 'delete',
    params
  })
}

// @Tags CommentScore
// @Summary 批量删除评分
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除评分"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /commentScore/deleteCommentScore [delete]
export const deleteCommentScoreByIds = (params) => {
  return service({
    url: '/commentScore/deleteCommentScoreByIds',
    method: 'delete',
    params
  })
}

// @Tags CommentScore
// @Summary 更新评分
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CommentScore true "更新评分"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /commentScore/updateCommentScore [put]
export const updateCommentScore = (data) => {
  return service({
    url: '/commentScore/updateCommentScore',
    method: 'put',
    data
  })
}

// @Tags CommentScore
// @Summary 用id查询评分
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CommentScore true "用id查询评分"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /commentScore/findCommentScore [get]
export const findCommentScore = (params) => {
  return service({
    url: '/commentScore/findCommentScore',
    method: 'get',
    params
  })
}

// @Tags CommentScore
// @Summary 分页获取评分列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取评分列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /commentScore/getCommentScoreList [get]
export const getCommentScoreList = (params) => {
  return service({
    url: '/commentScore/getCommentScoreList',
    method: 'get',
    params
  })
}
