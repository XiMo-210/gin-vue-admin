<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="种类:" prop="category">
          <el-switch v-model="formData.category" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="评论对象ID:" prop="targetId">
          <el-input v-model.number="formData.targetId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="一星数量:" prop="one">
          <el-input v-model.number="formData.one" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="二星数量:" prop="two">
          <el-input v-model.number="formData.two" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="三星数量:" prop="three">
          <el-input v-model.number="formData.three" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="四星数量:" prop="four">
          <el-input v-model.number="formData.four" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="五星数量:" prop="five">
          <el-input v-model.number="formData.five" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="评分:" prop="score">
          <el-input-number v-model="formData.score" :precision="2" :clearable="true"></el-input-number>
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
  createCommentScore,
  updateCommentScore,
  findCommentScore
} from '@/api/commentScore'

defineOptions({
    name: 'CommentScoreForm'
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
            category: false,
            targetId: 0,
            one: 0,
            two: 0,
            three: 0,
            four: 0,
            five: 0,
            score: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCommentScore({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recommentScore
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
               res = await createCommentScore(formData.value)
               break
             case 'update':
               res = await updateCommentScore(formData.value)
               break
             default:
               res = await createCommentScore(formData.value)
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
