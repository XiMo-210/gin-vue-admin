<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item
          label="用户名"
          prop="username"
        >
          <el-input
            v-model="searchInfo.username"
            placeholder="用户名"
            style="width: 140px"
          />

        </el-form-item>
        <el-form-item
          label="智慧种子"
          prop="exp"
        >

          <el-input
            v-model.number="searchInfo.startExp"
            placeholder="最小值"
            style="width: 100px"
          />
          —
          <el-input
            v-model.number="searchInfo.endExp"
            placeholder="最大值"
            style="width: 100px"
          />

        </el-form-item>
        <el-form-item
          label="求索石"
          prop="points"
        >

          <el-input
            v-model.number="searchInfo.startPoints"
            placeholder="最小值"
            style="width: 100px"
          />
          —
          <el-input
            v-model.number="searchInfo.endPoints"
            placeholder="最大值"
            style="width: 100px"
          />

        </el-form-item>
        <el-form-item
          label="是否完成主线任务"
          prop="isCompletedMain"
        >
          <el-select
            v-model="searchInfo.isCompletedMain"
            clearable
            placeholder="请选择"
            style="width: 100px;"
          >
            <el-option
              key="true"
              label="是"
              value="true"
            />
            <el-option
              key="false"
              label="否"
              value="false"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="正在进行的任务"
          prop="curTask"
        >

          <el-input
            v-model.number="searchInfo.curTask"
            placeholder="任务ID"
            style="width: 100px"
          />

        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            icon="search"
            @click="onSubmit"
          >查询</el-button>
          <el-button
            icon="refresh"
            @click="onReset"
          >重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-popover
          v-model:visible="deleteVisible"
          :disabled="!multipleSelection.length"
          placement="top"
          width="160"
        >
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button
              type="primary"
              link
              @click="deleteVisible = false"
            >取消</el-button>
            <el-button
              type="primary"
              @click="onDelete"
            >确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              style="margin-left: 10px;"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
            >删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        :header-cell-style="{ 'text-align': 'center' }"
        :cell-style="{ textAlign: 'center' }"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
      >
        <el-table-column
          type="selection"
          width="55"
        />

        <el-table-column
          align="left"
          label="注册时间"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column
          align="left"
          label="用户名"
          prop="username"
          width="120"
        />
        <el-table-column
          label="头像"
          width="200"
        >
          <template #default="scope">
            <el-image
              style="width: 100px; height: 100px"
              :src="getUrl(scope.row.avatar)"
              fit="cover"
            />
          </template>
        </el-table-column>
        <el-table-column
          sortable
          align="left"
          label="智慧种子"
          prop="exp"
          width="120"
        />
        <el-table-column
          sortable
          align="left"
          label="求索石"
          prop="points"
          width="120"
        />
        <el-table-column
          align="left"
          label="是否完成主线任务"
          prop="isCompletedMain"
          width="160"
        >
          <template #default="scope">{{ formatBoolean(scope.row.isCompletedMain) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="正在进行的任务"
          prop="curTask"
          width="160"
        />
        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="240"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              class="table-button"
              @click="getDetails(scope.row)"
            >
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              查看详情
            </el-button>
            <el-button
              type="primary"
              link
              class="table-button"
              @click="getStudentInfoDetails(scope.row)"
            >
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              关联学生
            </el-button>
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
            >删除</el-button>
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

    <el-dialog
      v-model="detailShow"
      style="width: 650px"
      lock-scroll
      :before-close="closeDetailShow"
      title="查看详情"
      destroy-on-close
    >
      <el-scrollbar height="500px">
        <el-descriptions
          column="1"
          border
        >
          <el-descriptions-item label="用户名">
            {{ formData.username }}
          </el-descriptions-item>
          <el-descriptions-item label="头像">
            <el-image
              style="width: 50px; height: 50px"
              :preview-src-list="ReturnArrImg(formData.avatar)"
              :src="getUrl(formData.avatar)"
              fit="cover"
            />
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

    <el-dialog
      v-model="studentInfoDetailShow"
      style="width: 500px"
      lock-scroll
      :before-close="closeStudentInfoDetailShow"
      title="关联学生"
      destroy-on-close
    >
      <el-scrollbar height="380px">
        <el-descriptions
          column="1"
          border
        >
          <el-descriptions-item label="姓名">
            {{ studentInfoFormData.name }}
          </el-descriptions-item>
          <el-descriptions-item label="性别">
            {{ studentInfoFormData.gender }}
          </el-descriptions-item>
          <el-descriptions-item label="入学年份">
            {{ studentInfoFormData.admissionYear }}
          </el-descriptions-item>
          <el-descriptions-item label="校区">
            {{ studentInfoFormData.campus }}
          </el-descriptions-item>
          <el-descriptions-item label="学号">
            {{ studentInfoFormData.studentId }}
          </el-descriptions-item>
          <el-descriptions-item label="学院">
            {{ studentInfoFormData.college }}
          </el-descriptions-item>
          <el-descriptions-item label="专业">
            {{ studentInfoFormData.major }}
          </el-descriptions-item>
          <el-descriptions-item label="班级">
            {{ studentInfoFormData.class }}
          </el-descriptions-item>
          <el-descriptions-item label="寝室">
            {{ studentInfoFormData.dormitory }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  deleteWxUser,
  deleteWxUserByIds,
  findWxUser,
  getWxUserList
} from '@/api/wxUser'
import {
  findStudentInfo,
} from '@/api/studentInfo'
import { getUrl } from '@/utils/image'

// 全量引入格式化工具 请按需保留
import { formatDate, formatBoolean, ReturnArrImg } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'WxUser'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  openid: '',
  studentInfoId: 0,
  faceId: '',
  username: '',
  password: '',
  avatar: '',
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
    if (searchInfo.value.isCompletedMain === '') {
      searchInfo.value.isCompletedMain = null
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
const setOptions = async() => {
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

// 删除行
const deleteWxUserFunc = async(row) => {
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

// 查看详情控制标记
const detailShow = ref(false)

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const getDetails = async(row) => {
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

const studentInfoFormData = ref({
  name: '',
  gender: '',
  birthday: new Date(),
  idCard: '',
  admissionLetterId: '',
  admissionYear: 0,
  originPlace: '',
  campus: '',
  studentId: '',
  college: '',
  major: '',
  class: '',
  dormitory: '',
  portrait: '',
  userId: 0,
  isRegister: false,
})

// 查看详情控制标记
const studentInfoDetailShow = ref(false)

// 打开详情弹窗
const openStudentInfoDetailShow = () => {
  studentInfoDetailShow.value = true
}

// 打开详情
const getStudentInfoDetails = async(row) => {
  // 打开弹窗
  const res = await findStudentInfo({ ID: row.studentInfoId })
  if (res.code === 0) {
    studentInfoFormData.value = res.data.restudentInfo
    openStudentInfoDetailShow()
  }
}

// 关闭详情弹窗
const closeStudentInfoDetailShow = () => {
  studentInfoDetailShow.value = false
  studentInfoFormData.value = {
    name: '',
    gender: '',
    birthday: new Date(),
    idCard: '',
    admissionLetterId: '',
    admissionYear: 0,
    originPlace: '',
    campus: '',
    studentId: '',
    college: '',
    major: '',
    class: '',
    dormitory: '',
    userId: 0,
    isRegister: false,
  }
}

</script>

<style>

</style>
