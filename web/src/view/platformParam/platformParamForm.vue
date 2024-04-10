<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="学校名称:" prop="schoolName">
          <el-input v-model="formData.schoolName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="校徽:" prop="schoolEmblem">
          <SelectImage v-model="formData.schoolEmblem" file-type="image"/>
       </el-form-item>
        <el-form-item label="入学年份:" prop="admissionYear">
          <el-input v-model.number="formData.admissionYear" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="报到开始日:" prop="registerStartDate">
          <el-date-picker v-model="formData.registerStartDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="报到结束日:" prop="registerEndDate">
          <el-date-picker v-model="formData.registerEndDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="当前学年:" prop="termYear">
          <el-input v-model.number="formData.termYear" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前学期:" prop="term">
          <el-input v-model="formData.term" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="学期开始时间:" prop="termStartDate">
          <el-date-picker v-model="formData.termStartDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="学期结束时间:" prop="termEndDate">
          <el-date-picker v-model="formData.termEndDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createPlatformParam,
  updatePlatformParam,
  findPlatformParam
} from '@/api/platformParam'

defineOptions({
    name: 'PlatformParamForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectImage from '@/components/selectImage/selectImage.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            schoolName: '',
            schoolEmblem: "",
            admissionYear: 0,
            registerStartDate: new Date(),
            registerEndDate: new Date(),
            termYear: 0,
            term: '',
            termStartDate: new Date(),
            termEndDate: new Date(),
        })
// 验证规则
const rule = reactive({
               schoolName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               schoolEmblem : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               admissionYear : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               registerStartDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               registerEndDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               termYear : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               term : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               termStartDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               termEndDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPlatformParam({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.replatformParam
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createPlatformParam(formData.value)
               break
             case 'update':
               res = await updatePlatformParam(formData.value)
               break
             default:
               res = await createPlatformParam(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
