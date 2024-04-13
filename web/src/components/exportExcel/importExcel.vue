<template>
  <el-upload
    v-loading.fullscreen.lock="fullscreenLoading"
    element-loading-text="学生信息导入中，请耐心等待..."
    :action="url"
    :show-file-list="false"
    :before-upload="beginLoad"
    :on-success="handleSuccess"
    :multiple="false"
  >
    <el-button
      type="primary"
      icon="upload"
    >导入</el-button>
  </el-upload>

</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

const baseUrl = import.meta.env.VITE_BASE_API

const props = defineProps({
  templateId: {
    type: String,
    required: true
  }
})

const emit = defineEmits(['on-success'])

const url = `${baseUrl}/sysExportTemplate/importExcel?templateID=${props.templateId}`

const fullscreenLoading = ref(false)

const beginLoad = () => {
  fullscreenLoading.value = true
}

const handleSuccess = (res) => {
  if (res.code === 0) {
    ElMessage.success('导入成功')
    emit('on-success')
  } else {
    ElMessage.error(res.msg)
  }
  fullscreenLoading.value = false
}

</script>
