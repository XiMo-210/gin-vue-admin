<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="姓名:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别:" prop="gender">
          <el-input v-model="formData.gender" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出生日期:" prop="birthday">
          <el-date-picker v-model="formData.birthday" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="身份证:" prop="idCard">
          <el-input v-model="formData.idCard" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="录取通知书编号:" prop="admissionLetterId">
          <el-input v-model="formData.admissionLetterId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="入学年份:" prop="admissionYear">
          <el-input v-model.number="formData.admissionYear" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="生源地:" prop="originPlace">
          <el-input v-model="formData.originPlace" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="校区:" prop="campus">
          <el-input v-model="formData.campus" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="学号:" prop="studentId">
          <el-input v-model="formData.studentId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="学院:" prop="college">
          <el-input v-model="formData.college" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="专业:" prop="major">
          <el-input v-model="formData.major" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="班级:" prop="class">
          <el-input v-model="formData.class" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="寝室:" prop="dormitory">
          <el-input v-model="formData.dormitory" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="人像:" prop="portrait">
          <SelectImage v-model="formData.portrait" file-type="image"/>
       </el-form-item>
        <el-form-item label="用户ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
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
  createStudentInfo,
  updateStudentInfo,
  findStudentInfo
} from '@/api/studentInfo'

defineOptions({
    name: 'StudentInfoForm'
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
            name: '',
            gender: '',
            birthday: new Date(),
            idCard: '',
            admissionLetterId: '',
            admissionYear: 0,
            originPlace: '',
            campus: '',
            studentId: '',
            college: '',
            major: '',
            class: '',
            dormitory: '',
            portrait: "",
            userId: 0,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               gender : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               birthday : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               idCard : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               admissionLetterId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               admissionYear : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               originPlace : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               campus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               studentId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               college : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               major : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               class : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               dormitory : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               portrait : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userId : [{
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
      const res = await findStudentInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.restudentInfo
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
               res = await createStudentInfo(formData.value)
               break
             case 'update':
               res = await updateStudentInfo(formData.value)
               break
             default:
               res = await createStudentInfo(formData.value)
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
