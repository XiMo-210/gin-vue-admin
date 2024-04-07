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

type OrganizationApi struct {
}

var organizationService = service.ServiceGroupApp.AdminServiceGroup.OrganizationService

// CreateOrganization 创建组织社团
// @Tags Organization
// @Summary 创建组织社团
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Organization true "创建组织社团"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /organization/createOrganization [post]
func (organizationApi *OrganizationApi) CreateOrganization(c *gin.Context) {
	var organization admin.Organization
	err := c.ShouldBindJSON(&organization)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, _ := c.Get("claims")
	customClaims, _ := claims.(*request.CustomClaims)
	organization.SysUserId = customClaims.BaseClaims.ID

	if err := organizationService.CreateOrganization(&organization); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrganization 删除组织社团
// @Tags Organization
// @Summary 删除组织社团
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Organization true "删除组织社团"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /organization/deleteOrganization [delete]
func (organizationApi *OrganizationApi) DeleteOrganization(c *gin.Context) {
	ID := c.Query("ID")
	if err := organizationService.DeleteOrganization(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrganizationByIds 批量删除组织社团
// @Tags Organization
// @Summary 批量删除组织社团
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /organization/deleteOrganizationByIds [delete]
func (organizationApi *OrganizationApi) DeleteOrganizationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := organizationService.DeleteOrganizationByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrganization 更新组织社团
// @Tags Organization
// @Summary 更新组织社团
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Organization true "更新组织社团"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /organization/updateOrganization [put]
func (organizationApi *OrganizationApi) UpdateOrganization(c *gin.Context) {
	var organization admin.Organization
	err := c.ShouldBindJSON(&organization)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := organizationService.UpdateOrganization(organization); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrganization 用id查询组织社团
// @Tags Organization
// @Summary 用id查询组织社团
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.Organization true "用id查询组织社团"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /organization/findOrganization [get]
func (organizationApi *OrganizationApi) FindOrganization(c *gin.Context) {
	claims, _ := c.Get("claims")
	customClaims, _ := claims.(*request.CustomClaims)

	if reorganization, err := organizationService.GetOrganization(customClaims.BaseClaims.ID); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"reorganization": reorganization}, c)
	}
}

// GetOrganizationList 分页获取组织社团列表
// @Tags Organization
// @Summary 分页获取组织社团列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.OrganizationSearch true "分页获取组织社团列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /organization/getOrganizationList [get]
func (organizationApi *OrganizationApi) GetOrganizationList(c *gin.Context) {
	var pageInfo adminReq.OrganizationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := organizationService.GetOrganizationInfoList(pageInfo); err != nil {
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

func (organizationApi *OrganizationApi) CreateOrganizationByAdmin(c *gin.Context) {
	var organization admin.Organization
	err := c.ShouldBindJSON(&organization)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := organizationService.CreateOrganization(&organization); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (organizationApi *OrganizationApi) UpdateOrganizationByAdmin(c *gin.Context) {
	var organization admin.Organization
	err := c.ShouldBindJSON(&organization)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := organizationService.UpdateOrganization(organization); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (organizationApi *OrganizationApi) FindOrganizationByAdmin(c *gin.Context) {
	ID := c.Query("ID")
	if reorganization, err := organizationService.GetOrganizationByAdmin(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorganization": reorganization}, c)
	}
}
