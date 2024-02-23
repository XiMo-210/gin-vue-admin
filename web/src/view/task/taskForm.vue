<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="任务类型:" prop="category">
          <el-input v-model.number="formData.category" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="任务名称:" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="任务描述:" prop="desc">
          <el-input v-model="formData.desc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="任务校区:" prop="campus">
          <el-input v-model="formData.campus" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="任务学院:" prop="college">
          <el-input v-model="formData.college" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="奖励积分:" prop="reward">
          <el-input v-model.number="formData.reward" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否需要完成主线任务:" prop="needMain">
          <el-switch v-model="formData.needMain" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="开始时间:" prop="startTime">
          <el-date-picker v-model="formData.startTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="结束时间:" prop="endTime">
          <el-date-picker v-model="formData.endTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  createTask,
  updateTask,
  findTask
} from '@/api/task'

defineOptions({
    name: 'TaskForm'
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
            category: 0,
            title: '',
            desc: '',
            campus: '',
            college: '',
            reward: 0,
            needMain: false,
            startTime: new Date(),
            endTime: new Date(),
        })
// 验证规则
const rule = reactive({
               category : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               desc : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               campus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               college : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               reward : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               needMain : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               startTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               endTime : [{
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
      const res = await findTask({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retask
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
               res = await createTask(formData.value)
               break
             case 'update':
               res = await updateTask(formData.value)
               break
             default:
               res = await createTask(formData.value)
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
