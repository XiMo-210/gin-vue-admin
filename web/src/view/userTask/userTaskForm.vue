<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="userId字段:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="taskId字段:" prop="taskId">
          <el-input v-model.number="formData.taskId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="pic字段:" prop="pic">
          <el-input v-model="formData.pic" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="loc字段:" prop="loc">
          <el-input v-model="formData.loc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="curStage字段:" prop="curStage">
          <el-switch v-model="formData.curStage" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createUserTask,
  updateUserTask,
  findUserTask
} from '@/api/userTask'

defineOptions({
    name: 'UserTaskForm'
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
            userId: 0,
            taskId: 0,
            pic: '',
            loc: '',
            curStage: false,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findUserTask({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reuserTask
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
               res = await createUserTask(formData.value)
               break
             case 'update':
               res = await updateUserTask(formData.value)
               break
             default:
               res = await createUserTask(formData.value)
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
