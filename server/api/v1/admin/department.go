package admin

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DepartmentApi struct {
}

var departmentService = service.ServiceGroupApp.AdminServiceGroup.DepartmentService

// CreateDepartment 创建部门
// @Tags Department
// @Summary 创建部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Department true "创建部门"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /department/createDepartment [post]
func (departmentApi *DepartmentApi) CreateDepartment(c *gin.Context) {
	var department admin.Department
	err := c.ShouldBindJSON(&department)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := departmentService.CreateDepartment(&department); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDepartment 删除部门
// @Tags Department
// @Summary 删除部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Department true "删除部门"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /department/deleteDepartment [delete]
func (departmentApi *DepartmentApi) DeleteDepartment(c *gin.Context) {
	ID := c.Query("ID")
	if err := departmentService.DeleteDepartment(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDepartmentByIds 批量删除部门
// @Tags Department
// @Summary 批量删除部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /department/deleteDepartmentByIds [delete]
func (departmentApi *DepartmentApi) DeleteDepartmentByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := departmentService.DeleteDepartmentByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDepartment 更新部门
// @Tags Department
// @Summary 更新部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Department true "更新部门"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /department/updateDepartment [put]
func (departmentApi *DepartmentApi) UpdateDepartment(c *gin.Context) {
	var department admin.Department
	err := c.ShouldBindJSON(&department)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := departmentService.UpdateDepartment(department); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDepartment 用id查询部门
// @Tags Department
// @Summary 用id查询部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.Department true "用id查询部门"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /department/findDepartment [get]
func (departmentApi *DepartmentApi) FindDepartment(c *gin.Context) {
	ID := c.Query("ID")
	if redepartment, err := departmentService.GetDepartment(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redepartment": redepartment}, c)
	}
}

// GetDepartmentList 分页获取部门列表
// @Tags Department
// @Summary 分页获取部门列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.DepartmentSearch true "分页获取部门列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /department/getDepartmentList [get]
func (departmentApi *DepartmentApi) GetDepartmentList(c *gin.Context) {
	var pageInfo adminReq.DepartmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, _ := c.Get("claims")
	customClaims, _ := claims.(*request.CustomClaims)

	organization, err := organizationService.GetOrganization(customClaims.BaseClaims.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	if list, total, err := departmentService.GetDepartmentInfoList(pageInfo, organization.ID); err != nil {
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
