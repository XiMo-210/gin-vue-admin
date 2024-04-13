<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elAdvertiserFormRef"
        :model="advertiserInfo"
        class="demo-form-inline"
        :rules="advertiserRule"
        :inline="true"
        label-width="80px"
      >
        <el-form-item
          label="标志"
          prop="logo"
        >
          <UploadImg
            v-model="advertiserInfo.logo"
          />
        </el-form-item>
        <p />
        <el-form-item
          label="简称"
          prop="name"
          style="width: 20%;"
        >
          <el-input
            v-model="advertiserInfo.name"
            placeholder="简称"
            clearable
          />
        </el-form-item>
        <el-form-item
          label="简介"
          prop="introduction"
          style="width: 70%;"
        >
          <el-input
            v-model="advertiserInfo.introduction"
            placeholder="简介"
            clearable
            :autosize="{maxRows:1}"
            type="textarea"
          />
        </el-form-item>
        <el-form-item
          label="主体信息"
          prop="subjectInfo"
          label-width="80px"
          style="width: 25%;"
        >
          <el-input
            v-model="advertiserInfo.subjectInfo"
            placeholder="主体信息"
            disabled
          />
        </el-form-item>
        <el-form-item
          label="营业执照注册号"
          prop="license"
          label-width="120px"
          style="width: 25%;"
        >
          <el-input
            v-model="advertiserInfo.license"
            placeholder="营业执照注册号"
            disabled
          />
        </el-form-item>
        <el-form-item
          label="所在省市"
          prop="location"
          label-width="80px"
          style="width: 20%;"
        >
          <el-input
            v-model="advertiserInfo.location"
            placeholder="所在省市"
            disabled
          />
        </el-form-item>
        <el-form-item>
          <el-button
            icon="mouse"
            type="primary"
            @click="onAdvertiserSubmit"
          >更新信息</el-button>
          <el-button
            icon="refresh"
            @click="onAdvertiserReset"
          >重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item
          label="广告名称"
          prop="name"
        >
          <el-input
            v-model="searchInfo.name"
            placeholder="广告名称"
          />
        </el-form-item>
        <el-form-item
          label="广告类型"
          prop="costCategory"
        >
          <el-select
            v-model="searchInfo.adCategory"
            placeholder="请选择"
            clearable
            style="width: 120px;"
          >
            <el-option
              key="infoStream"
              label="信息流广告"
              :value="1"
            />
            <el-option
              key="keywords"
              label="关键词广告"
              :value="2"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="计费类型"
          prop="costCategory"
        >
          <el-select
            v-model="searchInfo.costCategory"
            placeholder="请选择"
            clearable
            style="width: 120px;"
          >
            <el-option
              key="CPM"
              label="CPM"
              :value="1"
            />
            <el-option
              key="CPC"
              label="CPC"
              :value="2"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="状态"
          prop="status"
        >
          <el-select
            v-model="searchInfo.status"
            placeholder="请选择"
            clearable
            style="width: 120px;"
          >
            <el-option
              key="un"
              label="计划投放"
              :value="1"
            />
            <el-option
              key="in"
              label="投放中"
              :value="2"
            />
            <el-option
              key="stop"
              label="暂停投放"
              :value="3"
            />
            <el-option
              key="finish"
              label="完成投放"
              :value="4"
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
        <el-popover
          width="200px"
        >
          <template #reference>
            <el-button
              type="primary"
              plain
            >未使用积分: {{ advertiserInfo.points }} 点</el-button>
          </template>
          <template #default>
            <div>
              <p>
                购买更多广告积分，请添加下方客服微信。
              </p>
              <el-image
                src="https://download.tooc.xlj0.com/Tooc/14/14/wx.png"
                fit="fit"
                style="width: 100px; margin-top: 10px"
                :preview-src-list="['https://download.tooc.xlj0.com/Tooc/14/14/wx.png']"
              />
            </div>
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
        @sort-change="sortChange"
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          label="创建日期"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="广告名称"
          prop="name"
          width="120"
          :show-overflow-tooltip="true"
        />
        <el-table-column
          align="left"
          label="广告类型"
          prop="adCategory"
          width="120"
        >
          <template #default="scope">{{ formatAdCategoty(scope.row.adCategory) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="计费类型"
          prop="costCategory"
          width="160"
        >
          <template #default="scope">{{ formatCostCategoty(scope.row.costCategory) }}</template>
        </el-table-column>
        <el-table-column
          sortable
          align="left"
          label="已展示量"
          prop="costImpressions"
          width="120"
        >
          <template #default="scope">
            {{ scope.row.costImpressions+' 次' }}
          </template>
        </el-table-column>
        <el-table-column
          sortable
          align="left"
          label="已点击数"
          prop="costClicks"
          width="120"
        >
          <template #default="scope">
            {{ scope.row.costClicks+' 次' }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="购买量"
          prop="buyAmount"
          width="120"
        >
          <template #default="scope">
            {{ (scope.row.costCategory===1?'展示 ':'点击 ')+scope.row.buyAmount+' 次' }}
          </template>
        </el-table-column>
        <el-table-column
          sortable
          align="left"
          label="花费积分"
          prop="cost"
          width="120"
        >
          <template #default="scope">
            {{ scope.row.cost+' 点' }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="投放日期"
          width="200"
        >
          <template #default="scope">
            {{ new Date(scope.row.startDate).Format('yyyy-MM-dd') }}&nbsp;—&nbsp;{{ new Date(scope.row.endDate).Format('yyyy-MM-dd') }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="投放时段"
          prop="startHour"
          width="120"
        >
          <template #default="scope">
            {{ scope.row.startHour+':00' }}&nbsp;—&nbsp;{{ scope.row.endHour+':00' }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="状态"
          prop="status"
          width="120"
        >
          <template #default="scope">
            {{ formatStatus(scope.row.status) }}
          </template>
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
              @click="updateAdFunc(scope.row)"
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
      style="width: 50%;"
      align-center
    >
      <el-scrollbar>
        <el-form
          ref="elFormRef"
          :model="formData"
          label-position="right"
          :rules="rule"
          inline="true"
          label-width="100px"
        >
          <el-text
            tag="b"
            size="large"
            style="color: black;"
          >
            广告信息
          </el-text>
          <el-divider style="margin-top: 10px;margin-bottom: 10px;" />
          <el-form-item
            label="名称:"
            prop="name"
            style="width: 40%"
          >
            <el-input
              v-model="formData.name"
              :clearable="true"
              placeholder="请输入名称"
            />
          </el-form-item>
          <br>
          <el-form-item
            label="文案:"
            prop="text"
            style="width: 90%;"
          >
            <el-input
              v-model="formData.text"
              :clearable="true"
              placeholder="请输入文案"
            />
          </el-form-item>
          <el-form-item
            label="添加素材:"
            style="width: 30%"
          >
            <el-select
              v-model="materialType"
              placeholder="请选择"
            >
              <el-option
                key="none"
                label="无"
                :value="1"
              />
              <el-option
                key="img"
                label="图片"
                :value="2"
              />
              <el-option
                key="video"
                label="视频"
                :value="3"
              />
            </el-select>
          </el-form-item>
          <el-form-item
            label="设置跳转链接:"
          >
            <el-switch
              v-model="setLink"
              inline-prompt
              active-text="是"
              inactive-text="否"
            />
          </el-form-item>
          <el-form-item
            v-show="setLink"
            prop="link"
            style="width: 30%;"
          >
            <el-input
              v-model="formData.link"
              :clearable="true"
              placeholder="请输入跳转链接"
            />
          </el-form-item>
          <br>
          <el-form-item
            v-show="materialType===2"
            label="图片:"
            prop="img"
          >
            <UploadImg
              v-model="formData.img"
            />
          </el-form-item>
          <el-form-item
            v-show="materialType===3"
            label="视频:"
            prop="video"
          >
            <UploadVideo
              v-model="formData.video"
            />
          </el-form-item>
          <el-divider style="margin-top: 10px;margin-bottom: 10px;" />
          <el-text
            tag="b"
            size="large"
            style="color: black;"
          >
            投放策略
          </el-text>
          <el-divider style="margin-top: 10px;margin-bottom: 10px;" />
          <el-form-item
            label="广告类型:"
            prop="costCategory"
            style="width: 35%;"
          >
            <el-select
              v-model.number="formData.adCategory"
              placeholder="请选择"
              :change="handleAdCategoryChange()"
              :disabled="type==='update'"
            >
              <el-option
                key="infoStream"
                label="信息流广告"
                :value="1"
              />
              <el-option
                key="keywords"
                label="关键词广告"
                :value="2"
              />
            </el-select>
          </el-form-item>
          <br>
          <el-form-item
            label="投放日期:"
            prop="startDate"
          >
            <el-date-picker
              v-model="formData.startDate"
              type="date"
              placeholder="选择日期"
              style="width: 140px;"
            />
            &nbsp;—&nbsp;
            <el-date-picker
              v-model="formData.endDate"
              type="date"
              placeholder="选择日期"
              :disabled-date="disabledDate"
              style="width: 140px;"
            />
          </el-form-item>
          <el-form-item
            label="投放时段:"
            prop="startHour"
          >
            <el-time-select
              v-model.number="formData.startHour"
              style="width: 80px"
              :max-time="endTime"
              placeholder=""
              start="00:00"
              step="01:00"
              end="24:00"
              :clearable="false"
            />
            &nbsp;—&nbsp;
            <el-time-select
              v-model.number="formData.endHour"
              style="width: 80px"
              :min-time="formData.startHour+':00'"
              placeholder=""
              start="00:00"
              step="01:00"
              end="24:00"
              :clearable="false"
            />
          </el-form-item>
          <el-form-item
            v-show="formData.adCategory===1"
            label="目标标签:"
            prop="target"
            style="width: 90%;"
          >
            <el-cascader
              v-model="formData.target"
              :options="targetOptions"
              :props="targetProps"
              collapse-tags
              collapse-tags-tooltip
              max-collapse-tags="10"
              :show-all-levels="false"
              clearable
              style="width: 100%;"
            />
          </el-form-item>
          <el-form-item
            v-show="formData.adCategory===2"
            label="关键词:"
            prop="keywords"
            style="width: 90%;"
          >
            <el-select
              v-model="keywords"
              multiple
              filterable
              remote
              reserve-keyword
              placeholder="请输入关键词"
              :remote-method="remoteMethod"
              :loading="loading"
              :disabled="type==='update'"
            >
              <el-option
                :key="option"
                :label="option"
                :value="option"
              >
                <div v-show="option">
                  <span style="float: left;">{{ option }}</span>
                  <span style="float: right;">
                    <el-tag type="primary">{{ level() }}</el-tag>
                    &nbsp;
                    <el-tag type="primary">{{ "今日搜索次数: "+keywordStats.today }}</el-tag>
                    &nbsp;
                    <el-tag type="primary">{{ "本周搜索次数: " +keywordStats.week }}</el-tag>
                  </span>
                </div>
              </el-option>
            </el-select>
          </el-form-item>

          <el-divider style="margin-top: 10px;margin-bottom: 10px;" />
          <el-text
            tag="b"
            size="large"
            style="color: black;"
          >
            计费方式
          </el-text>
          <el-divider style="margin-top: 10px;margin-bottom: 10px;" />
          <el-form-item
            label="计费类型:"
            prop="costCategory"
            style="width: 35%;"
          >
            <el-select
              v-model.number="formData.costCategory"
              placeholder="请选择"
              :disabled="type==='update'"
            >
              <el-option
                v-if="formData.adCategory===1"
                key="CPM"
                label="CPM (千次展示)"
                :value="1"
              />
              <el-option
                key="CPC"
                label="CPC (单次点击)"
                :value="2"
              />
            </el-select>
          </el-form-item>
          <el-form-item
            label="购买量（次数）:"
            prop="buyAmount"
            style="width: 35%;"
            label-width="140px"
          >
            <el-input-number
              v-if="formData.costCategory===1"
              v-model="formData.buyAmount"
              controls-position="right"
              :min="1000"
              :step="1000"
              step-strictly
              :disabled="type==='update'"
            />
            <el-input-number
              v-else
              v-model="formData.buyAmount"
              controls-position="right"
              :min="100"
              :step="100"
              step-strictly
              :disabled="type==='update'"
            />
          </el-form-item>
          <br>
          <br>
          <el-form-item
            v-if="type==='create'"
            label="所需积分:"
            prop="cost"
            style="float: right;width: 30%;"
          >
            <el-text
              size="large"
              type="primary"
            >
              {{ calCost }}
            </el-text>
            &nbsp;&nbsp;&nbsp;&nbsp;点
          </el-form-item>
          <el-form-item
            v-else
            label="花费积分:"
            prop="cost"
            style="float: right;width: 30%;"
          >
            <el-text
              size="large"
              type="primary"
            >
              {{ formData.cost }}
            </el-text>
            &nbsp;&nbsp;&nbsp;&nbsp;点
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
  createAd,
  deleteAd,
  deleteAdByIds,
  updateAd,
  findAd,
  getAdList,
  getKeywordStats
} from '@/api/ad'
import {
  createAdvertiser,
  updateAdvertiser,
  findAdvertiser,
} from '@/api/advertiser'
// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, computed } from 'vue'
import UploadImg from '@/components/uploadImg/uploadImg.vue'
import UploadVideo from '@/components/uploadVideo/uploadVideo.vue'

defineOptions({
  name: 'Ad'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  advertiserId: 0,
  name: '',
  text: '',
  img: '',
  video: '',
  link: '',
  adCategory: 1,
  startDate: new Date(),
  endDate: new Date(),
  startHour: 0,
  endHour: 0,
  target: [],
  keywords: '',
  costCategory: 1,
  buyAmount: 0,
  costImpressions: 0,
  costClicks: 0,
  cost: 0,
  status: 0,
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
  text: [{
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
  link: [
    {
      whitespace: true,
      message: '不能只输入空格',
      trigger: ['input', 'blur'],
    }
  ],
  startDate: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  startHour: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  costCategory: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  buyAmount: [{
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
  const table = await getAdList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteAdFunc(row)
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
  const res = await deleteAdByIds({ IDs })
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
const updateAdFunc = async(row) => {
  const res = await findAd({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.read
    if (formData.value.img !== '') {
      materialType.value = 2
    } else if (formData.value.video !== '') {
      materialType.value = 3
    } else {
      materialType.value = 1
    }
    if (formData.value.link !== '') {
      setLink.value = true
    } else {
      setLink.value = false
    }
    keywords.value = JSON.parse(formData.value.keywords)
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteAdFunc = async(row) => {
  const res = await deleteAd({ ID: row.ID })
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
  if (advertiserInfo.value.ID === 0) {
    ElMessage({
      'type': 'warning',
      'message': '请先填写广告主信息',
    })
    return
  }
  type.value = 'create'
  materialType.value = 1
  setLink.value = false
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    advertiserId: 0,
    name: '',
    text: '',
    img: '',
    video: '',
    link: '',
    adCategory: 1,
    startDate: new Date(),
    endDate: new Date(),
    startHour: 0,
    endHour: 0,
    target: [],
    keywords: '',
    costCategory: 1,
    buyAmount: 0,
    costImpressions: 0,
    costClicks: 0,
    cost: 0,
    status: 0,
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    formData.value.advertiserId = advertiserInfo.value.ID
    formData.value.keywords = JSON.stringify(keywords.value)
    switch (type.value) {
      case 'create':
        formData.value.cost = calCost.value
        res = await createAd(formData.value)
        break
      case 'update':
        res = await updateAd(formData.value)
        break
      default:
        res = await createAd(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
      getAdvertiserInfo()
    }
  })
}

const advertiserType = ref('create')

const advertiserInfo = ref({
  ID: 0,
  name: '',
  logo: '',
  introduction: '',
  subjectInfo: '',
  license: '',
  location: '',
  points: 0,
})

// 验证规则
const advertiserRule = reactive({
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
  logo: [{
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
  subjectInfo: [{
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
  license: [{
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
  location: [{
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

const elAdvertiserFormRef = ref()

const onAdvertiserReset = () => {
  getAdvertiserInfo()
}

const onAdvertiserSubmit = () => {
  elAdvertiserFormRef.value?.validate(async(valid) => {
    if (!valid) {
      ElMessage({
        type: 'warning',
        message: '广告主信息不能为空'
      })
      return
    }
    let res
    switch (advertiserType.value) {
      case 'create':
        res = await createAdvertiser(advertiserInfo.value)
        break
      case 'update':
        res = await updateAdvertiser(advertiserInfo.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '更新成功'
      })
      getAdvertiserInfo()
    }
  })
}

const getAdvertiserInfo = async() => {
  const info = await findAdvertiser()
  if (info.code === 0) {
    advertiserInfo.value = info.data.readvertiser
  }
  if (info.data.readvertiser.ID === 0) {
    advertiserType.value = 'create'
  } else {
    advertiserType.value = 'update'
  }
}

getAdvertiserInfo()

const materialType = ref(1)

const setLink = ref(false)

const disabledDate = (date) => {
  return date < formData.value.startDate
}

const targetProps = { multiple: true }

const targetOptions = [
  {
    value: 'sex',
    label: '性别',
    children: [
      { value: '男', label: '男' },
      { value: '女', label: '女' },
    ],
  },
  {
    value: 'campus',
    label: '校区',
    children: [
      { value: '朝晖校区', label: '朝晖校区' },
      { value: '屏峰校区', label: '屏峰校区' },
      { value: '莫干山校区', label: '莫干山校区' },
    ],
  },
  {
    value: 'college',
    label: '学院',
    children: [
      { value: '化学工程学院', label: '化学工程学院' },
      { value: '生物工程学院', label: '生物工程学院' },
      { value: '环境学院', label: '环境学院' },
      { value: '材料科学与工程学院', label: '材料科学与工程学院' },
      { value: '食品科学与工程学院', label: '食品科学与工程学院' },
      { value: '机械工程学院', label: '机械工程学院' },
      { value: '信息工程学院', label: '信息工程学院' },
      { value: '计算机科学与技术学院、软件学院', label: '计算机科学与技术学院、软件学院' },
      { value: '土木工程学院', label: '土木工程学院' },
      { value: '理学院', label: '理学院' },
      { value: '管理学院', label: '管理学院' },
      { value: '经济学院', label: '经济学院' },
      { value: '教育科学与技术学院', label: '教育科学与技术学院' },
      { value: '外国语学院', label: '外国语学院' },
      { value: '人文学院', label: '人文学院' },
      { value: '设计与建筑学院', label: '设计与建筑学院' },
      { value: '法学院', label: '法学院' },
      { value: '公共管理学院', label: '公共管理学院' },
      { value: '马克思主义学院', label: '马克思主义学院' },
      { value: '国际学院', label: '国际学院' },
      { value: '健行学院', label: '健行学院' },
    ]
  },
  {
    value: 'hobby',
    label: '爱好',
    children: [
      { value: '音乐', label: '音乐' },
      { value: '运动', label: '运动' },
      { value: '艺术', label: '艺术' },
      { value: '阅读', label: '阅读' },
      { value: '旅行', label: '旅行' },
      { value: '影视', label: '影视' },
      { value: '游戏', label: '游戏' },
      { value: '手工艺', label: '手工艺' },
      { value: '烹饪', label: '烹饪' },
      { value: '摄影', label: '摄影' },
      { value: '写作', label: '写作' },
      { value: '户外', label: '户外' },
      { value: '跳舞', label: '跳舞' },
      { value: '棋牌', label: '棋牌' },
      { value: 'ACG', label: 'ACG' },
      { value: '宠物', label: '宠物' },
    ]
  },
  {
    value: 'originPlace',
    label: '生源地',
    children: [
      { value: '北京市', label: '北京市' },
      { value: '天津市', label: '天津市' },
      { value: '河北省', label: '河北省' },
      { value: '山西省', label: '山西省' },
      { value: '内蒙古自治区', label: '内蒙古自治区' },
      { value: '辽宁省', label: '辽宁省' },
      { value: '吉林省', label: '吉林省' },
      { value: '黑龙江省', label: '黑龙江省' },
      { value: '上海市', label: '上海市' },
      { value: '江苏省', label: '江苏省' },
      { value: '浙江省', label: '浙江省' },
      { value: '安徽省', label: '安徽省' },
      { value: '福建省', label: '福建省' },
      { value: '江西省', label: '江西省' },
      { value: '山东省', label: '山东省' },
      { value: '河南省', label: '河南省' },
      { value: '湖北省', label: '湖北省' },
      { value: '湖南省', label: '湖南省' },
      { value: '广东省', label: '广东省' },
      { value: '广西壮族自治区', label: '广西壮族自治区' },
      { value: '海南省', label: '海南省' },
      { value: '重庆市', label: '重庆市' },
      { value: '四川省', label: '四川省' },
      { value: '贵州省', label: '贵州省' },
      { value: '云南省', label: '云南省' },
      { value: '西藏自治区', label: '西藏自治区' },
      { value: '陕西省', label: '陕西省' },
      { value: '甘肃省', label: '甘肃省' },
      { value: '青海省', label: '青海省' },
      { value: '宁夏回族自治区', label: '宁夏回族自治区' },
      { value: '新疆维吾尔自治区', label: '新疆维吾尔自治区' },
      { value: '台湾省', label: '台湾省' },
      { value: '香港特别行政区', label: '香港特别行政区' },
      { value: '澳门特别行政区', label: '澳门特别行政区' },
    ]
  },
]

const formatAdCategoty = (category) => {
  switch (category) {
    case 1:
      return '信息流广告'
    case 2:
      return '关键词广告'
  }
}

const formatCostCategoty = (category) => {
  switch (category) {
    case 1:
      return 'CPM (千次展示)'
    case 2:
      return 'CPC (单次点击)'
  }
}

const formatStatus = (status) => {
  switch (status) {
    case 1:
      return '计划投放'
    case 2:
      return '投放中'
    case 3:
      return '暂停投放'
    case 4:
      return '完成投放'
  }
}

const loading = ref(false)
const option = ref('')
const keywordStats = ref({
  today: 0,
  week: 0
})
const remoteMethod = async(query) => {
  if (query) {
    loading.value = true
    const result = await getKeywordStats({ keyword: query })
    if (result.code === 0) {
      keywordStats.value = result.data
      option.value = query
      handleChange()
    }
    loading.value = false
  } else {
    option.value = ''
  }
}
const handleChange = () => {
  let level
  if (keywordStats.value.today >= 500 || keywordStats.value.week >= 5000) {
    level = 1
  } else if (keywordStats.value.today >= 200 || keywordStats.value.week >= 2000) {
    level = 2
  } else if (keywordStats.value.today >= 50 || keywordStats.value.week >= 500) {
    level = 3
  } else if (keywordStats.value.today >= 10 || keywordStats.value.week >= 100) {
    level = 4
  } else {
    level = 5
  }
  choice.value = {
    option: option.value,
    level: level
  }
  choices.value.push(choice.value)
}
const level = () => {
  let level
  if (keywordStats.value.today >= 500 || keywordStats.value.week >= 5000) {
    level = '一级 10点/次'
  } else if (keywordStats.value.today >= 200 || keywordStats.value.week >= 2000) {
    level = '二级 8点/次'
  } else if (keywordStats.value.today >= 50 || keywordStats.value.week >= 500) {
    level = '三级 5点/次'
  } else if (keywordStats.value.today >= 10 || keywordStats.value.week >= 100) {
    level = '四级 2点/次'
  } else {
    level = '五级 1点/次'
  }
  return level
}
const choice = ref({
  option: '',
  level: 0
})
const choices = ref([])
const calCost = computed(() => {
  let per = 0
  if (formData.value.adCategory === 2) {
    per = 30
    for (let i = 0; i < choices.value.length; i++) {
      if (formData.value.keywords.includes(choices.value[i].option)) {
        switch (choices.value[i].level) {
          case 1:
            per += 10
            break
          case 2:
            per += 8
            break
          case 3:
            per += 5
            break
          case 4:
            per += 2
            break
          default:
            per += 1
            break
        }
      }
    }
  } else {
    if (formData.value.costCategory === 1) {
      per = 1.5
    } else {
      per = 50
    }
  }
  return per * formData.value.buyAmount
})
function handleAdCategoryChange() {
  formData.value.costCategory = formData.value.adCategory
}

const keywords = ref([])

</script>

<style>

</style>
