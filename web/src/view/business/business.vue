<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elBusinessFormRef"
        :model="businessInfo"
        class="demo-form-inline"
        :rules="businessRule"
        :inline="true"
        label-width="80px"
      >
        <el-form-item
          label="封面图"
          prop="pic"
        >
          <UploadImg
            v-model="businessInfo.pic"
          />
        </el-form-item>
        <p />
        <el-form-item
          label="名称"
          prop="name"
          style="width: 40%;"
        >
          <el-input
            v-model="businessInfo.name"
            placeholder="店铺商家名称"
            clearable
          />
        </el-form-item>
        <p />
        <el-form-item
          label="地址"
          prop="address"
          style="width: 40%;"
        >
          <el-input
            v-model="businessInfo.address"
            placeholder="地址"
            clearable
          />
        </el-form-item>
        <el-form-item
          label="导航地址"
          prop="loc"
          style="width: 20%;"
        >
          <el-input
            v-model="businessInfo.loc"
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
          style="width: 40%;"
        >
          <el-input
            v-model="businessInfo.businessHours"
            placeholder="营业时间"
            clearable
          />
        </el-form-item>
        <el-form-item
          label="联系方式"
          prop="contact"
          style="width: 40%;"
        >
          <el-input
            v-model="businessInfo.contact"
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
            v-model="businessInfo.introduction"
            placeholder="介绍"
            clearable
            :autosize="{minRows:2,maxRows:8}"
            type="textarea"
          />
        </el-form-item>
        <p />
        <el-form-item>
          <el-button
            icon="mouse"
            type="primary"
            @click="onBusinessSubmit"
          >更新信息</el-button>
          <el-button
            icon="refresh"
            @click="onBusinessReset"
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
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item
          label="优惠券名称"
          prop="name"
        >
          <el-input
            v-model="searchInfo.name"
            placeholder="优惠券名称"
            style="width: 160px;"
          />
        </el-form-item>
        <el-form-item
          label="优惠券类型"
          prop="couponCategory"
        >
          <el-select
            v-model.number="searchInfo.couponCategory"
            placeholder="请选择"
            clearable
            style="width: 120px;"
          >
            <el-option
              key="fullMinus"
              label="满减券"
              :value="1"
            />
            <el-option
              key="cash"
              label="现金券"
              :value="2"
            />
            <el-option
              key="discount"
              label="折扣券"
              :value="3"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="兑换所需积分"
          prop="redeemPoints"
        >
          <el-input
            v-model.number="searchInfo.startRedeemPoints"
            placeholder="最小值"
            style="width: 120px ;"
          />
          —
          <el-input
            v-model.number="searchInfo.endRedeemPoints"
            placeholder="最大值"
            style="width: 120px;"
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
          label="优惠券名称"
          prop="name"
          width="160"
        />
        <el-table-column
          sortable
          align="left"
          label="总数量"
          prop="totalCount"
          width="120"
        />
        <el-table-column
          sortable
          align="left"
          label="剩余数量"
          prop="remainCount"
          width="120"
        />
        <el-table-column
          sortable
          align="left"
          label="已使用数量"
          prop="usedCount"
          width="120"
        />
        <el-table-column
          align="left"
          label="优惠券类型"
          prop="couponCategory"
          width="120"
        >
          <template #default="scope">
            {{ formatCouponCategory(scope.row.couponCategory) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="使用说明"
          prop="usageInstructions"
          width="120"
          show-overflow-tooltip="true"
        />
        <el-table-column
          align="left"
          label="兑换上限"
          prop="redeemCount"
          width="120"
        />
        <el-table-column
          align="left"
          label="兑换开始时间"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.redeemStartTime) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="兑换结束时间"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.redeemEndTime) }}</template>
        </el-table-column>
        <el-table-column
          sortable
          align="left"
          label="兑换所需积分"
          prop="redeemPoints"
          width="140"
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
              @click="updateCouponFunc(scope.row)"
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
        label-width="120px"
      >
        <el-form-item
          label="优惠券名称:"
          prop="name"
          style="width: 40%;"
        >
          <el-input
            v-model="formData.name"
            :clearable="true"
            placeholder="请输入优惠券名称"
          />
        </el-form-item>
        <el-form-item
          label="总数量:"
          prop="totalCount"
          style="width: 45%;"
        >
          <el-input
            v-model.number="formData.totalCount"
            placeholder="请输入总数量"
          >
            <template #append>0表示无数量限制</template>
          </el-input>
        </el-form-item>
        <div v-show="type==='update'">
          <el-form-item
            label="剩余数量:"
            prop="remainCount"
            style="width: 40%;"
          >
            <el-input
              v-model.number="formData.remainCount"
              disabled
            />
          </el-form-item>
          <el-form-item
            label="已使用数量:"
            prop="usedCount"
            style="width: 40%;"
          >
            <el-input
              v-model.number="formData.usedCount"
              disabled
            />
          </el-form-item>
        </div>
        <el-form-item
          label="优惠券类型:"
          prop="couponCategory"
          style="width: 40%;"
        >
          <el-select
            v-model.number="formData.couponCategory"
            placeholder="请选择"
          >
            <el-option
              key="fullMinus"
              label="满减券"
              :value="1"
            />
            <el-option
              key="cash"
              label="现金券"
              :value="2"
            />
            <el-option
              key="discount"
              label="折扣券"
              :value="3"
            />
          </el-select>
        </el-form-item>
        <span v-show="formData.couponCategory===1">
          <el-form-item
            label="满"
            prop="fullMoney"
          >
            <el-input
              v-model.number="formData.fullMoney"
              placeholder="满"
              style="width: 110px;"
            >
              <template #append>元</template>
            </el-input>
          </el-form-item>
          <el-form-item
            label="减"
            prop="minusMoney"
            label-width="0px"
          >
            <el-input
              v-model.number="formData.minusMoney"
              placeholder="减"
              style="width: 110px;"
            >
              <template #append>元</template>
            </el-input>
          </el-form-item>
        </span>
        <el-form-item
          v-show="formData.couponCategory===2"
          label="现金金额:"
          prop="cash"
          style="width: 40%;"
        >
          <el-input
            v-model.number="formData.cash"
            :clearable="true"
            placeholder="请输入现金金额"
          />
        </el-form-item>
        <el-form-item
          v-show="formData.couponCategory===3"
          label="折扣(百分制):"
          prop="discount"
          style="width: 50%;"
        >
          <el-input
            v-model.number="formData.discount"
            :clearable="true"
            placeholder="请输入折扣"
          >
            <template #append>如: 八折=>80</template>
          </el-input>
        </el-form-item>
        <el-form-item
          label="兑换上限:"
          prop="redeemCount"
          style="width: 40%;"
        >
          <el-input
            v-model.number="formData.redeemCount"
            :clearable="true"
            placeholder="请输入每个用户可兑换数量"
          >
            <template #append>张</template>
          </el-input>
        </el-form-item>
        <el-form-item
          label="兑换所需积分:"
          prop="redeemPoints"
          style="width: 40%;"
        >
          <el-input
            v-model.number="formData.redeemPoints"
            :clearable="true"
            placeholder="请输入兑换所需积分"
          />
        </el-form-item>
        <el-form-item
          label="兑换开始时间:"
          prop="redeemStartTime"
          style="width: 40%;"
        >
          <el-date-picker
            v-model="formData.redeemStartTime"
            type="date"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item
          label="兑换结束时间:"
          prop="redeemEndTime"
          style="width: 40%;"
        >
          <el-date-picker
            v-model="formData.redeemEndTime"
            type="date"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item
          label="有效期类型:"
          prop="expCategory"
          style="width: 40%;"
        >
          <el-select
            v-model.number="formData.expCategory"
            placeholder="请选择"
          >
            <el-option
              key="fixed"
              label="固定日期"
              :value="1"
            />
            <el-option
              key="valid"
              label="有效天数"
              :value="2"
            />
          </el-select>
        </el-form-item>
        <div v-show="formData.expCategory===1">
          <el-form-item
            label="固定开始时间:"
            prop="fixedStartTime"
            style="width: 40%;"
          >
            <el-date-picker
              v-model="formData.fixedStartTime"
              type="date"
              placeholder="选择日期"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item
            label="固定结束时间:"
            prop="fixedEndTime"
            style="width: 40%;"
          >
            <el-date-picker
              v-model="formData.fixedEndTime"
              type="date"
              placeholder="选择日期"
              :clearable="true"
            />
          </el-form-item>
        </div>
        <el-form-item
          v-show="formData.expCategory===2"
          label="有效天数:"
          prop="validDay"
          style="width: 40%;"
        >
          <el-input
            v-model.number="formData.validDay"
            :clearable="true"
            placeholder="请输入有效天数"
          ><template #append>天</template>
          </el-input>
        </el-form-item>
        <el-form-item
          label="使用说明:"
          prop="usageInstructions"
          style="width: 80%;"
        >
          <el-input
            v-model="formData.usageInstructions"
            :clearable="true"
            placeholder="请输入使用说明"
            :autosize="{minRows:2,maxRows:8}"
            type="textarea"
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
  createBusiness,
  updateBusiness,
  findBusiness,
} from '@/api/business'
import {
  createCoupon,
  deleteCoupon,
  deleteCouponByIds,
  updateCoupon,
  findCoupon,
  getCouponList
} from '@/api/coupon'

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
  businessId: 0,
  name: '',
  totalCount: 0,
  remainCount: 0,
  usedCount: 0,
  couponCategory: 1,
  fullMoney: 0,
  minusMoney: 0,
  cash: 0,
  discount: 0,
  usageInstructions: '',
  redeemCount: 0,
  redeemStartTime: new Date(),
  redeemEndTime: new Date(),
  expCategory: 1,
  fixedStartTime: new Date(),
  fixedEndTime: new Date(),
  validDay: 0,
  redeemPoints: 0,
})

// 验证规则
const rule = reactive({
  businessId: [{
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
  totalCount: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  remainCount: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  usedCount: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  couponCategory: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  fullMoney: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  minusMoney: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  cash: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  discount: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  usageInstructions: [{
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
  redeemCount: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  redeemStartTime: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  redeemEndTime: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  expCategory: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  fixedStartTime: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  fixedEndTime: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  validDay: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  redeemPoints: [{
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
  const table = await getCouponList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteCouponFunc(row)
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
  const res = await deleteCouponByIds({ IDs })
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
const updateCouponFunc = async(row) => {
  const res = await findCoupon({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.recoupon
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteCouponFunc = async(row) => {
  const res = await deleteCoupon({ ID: row.ID })
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
  if (businessInfo.value.ID === 0) {
    ElMessage({
      'type': 'warning',
      'message': '请先填写店铺商家信息',
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
    businessId: 0,
    name: '',
    totalCount: 0,
    remainCount: 0,
    usedCount: 0,
    couponCategory: 1,
    fullMoney: 0,
    minusMoney: 0,
    cash: 0,
    discount: 0,
    usageInstructions: '',
    redeemCount: 0,
    redeemStartTime: new Date(),
    redeemEndTime: new Date(),
    expCategory: 1,
    fixedStartTime: new Date(),
    fixedEndTime: new Date(),
    validDay: 0,
    redeemPoints: 0,
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    formData.value.businessId = businessInfo.value.ID
    switch (type.value) {
      case 'create':
        res = await createCoupon(formData.value)
        break
      case 'update':
        res = await updateCoupon(formData.value)
        break
      default:
        res = await createCoupon(formData.value)
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

const businessType = ref('create')

const businessInfo = ref({
  ID: 0,
  name: '',
  pic: '',
  address: '',
  loc: '',
  businessHours: '',
  contact: '',
  introduction: '',
})

// 验证规则
const businessRule = reactive({
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

const elBusinessFormRef = ref()

const onBusinessReset = () => {
  getBusinessInfo()
}

const onBusinessSubmit = () => {
  elBusinessFormRef.value?.validate(async(valid) => {
    if (!valid) {
      ElMessage({
        type: 'warning',
        message: '店铺商家信息不能为空'
      })
      return
    }
    let res
    switch (businessType.value) {
      case 'create':
        res = await createBusiness(businessInfo.value)
        break
      case 'update':
        res = await updateBusiness(businessInfo.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      getBusinessInfo()
    }
  })
}

const getBusinessInfo = async() => {
  const info = await findBusiness()
  if (info.code === 0) {
    businessInfo.value = info.data.rebusiness
  }
  if (info.data.rebusiness.ID === 0) {
    businessType.value = 'create'
    amapLoc.value = {
      longitude: 0,
      latitude: 0,
    }
  } else {
    businessType.value = 'update'
    const locParts = businessInfo.value.loc.split(',')
    amapLoc.value = {
      longitude: locParts[0],
      latitude: locParts[1],
    }
  }
}

getBusinessInfo()

function formatCouponCategory(category) {
  switch (category) {
    case 1:
      return '满减券'
    case 2:
      return '现金券'
    case 3:
      return '折扣券'
  }
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
  businessInfo.value.loc = amapLoc.value.longitude + ',' + amapLoc.value.latitude
  amapDialog.value = false
}

</script>

<style>

</style>
