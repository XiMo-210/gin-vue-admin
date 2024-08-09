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
          label="校区"
          prop="name"
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
          label="居住性别"
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
          label="所属寝室楼"
          prop="name"
        >
          <el-input
            v-model="searchInfo.name"
            placeholder="所属寝室楼"
          />
        </el-form-item>
        <el-form-item
          label="宿舍号"
          prop="roomNum"
        >
          <el-input
            v-model.number="searchInfo.roomNum"
            placeholder="宿舍号"
          />
        </el-form-item>
        <el-form-item
          label="容纳人数"
          prop="capacity"
        >
          <el-input
            v-model.number="searchInfo.capacity"
            placeholder="容纳人数"
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
        <el-button
          type="primary"
          @click="openQuestionDialog"
        >发布问卷</el-button>
        <el-button
          type="primary"
          @click="openAllocDialog"
        >智能分寝</el-button>
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
          width="60"
        />
        <el-table-column
          label="校区"
          prop="campus"
          width="120"
        />
        <el-table-column
          label="居住性别"
          prop="gender"
          width="120"
        />
        <el-table-column
          label="所属寝室楼"
          prop="name"
          width="120"
        />
        <el-table-column
          label="宿舍房间号"
          prop="roomNum"
          width="120"
        />
        <el-table-column
          label="容纳人数"
          prop="capacity"
          width="120"
        />
        <el-table-column
          label="已入住人数"
          prop="occupied"
          width="120"
        />
        <el-table-column
          label="备注"
          prop="remark"
          width="120"
        >
          <template #default="scope">{{ scope.row.remark===''?"无":scope.row.remark }}</template>
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
              @click="updateDormRoomFunc(scope.row)"
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
      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="right"
        :rules="rule"
        label-width="120px"
        inline="true"
      >
        <el-form-item
          label="所属寝室楼:"
          prop="dormitoryBuildingId"
          style="width: 60%;"
        >
          <el-select
            v-model.number="formData.dormitoryBuildingId"
            placeholder="请选择"
            clearable
            filterable
            :disabled="type==='update'"
          >
            <el-option
              v-for="item in buildingData"
              :key="item.ID"
              :label="item.campus+' '+item.name+' '+item.gender"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          v-show="type==='create'"
          label="宿舍号开始:"
          prop="roomNumStart"
        >
          <el-input-number
            v-model.number="formData.roomNumStart"
            :clearable="true"
            placeholder="请输入宿舍号开始"
          />
        </el-form-item>
        <el-form-item
          v-show="type==='create'"
          label="宿舍号结束:"
          prop="roomNumEnd"
        >
          <el-input-number
            v-model.number="formData.roomNumEnd"
            :clearable="true"
            placeholder="请输入宿舍号结束"
          />
        </el-form-item>
        <el-form-item
          v-show="type!=='create'"
          label="宿舍房间号:"
          prop="roomNum"
        >
          <el-input-number
            v-model.number="formData.roomNum"
            :clearable="true"
            placeholder="请输入宿舍号"
            disabled
          />
        </el-form-item>
        <el-form-item
          label="容纳人数:"
          prop="capacity"
        >
          <el-input-number
            v-model.number="formData.capacity"
            :clearable="true"
            placeholder="请输入容纳人数"
            :disabled="type==='update'"
          />
        </el-form-item>
        <el-form-item
          v-show="type==='update'"
          label="备注:"
          prop="remark"
          style="width: 75%;"
        >
          <el-input
            v-model="formData.remark"
            :clearable="true"
            placeholder="请输入备注"
            type="textarea"
            :autosize="{ minRows: 1, maxRows: 4 }"
          />
        </el-form-item>
        <el-table
          v-show="type==='update'"
          style="width: 100%"
          :data="studentData"
          row-key="ID"
          :header-cell-style="{ 'text-align': 'center' }"
          :cell-style="{ textAlign: 'center' }"
        >
          <el-table-column
            label="学号"
            prop="studentId"
            width="120"
          />
          <el-table-column
            label="姓名"
            prop="name"
            width="80"
          />
          <el-table-column
            label="学院"
            prop="college"
            width="120"
          />
          <el-table-column
            label="专业"
            prop="major"
            width="240"
          />
          <el-table-column
            label="班级"
            prop="class"
            width="280"
          />
        </el-table>
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
    <el-dialog
      v-model="questionVisible"
      :before-close="closeQuestionDialog"
      title="问卷"
      destroy-on-close
      style="width: 500px;"
    >
      <el-text
        tag="b"
        type="large"
      >问卷收集中······ 截至时间2024年8月9日</el-text>
      <br><br>
      <el-text
        tag="b"
        type="large"
      >已填写人数: 4382人  未填写人数: 636人</el-text>
      <el-button
        type="warning"
        style="float: right;"
        @click="remindUnfill"
      >提醒</el-button>
      <br>
      <br>
      <el-collapse>
        <el-collapse-item
          title="查看问卷内容"
          name="1"
        >
          <el-card>
            <template #header>
              <div>
                <span>1. 是否玩游戏？</span>
              </div>
            </template>
            <el-radio-group v-model="q1">
              <el-radio
                label="很少"
                :value="0"
                border
              >很少</el-radio>
              <el-radio
                label="偶尔"
                :value="1"
                border
              >偶尔</el-radio>
              <el-radio
                label="经常"
                :value="2"
                border
              >经常</el-radio>
            </el-radio-group>
          </el-card>
          <el-card style="margin-top: 20px;">
            <template #header>
              <div>
                <span>2. 起床时间？</span>
              </div>
            </template>
            <el-radio-group v-model="q2">
              <el-radio
                label="较早"
                :value="0"
                border
              >较早</el-radio>
              <el-radio
                label="正常"
                :value="1"
                border
              >正常</el-radio>
              <el-radio
                label="较晚"
                :value="2"
                border
              >较晚</el-radio>
            </el-radio-group>
          </el-card>
          <el-card style="margin-top: 20px;">
            <template #header>
              <div>
                <span>3. 睡觉时间？</span>
              </div>
            </template>
            <el-radio-group v-model="q3">
              <el-radio
                label="较早"
                :value="0"
                border
              >较早</el-radio>
              <el-radio
                label="正常"
                :value="1"
                border
              >正常</el-radio>
              <el-radio
                label="较晚"
                :value="2"
                border
              >较晚</el-radio>
            </el-radio-group>
          </el-card>
          <el-card style="margin-top: 20px;">
            <template #header>
              <div>
                <span>4. 爱好偏向？</span>
              </div>
            </template>
            <el-radio-group v-model="q4">
              <el-radio
                label="偏静"
                :value="0"
                border
              >偏静</el-radio>
              <el-radio
                label="不确定"
                :value="1"
                border
              >不确定</el-radio>
              <el-radio
                label="偏动"
                :value="2"
                border
              >偏动</el-radio>
            </el-radio-group>
          </el-card>
          <el-card style="margin-top: 20px;">
            <template #header>
              <div>
                <span>5. 社交偏向</span>
              </div>
            </template>
            <el-radio-group v-model="q5">
              <el-radio
                label="社恐"
                :value="0"
                border
              >社恐</el-radio>
              <el-radio
                label="不确定"
                :value="1"
                border
              >不确定</el-radio>
              <el-radio
                label="社牛"
                :value="2"
                border
              >社牛</el-radio>
            </el-radio-group>
          </el-card>
        </el-collapse-item>
      </el-collapse>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeQuestionDialog">取 消</el-button>
          <el-button
            type="danger"
            @click="enterQuestionDialog"
          >停止发布</el-button>
        </div>
      </template>
    </el-dialog>
    <el-dialog
      v-model="allocVisible"
      :before-close="closeAllocDialog"
      title="智能分寝"
      destroy-on-close
      style="width: 500px;"
    >
      将根据收集的问卷信息进行智能分寝，若问卷在收集中则会提前截止，是否确定？
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeAllocDialog">取 消</el-button>
          <el-button
            type="primary"
            @click="enterAllocDialog"
          >确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createDormRoom,
  deleteDormRoom,
  deleteDormRoomByIds,
  updateDormRoom,
  findDormRoom,
  getDormRoomList,
  dormAlloc,
  getDormStudents,
  remind
} from '@/api/dormRoom'
import {
  getDormitoryBuildingList
} from '@/api/dormitoryBuilding'

// 全量引入格式化工具 请按需保留
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'DormRoom'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  dormitoryBuildingId: 1,
  campus: '',
  gender: '',
  name: '',
  roomNum: 0,
  roomNumStart: 101,
  roomNumEnd: 130,
  capacity: 4,
  occupied: 0,
  remark: '',
})

const buildingData = ref({
  ID: 0,
  campus: '',
  name: '',
  gender: '',
})

// 验证规则
const rule = reactive({
  dormitoryBuildingId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
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
  const table = await getDormRoomList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteDormRoomFunc(row)
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
  const res = await deleteDormRoomByIds({ IDs })
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
const updateDormRoomFunc = async(row) => {
  const res = await findDormRoom({ ID: row.ID })
  getBuildingList()
  getDormStudent(row)
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.redormRoom
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteDormRoomFunc = async(row) => {
  const res = await deleteDormRoom({ ID: row.ID })
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
  getBuildingList()
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    dormitoryBuildingId: 1,
    roomNum: 0,
    roomNumStart: 101,
    roomNumEnd: 130,
    capacity: 4,
    occupied: 0,
    remark: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createDormRoom(formData.value)
        break
      case 'update':
        res = await updateDormRoom(formData.value)
        break
      default:
        res = await createDormRoom(formData.value)
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

const getBuildingList = async() => {
  const list = await getDormitoryBuildingList({ page: 1, pageSize: 1000 })
  if (list.code === 0) {
    buildingData.value = list.data.list
  }
}

const studentData = ref([])

const getDormStudent = async(row) => {
  const students = await getDormStudents({ ID: row.ID })
  if (students.code === 0) {
    studentData.value = students.data
  }
}

// 弹窗控制标记
const questionVisible = ref(false)

// 打开弹窗
const openQuestionDialog = () => {
  questionVisible.value = true
}

// 关闭弹窗
const closeQuestionDialog = () => {
  questionVisible.value = false
}

// 弹窗确定
const enterQuestionDialog = async() => {
}

// 弹窗控制标记
const allocVisible = ref(false)

// 打开弹窗
const openAllocDialog = () => {
  allocVisible.value = true
}

// 关闭弹窗
const closeAllocDialog = () => {
  allocVisible.value = false
}

// 弹窗确定
const enterAllocDialog = async() => {
  const res = await dormAlloc()
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '分配完毕'
    })
    closeAllocDialog()
    getTableData()
  }
}

const q1 = ref(0)
const q2 = ref(0)
const q3 = ref(0)
const q4 = ref(0)
const q5 = ref(0)

const remindUnfill = async() => {
  const res = await remind()
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '已提醒未填写学生'
    })
  }
}

</script>

<style>

</style>
