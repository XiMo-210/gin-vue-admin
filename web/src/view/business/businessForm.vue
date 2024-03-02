<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="关联的后台管理用户:" prop="sysUserId">
          <el-input v-model.number="formData.sysUserId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="封面图:" prop="pic">
          <el-input v-model="formData.pic" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="导航地址:" prop="loc">
          <el-input v-model="formData.loc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="营业时间:" prop="businessHours">
          <el-input v-model="formData.businessHours" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="联系方式:" prop="contact">
          <el-input v-model="formData.contact" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="介绍:" prop="introduction">
          <el-input v-model="formData.introduction" :clearable="true" placeholder="请输入" />
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
  createBusiness,
  updateBusiness,
  findBusiness
} from '@/api/business'

defineOptions({
    name: 'BusinessForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
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
               sysUserId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               pic : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               address : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               loc : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               businessHours : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               contact : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               introduction : [{
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
      const res = await findBusiness({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rebusiness
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
               res = await createBusiness(formData.value)
               break
             case 'update':
               res = await updateBusiness(formData.value)
               break
             default:
               res = await createBusiness(formData.value)
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
