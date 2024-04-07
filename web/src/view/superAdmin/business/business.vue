
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
          label="名称"
          prop="name"
        >
          <el-input
            v-model="searchInfo.name"
            placeholder="名称"
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
          label="创建日期"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column
          label="名称"
          prop="name"
          width="120"
        />
        <el-table-column
          label="封面图"
          width="100"
        >
          <template #default="scope">
            <el-image
              style="width: 50px; height: 50px"
              :src="scope.row.pic"
              fit="cover"
            />
          </template>
        </el-table-column>
        <el-table-column
          label="地址"
          prop="address"
          width="160"
          :show-overflow-tooltip="true"
        />
        <el-table-column
          label="营业时间"
          prop="businessHours"
          width="160"
          :show-overflow-tooltip="true"
        />
        <el-table-column
          label="联系方式"
          prop="contact"
          width="120"
        />
        <el-table-column
          label="介绍"
          prop="introduction"
          width="300"
          :show-overflow-tooltip="true"
        />
        <el-table-column
          label="关联用户ID"
          prop="sysUserId"
          width="120"
        >
          <template #default="scope">{{ scope.row.sysUserId===0?'暂未关联':scope.row.sysUserId }}</template>
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
              @click="updateBusinessFunc(scope.row)"
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
        :inline="true"
        label-width="100px"
      >

        <el-form-item
          label="封面图"
          prop="pic"
        >
          <UploadImg
            v-model="formData.pic"
          />
        </el-form-item>
        <p />
        <el-form-item
          label="名称"
          prop="name"
          style="width: 50%;"
        >
          <el-input
            v-model="formData.name"
            placeholder="店铺商家名称"
            clearable
          />
        </el-form-item>
        <p />
        <el-form-item
          label="地址"
          prop="address"
          style="width: 50%;"
        >
          <el-input
            v-model="formData.address"
            placeholder="地址"
            clearable
          />
        </el-form-item>
        <p />
        <el-form-item
          label="导航地址"
          prop="loc"
        >
          <el-input
            v-model="formData.loc"
            placeholder="导航位置"
            disabled
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            size="default"
            @click="openAmapDialog"
          >设置导航地址
          </el-button>
        </el-form-item>
        <p />
        <el-form-item
          label="营业时间"
          prop="businessHours"
          style="width: 43%;"
        >
          <el-input
            v-model="formData.businessHours"
            placeholder="营业时间"
            clearable
          />
        </el-form-item>
        <el-form-item
          label="联系方式"
          prop="contact"
          style="width: 43%;"
        >
          <el-input
            v-model="formData.contact"
            placeholder="联系方式"
            clearable
          />
        </el-form-item>
        <p />
        <el-form-item
          label="介绍"
          prop="introduction"
          style="width: 90%;"
        >
          <el-input
            v-model="formData.introduction"
            placeholder="介绍"
            clearable
            :autosize="{minRows:2,maxRows:8}"
            type="textarea"
          />
        </el-form-item>
        <el-form-item
          label="关联用户ID:"
          prop="sysUserId"
        >
          <el-input
            v-model.number="formData.sysUserId"
            placeholder="请输入关联用户ID"
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
    <el-dialog
      v-model="amapDialog"
      :before-close="closeAmapDialog"
      destroy-on-close
    >
      <Amap
        v-model="amapLoc"
      />
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createBusinessByAdmin,
  deleteBusiness,
  deleteBusinessByIds,
  updateBusinessByAdmin,
  findBusinessByAdmin,
  getBusinessList
} from '@/api/business'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import UploadImg from '@/components/uploadImg/uploadImg.vue'
import Amap from '@/components/amap/amap.vue'

defineOptions({
  name: 'Business'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  sysUserId: 0,
  name: '',
  pic: '',
  address: '',
  loc: '',
  businessHours: '',
  contact: '',
  introduction: '',
})

// 验证规则
const rule = reactive({
  sysUserId: [{
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
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  address: [{
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
  loc: [{
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
  businessHours: [{
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
  contact: [{
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
  const table = await getBusinessList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteBusinessFunc(row)
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
  const res = await deleteBusinessByIds({ IDs })
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
const updateBusinessFunc = async(row) => {
  const res = await findBusinessByAdmin({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.rebusiness
    dialogFormVisible.value = true
    const locParts = formData.value.loc.split(',')
    amapLoc.value = {
      longitude: locParts[0],
      latitude: locParts[1],
    }
  }
}

// 删除行
const deleteBusinessFunc = async(row) => {
  const res = await deleteBusiness({ ID: row.ID })
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
  amapLoc.value = {
    longitude: 0,
    latitude: 0,
  }
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    sysUserId: 0,
    name: '',
    pic: '',
    address: '',
    loc: '',
    businessHours: '',
    contact: '',
    introduction: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createBusinessByAdmin(formData.value)
        break
      case 'update':
        res = await updateBusinessByAdmin(formData.value)
        break
      default:
        res = await createBusinessByAdmin(formData.value)
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

const amapDialog = ref(false)

const amapLoc = ref({
  longitude: 0,
  latitude: 0,
})

const openAmapDialog = () => {
  amapDialog.value = true
}

const closeAmapDialog = () => {
  formData.value.loc = amapLoc.value.longitude + ',' + amapLoc.value.latitude
  amapDialog.value = false
}

</script>

  <style>

  </style>

