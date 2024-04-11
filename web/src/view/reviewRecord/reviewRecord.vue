<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item
          label="用户名"
          prop="username"
        >
          <el-input
            v-model.number="searchInfo.username"
            placeholder="用户名"
            style="width: 140px"
          />
        </el-form-item>
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
          prop="taskTitle"
        >
          <el-input
            v-model.number="searchInfo.taskTitle"
            placeholder="任务名称"
            style="width: 140px"
          />
        </el-form-item>
        <el-form-item
          label="状态"
          prop="status"
        >
          <el-select
            v-model.number="searchInfo.status"
            clearable
            placeholder="请选择"
            style="width: 120px"
          >
            <el-option
              key="notreview"
              label="未审核"
              :value="1"
            />
            <el-option
              key="pass"
              label="通过"
              :value="2"
            />
            <el-option
              key="reject"
              label="驳回"
              :value="3"
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
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        :header-cell-style="{ 'text-align': 'center' }"
        :cell-style="{ textAlign: 'center' }"
      >
        <el-table-column
          label="ID"
          prop="ID"
          width="60"
        />
        <el-table-column
          label="申请日期"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column
          label="用户名"
          prop="username"
          width="120"
        />
        <el-table-column
          label="任务类型"
          prop="category"
          width="120"
        >
          <template #default="scope">
            {{ formatCategory(scope.row.category) }}
          </template>
        </el-table-column>
        <el-table-column
          label="任务名称"
          prop="taskTitle"
          width="160"
          show-overflow-tooltip="true"
        />
        <el-table-column
          label="任务阶段名称"
          prop="taskStageTitle"
          width="160"
          show-overflow-tooltip="true"
        />
        <el-table-column
          label="任务截止时间"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.endTime) }}</template>
        </el-table-column>
        <el-table-column
          label="状态"
          prop="status"
          width="120"
        >
          <template #default="scope">
            {{ formatStatus(scope.row.status) }}
          </template>
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
              @click="updateReviewRecordsFunc(scope.row)"
            >{{ scope.row.status===1? '审核':'查看' }}</el-button>
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
      title="审核"
      align-center
      destroy-on-close
    >
      <el-scrollbar>
        <el-form
          ref="elFormRef"
          :inline="true"
          :model="formData"
          label-position="right"
        ><el-collapse>
           <el-collapse-item>
             <template #title>
               <el-text
                 tag="b"
                 size="large"
               >
                 {{ "任务类型: "+formatCategory( formData.category ) }}
                &nbsp;&nbsp;&nbsp;&nbsp;
                 {{ "任务名称: "+formData.taskTitle }}
               </el-text>
             </template>
             <el-form-item
               label="任务校区:"
               prop="campus"
               style="width: 40%;"
             >
               <el-select
                 v-model="formData.campus"
                 disabled
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
               prop="college"
               style="width: 40%;"
             >
               <el-input
                 v-model="formData.college"
                 disabled
               />
             </el-form-item>
             <el-form-item
               label="开始时间:"
               prop="startTime"
               style="width: 30%;"
             >
               <el-date-picker
                 v-model="formData.startTime"
                 type="date"
                 disabled
                 :clearable="true"
               />
             </el-form-item>
             <el-form-item
               label="结束时间:"
               prop="endTime"
               style="width: 30%;"
             >
               <el-date-picker
                 v-model="formData.endTime"
                 type="date"
                 disabled
                 :clearable="true"
               />
             </el-form-item>
             <el-form-item
               label="任务描述:"
               prop="taskDesc"
               style="width: 80%;"
             >
               <el-input
                 v-model="formData.taskDesc"
                 disabled
                 :autosize="{minRows:2,maxRows:4}"
                 type="textarea"
               />
             </el-form-item>
           </el-collapse-item>
           <el-collapse-item>
             <template #title>
               <el-text
                 tag="b"
                 size="large"
               >{{ "阶段任务名称: "+formData.taskStageTitle }}
               </el-text>
             </template>
             <el-form-item
               v-show="formData.requiredItem"
               label="所需物品:"
               prop="requiredItem"
               style="width: 40%;"
             >
               <el-input
                 v-model="formData.requiredItem"
                 disabled
               />
             </el-form-item>
             <p />
             <el-form-item
               label="阶段任务描述:"
               :prop="'desc'"
               style="width: 80%;"
             >
               <el-input
                 v-model="formData.taskStageDesc"
                 disabled
                 :autosize="{minRows:1,maxRows:2}"
                 type="textarea"
               />
             </el-form-item>
             <el-form-item
               v-show="formData.imgs.length"
               label="图片描述:"
               prop="imgs"
             >
               <span style="display: flex;flex-wrap: wrap; margin-top: 5px; margin-bottom: -20px;">
                 <span
                   v-for="(img, index) in formData.imgs"
                   :key="index"
                   style="position: relative; margin-right: 10px;"
                 >
                   <el-image
                     :src="img"
                     fit="contain"
                     :zoom-rate="1.2"
                     :max-scale="7"
                     :min-scale="0.2"
                     :preview-src-list="formData.imgs"
                     :initial-index="index"
                     style="width: 80px; height: 80px;"
                   />
                 </span>
               </span>
             </el-form-item>
           </el-collapse-item>
         </el-collapse>
          <el-form-item
            label="任务要求照片:"
            prop="taskStagePic"
            style="width: 28%;margin-top: 20px;"
          >
            <el-image
              :src="formData.taskStagePic"
              fit="fill"
              :zoom-rate="1.2"
              :max-scale="7"
              :min-scale="0.2"
              :preview-src-list="[formData.taskStagePic]"
              :initial-index="0"
              style="width: 130px; height: 130px;"
            />
          </el-form-item>
          <el-form-item
            v-show="formData.needFace"
            label="需要用户人脸:"
            prop="fill"
            style="width: 28%;margin-top: 20px;"
          >
            <el-image
              :src="formData.faceId"
              fit="contain"
              :zoom-rate="1.2"
              :max-scale="7"
              :min-scale="0.2"
              :preview-src-list="[formData.faceId]"
              :initial-index="0"
              style="width: 130px; height: 130px;"
            />
          </el-form-item>
          <el-form-item
            label="用户提交照片:"
            prop="fill"
            style="width: 28%;margin-top: 20px;"
          >
            <el-image
              :src="formData.pic"
              fit="contain"
              :zoom-rate="1.2"
              :max-scale="7"
              :min-scale="0.2"
              :preview-src-list="[formData.pic]"
              :initial-index="0"
              style="width: 130px; height: 130px;"
            />
          </el-form-item>
          <el-form-item
            v-show="formData.status===3"
            label="驳回原因"
            style="width: 80%;"
          >
            <el-input
              v-model="formData.reply"
              disabled
              :autosize="{minRows:1,maxRows:2}"
              type="textarea"
            />
          </el-form-item>
        </el-form>
      </el-scrollbar>
      <template #footer>
        <div
          v-if="formData.status===1"
          class="dialog-footer"
        >
          <el-button
            @click="closeDialog"
          >取 消</el-button>
          <el-button
            type="danger"
            @click="openRejectDialog"
          >驳 回</el-button>
          <el-button
            type="success"
            @click="enterDialog"
          >通 过</el-button>
        </div>
        <div
          v-else
          class="dialog-footer"
        >
          <el-button
            v-if="formData.status===2"
            type="success"
          >状态: 已通过</el-button>
          <el-button
            v-else-if="formData.status===3"
            type="danger"
          >状态: 已驳回</el-button>
        </div>
      </template>
    </el-dialog>
    <el-dialog
      v-model="rejectDialogVisiable"
      title="驳回"
      :before-close="closeRejectDialog"
      width="30%"
      align-center
      destroy-on-close
    >
      <el-form
        ref="rejectRef"
        :model="formData"
        :rules="rule"
      >
        <el-form-item
          prop="reply"
          label="原因"
        >
          <el-input
            v-model="formData.reply"
            placeholder="请输入驳回原因"
            :autosize="{minRows:2,maxRows:4}"
            type="textarea"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <span>
          <el-button @click="closeRejectDialog">取消</el-button>
          <el-button
            type="primary"
            @click="enterRejectDialog"
          >确认</el-button>
        </span>
      </template>
    </el-dialog>

  </div>
</template>

<script setup>
import {
  deleteReviewRecords,
  updateReviewRecords,
  findReviewRecords,
  getReviewRecordsList
} from '@/api/reviewRecord'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'ReviewRecords'
})

// 验证规则
const rule = reactive({
  'reply': [{
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
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  username: '',
  faceId: 0,
  category: 0,
  taskTitle: '',
  taskDesc: '',
  campus: '',
  college: '',
  startTime: new Date(),
  endTime: new Date(),
  taskStageTitle: '',
  taskStageDesc: '',
  requiredItem: '',
  imgs: [],
  needFace: false,
  taskStagePic: '',
  stage: 0,
  pic: '',
  reply: '',
  status: 0,
})

const elFormRef = ref()
const elSearchFormRef = ref()
const rejectRef = ref()

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
  const table = await getReviewRecordsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteReviewRecordsFunc(row)
  })
}

// 更新行
const updateReviewRecordsFunc = async(row) => {
  const res = await findReviewRecords({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.rereviewRecord
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteReviewRecordsFunc = async(row) => {
  const res = await deleteReviewRecords({ ID: row.ID })
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

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
}

// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    formData.value.status = 2
    const res = await updateReviewRecords(formData.value)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '通过'
      })
      closeDialog()
      getTableData()
    } else {
      formData.value.status = 1
    }
  })
}

const rejectDialogVisiable = ref(false)

const openRejectDialog = () => {
  rejectDialogVisiable.value = true
}

const closeRejectDialog = () => {
  rejectDialogVisiable.value = false
}

const enterRejectDialog = async() => {
  rejectRef.value?.validate(async(valid) => {
    if (!valid) return
    formData.value.status = 3
    const res = await updateReviewRecords(formData.value)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '驳回'
      })
      closeRejectDialog()
      closeDialog()
      getTableData()
    } else {
      formData.value.status = 1
      closeRejectDialog()
    }
  })
}

const formatCategory = (category) => {
  switch (category) {
    case 1:
      return '主线任务'
    case 2:
      return '支线任务'
    case 3:
      return '隐藏任务'
  }
}

const formatStatus = (status) => {
  switch (status) {
    case 1:
      return '未审核'
    case 2:
      return '通过'
    case 3:
      return '驳回'
  }
}

</script>

  <style>

  </style>
