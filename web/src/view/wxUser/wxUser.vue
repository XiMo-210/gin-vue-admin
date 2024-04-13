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
          label="ID"
          prop="ID"
          width="60"
        />
        <el-table-column
          label="注册时间"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column
          label="头像"
          width="80"
        >
          <template #default="scope">
            <el-image
              style="width: 50px; height: 50px"
              :src="getUrl(scope.row.avatar)"
              fit="cover"
            />
          </template>
        </el-table-column>
        <el-table-column
          label="用户名"
          prop="username"
          width="120"
        />
        <el-table-column
          sortable
          label="智慧种子"
          prop="exp"
          width="120"
        />
        <el-table-column
          sortable
          label="求索石"
          prop="points"
          width="120"
        />
        <el-table-column
          label="是否完成主线任务"
          prop="isCompletedMain"
          width="160"
        >
          <template #default="scope">{{ formatBoolean(scope.row.isCompletedMain) }}</template>
        </el-table-column>
        <el-table-column
          label="正在进行的任务ID"
          prop="curTask"
          width="140"
        >
          <template #default="scope">{{ scope.row.curTask===0?"暂无进行任务":scope.row.curTask }}</template>
        </el-table-column>
        <el-table-column
          label="操作"
          fixed="right"
          min-width="360"
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
              icon="link"
              :disabled="scope.row.studentInfoId===0"
              @click="getStudentInfoDetails(scope.row)"
            >
              {{ scope.row.studentInfoId===0?"暂未认证":"关联学生" }}
            </el-button>
            <el-button
              type="primary"
              link
              class="table-button"
              icon="PieChart"
              :disabled="scope.row.studentInfoId===0"
              @click="getUserTaskConditionInfo(scope.row)"
            >
              任务情况
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
      lock-scroll
      :before-close="closeDetailShow"
      title="查看详情"
      destroy-on-close
    >
      <el-form
        :inline="true"
        :model="formData"
        label-position="right"
        label-width="80px"
      >
        <el-form-item
          label="头像:"
          prop="avatar"
        >
          <el-image
            style="width: 130px; height: 130px"
            :src="formData.avatar"
            fit="cover"
          />
        </el-form-item>
        <br>
        <el-form-item
          label="用户名:"
          prop="username"
          style="width: 40%;"
        >
          <el-input
            v-model="formData.username"
            disabled
          />
        </el-form-item>
        <el-form-item
          label="个性签名:"
          prop="signature"
          style="width: 90%;"
        >
          <el-input
            v-model="formData.signature"
            disabled
          />
        </el-form-item>
        <el-form-item
          label="手机:"
          prop="phone"
          style="width: 40%;"
        >
          <el-input
            v-model="formData.phone"
            disabled
          />
        </el-form-item>
        <el-form-item
          label="邮箱:"
          prop="email"
          style="width: 40%;"
        >
          <el-input
            v-model="formData.email"
            disabled
          />
        </el-form-item>
        <el-form-item
          label="智慧种子:"
          prop="exp"
          style="width: 40%;"
        >
          <el-input
            v-model.number="formData.exp"
            disabled
          />
        </el-form-item>
        <el-form-item
          label="求索石:"
          prop="points"
          style="width: 40%;"
        >
          <el-input
            v-model.number="formData.points"
            disabled
          />
        </el-form-item>
        <el-form-item
          label="主线任务:"
          prop="isCompletedMain"
        >
          <el-text tag="b">
            {{ formData.isCompletedMain?"已完成":"未完成" }}
          </el-text>
        </el-form-item>
      </el-form>
    </el-dialog>
    <el-dialog
      v-model="studentInfoDetailShow"
      style="width: 450px;scale: 1.2;"
      lock-scroll
      :before-close="closeStudentInfoDetailShow"
      title="关联学生"
      destroy-on-close
    >
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
    </el-dialog>
    <el-dialog
      v-model="userTaskConditionShow"
      lock-scroll
      :before-close="closeUserTaskConditionShow"
      title="任务情况"
      destroy-on-close
      align-center
    >
      <el-scrollbar max-height="600px">
        <div v-show="userTaskCondition.curTask.task.ID===0">
          <el-text
            tag="b"
            size="large"
            style="color: black;"
          >
            暂无进行中任务
          </el-text>
          <el-divider style="margin-top: 10px;margin-bottom: 10px;" />
        </div>
        <div v-show="userTaskCondition.curTask.task.ID!==0">
          <el-descriptions
            title="正在进行的任务"
            :column="2"
            border
          >
            <el-descriptions-item
              label="任务类型"
              min-width="80px"
            >{{ formatCategory(userTaskCondition.curTask.task.category) }}</el-descriptions-item>
            <el-descriptions-item
              label="奖励积分"
              min-width="80px"
            >{{ userTaskCondition.curTask.task.reward }}</el-descriptions-item>
            <el-descriptions-item label="任务校区">{{ userTaskCondition.curTask.task.campus }}</el-descriptions-item>
            <el-descriptions-item label="任务学院">{{ userTaskCondition.curTask.task.college }}</el-descriptions-item>
            <el-descriptions-item label="任务名称">{{ userTaskCondition.curTask.task.title }}</el-descriptions-item>
            <el-descriptions-item label="任务描述">{{ userTaskCondition.curTask.task.desc }}</el-descriptions-item>
          </el-descriptions>
          <el-collapse>
            <el-collapse-item>
              <template #title>
                <el-text
                  tag="b"
                  style="color:cadetblue"
                >
                  阶段详情
                </el-text>
              </template>
              <el-steps
                direction="vertical"
                :active="userTaskCondition.curUserTask.curStage-1"
                finish-status="success"
              >
                <el-step
                  v-for="(taskStage, index) in userTaskCondition.curTask.taskStages"
                  :key="index"
                  :title="'Stage'+(index+1)"
                >
                  <template #description>
                    {{ taskStage.title }}
                    <br>
                    {{ taskStage.desc }}
                    <br>
                    {{ (index+2===userTaskCondition.curUserTask.curStage? "完成时间: "+formatDate(userTaskCondition.curUserTask.UpdatedAt):'') }}
                  </template>
                </el-step>
              </el-steps>
            </el-collapse-item>
          </el-collapse>
        </div>
        <div style="margin-top: 10px;margin-bottom: 10px;">
          <el-text
            tag="b"
            size="large"
            style="color: black;"
          >
            已完成任务
          </el-text>
        </div>
        <el-table
          :data="userTaskCondition.completedTasks"
          stripe
          :header-cell-style="{ 'text-align': 'center' }"
          :cell-style="{ textAlign: 'center' }"
        >
          <el-table-column
            prop="finishTime"
            label="完成时间"
            width="160"
          >
            <template #default="scope">
              {{ formatDate(scope.row.finishTime) }}
            </template>
          </el-table-column>
          <el-table-column
            prop="category"
            label="任务类型"
            width="120"
          >
            <template #default="scope">
              {{ formatCategory(scope.row.category) }}
            </template>
          </el-table-column>
          <el-table-column
            prop="title"
            label="任务名称"
            width="120"
            show-overflow-tooltip="true"
          />
          <el-table-column
            prop="desc"
            label="任务描述"
            width="160"
            show-overflow-tooltip="true"
          />
          <el-table-column
            prop="campus"
            label="任务校区"
            width="120"
          />
          <el-table-column
            prop="college"
            label="任务学院"
            width="120"
          />
          <el-table-column
            prop="reward"
            label="奖励积分"
            width="100"
          />
        </el-table>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  deleteWxUser,
  findWxUser,
  getWxUserList
} from '@/api/wxUser'
import {
  findStudentInfo,
} from '@/api/studentInfo'
import {
  getUserTaskCondition,
} from '@/api/userTask'
import { getUrl } from '@/utils/image'
// 全量引入格式化工具 请按需保留
import { formatDate, formatBoolean } from '@/utils/format'
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

const userTaskCondition = ref({
  curTask: {
    task: {
      ID: 0,
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
  },
  curUserTask: {
    curStage: 1,
    UpdatedAt: new Date(),
  },
  completedTasks: [{
    category: 1,
    title: '',
    desc: '',
    campus: '全部',
    college: '全部',
    reward: '',
    finishTime: new Date(),
  }]
})

const userTaskConditionShow = ref(false)

const openUserTaskConditionShow = () => {
  userTaskConditionShow.value = true
}

const getUserTaskConditionInfo = async(row) => {
  const res = await getUserTaskCondition({ ID: row.ID })
  if (res.code === 0) {
    userTaskCondition.value = res.data
    openUserTaskConditionShow()
  }
}

const closeUserTaskConditionShow = () => {
  userTaskConditionShow.value = false
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
