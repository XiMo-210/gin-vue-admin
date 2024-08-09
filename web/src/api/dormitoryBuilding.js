import service from '@/utils/request'

// @Tags DormitoryBuilding
// @Summary 创建寝室楼
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DormitoryBuilding true "创建寝室楼"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dormitoryBuilding/createDormitoryBuilding [post]
export const createDormitoryBuilding = (data) => {
  return service({
    url: '/dormitoryBuilding/createDormitoryBuilding',
    method: 'post',
    data
  })
}

// @Tags DormitoryBuilding
// @Summary 删除寝室楼
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DormitoryBuilding true "删除寝室楼"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dormitoryBuilding/deleteDormitoryBuilding [delete]
export const deleteDormitoryBuilding = (params) => {
  return service({
    url: '/dormitoryBuilding/deleteDormitoryBuilding',
    method: 'delete',
    params
  })
}

// @Tags DormitoryBuilding
// @Summary 批量删除寝室楼
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除寝室楼"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dormitoryBuilding/deleteDormitoryBuilding [delete]
export const deleteDormitoryBuildingByIds = (params) => {
  return service({
    url: '/dormitoryBuilding/deleteDormitoryBuildingByIds',
    method: 'delete',
    params
  })
}

// @Tags DormitoryBuilding
// @Summary 更新寝室楼
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DormitoryBuilding true "更新寝室楼"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dormitoryBuilding/updateDormitoryBuilding [put]
export const updateDormitoryBuilding = (data) => {
  return service({
    url: '/dormitoryBuilding/updateDormitoryBuilding',
    method: 'put',
    data
  })
}

// @Tags DormitoryBuilding
// @Summary 用id查询寝室楼
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DormitoryBuilding true "用id查询寝室楼"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dormitoryBuilding/findDormitoryBuilding [get]
export const findDormitoryBuilding = (params) => {
  return service({
    url: '/dormitoryBuilding/findDormitoryBuilding',
    method: 'get',
    params
  })
}

// @Tags DormitoryBuilding
// @Summary 分页获取寝室楼列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取寝室楼列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dormitoryBuilding/getDormitoryBuildingList [get]
export const getDormitoryBuildingList = (params) => {
  return service({
    url: '/dormitoryBuilding/getDormitoryBuildingList',
    method: 'get',
    params
  })
}
