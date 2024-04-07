<template>
  <div
    ref="charts"
    class="chart"
    style="width: 800px;height:400px;"
  />
</template>

<script setup>
import {
  completeCondition
} from '@/api/task'
import { ref, onMounted, toRefs } from 'vue'
import * as echarts from 'echarts'
import 'echarts-liquidfill'

const props = defineProps({
  id: Number
})
const { id } = toRefs(props)

const charts = ref()

const total = ref(0)
const complete = ref(0)

onMounted(async() => {
  const mychart = echarts.init(charts.value)
  const res = await completeCondition({ ID: id.value })
  if (res.code === 0) {
    total.value = res.data.total
    complete.value = res.data.complete

    const data1 = complete.value / total.value
    const data2 = data1 * 0.8
    const data3 = data2 * 0.8
    mychart.setOption({
      graphic: [{
        type: 'group',
        left: 'center',
        top: '60%',
        children: [{
          type: 'text',
          z: 100,
          left: '40',
          top: 'middle',
          style: {
            fill: '#fff',
            text: '完成人数: ' + complete.value + '人\n\n任务人数: ' + total.value + '人',
            font: '18px Microsoft YaHei'
          }
        }]
      }],
      series: {
        type: 'liquidFill',
        radius: '80%',
        center: ['50%', '50%'],
        data: [
          data1,
          {
            value: data2,
            direction: 'left',
          },
          data3],
        backgroundStyle: {
          borderWidth: 1,
          color: 'RGBA(51, 66, 127, 0.7)',
        },
        label: {
          position: ['50%', '40%'],
          formatter: (data1 * 100).toFixed(1) + '%',
          textStyle: {
            color: '#fff',
            fontSize: 30,
          },
        },
        outline: {
          show: true,
          itemStyle: {
            borderWidth: 2,
            borderColor: '#F1FAFA',
          },
        },
      },
    })
  }
})
</script>
<style scoped>
</style>
