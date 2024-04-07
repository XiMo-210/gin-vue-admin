<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elOrganizationFormRef"
        :model="organizationInfo"
        class="demo-form-inline"
        :rules="orgRule"
        label-width="80px"
      >
        <el-form-item
          label="封面图"
          prop="pic"
        >
          <UploadImg
            v-model="organizationInfo.pic"
          />
        </el-form-item>
        <el-form-item
          label="类别"
          prop="category"
          style="width: 20%;"
        >
          <el-select
            v-model.number="organizationInfo.category"
            placeholder="请选择"
          >
            <el-option
              key="org"
              label="学生组织"
              :value="1"
            />
            <el-option
              key="club"
              label="学生社团"
              :value="2"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="名称"
          prop="name"
        >
          <el-input
            v-model="organizationInfo.name"
            placeholder="组织社团名称"
            clearable
            style="width: 50%;"
          />
        </el-form-item>
        <el-form-item
          label="介绍"
          prop="introduction"
        >
          <el-input
            v-model="organizationInfo.introduction"
            placeholder="介绍"
            clearable
            :autosize="{minRows:2,maxRows:8}"
            type="textarea"
            style="width: 50%;"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            icon="mouse"
            type="primary"
            @click="onSubmit"
          >更新信息</el-button>
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
          label="部门名称"
          prop="name"
          width="160"
        />
        <el-table-column
          align="left"
          label="部门描述"
          prop="desc"
          width="800"
          show-overflow-tooltip="true"
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
              icon="edit"
              class="table-button"
              @click="updateDepartmentFunc(scope.row)"
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
      width="600"
      align-center
    >
      <el-scrollbar height="300px">
        <el-form
          ref="elFormRef"
          :model="formData"
          label-position="right"
          :rules="rule"
        >
          <el-form-item
            label="部门名称:"
            prop="name"
          >
            <el-input
              v-model="formData.name"
              :clearable="true"
              placeholder="请输入部门名称"
            />
          </el-form-item>
          <el-form-item
            label="部门描述:"
            prop="desc"
          >
            <el-input
              v-model="formData.desc"
              placeholder="请输入部门描述"
              clearable
              :autosize="{minRows:10 ,maxRows:10}"
              type="textarea"
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
  </div>
</template>

<script setup>
import {
  createOrganization,
  updateOrganization,
  findOrganization,
} from '@/api/organization'
import {
  createDepartment,
  deleteDepartment,
  deleteDepartmentByIds,
  updateDepartment,
  findDepartment,
  getDepartmentList
} from '@/api/department'
// 全量引入格式化工具 请按需保留
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import UploadImg from '@/components/uploadImg/uploadImg.vue'

defineOptions({
  name: 'Organization'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  organizationId: 0,
  name: '',
  desc: '',
})

// 验证规则
const rule = reactive({
  organizationId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
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
})

const elFormRef = ref()
const elOrganizationFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

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
  const table = await getDepartmentList({ page: page.value, pageSize: pageSize.value })
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
    deleteDepartmentFunc(row)
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
  const res = await deleteDepartmentByIds({ IDs })
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
const updateDepartmentFunc = async(row) => {
  const res = await findDepartment({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.redepartment
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteDepartmentFunc = async(row) => {
  const res = await deleteDepartment({ ID: row.ID })
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
  if (organizationInfo.value.ID === 0) {
    ElMessage({
      'type': 'warning',
      'message': '请先填写组织部门信息',
    })
    return
  }
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    organizationId: 0,
    name: '',
    desc: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    formData.value.organizationId = organizationInfo.value.ID
    switch (type.value) {
      case 'create':
        res = await createDepartment(formData.value)
        break
      case 'update':
        res = await updateDepartment(formData.value)
        break
      default:
        res = await createDepartment(formData.value)
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

const orgType = ref('create')

const organizationInfo = ref({
  ID: 0,
  category: 0,
  name: '',
  pic: '',
  introduction: ''
})

const orgRule = reactive({
  introduction: [{
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
  category: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
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
  pic: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
})

const onReset = () => {
  getOrganizationInfo()
}

const onSubmit = () => {
  elOrganizationFormRef.value?.validate(async(valid) => {
    if (!valid) {
      ElMessage({
        type: 'warning',
        message: '组织社团信息不能为空'
      })
      return
    }
    let res
    switch (orgType.value) {
      case 'create':
        res = await createOrganization(organizationInfo.value)
        break
      case 'update':
        res = await updateOrganization(organizationInfo.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      getOrganizationInfo()
    }
  })
}

const getOrganizationInfo = async() => {
  const info = await findOrganization()
  if (info.code === 0) {
    organizationInfo.value = info.data.reorganization
  }
  if (info.data.reorganization.ID === 0) {
    orgType.value = 'create'
    organizationInfo.value.category = 1
  } else {
    orgType.value = 'update'
  }
}

getOrganizationInfo()

</script>

<style>

</style>
