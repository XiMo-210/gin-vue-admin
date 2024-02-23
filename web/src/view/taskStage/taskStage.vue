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
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="所属任务ID" prop="taskId" width="120" />
        <el-table-column align="left" label="阶段" prop="stage" width="120" />
        <el-table-column align="left" label="阶段任务名称" prop="title" width="120" />
        <el-table-column align="left" label="阶段任务描述" prop="desc" width="120" />
        <el-table-column align="left" label="需要准备的物品" prop="requiredItem" width="120" />
           <el-table-column label="图片" width="200">
              <template #default="scope">
                 <div class="multiple-img-box">
                    <el-image v-for="(item,index) in scope.row.imgs" style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover"/>
                </div>
              </template>
           </el-table-column>
        <el-table-column align="left" label="是否需要上传指定内容图片" prop="needPic" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.needPic) }}</template>
        </el-table-column>
        <el-table-column align="left" label="图片是否需要人脸" prop="needFace" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.needFace) }}</template>
        </el-table-column>
        <el-table-column align="left" label="指定内容图片" prop="pic" width="120" />
        <el-table-column align="left" label="是否需要指定位置" prop="needLoc" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.needLoc) }}</template>
        </el-table-column>
        <el-table-column align="left" label="允许距离" prop="allowDist" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.allowDist) }}</template>
        </el-table-column>
        <el-table-column align="left" label="是否需要导航" prop="needNav" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.needNav) }}</template>
        </el-table-column>
        <el-table-column align="left" label="指定位置" prop="loc" width="120" />
        <el-table-column align="left" label="是否关联摄像头" prop="needCamera" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.needCamera) }}</template>
        </el-table-column>
        <el-table-column align="left" label="摄像头编号" prop="cameraId" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateTaskStageFunc(scope.row)">变更</el-button>
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
            <el-form-item label="所属任务ID:"  prop="taskId" >
              <el-input v-model.number="formData.taskId" :clearable="true" placeholder="请输入所属任务ID" />
            </el-form-item>
            <el-form-item label="阶段:"  prop="stage" >
              <el-input v-model.number="formData.stage" :clearable="true" placeholder="请输入阶段" />
            </el-form-item>
            <el-form-item label="阶段任务名称:"  prop="title" >
              <el-input v-model="formData.title" :clearable="true"  placeholder="请输入阶段任务名称" />
            </el-form-item>
            <el-form-item label="阶段任务描述:"  prop="desc" >
              <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入阶段任务描述" />
            </el-form-item>
            <el-form-item label="需要准备的物品:"  prop="requiredItem" >
              <el-input v-model="formData.requiredItem" :clearable="true"  placeholder="请输入需要准备的物品" />
            </el-form-item>
            <el-form-item label="图片:"  prop="imgs" >
                <SelectImage
                 multiple
                 v-model="formData.imgs"
                 file-type="image"
                 />
            </el-form-item>
            <el-form-item label="是否需要上传指定内容图片:"  prop="needPic" >
              <el-switch v-model="formData.needPic" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="图片是否需要人脸:"  prop="needFace" >
              <el-switch v-model="formData.needFace" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="指定内容图片:"  prop="pic" >
              <el-input v-model="formData.pic" :clearable="true"  placeholder="请输入指定内容图片" />
            </el-form-item>
            <el-form-item label="是否需要指定位置:"  prop="needLoc" >
              <el-switch v-model="formData.needLoc" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="允许距离:"  prop="allowDist" >
              <el-switch v-model="formData.allowDist" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="是否需要导航:"  prop="needNav" >
              <el-switch v-model="formData.needNav" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="指定位置:"  prop="loc" >
              <el-input v-model="formData.loc" :clearable="true"  placeholder="请输入指定位置" />
            </el-form-item>
            <el-form-item label="是否关联摄像头:"  prop="needCamera" >
              <el-switch v-model="formData.needCamera" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="摄像头编号:"  prop="cameraId" >
              <el-input v-model.number="formData.cameraId" :clearable="true" placeholder="请输入摄像头编号" />
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
                <el-descriptions-item label="所属任务ID">
                        {{ formData.taskId }}
                </el-descriptions-item>
                <el-descriptions-item label="阶段">
                        {{ formData.stage }}
                </el-descriptions-item>
                <el-descriptions-item label="阶段任务名称">
                        {{ formData.title }}
                </el-descriptions-item>
                <el-descriptions-item label="阶段任务描述">
                        {{ formData.desc }}
                </el-descriptions-item>
                <el-descriptions-item label="需要准备的物品">
                        {{ formData.requiredItem }}
                </el-descriptions-item>
                <el-descriptions-item label="图片">
                        <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="ReturnArrImg(formData.imgs)" :initial-index="index" v-for="(item,index) in formData.imgs" :key="index" :src="getUrl(item)" fit="cover" />
                </el-descriptions-item>
                <el-descriptions-item label="是否需要上传指定内容图片">
                    {{ formatBoolean(formData.needPic) }}
                </el-descriptions-item>
                <el-descriptions-item label="图片是否需要人脸">
                    {{ formatBoolean(formData.needFace) }}
                </el-descriptions-item>
                <el-descriptions-item label="指定内容图片">
                        {{ formData.pic }}
                </el-descriptions-item>
                <el-descriptions-item label="是否需要指定位置">
                    {{ formatBoolean(formData.needLoc) }}
                </el-descriptions-item>
                <el-descriptions-item label="允许距离">
                    {{ formatBoolean(formData.allowDist) }}
                </el-descriptions-item>
                <el-descriptions-item label="是否需要导航">
                    {{ formatBoolean(formData.needNav) }}
                </el-descriptions-item>
                <el-descriptions-item label="指定位置">
                        {{ formData.loc }}
                </el-descriptions-item>
                <el-descriptions-item label="是否关联摄像头">
                    {{ formatBoolean(formData.needCamera) }}
                </el-descriptions-item>
                <el-descriptions-item label="摄像头编号">
                        {{ formData.cameraId }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createTaskStage,
  deleteTaskStage,
  deleteTaskStageByIds,
  updateTaskStage,
  findTaskStage,
  getTaskStageList
} from '@/api/taskStage'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'TaskStage'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        taskId: 0,
        stage: 0,
        title: '',
        desc: '',
        requiredItem: '',
        imgs: [],
        needPic: false,
        needFace: false,
        pic: '',
        needLoc: false,
        allowDist: false,
        needNav: false,
        loc: '',
        needCamera: false,
        cameraId: 0,
        })


// 验证规则
const rule = reactive({
               taskId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               stage : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               title : [{
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
               needPic : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               needFace : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               pic : [{
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
               needLoc : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               allowDist : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               needNav : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               loc : [{
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
               needCamera : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               cameraId : [{
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
    if (searchInfo.value.needPic === ""){
        searchInfo.value.needPic=null
    }
    if (searchInfo.value.needFace === ""){
        searchInfo.value.needFace=null
    }
    if (searchInfo.value.needLoc === ""){
        searchInfo.value.needLoc=null
    }
    if (searchInfo.value.allowDist === ""){
        searchInfo.value.allowDist=null
    }
    if (searchInfo.value.needNav === ""){
        searchInfo.value.needNav=null
    }
    if (searchInfo.value.needCamera === ""){
        searchInfo.value.needCamera=null
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
  const table = await getTaskStageList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteTaskStageFunc(row)
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
      const res = await deleteTaskStageByIds({ IDs })
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
const updateTaskStageFunc = async(row) => {
    const res = await findTaskStage({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.retaskStage
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteTaskStageFunc = async (row) => {
    const res = await deleteTaskStage({ ID: row.ID })
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
  const res = await findTaskStage({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.retaskStage
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          taskId: 0,
          stage: 0,
          title: '',
          desc: '',
          requiredItem: '',
          needPic: false,
          needFace: false,
          pic: '',
          needLoc: false,
          allowDist: false,
          needNav: false,
          loc: '',
          needCamera: false,
          cameraId: 0,
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
        taskId: 0,
        stage: 0,
        title: '',
        desc: '',
        requiredItem: '',
        needPic: false,
        needFace: false,
        pic: '',
        needLoc: false,
        allowDist: false,
        needNav: false,
        loc: '',
        needCamera: false,
        cameraId: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createTaskStage(formData.value)
                  break
                case 'update':
                  res = await updateTaskStage(formData.value)
                  break
                default:
                  res = await createTaskStage(formData.value)
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
