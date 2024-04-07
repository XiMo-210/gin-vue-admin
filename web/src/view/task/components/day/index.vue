<template>
  <div
    ref="charts"
    class="chart"
    style="width: 800px;height:400px;"
  />
</template>

<script setup>
import {
  dayCompleteNum
} from '@/api/task'
import { ref, onMounted, toRefs } from 'vue'
import * as echarts from 'echarts'

const props = defineProps({
  id: Number
})
const { id } = toRefs(props)

const charts = ref()

const date = ref([])
const data = ref([])

onMounted(async() => {
  const mychart = echarts.init(charts.value)
  const res = await dayCompleteNum({ ID: id.value })
  if (res.code === 0) {
    date.value = res.data.dates
    data.value = res.data.nums
    const option = {
      tooltip: {
        trigger: 'axis',
        position: function(pt) {
          return [pt[0], '10%']
        }
      },
      xAxis: {
        type: 'category',
        boundaryGap: false,
        name: '日期',
        data: date.value
      },
      yAxis: {
        type: 'value',
        boundaryGap: [0, '100%'],
        name: '日完成数',
        axisLabel: {
          formatter: function(value) {
            if (value % 1 === 0) {
              return value
            } else {
              return ''
            }
          }
        }
      },
      dataZoom: [
        {
          type: 'inside',
          start: 0,
          end: 100
        },
        {
          start: 0,
          end: 10
        }
      ],
      series: [
        {
          name: '日完成数',
          type: 'line',
          symbol: 'none',
          smooth: true,
          areaStyle: {},
          data: data.value
        }
      ]
    }
    mychart.setOption(option)
  }
})

// const getDateArr = () => {
//   const currentDate = new Date()
//   const dateArr = []

//   for (let i = 6; i >= 0; i--) {
//     const previousDate = new Date(currentDate.getTime() - i * 24 * 60 * 60 * 1000)
//     const previousMonth = previousDate.getMonth() + 1
//     const previousDay = previousDate.getDate()

//     const previousDateString = `${previousMonth}-${previousDay}`

//     dateArr.push(previousDateString)
//   }

//   return dateArr
// }

</script>
  <style scoped>
  </style>
