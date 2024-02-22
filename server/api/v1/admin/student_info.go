package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StudentInfoApi struct {
}

var studentInfoService = service.ServiceGroupApp.AdminServiceGroup.StudentInfoService

// CreateStudentInfo 创建学生信息
// @Tags StudentInfo
// @Summary 创建学生信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.StudentInfo true "创建学生信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /studentInfo/createStudentInfo [post]
func (studentInfoApi *StudentInfoApi) CreateStudentInfo(c *gin.Context) {
	var studentInfo admin.StudentInfo
	err := c.ShouldBindJSON(&studentInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := studentInfoService.CreateStudentInfo(&studentInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStudentInfo 删除学生信息
// @Tags StudentInfo
// @Summary 删除学生信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.StudentInfo true "删除学生信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /studentInfo/deleteStudentInfo [delete]
func (studentInfoApi *StudentInfoApi) DeleteStudentInfo(c *gin.Context) {
	ID := c.Query("ID")
	if err := studentInfoService.DeleteStudentInfo(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStudentInfoByIds 批量删除学生信息
// @Tags StudentInfo
// @Summary 批量删除学生信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /studentInfo/deleteStudentInfoByIds [delete]
func (studentInfoApi *StudentInfoApi) DeleteStudentInfoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := studentInfoService.DeleteStudentInfoByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStudentInfo 更新学生信息
// @Tags StudentInfo
// @Summary 更新学生信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.StudentInfo true "更新学生信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /studentInfo/updateStudentInfo [put]
func (studentInfoApi *StudentInfoApi) UpdateStudentInfo(c *gin.Context) {
	var studentInfo admin.StudentInfo
	err := c.ShouldBindJSON(&studentInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := studentInfoService.UpdateStudentInfo(studentInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStudentInfo 用id查询学生信息
// @Tags StudentInfo
// @Summary 用id查询学生信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.StudentInfo true "用id查询学生信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /studentInfo/findStudentInfo [get]
func (studentInfoApi *StudentInfoApi) FindStudentInfo(c *gin.Context) {
	ID := c.Query("ID")
	if restudentInfo, err := studentInfoService.GetStudentInfo(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restudentInfo": restudentInfo}, c)
	}
}

// GetStudentInfoList 分页获取学生信息列表
// @Tags StudentInfo
// @Summary 分页获取学生信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.StudentInfoSearch true "分页获取学生信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /studentInfo/getStudentInfoList [get]
func (studentInfoApi *StudentInfoApi) GetStudentInfoList(c *gin.Context) {
	var pageInfo adminReq.StudentInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := studentInfoService.GetStudentInfoInfoList(pageInfo); err != nil {
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
