import service from '@/utils/request'

// @Tags StudentInfo
// @Summary 创建学生信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StudentInfo true "创建学生信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /studentInfo/createStudentInfo [post]
export const createStudentInfo = (data) => {
  return service({
    url: '/studentInfo/createStudentInfo',
    method: 'post',
    data
  })
}

// @Tags StudentInfo
// @Summary 删除学生信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StudentInfo true "删除学生信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /studentInfo/deleteStudentInfo [delete]
export const deleteStudentInfo = (params) => {
  return service({
    url: '/studentInfo/deleteStudentInfo',
    method: 'delete',
    params
  })
}

// @Tags StudentInfo
// @Summary 批量删除学生信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除学生信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /studentInfo/deleteStudentInfo [delete]
export const deleteStudentInfoByIds = (params) => {
  return service({
    url: '/studentInfo/deleteStudentInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags StudentInfo
// @Summary 更新学生信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StudentInfo true "更新学生信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /studentInfo/updateStudentInfo [put]
export const updateStudentInfo = (data) => {
  return service({
    url: '/studentInfo/updateStudentInfo',
    method: 'put',
    data
  })
}

// @Tags StudentInfo
// @Summary 用id查询学生信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.StudentInfo true "用id查询学生信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /studentInfo/findStudentInfo [get]
export const findStudentInfo = (params) => {
  return service({
    url: '/studentInfo/findStudentInfo',
    method: 'get',
    params
  })
}

// @Tags StudentInfo
// @Summary 分页获取学生信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取学生信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /studentInfo/getStudentInfoList [get]
export const getStudentInfoList = (params) => {
  return service({
    url: '/studentInfo/getStudentInfoList',
    method: 'get',
    params
  })
}
