<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="所属店铺商家:" prop="businessId">
          <el-input v-model.number="formData.businessId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="优惠券名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总数量:" prop="totalCount">
          <el-input v-model.number="formData.totalCount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="剩余数量:" prop="remainCount">
          <el-input v-model.number="formData.remainCount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="已使用数量:" prop="usedCount">
          <el-input v-model.number="formData.usedCount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="优惠券类型:" prop="couponCategory">
          <el-input v-model.number="formData.couponCategory" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="满:" prop="fullMoney">
          <el-input v-model.number="formData.fullMoney" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="减:" prop="minusMoney">
          <el-input v-model.number="formData.minusMoney" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="现金金额:" prop="cash">
          <el-input v-model.number="formData.cash" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="折扣:" prop="discount">
          <el-input v-model.number="formData.discount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="使用说明:" prop="usageInstructions">
          <el-input v-model="formData.usageInstructions" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="每个用户可兑换数量:" prop="redeemCount">
          <el-input v-model.number="formData.redeemCount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="兑换开始时间:" prop="redeemStartTime">
          <el-date-picker v-model="formData.redeemStartTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="兑换结束时间:" prop="redeemEndTime">
          <el-date-picker v-model="formData.redeemEndTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="有效期类型:" prop="expCategory">
          <el-input v-model.number="formData.expCategory" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="固定开始时间:" prop="fixedStartTime">
          <el-date-picker v-model="formData.fixedStartTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="固定结束时间:" prop="fixedEndTime">
          <el-date-picker v-model="formData.fixedEndTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="有效天数:" prop="validDay">
          <el-input v-model.number="formData.validDay" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="兑换所需积分:" prop="redeemPoints">
          <el-input v-model.number="formData.redeemPoints" :clearable="true" placeholder="请输入" />
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
  createCoupon,
  updateCoupon,
  findCoupon
} from '@/api/coupon'

defineOptions({
    name: 'CouponForm'
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
            businessId: 0,
            name: '',
            totalCount: 0,
            remainCount: 0,
            usedCount: 0,
            couponCategory: 0,
            fullMoney: 0,
            minusMoney: 0,
            cash: 0,
            discount: 0,
            usageInstructions: '',
            redeemCount: 0,
            redeemStartTime: new Date(),
            redeemEndTime: new Date(),
            expCategory: 0,
            fixedStartTime: new Date(),
            fixedEndTime: new Date(),
            validDay: 0,
            redeemPoints: 0,
        })
// 验证规则
const rule = reactive({
               businessId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               totalCount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               remainCount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               usedCount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               couponCategory : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               fullMoney : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               minusMoney : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cash : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               discount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               usageInstructions : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               redeemCount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               redeemStartTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               redeemEndTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               expCategory : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               fixedStartTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               fixedEndTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               validDay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               redeemPoints : [{
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
      const res = await findCoupon({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recoupon
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
               res = await createCoupon(formData.value)
               break
             case 'update':
               res = await updateCoupon(formData.value)
               break
             default:
               res = await createCoupon(formData.value)
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
