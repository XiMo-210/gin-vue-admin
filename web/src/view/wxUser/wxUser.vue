<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item label="关联学生信息" prop="studentInfoId">
            
             <el-input v-model.number="searchInfo.studentInfoId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="用户名" prop="username">
         <el-input v-model="searchInfo.username" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="智慧种子" prop="exp">
            
            <el-input v-model.number="searchInfo.startExp" placeholder="最小值" />
            —
            <el-input v-model.number="searchInfo.endExp" placeholder="最大值" />

        </el-form-item>
        <el-form-item label="求索石" prop="points">
            
            <el-input v-model.number="searchInfo.startPoints" placeholder="最小值" />
            —
            <el-input v-model.number="searchInfo.endPoints" placeholder="最大值" />

        </el-form-item>
            <el-form-item label="是否完成主线任务" prop="isCompletedMain">
            <el-select v-model="searchInfo.isCompletedMain" clearable placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>
        <el-form-item label="正在进行的任务" prop="curTask">
            
             <el-input v-model.number="searchInfo.curTask" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="微信用户唯一标识" prop="openid" width="120" />
        <el-table-column align="left" label="关联学生信息" prop="studentInfoId" width="120" />
          <el-table-column label="人脸识别标识" width="200">
              <template #default="scope">
                <el-image style="width: 100px; height: 100px" :src="getUrl(scope.row.faceId)" fit="cover"/>
              </template>
          </el-table-column>
        <el-table-column align="left" label="用户名" prop="username" width="120" />
        <el-table-column align="left" label="密码" prop="password" width="120" />
          <el-table-column label="人像" width="200">
              <template #default="scope">
                <el-image style="width: 100px; height: 100px" :src="getUrl(scope.row.avatar)" fit="cover"/>
              </template>
          </el-table-column>
        <el-table-column align="left" label="个性签名" prop="signature" width="120" />
        <el-table-column align="left" label="爱好" prop="hobby" width="120" />
        <el-table-column align="left" label="联系方式" prop="contact" width="120" />
        <el-table-column align="left" label="手机" prop="phone" width="120" />
        <el-table-column align="left" label="邮箱" prop="email" width="120" />
        <el-table-column sortable align="left" label="智慧种子" prop="exp" width="120" />
        <el-table-column sortable align="left" label="求索石" prop="points" width="120" />
        <el-table-column align="left" label="是否完成主线任务" prop="isCompletedMain" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.isCompletedMain) }}</template>
        </el-table-column>
        <el-table-column align="left" label="正在进行的任务" prop="curTask" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateWxUserFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-scrollbar height="500px">
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="微信用户唯一标识:"  prop="openid" >
              <el-input v-model="formData.openid" :clearable="true"  placeholder="请输入微信用户唯一标识" />
            </el-form-item>
            <el-form-item label="关联学生信息:"  prop="studentInfoId" >
              <el-input v-model.number="formData.studentInfoId" :clearable="true" placeholder="请输入关联学生信息" />
            </el-form-item>
            <el-form-item label="人脸识别标识:"  prop="faceId" >
                <SelectImage
                 v-model="formData.faceId"
                 file-type="image"
                />
            </el-form-item>
            <el-form-item label="用户名:"  prop="username" >
              <el-input v-model="formData.username" :clearable="true"  placeholder="请输入用户名" />
            </el-form-item>
            <el-form-item label="密码:"  prop="password" >
              <el-input v-model="formData.password" :clearable="true"  placeholder="请输入密码" />
            </el-form-item>
            <el-form-item label="人像:"  prop="avatar" >
                <SelectImage
                 v-model="formData.avatar"
                 file-type="image"
                />
            </el-form-item>
            <el-form-item label="个性签名:"  prop="signature" >
              <el-input v-model="formData.signature" :clearable="true"  placeholder="请输入个性签名" />
            </el-form-item>
            <el-form-item label="爱好:"  prop="hobby" >
              <el-input v-model="formData.hobby" :clearable="true"  placeholder="请输入爱好" />
            </el-form-item>
            <el-form-item label="联系方式:"  prop="contact" >
              <el-input v-model="formData.contact" :clearable="true"  placeholder="请输入联系方式" />
            </el-form-item>
            <el-form-item label="手机:"  prop="phone" >
              <el-input v-model="formData.phone" :clearable="true"  placeholder="请输入手机" />
            </el-form-item>
            <el-form-item label="邮箱:"  prop="email" >
              <el-input v-model="formData.email" :clearable="true"  placeholder="请输入邮箱" />
            </el-form-item>
            <el-form-item label="智慧种子:"  prop="exp" >
              <el-input v-model.number="formData.exp" :clearable="true" placeholder="请输入智慧种子" />
            </el-form-item>
            <el-form-item label="求索石:"  prop="points" >
              <el-input v-model.number="formData.points" :clearable="true" placeholder="请输入求索石" />
            </el-form-item>
            <el-form-item label="是否完成主线任务:"  prop="isCompletedMain" >
              <el-switch v-model="formData.isCompletedMain" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="正在进行的任务:"  prop="curTask" >
              <el-input v-model.number="formData.curTask" :clearable="true" placeholder="请输入正在进行的任务" />
            </el-form-item>
          </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
                <el-descriptions-item label="微信用户唯一标识">
                        {{ formData.openid }}
                </el-descriptions-item>
                <el-descriptions-item label="关联学生信息">
                        {{ formData.studentInfoId }}
                </el-descriptions-item>
                <el-descriptions-item label="人脸识别标识">
                        <el-image style="width: 50px; height: 50px" :preview-src-list="ReturnArrImg(formData.faceId)" :src="getUrl(formData.faceId)" fit="cover" />
                </el-descriptions-item>
                <el-descriptions-item label="用户名">
                        {{ formData.username }}
                </el-descriptions-item>
                <el-descriptions-item label="密码">
                        {{ formData.password }}
                </el-descriptions-item>
                <el-descriptions-item label="人像">
                        <el-image style="width: 50px; height: 50px" :preview-src-list="ReturnArrImg(formData.avatar)" :src="getUrl(formData.avatar)" fit="cover" />
                </el-descriptions-item>
                <el-descriptions-item label="个性签名">
                        {{ formData.signature }}
                </el-descriptions-item>
                <el-descriptions-item label="爱好">
                        {{ formData.hobby }}
                </el-descriptions-item>
                <el-descriptions-item label="联系方式">
                        {{ formData.contact }}
                </el-descriptions-item>
                <el-descriptions-item label="手机">
                        {{ formData.phone }}
                </el-descriptions-item>
                <el-descriptions-item label="邮箱">
                        {{ formData.email }}
                </el-descriptions-item>
                <el-descriptions-item label="智慧种子">
                        {{ formData.exp }}
                </el-descriptions-item>
                <el-descriptions-item label="求索石">
                        {{ formData.points }}
                </el-descriptions-item>
                <el-descriptions-item label="是否完成主线任务">
                    {{ formatBoolean(formData.isCompletedMain) }}
                </el-descriptions-item>
                <el-descriptions-item label="正在进行的任务">
                        {{ formData.curTask }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createWxUser,
  deleteWxUser,
  deleteWxUserByIds,
  updateWxUser,
  findWxUser,
  getWxUserList
} from '@/api/wxUser'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'WxUser'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        openid: '',
        studentInfoId: 0,
        faceId: "",
        username: '',
        password: '',
        avatar: "",
        signature: '',
        hobby: '',
        contact: '',
        phone: '',
        email: '',
        exp: 0,
        points: 0,
        isCompletedMain: false,
        curTask: 0,
        })


// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  searchInfo.value.sort = prop
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    if (searchInfo.value.isCompletedMain === ""){
        searchInfo.value.isCompletedMain=null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getWxUserList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteWxUserFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteWxUserByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateWxUserFunc = async(row) => {
    const res = await findWxUser({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rewxUser
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWxUserFunc = async (row) => {
    const res = await deleteWxUser({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findWxUser({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.rewxUser
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          openid: '',
          studentInfoId: 0,
          username: '',
          password: '',
          signature: '',
          hobby: '',
          contact: '',
          phone: '',
          email: '',
          exp: 0,
          points: 0,
          isCompletedMain: false,
          curTask: 0,
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        openid: '',
        studentInfoId: 0,
        username: '',
        password: '',
        signature: '',
        hobby: '',
        contact: '',
        phone: '',
        email: '',
        exp: 0,
        points: 0,
        isCompletedMain: false,
        curTask: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createWxUser(formData.value)
                  break
                case 'update':
                  res = await updateWxUser(formData.value)
                  break
                default:
                  res = await createWxUser(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
