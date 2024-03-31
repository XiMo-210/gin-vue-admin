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
      <video
        v-if="model"
        :src="model"
        style="height: 150px"
        controls
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
import { isVideoMime } from '@/utils/image'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'UploadVideo',
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
  const isLt10M = file.size / 1024 / 1024 < 10
  const isVideo = isVideoMime(file.type)
  let pass = true
  if (!isVideo) {
    ElMessage.error('上传视频只能是 mp4,webm 格式')
    pass = false
  }
  if (!isLt10M) {
    ElMessage.error('上传视频大小不能超过 100MB')
    pass = false
  }
  return pass
}

</script>

<style scoped>
</style>
