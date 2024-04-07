import service from '@/utils/request'

// @Tags Business
// @Summary 创建店铺商家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Business true "创建店铺商家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /business/createBusiness [post]
export const createBusiness = (data) => {
  return service({
    url: '/business/createBusiness',
    method: 'post',
    data
  })
}

// @Tags Business
// @Summary 删除店铺商家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Business true "删除店铺商家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /business/deleteBusiness [delete]
export const deleteBusiness = (params) => {
  return service({
    url: '/business/deleteBusiness',
    method: 'delete',
    params
  })
}

// @Tags Business
// @Summary 批量删除店铺商家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除店铺商家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /business/deleteBusiness [delete]
export const deleteBusinessByIds = (params) => {
  return service({
    url: '/business/deleteBusinessByIds',
    method: 'delete',
    params
  })
}

// @Tags Business
// @Summary 更新店铺商家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Business true "更新店铺商家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /business/updateBusiness [put]
export const updateBusiness = (data) => {
  return service({
    url: '/business/updateBusiness',
    method: 'put',
    data
  })
}

// @Tags Business
// @Summary 用id查询店铺商家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Business true "用id查询店铺商家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /business/findBusiness [get]
export const findBusiness = (params) => {
  return service({
    url: '/business/findBusiness',
    method: 'get',
    params
  })
}

// @Tags Business
// @Summary 分页获取店铺商家列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取店铺商家列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /business/getBusinessList [get]
export const getBusinessList = (params) => {
  return service({
    url: '/business/getBusinessList',
    method: 'get',
    params
  })
}

export const createBusinessByAdmin = (data) => {
  return service({
    url: '/business/createBusinessByAdmin',
    method: 'post',
    data
  })
}

export const updateBusinessByAdmin = (data) => {
  return service({
    url: '/business/updateBusinessByAdmin',
    method: 'put',
    data
  })
}

export const findBusinessByAdmin = (params) => {
  return service({
    url: '/business/findBusinessByAdmin',
    method: 'get',
    params
  })
}
