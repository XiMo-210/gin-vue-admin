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
      
        <el-form-item label="名称" prop="name">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="计费类型" prop="costCategory">
            
             <el-input v-model.number="searchInfo.costCategory" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="状态" prop="status">
            
             <el-input v-model.number="searchInfo.status" placeholder="搜索条件" />

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
        
        <el-table-column align="left" label="所属广告主" prop="advertiserId" width="120" />
        <el-table-column align="left" label="名称" prop="name" width="120" />
        <el-table-column align="left" label="文本" prop="text" width="120" />
          <el-table-column label="图片" width="200">
              <template #default="scope">
                <el-image style="width: 100px; height: 100px" :src="getUrl(scope.row.img)" fit="cover"/>
              </template>
          </el-table-column>
                    <el-table-column label="视频" width="200">
                        <template #default="scope">
                             <div class="file-list">
                               <el-tag v-for="file in scope.row.video" :key="file.uid">{{file.name}}</el-tag>
                             </div>
                        </template>
                    </el-table-column>
        <el-table-column align="left" label="跳转链接" prop="link" width="120" />
         <el-table-column align="left" label="开始日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.startDate) }}</template>
         </el-table-column>
         <el-table-column align="left" label="结束日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.endDate) }}</template>
         </el-table-column>
        <el-table-column align="left" label="开始时段" prop="startHour" width="120" />
        <el-table-column align="left" label="结束时段" prop="endHour" width="120" />
           <el-table-column label="目标受众" width="200">
              <template #default="scope">
                 <div class="multiple-img-box">
                    <el-image v-for="(item,index) in scope.row.target" style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover"/>
                </div>
              </template>
           </el-table-column>
           <el-table-column label="关键词" width="200">
              <template #default="scope">
                 <div class="multiple-img-box">
                    <el-image v-for="(item,index) in scope.row.keywords" style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover"/>
                </div>
              </template>
           </el-table-column>
        <el-table-column align="left" label="计费类型" prop="costCategory" width="120" />
        <el-table-column align="left" label="购买量" prop="bugAmount" width="120" />
        <el-table-column sortable align="left" label="已展示量" prop="costImpressions" width="120" />
        <el-table-column sortable align="left" label="已点击数" prop="costClicks" width="120" />
        <el-table-column sortable align="left" label="花费费用" prop="cost" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateAdFunc(scope.row)">变更</el-button>
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
            <el-form-item label="所属广告主:"  prop="advertiserId" >
              <el-input v-model.number="formData.advertiserId" :clearable="true" placeholder="请输入所属广告主" />
            </el-form-item>
            <el-form-item label="名称:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
            </el-form-item>
            <el-form-item label="文本:"  prop="text" >
              <el-input v-model="formData.text" :clearable="true"  placeholder="请输入文本" />
            </el-form-item>
            <el-form-item label="图片:"  prop="img" >
                <SelectImage
                 v-model="formData.img"
                 file-type="image"
                />
            </el-form-item>
            <el-form-item label="视频:"  prop="video" >
                <SelectFile v-model="formData.video" />
            </el-form-item>
            <el-form-item label="跳转链接:"  prop="link" >
              <el-input v-model="formData.link" :clearable="true"  placeholder="请输入跳转链接" />
            </el-form-item>
            <el-form-item label="开始日期:"  prop="startDate" >
              <el-date-picker v-model="formData.startDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="结束日期:"  prop="endDate" >
              <el-date-picker v-model="formData.endDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="开始时段:"  prop="startHour" >
              <el-input v-model.number="formData.startHour" :clearable="true" placeholder="请输入开始时段" />
            </el-form-item>
            <el-form-item label="结束时段:"  prop="endHour" >
              <el-input v-model.number="formData.endHour" :clearable="true" placeholder="请输入结束时段" />
            </el-form-item>
            <el-form-item label="目标受众:"  prop="target" >
                <SelectImage
                 multiple
                 v-model="formData.target"
                 file-type="image"
                 />
            </el-form-item>
            <el-form-item label="关键词:"  prop="keywords" >
                <SelectImage
                 multiple
                 v-model="formData.keywords"
                 file-type="image"
                 />
            </el-form-item>
            <el-form-item label="计费类型:"  prop="costCategory" >
              <el-input v-model.number="formData.costCategory" :clearable="true" placeholder="请输入计费类型" />
            </el-form-item>
            <el-form-item label="购买量:"  prop="bugAmount" >
              <el-input v-model.number="formData.bugAmount" :clearable="true" placeholder="请输入购买量" />
            </el-form-item>
            <el-form-item label="已展示量:"  prop="costImpressions" >
              <el-input v-model.number="formData.costImpressions" :clearable="true" placeholder="请输入已展示量" />
            </el-form-item>
            <el-form-item label="已点击数:"  prop="costClicks" >
              <el-input v-model.number="formData.costClicks" :clearable="true" placeholder="请输入已点击数" />
            </el-form-item>
            <el-form-item label="花费费用:"  prop="cost" >
              <el-input v-model.number="formData.cost" :clearable="true" placeholder="请输入花费费用" />
            </el-form-item>
            <el-form-item label="状态:"  prop="status" >
              <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入状态" />
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
                <el-descriptions-item label="所属广告主">
                        {{ formData.advertiserId }}
                </el-descriptions-item>
                <el-descriptions-item label="名称">
                        {{ formData.name }}
                </el-descriptions-item>
                <el-descriptions-item label="文本">
                        {{ formData.text }}
                </el-descriptions-item>
                <el-descriptions-item label="图片">
                        <el-image style="width: 50px; height: 50px" :preview-src-list="ReturnArrImg(formData.img)" :src="getUrl(formData.img)" fit="cover" />
                </el-descriptions-item>
                <el-descriptions-item label="视频">
                        <div class="fileBtn" v-for="(item,index) in formData.video" :key="index">
                          <el-button type="primary" text bg @click="onDownloadFile(item.url)">
                            <el-icon style="margin-right: 5px"><Download /></el-icon>
                            {{ item.name }}
                          </el-button>
                        </div>
                </el-descriptions-item>
                <el-descriptions-item label="跳转链接">
                        {{ formData.link }}
                </el-descriptions-item>
                <el-descriptions-item label="开始日期">
                      {{ formatDate(formData.startDate) }}
                </el-descriptions-item>
                <el-descriptions-item label="结束日期">
                      {{ formatDate(formData.endDate) }}
                </el-descriptions-item>
                <el-descriptions-item label="开始时段">
                        {{ formData.startHour }}
                </el-descriptions-item>
                <el-descriptions-item label="结束时段">
                        {{ formData.endHour }}
                </el-descriptions-item>
                <el-descriptions-item label="目标受众">
                        <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="ReturnArrImg(formData.target)" :initial-index="index" v-for="(item,index) in formData.target" :key="index" :src="getUrl(item)" fit="cover" />
                </el-descriptions-item>
                <el-descriptions-item label="关键词">
                        <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="ReturnArrImg(formData.keywords)" :initial-index="index" v-for="(item,index) in formData.keywords" :key="index" :src="getUrl(item)" fit="cover" />
                </el-descriptions-item>
                <el-descriptions-item label="计费类型">
                        {{ formData.costCategory }}
                </el-descriptions-item>
                <el-descriptions-item label="购买量">
                        {{ formData.bugAmount }}
                </el-descriptions-item>
                <el-descriptions-item label="已展示量">
                        {{ formData.costImpressions }}
                </el-descriptions-item>
                <el-descriptions-item label="已点击数">
                        {{ formData.costClicks }}
                </el-descriptions-item>
                <el-descriptions-item label="花费费用">
                        {{ formData.cost }}
                </el-descriptions-item>
                <el-descriptions-item label="状态">
                        {{ formData.status }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
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
  getAdList
} from '@/api/ad'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'Ad'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        advertiserId: 0,
        name: '',
        text: '',
        img: "",
        video: [],
        link: '',
        startDate: new Date(),
        endDate: new Date(),
        startHour: 0,
        endHour: 0,
        target: [],
        keywords: [],
        costCategory: 0,
        bugAmount: 0,
        costImpressions: 0,
        costClicks: 0,
        cost: 0,
        status: 0,
        })


// 验证规则
const rule = reactive({
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
               text : [{
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
               img : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               video : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               link : [{
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
               startDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               endDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               startHour : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               endHour : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               target : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               keywords : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               costCategory : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               bugAmount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               cost : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               status : [{
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
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAdFunc = async (row) => {
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


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findAd({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.read
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          advertiserId: 0,
          name: '',
          text: '',
          link: '',
          startDate: new Date(),
          endDate: new Date(),
          startHour: 0,
          endHour: 0,
          costCategory: 0,
          bugAmount: 0,
          costImpressions: 0,
          costClicks: 0,
          cost: 0,
          status: 0,
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
        advertiserId: 0,
        name: '',
        text: '',
        link: '',
        startDate: new Date(),
        endDate: new Date(),
        startHour: 0,
        endHour: 0,
        costCategory: 0,
        bugAmount: 0,
        costImpressions: 0,
        costClicks: 0,
        cost: 0,
        status: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
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
              }
      })
}

const downloadFile = (url) => {
    window.open(getUrl(url), '_blank')
}

</script>

<style>

.file-list{
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}

.fileBtn{
  margin-bottom: 10px;
}

.fileBtn:last-child{
  margin-bottom: 0;
}

</style>
