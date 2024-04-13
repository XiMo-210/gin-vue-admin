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
          label="姓名"
          prop="name"
        >
          <el-input
            v-model="searchInfo.name"
            placeholder="姓名"
            style="width: 140px"
          />
        </el-form-item>
        <el-form-item
          label="性别"
          prop="gender"
        >
          <el-select
            v-model="searchInfo.gender"
            clearable
            placeholder="请选择"
            style="width: 100px"
          >
            <el-option
              key="man"
              value="男"
            />
            <el-option
              key="woman"
              value="女"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="入学年份"
          prop="admissionYear"
        >
          <el-input
            v-model.number="searchInfo.admissionYear"
            placeholder="入学年份"
            style="width: 140px"
          />
        </el-form-item>
        <el-form-item
          label="生源地"
          prop="originPlace"
        >
          <el-input
            v-model="searchInfo.originPlace"
            placeholder="生源地"
            style="width: 140px"
          />
        </el-form-item>
        <el-form-item
          label="校区"
          prop="campus"
        >
          <el-select
            v-model="searchInfo.campus"
            clearable
            placeholder="请选择"
            style="width: 140px"
          >
            <el-option
              key="zh"
              value="朝晖校区"
            />
            <el-option
              key="pf"
              value="屏峰校区"
            />
            <el-option
              key="mgs"
              value="莫干山校区"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="学号"
          prop="studentId"
        >
          <el-input
            v-model="searchInfo.studentId"
            placeholder="学号"
            style="width: 140px"
          />
        </el-form-item>
        <el-form-item
          label="学院"
          prop="college"
        >
          <el-input
            v-model="searchInfo.college"
            placeholder="学院"
            style="width: 260px"
          />
        </el-form-item>
        <el-form-item
          label="专业"
          prop="major"
        >
          <el-input
            v-model="searchInfo.major"
            placeholder="专业"
            style="width: 160px"
          />
        </el-form-item>
        <el-form-item
          label="班级"
          prop="class"
        >
          <el-input
            v-model="searchInfo.class"
            placeholder="班级"
            style="width: 214px"
          />
        </el-form-item>
        <el-form-item
          label="寝室"
          prop="dormitory"
        >
          <el-input
            v-model="searchInfo.dormitory"
            placeholder="寝室"
            style="width: 180px"
          />
        </el-form-item>
        <el-form-item
          label="是否注册用户"
          prop="isRegister"
        >
          <el-select
            v-model="searchInfo.isRegister"
            clearable
            placeholder="请选择"
            style="width: 100px"
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
        <el-button
          type="primary"
          icon="plus"
          @click="openDialog"
        >新增</el-button>
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
        <ImportExcel
          template-id="studentInfo"
          @on-success="getTableData"
        />
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
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          label="ID"
          prop="ID"
          width="65"
        />
        <el-table-column
          label="姓名"
          prop="name"
          width="120"
          :show-overflow-tooltip="true"
        />
        <el-table-column
          label="性别"
          prop="gender"
          width="80"
        />
        <el-table-column
          label="入学年份"
          prop="admissionYear"
          width="100"
        />
        <el-table-column
          label="校区"
          prop="campus"
          width="120"
        />
        <el-table-column
          label="学号"
          prop="studentId"
          width="120"
        />
        <el-table-column
          label="学院"
          prop="college"
          width="160"
          :show-overflow-tooltip="true"
        />
        <el-table-column
          label="专业"
          prop="major"
          width="160"
          :show-overflow-tooltip="true"
        />
        <el-table-column
          label="班级"
          prop="class"
          width="160"
          :show-overflow-tooltip="true"
        />
        <el-table-column
          label="寝室"
          prop="dormitory"
          width="160"
          :show-overflow-tooltip="true"
        />
        <el-table-column
          label="是否注册"
          prop="userId"
          width="80"
        >
          <template #default="scope">{{ scope.row.userId!==0 ? "是":"否" }}</template>
        </el-table-column>
        <el-table-column
          label="操作"
          fixed="right"
          min-width="240"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateStudentInfoFunc(scope.row)"
            >查看 / 变更</el-button>
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
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="type==='create'?'添加':'修改'"
      destroy-on-close
      align-center
    >
      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="right"
        :rules="rule"
        label-width="100px"
        inline="true"
      >
        <el-text
          tag="b"
          size="large"
          style="color: black;"
        >
          基本信息
        </el-text>
        <el-divider style="margin-top: 10px;margin-bottom: 10px;" />
        <el-form-item
          label="姓名:"
          prop="name"
          style="width: 40%;"
        >
          <el-input
            v-model="formData.name"
            placeholder="请输入姓名"
          />
        </el-form-item>
        <el-form-item
          label="性别:"
          prop="gender"
          style="width: 25%;"
        >
          <el-select
            v-model="formData.gender"
          >
            <el-option
              key="man"
              value="男"
            />
            <el-option
              key="woman"
              value="女"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="出生日期:"
          prop="birthday"
          style="width: 40%;"
        >
          <el-date-picker
            v-model="formData.birthday"
            type="date"
            placeholder="选择日期"
          />
        </el-form-item>
        <el-form-item
          label="入学年份:"
          prop="admissionYear"
        >
          <el-input-number
            v-model="formData.admissionYear"
            controls-position="right"
            :step="1"
            step-strictly
          />
        </el-form-item>
        <el-form-item
          label="身份证:"
          prop="idCard"
          style="width: 40%;"
        >
          <el-input
            v-model="formData.idCard"
            placeholder="请输入身份证"
          />
        </el-form-item>
        <el-form-item
          label="录取编号:"
          prop="admissionLetterId"
          style="width: 40%;"
        >
          <el-input
            v-model="formData.admissionLetterId"
            placeholder="请输入录取通知书编号"
          />
        </el-form-item>
        <el-form-item
          label="生源地:"
          prop="originPlace"
          style="width: 40%;"
        >
          <el-input
            v-model="formData.originPlace"
            placeholder="请输入生源地"
          />
        </el-form-item>
        <br>
        <el-text
          tag="b"
          size="large"
          style="color: black;"
        >
          学生信息
        </el-text>
        <el-divider style="margin-top: 10px;margin-bottom: 10px;" />
        <el-form-item
          label="学号:"
          prop="studentId"
          style="width: 40%;"
        >
          <el-input
            v-model="formData.studentId"
            placeholder="请输入学号"
          />
        </el-form-item>
        <el-form-item
          label="校区:"
          prop="campus"
          style="width: 30%;"
        >
          <el-select
            v-model="formData.campus"
          >
            <el-option
              key="zh"
              value="朝晖校区"
            />
            <el-option
              key="pf"
              value="屏峰校区"
            />
            <el-option
              key="mgs"
              value="莫干山校区"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="学院:"
          prop="college"
          style="width: 60%;"
        >
          <el-input
            v-model="formData.college"
            placeholder="请输入学院"
          />
        </el-form-item>
        <el-form-item
          label="专业:"
          prop="major"
          style="width: 60%;"
        >
          <el-input
            v-model="formData.major"
            placeholder="请输入专业"
          />
        </el-form-item>
        <el-form-item
          label="班级:"
          prop="class"
          style="width: 60%;"
        >
          <el-input
            v-model="formData.class"
            placeholder="请输入班级"
          />
        </el-form-item>
        <el-form-item
          label="寝室:"
          prop="dormitory"
          style="width: 40%;"
        >
          <el-input
            v-model="formData.dormitory"
            placeholder="请输入寝室"
          />
        </el-form-item>
        <br>
        <el-text
          tag="b"
          size="large"
          style="color: black;"
        >
          人脸信息
        </el-text>
        <el-divider style="margin-top: 10px;margin-bottom: 10px;" />
        <el-form-item
          label="人像:"
          prop="portrait"
        >
          <UploadImg
            v-model="formData.portrait"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button
            type="primary"
            @click="enterDialog"
          >确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createStudentInfo,
  deleteStudentInfo,
  deleteStudentInfoByIds,
  updateStudentInfo,
  findStudentInfo,
  getStudentInfoList
} from '@/api/studentInfo'
import UploadImg from '@/components/uploadImg/uploadImg.vue'

import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'

defineOptions({
  name: 'StudentInfo'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: '',
  gender: '男',
  birthday: new Date(),
  idCard: '',
  admissionLetterId: '',
  admissionYear: new Date().getFullYear(),
  originPlace: '',
  campus: '朝晖校区',
  studentId: '',
  college: '',
  major: '',
  class: '',
  dormitory: '',
  portrait: '',
  userId: 0,
  isRegister: false,
})

// 验证规则
const rule = reactive({
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  gender: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  birthday: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  idCard: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  admissionLetterId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  admissionYear: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  originPlace: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  campus: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  studentId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  college: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  major: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  class: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  dormitory: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  portrait: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
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
    if (searchInfo.value.isRegister === '') {
      searchInfo.value.isRegister = null
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
  const table = await getStudentInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteStudentInfoFunc(row)
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
  const res = await deleteStudentInfoByIds({ IDs })
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
const updateStudentInfoFunc = async(row) => {
  const res = await findStudentInfo({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.restudentInfo
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteStudentInfoFunc = async(row) => {
  const res = await deleteStudentInfo({ ID: row.ID })
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

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: '',
    gender: '男',
    birthday: new Date(),
    idCard: '',
    admissionLetterId: '',
    admissionYear: new Date().getFullYear(),
    originPlace: '',
    campus: '朝晖校区',
    studentId: '',
    college: '',
    major: '',
    class: '',
    dormitory: '',
    userId: 0,
    isRegister: false,
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createStudentInfo(formData.value)
        break
      case 'update':
        res = await updateStudentInfo(formData.value)
        break
      default:
        res = await createStudentInfo(formData.value)
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
