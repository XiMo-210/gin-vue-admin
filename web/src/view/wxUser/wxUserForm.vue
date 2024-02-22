<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="微信用户唯一标识:" prop="openid">
          <el-input v-model="formData.openid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="关联学生信息:" prop="studentInfoId">
          <el-input v-model.number="formData.studentInfoId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="人脸识别标识:" prop="faceId">
          <SelectImage v-model="formData.faceId" file-type="image"/>
       </el-form-item>
        <el-form-item label="用户名:" prop="username">
          <el-input v-model="formData.username" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="人像:" prop="avatar">
          <SelectImage v-model="formData.avatar" file-type="image"/>
       </el-form-item>
        <el-form-item label="个性签名:" prop="signature">
          <el-input v-model="formData.signature" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="爱好:" prop="hobby">
          <el-input v-model="formData.hobby" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="联系方式:" prop="contact">
          <el-input v-model="formData.contact" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手机:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="智慧种子:" prop="exp">
          <el-input v-model.number="formData.exp" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="求索石:" prop="points">
          <el-input v-model.number="formData.points" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否完成主线任务:" prop="isCompletedMain">
          <el-switch v-model="formData.isCompletedMain" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="正在进行的任务:" prop="curTask">
          <el-input v-model.number="formData.curTask" :clearable="true" placeholder="请输入" />
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
  createWxUser,
  updateWxUser,
  findWxUser
} from '@/api/wxUser'

defineOptions({
    name: 'WxUserForm'
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
            openid: '',
            studentInfoId: 0,
            faceId: "",
            username: '',
            password: '',
            avatar: "",
            signature: '',
            hobby: '',
            contact: '',
            phone: '',
            email: '',
            exp: 0,
            points: 0,
            isCompletedMain: false,
            curTask: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWxUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewxUser
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
               res = await createWxUser(formData.value)
               break
             case 'update':
               res = await updateWxUser(formData.value)
               break
             default:
               res = await createWxUser(formData.value)
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
