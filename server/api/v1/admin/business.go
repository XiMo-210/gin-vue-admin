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

type BusinessApi struct {
}

var businessService = service.ServiceGroupApp.AdminServiceGroup.BusinessService

// CreateBusiness 创建店铺商家
// @Tags Business
// @Summary 创建店铺商家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Business true "创建店铺商家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /business/createBusiness [post]
func (businessApi *BusinessApi) CreateBusiness(c *gin.Context) {
	var business admin.Business
	err := c.ShouldBindJSON(&business)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, _ := c.Get("claims")
	customClaims, _ := claims.(*request.CustomClaims)
	business.SysUserId = customClaims.BaseClaims.ID

	if err := businessService.CreateBusiness(&business); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBusiness 删除店铺商家
// @Tags Business
// @Summary 删除店铺商家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Business true "删除店铺商家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /business/deleteBusiness [delete]
func (businessApi *BusinessApi) DeleteBusiness(c *gin.Context) {
	ID := c.Query("ID")
	if err := businessService.DeleteBusiness(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBusinessByIds 批量删除店铺商家
// @Tags Business
// @Summary 批量删除店铺商家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /business/deleteBusinessByIds [delete]
func (businessApi *BusinessApi) DeleteBusinessByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := businessService.DeleteBusinessByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBusiness 更新店铺商家
// @Tags Business
// @Summary 更新店铺商家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Business true "更新店铺商家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /business/updateBusiness [put]
func (businessApi *BusinessApi) UpdateBusiness(c *gin.Context) {
	var business admin.Business
	err := c.ShouldBindJSON(&business)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := businessService.UpdateBusiness(business); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBusiness 用id查询店铺商家
// @Tags Business
// @Summary 用id查询店铺商家
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.Business true "用id查询店铺商家"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /business/findBusiness [get]
func (businessApi *BusinessApi) FindBusiness(c *gin.Context) {
	claims, _ := c.Get("claims")
	customClaims, _ := claims.(*request.CustomClaims)

	if rebusiness, err := businessService.GetBusiness(customClaims.BaseClaims.ID); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"rebusiness": rebusiness}, c)
	}
}

// GetBusinessList 分页获取店铺商家列表
// @Tags Business
// @Summary 分页获取店铺商家列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.BusinessSearch true "分页获取店铺商家列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /business/getBusinessList [get]
func (businessApi *BusinessApi) GetBusinessList(c *gin.Context) {
	var pageInfo adminReq.BusinessSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := businessService.GetBusinessInfoList(pageInfo); err != nil {
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

func (businessApi *BusinessApi) CreateBusinessByAdmin(c *gin.Context) {
	var business admin.Business
	err := c.ShouldBindJSON(&business)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := businessService.CreateBusiness(&business); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (businessApi *BusinessApi) UpdateBusinessByAdmin(c *gin.Context) {
	var business admin.Business
	err := c.ShouldBindJSON(&business)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := businessService.UpdateBusiness(business); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (businessApi *BusinessApi) FindBusinessByAdmin(c *gin.Context) {
	ID := c.Query("ID")
	if rebusiness, err := businessService.GetBusinessByAdmin(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebusiness": rebusiness}, c)
	}
}
