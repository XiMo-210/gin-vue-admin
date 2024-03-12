import service from '@/utils/request'

// @Tags Advertiser
// @Summary 创建广告主
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Advertiser true "创建广告主"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /advertiser/createAdvertiser [post]
export const createAdvertiser = (data) => {
  return service({
    url: '/advertiser/createAdvertiser',
    method: 'post',
    data
  })
}

// @Tags Advertiser
// @Summary 删除广告主
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Advertiser true "删除广告主"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /advertiser/deleteAdvertiser [delete]
export const deleteAdvertiser = (params) => {
  return service({
    url: '/advertiser/deleteAdvertiser',
    method: 'delete',
    params
  })
}

// @Tags Advertiser
// @Summary 批量删除广告主
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除广告主"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /advertiser/deleteAdvertiser [delete]
export const deleteAdvertiserByIds = (params) => {
  return service({
    url: '/advertiser/deleteAdvertiserByIds',
    method: 'delete',
    params
  })
}

// @Tags Advertiser
// @Summary 更新广告主
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Advertiser true "更新广告主"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /advertiser/updateAdvertiser [put]
export const updateAdvertiser = (data) => {
  return service({
    url: '/advertiser/updateAdvertiser',
    method: 'put',
    data
  })
}

// @Tags Advertiser
// @Summary 用id查询广告主
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Advertiser true "用id查询广告主"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /advertiser/findAdvertiser [get]
export const findAdvertiser = (params) => {
  return service({
    url: '/advertiser/findAdvertiser',
    method: 'get',
    params
  })
}

// @Tags Advertiser
// @Summary 分页获取广告主列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取广告主列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /advertiser/getAdvertiserList [get]
export const getAdvertiserList = (params) => {
  return service({
    url: '/advertiser/getAdvertiserList',
    method: 'get',
    params
  })
}
