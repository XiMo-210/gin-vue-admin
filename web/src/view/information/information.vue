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
          label="信息类别"
          prop="category"
        >
          <el-select
            v-model.number="searchInfo.category"
            placeholder="请选择"
            clearable
            style="width: 100px;"
          >
            <el-option
              key="notice"
              label="通知"
              :value="1"
            />
            <el-option
              key="news"
              label="新闻"
              :value="2"
            />
            <el-option
              key="announcement"
              label="公告"
              :value="3"
            />
          </el-select>

        </el-form-item>
        <el-form-item
          label="标题"
          prop="title"
        >
          <el-input
            v-model="searchInfo.title"
            placeholder="标题"
          />

        </el-form-item>
        <el-form-item
          label="来源"
          prop="source"
        >
          <el-input
            v-model="searchInfo.source"
            placeholder="来源"
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
          label="发布时间"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column
          align="left"
          label="信息类别"
          prop="category"
          width="120"
        >
          <template #default="scope">
            {{ formatCategory(scope.row.category) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="标题"
          prop="title"
          width="160"
          show-overflow-tooltip="true"
        />
        <el-table-column
          align="left"
          label="来源"
          prop="source"
          width="160"
          show-overflow-tooltip="true"
        />
        <el-table-column
          align="left"
          label="内容"
          prop="content"
          width="440"
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
              @click="updateInformationFunc(scope.row)"
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
      style="width: 40%;"
    >
      <el-scrollbar height="550px">
        <el-form
          ref="elFormRef"
          :model="formData"
          label-position="right"
          :rules="rule"
          label-width="100px"
        >
          <el-form-item
            label="信息类别:"
            prop="category"
          >
            <el-select
              v-model="formData.category"
              placeholder="请选择"
              clearable
            >
              <el-option
                key="notice"
                label="通知"
                :value="1"
              />
              <el-option
                key="news"
                label="新闻"
                :value="2"
              />
              <el-option
                key="announcement"
                label="公告"
                :value="3"
              />
            </el-select>
          </el-form-item>
          <el-form-item
            label="标题:"
            prop="title"
          >
            <el-input
              v-model="formData.title"
              :clearable="true"
              placeholder="请输入标题"
            />
          </el-form-item>
          <el-form-item
            label="来源:"
            prop="source"
          >
            <el-input
              v-model="formData.source"
              :clearable="true"
              placeholder="请输入来源"
            />
          </el-form-item>
          <el-form-item
            label="内容:"
            prop="content"
          >
            <el-input
              v-model="formData.content"
              :clearable="true"
              placeholder="请输入内容"
              :autosize="{minRows:2,maxRows:8}"
              type="textarea"
            />
          </el-form-item>
          <el-form-item
            label="图片:"
            prop="imgs"
          >
            <div>
              <el-upload
                class="upload-demo"
                action="https://api.lonesome.cn/api/wx/upload"
                :show-file-list="false"
                :on-success="uploadSuccess"
              >
                <el-button
                  type="primary"
                >点击上传</el-button>
              </el-upload>

              <div class="image-list">
                <div
                  v-for="(img, index) in formData.imgs"
                  :key="index"
                  class="thumbnail-item"
                >
                  <el-image
                    :src="img"
                    fit="cover"
                    style="width: 100px; height: 100px;"
                  />
                </div>
              </div>
            </div>
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
  createInformation,
  deleteInformation,
  deleteInformationByIds,
  updateInformation,
  findInformation,
  getInformationList
} from '@/api/information'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'Information'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  category: 1,
  title: '',
  source: '',
  content: '',
  imgs: [],
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
  source: [{
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
  content: [{
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
  const table = await getInformationList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteInformationFunc(row)
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
  const res = await deleteInformationByIds({ IDs })
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
const updateInformationFunc = async(row) => {
  const res = await findInformation({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reinformation
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteInformationFunc = async(row) => {
  const res = await deleteInformation({ ID: row.ID })
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
    category: 1,
    title: '',
    source: '',
    content: '',
    imgs: [],
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createInformation(formData.value)
        break
      case 'update':
        res = await updateInformation(formData.value)
        break
      default:
        res = await createInformation(formData.value)
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

function formatCategory(category) {
  switch (category) {
    case 1:
      return '通知'
    case 2:
      return '新闻'
    case 3:
      return '公告'
  }
}

const uploadSuccess = (res) => {
  const { data } = res
  if (data) {
    formData.value.imgs.push(data)
  }
  console.log(formData.value.imgs)
  ElMessage({
    type: 'success',
    message: '上传成功'
  })
}

</script>

<style>
.upload-demo {
  display: inline-block;
  margin-bottom: 20px;
}

.image-list {
  display: flex;
  flex-wrap: wrap;
  margin-top: 10px;
}

.thumbnail-item {
  position: relative;
  margin-right: 10px;
  margin-bottom: 10px;
}
</style>
