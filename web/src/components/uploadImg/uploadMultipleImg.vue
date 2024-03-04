<template>
  <div>
    <el-upload
      action="https://api.lonesome.cn/api/wx/upload"
      :auto-upload="true"
      :show-file-list="false"
      :on-success="uploadSuccess"
      :on-error="uploadError"
      :before-upload="beforeUpload"
    >
      <el-button
        type="primary"
        plain
        style="margin-bottom: 12px;"
      >上传</el-button>
    </el-upload>
    <span class="image-list">
      <span
        v-for="(img, index) in model"
        :key="index"
        class="thumbnail-item"
      >
        <el-image
          :src="img"
          fit="contain"
          :zoom-rate="1.2"
          :max-scale="7"
          :min-scale="0.2"
          :preview-src-list="model"
          :initial-index="index"
          style="width: 80px; height: 80px;"
        />
      </span>
    </span>
  </div>
</template>

<script lang="ts" setup>
import { isImageMime } from '@/utils/image'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'UploadImg',
})

const model = defineModel < Array < string >> ({ required: true })

const uploadSuccess = (res) => {
  const { data } = res
  if (data) {
    model.value.push(data)
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

<style>
.image-list {
  display: flex;
  flex-wrap: wrap;
  margin-top: 5px;
  margin-bottom: -20px;
}

.thumbnail-item {
  position: relative;
  margin-right: 10px;
}
</style>
