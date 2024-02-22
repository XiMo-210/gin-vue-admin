package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
)

type StudentInfoService struct {
}

// CreateStudentInfo 创建学生信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentInfoService *StudentInfoService) CreateStudentInfo(studentInfo *admin.StudentInfo) (err error) {
	err = global.GVA_DB.Create(studentInfo).Error
	return err
}

// DeleteStudentInfo 删除学生信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentInfoService *StudentInfoService) DeleteStudentInfo(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.StudentInfo{}, "id = ?", ID).Error
	return err
}

// DeleteStudentInfoByIds 批量删除学生信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentInfoService *StudentInfoService) DeleteStudentInfoByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.StudentInfo{}, "id in ?", IDs).Error
	return err
}

// UpdateStudentInfo 更新学生信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentInfoService *StudentInfoService) UpdateStudentInfo(studentInfo admin.StudentInfo) (err error) {
	err = global.GVA_DB.Save(&studentInfo).Error
	return err
}

// GetStudentInfo 根据ID获取学生信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentInfoService *StudentInfoService) GetStudentInfo(ID string) (studentInfo admin.StudentInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&studentInfo).Error
	return
}

// GetStudentInfoInfoList 分页获取学生信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentInfoService *StudentInfoService) GetStudentInfoInfoList(info adminReq.StudentInfoSearch) (list []admin.StudentInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.StudentInfo{})
	var studentInfos []admin.StudentInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.Gender != "" {
		db = db.Where("gender = ?", info.Gender)
	}
	if info.AdmissionYear != nil {
		db = db.Where("admission_year = ?", info.AdmissionYear)
	}
	if info.OriginPlace != "" {
		db = db.Where("origin_place = ?", info.OriginPlace)
	}
	if info.Campus != "" {
		db = db.Where("campus = ?", info.Campus)
	}
	if info.StudentId != "" {
		db = db.Where("student_id = ?", info.StudentId)
	}
	if info.College != "" {
		db = db.Where("college = ?", info.College)
	}
	if info.Major != "" {
		db = db.Where("major = ?", info.Major)
	}
	if info.Class != "" {
		db = db.Where("class = ?", info.Class)
	}
	if info.Dormitory != "" {
		db = db.Where("dormitory = ?", info.Dormitory)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&studentInfos).Error
	return studentInfos, total, err
}
