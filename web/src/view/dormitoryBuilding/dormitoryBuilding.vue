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
          label="所属校区"
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
          label="寝室楼名称"
          prop="name"
        >
          <el-input
            v-model="searchInfo.name"
            placeholder="寝室楼名称"
          />

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
          label="管理员姓名"
          prop="managerName"
        >
          <el-input
            v-model="searchInfo.managerName"
            placeholder="管理员姓名"
            style="width: 140px"
          />

        </el-form-item>
        <el-form-item
          label="管理员联系方式"
          prop="managerContact"
        >
          <el-input
            v-model="searchInfo.managerContact"
            placeholder="管理员联系方式"
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
          label="ID"
          prop="ID"
          width="60"
        />
        <el-table-column
          label="所属校区"
          prop="campus"
          width="160"
        />
        <el-table-column
          label="寝室楼名称"
          prop="name"
          width="160"
        />
        <el-table-column
          label="居住性别"
          prop="gender"
          width="120"
        />
        <el-table-column
          label="管理员姓名"
          prop="managerName"
          width="160"
        />
        <el-table-column
          label="管理员联系方式"
          prop="managerContact"
          width="160"
        />
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
              @click="updateDormitoryBuildingFunc(scope.row)"
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
      style="width: 500px;"
      destroy-on-close
    >
      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="right"
        :rules="rule"
        label-width="130px"
      >
        <el-form-item
          label="所属校区:"
          prop="campus"
        >
          <el-select
            v-model="formData.campus"
            placeholder="请选择"
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
          label="寝室楼名称:"
          prop="name"
        >
          <el-input
            v-model="formData.name"
            :clearable="true"
            placeholder="请输入寝室楼名称"
          />
        </el-form-item>
        <el-form-item
          label="居住性别:"
          prop="gender"
        >
          <el-select
            v-model="formData.gender"
            placeholder="请选择"
          >
            <el-option
              key="male"
              value="男"
            />
            <el-option
              key="female"
              value="女"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="管理员姓名:"
          prop="managerName"
        >
          <el-input
            v-model="formData.managerName"
            :clearable="true"
            placeholder="请输入管理员姓名"
          />
        </el-form-item>
        <el-form-item
          label="管理员联系方式:"
          prop="managerContact"
        >
          <el-input
            v-model="formData.managerContact"
            :clearable="true"
            placeholder="请输入管理员联系方式"
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
  createDormitoryBuilding,
  deleteDormitoryBuilding,
  deleteDormitoryBuildingByIds,
  updateDormitoryBuilding,
  findDormitoryBuilding,
  getDormitoryBuildingList
} from '@/api/dormitoryBuilding'

import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'DormitoryBuilding'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  campus: '',
  name: '',
  gender: '',
  managerName: '',
  managerContact: '',
})

// 验证规则
const rule = reactive({
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
  managerName: [{
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
  managerContact: [{
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
  const table = await getDormitoryBuildingList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteDormitoryBuildingFunc(row)
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
  const res = await deleteDormitoryBuildingByIds({ IDs })
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
const updateDormitoryBuildingFunc = async(row) => {
  const res = await findDormitoryBuilding({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.redormitoryBuilding
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteDormitoryBuildingFunc = async(row) => {
  const res = await deleteDormitoryBuilding({ ID: row.ID })
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
    campus: '',
    name: '',
    gender: '',
    managerName: '',
    managerContact: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createDormitoryBuilding(formData.value)
        break
      case 'update':
        res = await updateDormitoryBuilding(formData.value)
        break
      default:
        res = await createDormitoryBuilding(formData.value)
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
