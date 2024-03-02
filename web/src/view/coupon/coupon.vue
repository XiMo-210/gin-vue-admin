<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item label="优惠券名称" prop="name">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="优惠券类型" prop="couponCategory">
            
             <el-input v-model.number="searchInfo.couponCategory" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="兑换所需积分" prop="redeemPoints">
            
            <el-input v-model.number="searchInfo.startRedeemPoints" placeholder="最小值" />
            —
            <el-input v-model.number="searchInfo.endRedeemPoints" placeholder="最大值" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="所属店铺商家" prop="businessId" width="120" />
        <el-table-column align="left" label="优惠券名称" prop="name" width="120" />
        <el-table-column sortable align="left" label="总数量" prop="totalCount" width="120" />
        <el-table-column sortable align="left" label="剩余数量" prop="remainCount" width="120" />
        <el-table-column sortable align="left" label="已使用数量" prop="usedCount" width="120" />
        <el-table-column align="left" label="优惠券类型" prop="couponCategory" width="120" />
        <el-table-column align="left" label="满" prop="fullMoney" width="120" />
        <el-table-column align="left" label="减" prop="minusMoney" width="120" />
        <el-table-column align="left" label="现金金额" prop="cash" width="120" />
        <el-table-column align="left" label="折扣" prop="discount" width="120" />
        <el-table-column align="left" label="使用说明" prop="usageInstructions" width="120" />
        <el-table-column align="left" label="每个用户可兑换数量" prop="redeemCount" width="120" />
         <el-table-column align="left" label="兑换开始时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.redeemStartTime) }}</template>
         </el-table-column>
         <el-table-column align="left" label="兑换结束时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.redeemEndTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="有效期类型" prop="expCategory" width="120" />
         <el-table-column align="left" label="固定开始时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.fixedStartTime) }}</template>
         </el-table-column>
         <el-table-column align="left" label="固定结束时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.fixedEndTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="有效天数" prop="validDay" width="120" />
        <el-table-column sortable align="left" label="兑换所需积分" prop="redeemPoints" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateCouponFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-scrollbar height="500px">
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="所属店铺商家:"  prop="businessId" >
              <el-input v-model.number="formData.businessId" :clearable="true" placeholder="请输入所属店铺商家" />
            </el-form-item>
            <el-form-item label="优惠券名称:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入优惠券名称" />
            </el-form-item>
            <el-form-item label="总数量:"  prop="totalCount" >
              <el-input v-model.number="formData.totalCount" :clearable="true" placeholder="请输入总数量" />
            </el-form-item>
            <el-form-item label="剩余数量:"  prop="remainCount" >
              <el-input v-model.number="formData.remainCount" :clearable="true" placeholder="请输入剩余数量" />
            </el-form-item>
            <el-form-item label="已使用数量:"  prop="usedCount" >
              <el-input v-model.number="formData.usedCount" :clearable="true" placeholder="请输入已使用数量" />
            </el-form-item>
            <el-form-item label="优惠券类型:"  prop="couponCategory" >
              <el-input v-model.number="formData.couponCategory" :clearable="true" placeholder="请输入优惠券类型" />
            </el-form-item>
            <el-form-item label="满:"  prop="fullMoney" >
              <el-input v-model.number="formData.fullMoney" :clearable="true" placeholder="请输入满" />
            </el-form-item>
            <el-form-item label="减:"  prop="minusMoney" >
              <el-input v-model.number="formData.minusMoney" :clearable="true" placeholder="请输入减" />
            </el-form-item>
            <el-form-item label="现金金额:"  prop="cash" >
              <el-input v-model.number="formData.cash" :clearable="true" placeholder="请输入现金金额" />
            </el-form-item>
            <el-form-item label="折扣:"  prop="discount" >
              <el-input v-model.number="formData.discount" :clearable="true" placeholder="请输入折扣" />
            </el-form-item>
            <el-form-item label="使用说明:"  prop="usageInstructions" >
              <el-input v-model="formData.usageInstructions" :clearable="true"  placeholder="请输入使用说明" />
            </el-form-item>
            <el-form-item label="每个用户可兑换数量:"  prop="redeemCount" >
              <el-input v-model.number="formData.redeemCount" :clearable="true" placeholder="请输入每个用户可兑换数量" />
            </el-form-item>
            <el-form-item label="兑换开始时间:"  prop="redeemStartTime" >
              <el-date-picker v-model="formData.redeemStartTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="兑换结束时间:"  prop="redeemEndTime" >
              <el-date-picker v-model="formData.redeemEndTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="有效期类型:"  prop="expCategory" >
              <el-input v-model.number="formData.expCategory" :clearable="true" placeholder="请输入有效期类型" />
            </el-form-item>
            <el-form-item label="固定开始时间:"  prop="fixedStartTime" >
              <el-date-picker v-model="formData.fixedStartTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="固定结束时间:"  prop="fixedEndTime" >
              <el-date-picker v-model="formData.fixedEndTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="有效天数:"  prop="validDay" >
              <el-input v-model.number="formData.validDay" :clearable="true" placeholder="请输入有效天数" />
            </el-form-item>
            <el-form-item label="兑换所需积分:"  prop="redeemPoints" >
              <el-input v-model.number="formData.redeemPoints" :clearable="true" placeholder="请输入兑换所需积分" />
            </el-form-item>
          </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
                <el-descriptions-item label="所属店铺商家">
                        {{ formData.businessId }}
                </el-descriptions-item>
                <el-descriptions-item label="优惠券名称">
                        {{ formData.name }}
                </el-descriptions-item>
                <el-descriptions-item label="总数量">
                        {{ formData.totalCount }}
                </el-descriptions-item>
                <el-descriptions-item label="剩余数量">
                        {{ formData.remainCount }}
                </el-descriptions-item>
                <el-descriptions-item label="已使用数量">
                        {{ formData.usedCount }}
                </el-descriptions-item>
                <el-descriptions-item label="优惠券类型">
                        {{ formData.couponCategory }}
                </el-descriptions-item>
                <el-descriptions-item label="满">
                        {{ formData.fullMoney }}
                </el-descriptions-item>
                <el-descriptions-item label="减">
                        {{ formData.minusMoney }}
                </el-descriptions-item>
                <el-descriptions-item label="现金金额">
                        {{ formData.cash }}
                </el-descriptions-item>
                <el-descriptions-item label="折扣">
                        {{ formData.discount }}
                </el-descriptions-item>
                <el-descriptions-item label="使用说明">
                        {{ formData.usageInstructions }}
                </el-descriptions-item>
                <el-descriptions-item label="每个用户可兑换数量">
                        {{ formData.redeemCount }}
                </el-descriptions-item>
                <el-descriptions-item label="兑换开始时间">
                      {{ formatDate(formData.redeemStartTime) }}
                </el-descriptions-item>
                <el-descriptions-item label="兑换结束时间">
                      {{ formatDate(formData.redeemEndTime) }}
                </el-descriptions-item>
                <el-descriptions-item label="有效期类型">
                        {{ formData.expCategory }}
                </el-descriptions-item>
                <el-descriptions-item label="固定开始时间">
                      {{ formatDate(formData.fixedStartTime) }}
                </el-descriptions-item>
                <el-descriptions-item label="固定结束时间">
                      {{ formatDate(formData.fixedEndTime) }}
                </el-descriptions-item>
                <el-descriptions-item label="有效天数">
                        {{ formData.validDay }}
                </el-descriptions-item>
                <el-descriptions-item label="兑换所需积分">
                        {{ formData.redeemPoints }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createCoupon,
  deleteCoupon,
  deleteCouponByIds,
  updateCoupon,
  findCoupon,
  getCouponList
} from '@/api/coupon'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'Coupon'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        businessId: 0,
        name: '',
        totalCount: 0,
        remainCount: 0,
        usedCount: 0,
        couponCategory: 0,
        fullMoney: 0,
        minusMoney: 0,
        cash: 0,
        discount: 0,
        usageInstructions: '',
        redeemCount: 0,
        redeemStartTime: new Date(),
        redeemEndTime: new Date(),
        expCategory: 0,
        fixedStartTime: new Date(),
        fixedEndTime: new Date(),
        validDay: 0,
        redeemPoints: 0,
        })


// 验证规则
const rule = reactive({
               businessId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               totalCount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               remainCount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               usedCount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               couponCategory : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               fullMoney : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               minusMoney : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               cash : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               discount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               usageInstructions : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               redeemCount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               redeemStartTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               redeemEndTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               expCategory : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               fixedStartTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               fixedEndTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               validDay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               redeemPoints : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
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
const setOptions = async () =>{
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
const deleteCouponFunc = async (row) => {
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


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findCoupon({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.recoupon
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          businessId: 0,
          name: '',
          totalCount: 0,
          remainCount: 0,
          usedCount: 0,
          couponCategory: 0,
          fullMoney: 0,
          minusMoney: 0,
          cash: 0,
          discount: 0,
          usageInstructions: '',
          redeemCount: 0,
          redeemStartTime: new Date(),
          redeemEndTime: new Date(),
          expCategory: 0,
          fixedStartTime: new Date(),
          fixedEndTime: new Date(),
          validDay: 0,
          redeemPoints: 0,
          }
}


// 打开弹窗
const openDialog = () => {
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
        couponCategory: 0,
        fullMoney: 0,
        minusMoney: 0,
        cash: 0,
        discount: 0,
        usageInstructions: '',
        redeemCount: 0,
        redeemStartTime: new Date(),
        redeemEndTime: new Date(),
        expCategory: 0,
        fixedStartTime: new Date(),
        fixedEndTime: new Date(),
        validDay: 0,
        redeemPoints: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
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

</script>

<style>

</style>
