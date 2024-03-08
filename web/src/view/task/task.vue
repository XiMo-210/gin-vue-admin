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
        >
          <template #default="scope">
            {{ formatCategory(scope.row.category) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="任务名称"
          prop="title"
          width="120"
          show-overflow-tooltip="true"
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
          width="120px"
          show-overflow-tooltip="true"
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
              icon="edit"
              class="table-button"
              @click="updateTaskFunc(scope.row)"
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
    >
      <el-scrollbar height="600px">
        <el-form
          ref="elFormRef"
          :inline="true"
          :model="taskWithStagesformData"
          label-position="right"
          :rules="rule"
        >
          <el-form-item
            label="任务类型:"
            prop="task.category"
            style="width: 40%;"
          >
            <el-select
              v-model.number="taskWithStagesformData.task.category"
              placeholder="请选择"
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
            label="任务名称:"
            prop="task.title"
            style="width: 40%;"
          >
            <el-input
              v-model="taskWithStagesformData.task.title"
              placeholder="请输入任务名称"
            />
          </el-form-item>

          <el-form-item
            label="任务校区:"
            prop="task.campus"
            style="width: 40%;"
          >
            <el-select
              v-model="taskWithStagesformData.task.campus"
              placeholder="请选择"
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
            label="任务学院:"
            prop="task.college"
            style="width: 40%;"
          >
            <el-input
              v-model="taskWithStagesformData.task.college"
              placeholder="请输入任务学院"
            />
          </el-form-item>
          <el-form-item
            label="奖励积分:"
            prop="task.reward"
            style="width: 40%;"
          >
            <el-input
              v-model.number="taskWithStagesformData.task.reward"
              placeholder="请输入奖励积分"
            />
          </el-form-item>
          <el-form-item
            v-show="taskWithStagesformData.task.category!==1"
            label="需要完成主线任务:"
            prop="task.needMain"
            style="width: 40%;"
          >
            <el-switch
              v-model="taskWithStagesformData.task.needMain"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
              clearable
            />
          </el-form-item>
          <p />
          <el-form-item>
            <el-button
              type="primary"
              @click="switchIsSetPublishTime"
            >{{ buttonContent }}</el-button>
          </el-form-item>
          <span v-show="publishTimeShow">
            <el-form-item
              label="开始时间:"
              prop="task.startTime"
              style="width: 30%;"
            >
              <el-date-picker
                v-model="taskWithStagesformData.task.startTime"
                type="date"
                placeholder="选择日期"
                :clearable="true"
              />
            </el-form-item>
            <el-form-item
              label="结束时间:"
              prop="task.endTime"
              style="width: 30%;"
            >
              <el-date-picker
                v-model="taskWithStagesformData.task.endTime"
                type="date"
                placeholder="选择日期"
                :clearable="true"
              />
            </el-form-item>
          </span>
          <el-form-item
            label="任务描述:"
            prop="task.desc"
            style="width: 80%;"
          >
            <el-input
              v-model="taskWithStagesformData.task.desc"
              placeholder="请输入任务描述"
              :autosize="{minRows:2,maxRows:4}"
              type="textarea"
            />
          </el-form-item>
          <el-divider style="margin: 0px 0px;top: -10px;" />
        </el-form>
        <el-form
          ref="elFormRef"
          :inline="true"
          :model="taskWithStagesformData"
          label-position="right"
          :rules="rule"
          label-width="100px"
        >
          <el-form-item>
            <el-button
              type="primary"
              @click="addTaskStage"
            >添加任务阶段</el-button>
          </el-form-item>
          <el-steps
            direction="vertical"
            :active="activeStep"
          >
            <div
              v-for="(taskStage,index) in taskWithStagesformData.taskStages"
              :key="taskStage.stage"
            >
              <el-step :title="'阶段'+(index+1)" />
              <el-form-item
                label="阶段任务名称:"
                :prop="'title'+index"
                style="width: 40%;"
              >
                <el-input
                  v-model="taskStage.title"
                  placeholder="阶段任务名称"
                />
              </el-form-item>
              <el-form-item
                label="所需物品:"
                :prop="'requiredItem'+index"
                style="width: 40%;"
              >
                <el-input
                  v-model="taskStage.requiredItem"
                  placeholder="所需物品"
                />
              </el-form-item>
              <el-form-item
                label="阶段任务描述:"
                :prop="'desc'+index"
                style="width: 80%;"
              >
                <el-input
                  v-model="taskStage.desc"
                  placeholder="阶段任务描述"
                  :autosize="{minRows:1,maxRows:2}"
                  type="textarea"
                />
              </el-form-item>
              <el-form-item
                label="图片描述:"
                :prop="'imgs'+index"
                style="width: 20%;"
              >
                <UploadMultipleImg
                  v-model="taskStage.imgs"
                />
              </el-form-item>
              <p />
              <el-form-item
                label="摄像头打卡:"
                :prop="'needCamera'+index"
                style="width: 20%;"
              >
                <el-switch
                  v-model="taskStage.needCamera"
                  active-color="#13ce66"
                  inactive-color="#ff4949"
                  active-text="是"
                  inactive-text="否"
                  @click="needCamera(taskStage)"
                />
              </el-form-item>
              <el-form-item
                v-show="taskStage.needCamera"
                label="摄像头编号:"
                :prop="'cameraId'+index"
                style="width: 35%;"
              >
                <el-input
                  v-model.number="taskStage.cameraId"
                  placeholder="摄像头编号"
                />
              </el-form-item>
              <p />
              <span v-show="!taskStage.needCamera">
                <el-form-item
                  label="图片打卡:"
                  :prop="'needPic'+index"
                  style="width: 20%;"
                >
                  <el-switch
                    v-model="taskStage.needPic"
                    active-color="#13ce66"
                    inactive-color="#ff4949"
                    active-text="是"
                    inactive-text="否"
                  />
                </el-form-item>
                <span v-show="taskStage.needPic">
                  <el-form-item
                    label="需要人脸:"
                    :prop="'needFace'+index"
                    style="width: 20%;"
                  >
                    <el-switch
                      v-model="taskStage.needFace"
                      active-color="#13ce66"
                      inactive-color="#ff4949"
                      active-text="是"
                      inactive-text="否"
                    />
                  </el-form-item>
                  <el-form-item
                    :prop="'pic'+index"
                  >
                    <UploadImg
                      v-model="taskStage.pic"
                      style="margin-left: 80px;"
                    />
                  </el-form-item>
                </span>
                <p />
                <el-form-item
                  label="位置打卡:"
                  :prop="'needLoc'+index"
                  style="width: 20%;"
                >
                  <el-switch
                    v-model="taskStage.needLoc"
                    active-color="#13ce66"
                    inactive-color="#ff4949"
                    active-text="是"
                    inactive-text="否"
                  />
                </el-form-item>
                <el-form-item
                  label="允许导航:"
                  :prop="'needNav'+index"
                  style="width: 20%;"
                >
                  <el-switch
                    v-model="taskStage.needNav"
                    active-color="#13ce66"
                    inactive-color="#ff4949"
                    active-text="是"
                    inactive-text="否"
                  />
                </el-form-item>
                <el-form-item
                  v-show="taskStage.needLoc||taskStage.needNav"
                  style="margin-left: 6% ;"
                >
                  <el-button
                    type="primary"
                    size="default"
                    @click="openAmapDialog(index)"
                  >选择位置
                  </el-button>
                </el-form-item>
                <el-form-item
                  v-show="taskStage.needLoc||taskStage.needNav"
                  :prop="'loc'+index"
                  style="width: 22%;margin-left: -3.5%;"
                >
                  <el-input
                    v-model="taskStage.loc"
                    placeholder="位置"
                    disabled
                  />
                </el-form-item>

              </span>
              <el-divider style="margin: 0px 0px;top: -10px;" />
            </div>
          </el-steps>
          <el-button
            type="primary"
            @click.prevent="removeTaskStage"
          >删除阶段</el-button>
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
      v-model="amapDialog"
      :before-close="closeAmapDialog"
      destroy-on-close
    >
      <Amap
        v-model="amapParam.loc"
      />
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

import UploadImg from '@/components/uploadImg/uploadImg.vue'
import UploadMultipleImg from '@/components/uploadImg/uploadMultipleImg.vue'
import Amap from '@/components/amap/amap.vue'

// 全量引入格式化工具 请按需保留
import { formatDate, formatBoolean } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'Task'
})

// 验证规则
const rule = reactive({
  'task.category': [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  'task.title': [{
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
  'task.desc': [{
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
  'task.campus': [{
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
  'task.college': [{
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
  'task.reward': [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  'task.needMain': [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  'task.startTime': [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  'task.endTime': [{
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

const taskWithStagesformData = ref({
  task: {
    id: 0,
    category: 1,
    title: '',
    desc: '',
    campus: '全部',
    college: '全部',
    reward: 0,
    needMain: false,
    startTime: new Date(),
    endTime: new Date(),
  },
  taskStages: [{
    stage: 1,
    title: '',
    desc: '',
    requiredItem: '',
    imgs: [],
    needPic: false,
    needFace: false,
    pic: '',
    needLoc: false,
    allowDist: 0,
    needNav: false,
    loc: '',
    needCamera: false,
    cameraId: 0
  }],
})

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateTaskFunc = async(row) => {
  setPublishTime()
  const res = await findTask({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    taskWithStagesformData.value = res.data.retask
    dialogFormVisible.value = true
  }
  activeStep.value = res.data.retask.taskStages.length
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
  setPublishTime()
  activeStep.value = 1
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  taskWithStagesformData.value = {
    task: {
      id: 0,
      category: 1,
      title: '',
      desc: '',
      campus: '全部',
      college: '全部',
      reward: 0,
      needMain: false,
      startTime: new Date(),
      endTime: new Date(),
    },
    taskStages: [{
      stage: 1,
      title: '',
      desc: '',
      requiredItem: '',
      imgs: [],
      needPic: false,
      needFace: false,
      pic: '',
      needLoc: false,
      allowDist: 0,
      needNav: false,
      loc: '',
      needCamera: false,
      cameraId: 0
    }],
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createTask(taskWithStagesformData.value)
        break
      case 'update':
        res = await updateTask(taskWithStagesformData.value)
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

const publishTimeShow = ref(true)
const buttonContent = ref('')
const activeStep = ref(0)

function notSetPublishTime() {
  publishTimeShow.value = false
  buttonContent.value = '设置任务时间'
  taskWithStagesformData.value.task.startTime = taskWithStagesformData.value.task.endTime = new Date('2000-01-01')
}

function setPublishTime() {
  publishTimeShow.value = true
  buttonContent.value = '暂不发布任务'
  taskWithStagesformData.value.task.startTime = taskWithStagesformData.value.task.endTime = new Date()
}

function switchIsSetPublishTime() {
  if (publishTimeShow.value) {
    notSetPublishTime()
  } else {
    setPublishTime()
  }
}

function addTaskStage() {
  taskWithStagesformData.value.taskStages.push({
    stage: taskWithStagesformData.value.taskStages.length + 1,
    title: '',
    desc: '',
    requiredItem: '',
    imgs: [],
    needPic: false,
    needFace: false,
    pic: '',
    needLoc: false,
    allowDist: 0,
    needNav: false,
    loc: '',
    needCamera: false,
    cameraId: 0
  })
  activeStep.value += 1
}

function needCamera(taskStage) {
  taskStage.needPic = false
  taskStage.needFace = false
  taskStage.pic = ''
  taskStage.needLoc = false
  taskStage.allowDist = 0
}

function removeTaskStage() {
  if (taskWithStagesformData.value.taskStages.length !== 1) {
    taskWithStagesformData.value.taskStages.pop()
    activeStep.value -= 1
  }
}

const amapDialog = ref(false)

const amapParam = ref({
  index: 0,
  loc: {
    longitude: 0,
    latitude: 0
  }
})

const openAmapDialog = (i) => {
  if (taskWithStagesformData.value.taskStages[i].loc === '') {
    amapParam.value = {
      index: i,
      loc: {
        longitude: 0,
        latitude: 0
      }
    }
  } else {
    const locParts = taskWithStagesformData.value.taskStages[i].loc.split(',')
    amapParam.value = {
      index: i,
      loc: {
        longitude: locParts[0],
        latitude: locParts[1],
      }
    }
  }
  amapDialog.value = true
}

const closeAmapDialog = () => {
  taskWithStagesformData.value.taskStages[amapParam.value.index].loc = amapParam.value.loc.longitude + ',' + amapParam.value.loc.latitude
  amapDialog.value = false
}

function formatCategory(category) {
  switch (category) {
    case 1:
      return '主线任务'
    case 2:
      return '支线任务'
    case 3:
      return '隐藏任务'
  }
}

</script>

<style>

</style>
