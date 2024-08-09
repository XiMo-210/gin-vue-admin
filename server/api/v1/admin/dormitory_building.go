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

type DormitoryBuildingApi struct {
}

var dormitoryBuildingService = service.ServiceGroupApp.AdminServiceGroup.DormitoryBuildingService


// CreateDormitoryBuilding 创建寝室楼
// @Tags DormitoryBuilding
// @Summary 创建寝室楼
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.DormitoryBuilding true "创建寝室楼"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dormitoryBuilding/createDormitoryBuilding [post]
func (dormitoryBuildingApi *DormitoryBuildingApi) CreateDormitoryBuilding(c *gin.Context) {
	var dormitoryBuilding admin.DormitoryBuilding
	err := c.ShouldBindJSON(&dormitoryBuilding)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dormitoryBuildingService.CreateDormitoryBuilding(&dormitoryBuilding); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDormitoryBuilding 删除寝室楼
// @Tags DormitoryBuilding
// @Summary 删除寝室楼
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.DormitoryBuilding true "删除寝室楼"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dormitoryBuilding/deleteDormitoryBuilding [delete]
func (dormitoryBuildingApi *DormitoryBuildingApi) DeleteDormitoryBuilding(c *gin.Context) {
	ID := c.Query("ID")
	if err := dormitoryBuildingService.DeleteDormitoryBuilding(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDormitoryBuildingByIds 批量删除寝室楼
// @Tags DormitoryBuilding
// @Summary 批量删除寝室楼
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dormitoryBuilding/deleteDormitoryBuildingByIds [delete]
func (dormitoryBuildingApi *DormitoryBuildingApi) DeleteDormitoryBuildingByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := dormitoryBuildingService.DeleteDormitoryBuildingByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDormitoryBuilding 更新寝室楼
// @Tags DormitoryBuilding
// @Summary 更新寝室楼
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.DormitoryBuilding true "更新寝室楼"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dormitoryBuilding/updateDormitoryBuilding [put]
func (dormitoryBuildingApi *DormitoryBuildingApi) UpdateDormitoryBuilding(c *gin.Context) {
	var dormitoryBuilding admin.DormitoryBuilding
	err := c.ShouldBindJSON(&dormitoryBuilding)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dormitoryBuildingService.UpdateDormitoryBuilding(dormitoryBuilding); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDormitoryBuilding 用id查询寝室楼
// @Tags DormitoryBuilding
// @Summary 用id查询寝室楼
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.DormitoryBuilding true "用id查询寝室楼"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dormitoryBuilding/findDormitoryBuilding [get]
func (dormitoryBuildingApi *DormitoryBuildingApi) FindDormitoryBuilding(c *gin.Context) {
	ID := c.Query("ID")
	if redormitoryBuilding, err := dormitoryBuildingService.GetDormitoryBuilding(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redormitoryBuilding": redormitoryBuilding}, c)
	}
}

// GetDormitoryBuildingList 分页获取寝室楼列表
// @Tags DormitoryBuilding
// @Summary 分页获取寝室楼列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.DormitoryBuildingSearch true "分页获取寝室楼列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dormitoryBuilding/getDormitoryBuildingList [get]
func (dormitoryBuildingApi *DormitoryBuildingApi) GetDormitoryBuildingList(c *gin.Context) {
	var pageInfo adminReq.DormitoryBuildingSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dormitoryBuildingService.GetDormitoryBuildingInfoList(pageInfo); err != nil {
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
