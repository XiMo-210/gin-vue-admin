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
          label="任务类型"
          prop="category"
        >

          <el-select
            v-model.number="searchInfo.category"
            clearable
            placeholder="请选择"
            style="width: 120px"
          >
            <el-option
              key="main"
              label="主线任务"
              :value="1"
            />
            <el-option
              key="side"
              label="支线任务"
              :value="2"
            />
            <el-option
              key="hide"
              label="隐藏任务"
              :value="3"
            />
          </el-select>

        </el-form-item>
        <el-form-item
          label="任务名称"
          prop="title"
        >
          <el-input
            v-model="searchInfo.title"
            placeholder="任务名称"
            style="width: 140px"
          />

        </el-form-item>
        <el-form-item
          label="任务校区"
          prop="campus"
        >
          <el-select
            v-model="searchInfo.campus"
            clearable
            placeholder="请选择"
            style="width: 140px"
          >
            <el-option
              key="all"
              value="全部"
            />
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
          label="任务学院"
          prop="college"
        >
          <el-input
            v-model="searchInfo.college"
            placeholder="任务学院"
            style="width: 260px"
          />
        </el-form-item>
        <p />
        <el-form-item
          label="开始时间"
          prop="startTime"
        >
          <el-date-picker
            v-model="searchInfo.startTime"
            type="date"
            placeholder="请选择"
            style="width: 140px"
          />
        </el-form-item>
        <el-form-item
          label="结束时间"
          prop="endTime"
        >
          <el-date-picker
            v-model="searchInfo.endTime"
            type="date"
            placeholder="请选择"
            style="width: 140px"
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
          align="left"
          label="ID"
          prop="ID"
          width="60"
        />
        <el-table-column
          align="left"
          label="任务类型"
          prop="category"
          width="120"
        />
        <el-table-column
          align="left"
          label="任务描述"
          prop="desc"
          width="160"
          show-overflow-tooltip="true"
        />
        <el-table-column
          align="left"
          label="任务校区"
          prop="campus"
          width="120"
        />
        <el-table-column
          align="left"
          label="任务学院"
          prop="college"
          width="240px"
        />
        <el-table-column
          align="left"
          label="奖励积分"
          prop="reward"
          width="100"
        />
        <el-table-column
          align="left"
          label="需要完成主线任务"
          prop="needMain"
          width="140"
        >
          <template #default="scope">{{ formatBoolean(scope.row.needMain) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="开始时间"
          width="160"
        >
          <template #default="scope">{{ formatDate(scope.row.startTime) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="结束时间"
          width="160"
        >
          <template #default="scope">{{ formatDate(scope.row.endTime) }}</template>
        </el-table-column>
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
              icon="edit"
              class="table-button"
              @click="updateTaskFunc(scope.row)"
            >变更</el-button>
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
    >
      <el-scrollbar height="500px">
        <el-form
          ref="elFormRef"
          :model="formData"
          label-position="right"
          :rules="rule"
          label-width="80px"
        >
          <el-form-item
            label="任务类型:"
            prop="category"
          >
            <el-input
              v-model.number="formData.category"
              :clearable="true"
              placeholder="请输入任务类型"
            />
          </el-form-item>
          <el-form-item
            label="任务名称:"
            prop="title"
          >
            <el-input
              v-model="formData.title"
              :clearable="true"
              placeholder="请输入任务名称"
            />
          </el-form-item>
          <el-form-item
            label="任务描述:"
            prop="desc"
          >
            <el-input
              v-model="formData.desc"
              :clearable="true"
              placeholder="请输入任务描述"
            />
          </el-form-item>
          <el-form-item
            label="任务校区:"
            prop="campus"
          >
            <el-input
              v-model="formData.campus"
              :clearable="true"
              placeholder="请输入任务校区"
            />
          </el-form-item>
          <el-form-item
            label="任务学院:"
            prop="college"
          >
            <el-input
              v-model="formData.college"
              :clearable="true"
              placeholder="请输入任务学院"
            />
          </el-form-item>
          <el-form-item
            label="奖励积分:"
            prop="reward"
          >
            <el-input
              v-model.number="formData.reward"
              :clearable="true"
              placeholder="请输入奖励积分"
            />
          </el-form-item>
          <el-form-item
            label="是否需要完成主线任务:"
            prop="needMain"
          >
            <el-switch
              v-model="formData.needMain"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
              clearable
            />
          </el-form-item>
          <el-form-item
            label="开始时间:"
            prop="startTime"
          >
            <el-date-picker
              v-model="formData.startTime"
              type="date"
              style="width:100%"
              placeholder="选择日期"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item
            label="结束时间:"
            prop="endTime"
          >
            <el-date-picker
              v-model="formData.endTime"
              type="date"
              style="width:100%"
              placeholder="选择日期"
              :clearable="true"
            />
          </el-form-item>
        </el-form>
      </el-scrollbar>
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

    <el-dialog
      v-model="detailShow"
      style="width: 800px"
      lock-scroll
      :before-close="closeDetailShow"
      title="查看详情"
      destroy-on-close
    >
      <el-scrollbar height="550px">
        <el-descriptions
          column="1"
          border
        >
          <el-descriptions-item label="任务类型">
            {{ formData.category }}
          </el-descriptions-item>
          <el-descriptions-item label="任务名称">
            {{ formData.title }}
          </el-descriptions-item>
          <el-descriptions-item label="任务描述">
            {{ formData.desc }}
          </el-descriptions-item>
          <el-descriptions-item label="任务校区">
            {{ formData.campus }}
          </el-descriptions-item>
          <el-descriptions-item label="任务学院">
            {{ formData.college }}
          </el-descriptions-item>
          <el-descriptions-item label="奖励积分">
            {{ formData.reward }}
          </el-descriptions-item>
          <el-descriptions-item label="是否需要完成主线任务">
            {{ formatBoolean(formData.needMain) }}
          </el-descriptions-item>
          <el-descriptions-item label="开始时间">
            {{ formatDate(formData.startTime) }}
          </el-descriptions-item>
          <el-descriptions-item label="结束时间">
            {{ formatDate(formData.endTime) }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createTask,
  deleteTask,
  deleteTaskByIds,
  updateTask,
  findTask,
  getTaskList
} from '@/api/task'

// 全量引入格式化工具 请按需保留
import { formatDate, formatBoolean } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'Task'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  category: 0,
  title: '',
  desc: '',
  campus: '',
  college: '',
  reward: 0,
  needMain: false,
  startTime: new Date(),
  endTime: new Date(),
})

// 验证规则
const rule = reactive({
  category: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  title: [{
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
  desc: [{
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
  reward: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  needMain: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  startTime: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  endTime: [{
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
    if (searchInfo.value.needMain === '') {
      searchInfo.value.needMain = null
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
  const table = await getTaskList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteTaskFunc(row)
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
  const res = await deleteTaskByIds({ IDs })
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
const updateTaskFunc = async(row) => {
  const res = await findTask({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.retask
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteTaskFunc = async(row) => {
  const res = await deleteTask({ ID: row.ID })
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
const getDetails = async(row) => {
  // 打开弹窗
  const res = await findTask({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.retask
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    category: 0,
    title: '',
    desc: '',
    campus: '',
    college: '',
    reward: 0,
    needMain: false,
    startTime: new Date(),
    endTime: new Date(),
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
    category: 0,
    title: '',
    desc: '',
    campus: '',
    college: '',
    reward: 0,
    needMain: false,
    startTime: new Date(),
    endTime: new Date(),
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createTask(formData.value)
        break
      case 'update':
        res = await updateTask(formData.value)
        break
      default:
        res = await createTask(formData.value)
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
