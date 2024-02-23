<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="所属任务ID:" prop="taskId">
          <el-input v-model.number="formData.taskId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="阶段:" prop="stage">
          <el-input v-model.number="formData.stage" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="阶段任务名称:" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="阶段任务描述:" prop="desc">
          <el-input v-model="formData.desc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="需要准备的物品:" prop="requiredItem">
          <el-input v-model="formData.requiredItem" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="图片:" prop="imgs">
           <SelectImage v-model="formData.imgs" multiple file-type="image"/>
       </el-form-item>
        <el-form-item label="是否需要上传指定内容图片:" prop="needPic">
          <el-switch v-model="formData.needPic" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="图片是否需要人脸:" prop="needFace">
          <el-switch v-model="formData.needFace" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="指定内容图片:" prop="pic">
          <el-input v-model="formData.pic" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否需要指定位置:" prop="needLoc">
          <el-switch v-model="formData.needLoc" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="允许距离:" prop="allowDist">
          <el-switch v-model="formData.allowDist" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="是否需要导航:" prop="needNav">
          <el-switch v-model="formData.needNav" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="指定位置:" prop="loc">
          <el-input v-model="formData.loc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否关联摄像头:" prop="needCamera">
          <el-switch v-model="formData.needCamera" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="摄像头编号:" prop="cameraId">
          <el-input v-model.number="formData.cameraId" :clearable="true" placeholder="请输入" />
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
  createTaskStage,
  updateTaskStage,
  findTaskStage
} from '@/api/taskStage'

defineOptions({
    name: 'TaskStageForm'
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
               }],
               stage : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               needPic : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               needFace : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               pic : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               needLoc : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               allowDist : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               needNav : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               loc : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               needCamera : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cameraId : [{
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
      const res = await findTaskStage({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retaskStage
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
