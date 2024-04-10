<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="right"
        :rules="rule"
        label-width="100px"
      >
        <el-form-item
          label="校徽:"
          prop="schoolEmblem"
        >
          <UploadImg
            v-model="formData.schoolEmblem"
          />
        </el-form-item>
        <el-form-item
          label="学校名称:"
          prop="schoolName"
          style="width: 30%;"
        >
          <el-input
            v-model="formData.schoolName"
            :clearable="true"
            placeholder="请输入学校名称"
          />
        </el-form-item>
        <el-form-item
          label="入学年份:"
          prop="admissionYear"
        >
          <el-input-number
            v-model="formData.admissionYear"
            controls-position="right"
            :step="2"
            step-strictly
          />
        </el-form-item>
        <el-form-item
          label="报到时间:"
          prop="registerDate"
          style="width: 30%;"
        >
          <el-date-picker
            v-model="formData.registerDate"
            type="daterange"
            range-separator="To"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
          />
        </el-form-item>
        <el-form-item
          label="当前学年:"
          prop="termYear"
        >
          <el-input-number
            v-model="formData.termYear"
            controls-position="right"
            :step="2"
            step-strictly
          />
        </el-form-item>
        <el-form-item
          label="当前学期:"
          prop="term"
          style="width: 20%;"
        >
          <el-select
            v-model="formData.term"
          >
            <el-option
              key="1"
              value="第一学期"
              label="第一学期"
            />
            <el-option
              key="2"
              value="第二学期"
              label="第二学期"
            />
            <el-option
              key="3"
              value="第三学期"
              label="第三学期"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="学期时间:"
          prop="termDate"
          style="width: 30%;"
        >
          <el-date-picker
            v-model="formData.termDate"
            type="daterange"
            range-separator="To"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            unlink-panels
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            icon="edit"
            @click="onSubmit"
          >更新平台参数</el-button>
          <el-button
            icon="refresh"
            @click="onReset"
          >重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createPlatformParam,
  updatePlatformParam,
  findPlatformParam
} from '@/api/platformParam'
import UploadImg from '@/components/uploadImg/uploadImg.vue'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'PlatformParam'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  ID: 0,
  schoolName: '',
  schoolEmblem: '',
  admissionYear: 0,
  registerDate: [],
  registerStartDate: new Date(),
  registerEndDate: new Date(),
  termYear: 0,
  term: '第一学期',
  termDate: [],
  termStartDate: new Date(),
  termEndDate: new Date(),
})

// 验证规则
const rule = reactive({
  schoolName: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  schoolEmblem: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  admissionYear: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  registerDate: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  termYear: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  term: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  termDate: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
})

const elFormRef = ref()

const type = ref('create')

// 弹窗确定
const onSubmit = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    formData.value.ID = 1
    formData.value.registerStartDate = formData.value.registerDate[0]
    formData.value.registerEndDate = formData.value.registerDate[1]
    formData.value.termStartDate = formData.value.termDate[0]
    formData.value.termEndDate = formData.value.termDate[1]
    switch (type.value) {
      case 'create':
        res = await createPlatformParam(formData.value)
        break
      case 'update':
        res = await updatePlatformParam(formData.value)
        break
      default:
        res = await createPlatformParam(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更新成功'
      })
    }
  })
}

const getPlatformParamInfo = async() => {
  const info = await findPlatformParam()
  if (info.code === 0) {
    formData.value = info.data.replatformParam
    if (formData.value.ID === 0) {
      type.value = 'create'
      var date = new Date()
      formData.value.admissionYear = date.getFullYear()
      formData.value.termYear = date.getFullYear()
      formData.value.term = '第一学期'
    } else {
      type.value = 'update'
      formData.value.registerDate = [formData.value.registerStartDate, formData.value.registerEndDate]
      formData.value.termDate = [formData.value.termStartDate, formData.value.termEndDate]
    }
  }
}

getPlatformParamInfo()

const onReset = () => {
  getPlatformParamInfo()
}

</script>

<style>

</style>
