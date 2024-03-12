<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="所属广告主:" prop="advertiserId">
          <el-input v-model.number="formData.advertiserId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文本:" prop="text">
          <el-input v-model="formData.text" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="图片:" prop="img">
          <SelectImage v-model="formData.img" file-type="image"/>
       </el-form-item>
        <el-form-item label="视频:" prop="video">
          <SelectFile v-model="formData.video" />
       </el-form-item>
        <el-form-item label="跳转链接:" prop="link">
          <el-input v-model="formData.link" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="开始日期:" prop="startDate">
          <el-date-picker v-model="formData.startDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="结束日期:" prop="endDate">
          <el-date-picker v-model="formData.endDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="开始时段:" prop="startHour">
          <el-input v-model.number="formData.startHour" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="结束时段:" prop="endHour">
          <el-input v-model.number="formData.endHour" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="目标受众:" prop="target">
           <SelectImage v-model="formData.target" multiple file-type="image"/>
       </el-form-item>
        <el-form-item label="关键词:" prop="keywords">
           <SelectImage v-model="formData.keywords" multiple file-type="image"/>
       </el-form-item>
        <el-form-item label="计费类型:" prop="costCategory">
          <el-input v-model.number="formData.costCategory" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="购买量:" prop="bugAmount">
          <el-input v-model.number="formData.bugAmount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="已展示量:" prop="costImpressions">
          <el-input v-model.number="formData.costImpressions" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="已点击数:" prop="costClicks">
          <el-input v-model.number="formData.costClicks" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="花费费用:" prop="cost">
          <el-input v-model.number="formData.cost" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
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
  createAd,
  updateAd,
  findAd
} from '@/api/ad'

defineOptions({
    name: 'AdForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectImage from '@/components/selectImage/selectImage.vue'
import SelectFile from '@/components/selectFile/selectFile.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
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
               }],
               text : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               img : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               video : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               link : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               startDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               endDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               startHour : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               endHour : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               target : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               keywords : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               costCategory : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               bugAmount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cost : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
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
      const res = await findAd({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.read
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
