<template>
  <div>
    <el-upload
      class="uploader"
      action="https://api.lonesome.cn/api/wx/upload"
      :show-file-list="false"
      :on-success="uploadSuccess"
      :on-error="uploadError"
      :before-upload="beforeUpload"
    >
      <el-image
        v-if="model"
        :src="model"
        style="height: 150px"
        :fit="contain"
      />
      <el-button
        v-else
        class="uploader-icon"
      >上传
      </el-button>
    </el-upload>
  </div>
</template>

<script setup>
import { isImageMime } from '@/utils/image'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'UploadImg',
})

const model = defineModel({ required: true })

const uploadSuccess = (res) => {
  const { data } = res
  if (data) {
    model.value = data
  }
  ElMessage({
    type: 'success',
    message: '上传成功'
  })
}

const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
}

const beforeUpload = (file) => {
  const isLt1M = file.size / 1024 / 1024 < 1
  const isImage = isImageMime(file.type)
  let pass = true
  if (!isImage) {
    ElMessage.error('上传图片只能是 jpg,png,svg,webp 格式')
    pass = false
  }
  if (!isLt1M) {
    ElMessage.error('上传图片大小不能超过1MB')
    pass = false
  }
  return pass
}

</script>

<style scoped>
.uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon .uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 150px;
  height: 150px;
  text-align: center;
}
</style>
