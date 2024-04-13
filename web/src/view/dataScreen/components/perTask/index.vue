<template>
  <div class="box">
    <BorderBox4
      :reverse="true"
      class="border"
    />
    <div
      ref="charts"
      class="chart"
    />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import * as echarts from 'echarts'

const charts = ref()

onMounted(() => {
  const mychart = echarts.init(charts.value)

  const data = [
    { value: 2.4, name: '健行学院' },
    { value: 7.6, name: '化学工程学院' },
    { value: 2.8, name: '材料科学与工程学院' },
    { value: 3.4, name: '生物工程学院' },
    { value: 2.3, name: '环境学院' },
    { value: 2.2, name: '食品科学与工程学院' },
    { value: 6.4, name: '药学院、绿色制药协同创新中心' },
    { value: 5.1, name: '机械工程学院' },
    { value: 4.2, name: '经济学院' },
    { value: 5.6, name: '管理学院' },
    { value: 1.8, name: '公共管理学院' },
    { value: 4.9, name: '法学院' },
    { value: 6.8, name: '信息工程学院' },
    { value: 7.7, name: '计算机科学与技术学院、软件学院' },
    { value: 7.3, name: '设计与建筑学院' },
    { value: 5.4, name: '土木工程学院' },
    { value: 5.2, name: '人文学院' },
    { value: 4.4, name: '外国语学院' },
    { value: 4.6, name: '理学院' },
    { value: 4.4, name: '教育科学与技术学院' }
  ]

  let index = 0

  function fun() {
    setInterval(function() {
      mychart.dispatchAction({
        type: 'hideTip',
        seriesIndex: 0,
        dataIndex: index
      })
      // 显示提示框
      mychart.dispatchAction({
        type: 'showTip',
        seriesIndex: 0,
        dataIndex: index
      })
      // 取消高亮指定的数据图形
      mychart.dispatchAction({
        type: 'downplay',
        seriesIndex: 0,
        dataIndex: index === 0 ? data.length : index - 1
      })
      mychart.dispatchAction({
        type: 'highlight',
        seriesIndex: 0,
        dataIndex: index
      })
      index++
      if (index > data.length) {
        index = 0
      }
    }, 3000)
  }

  fun()

  mychart.setOption({
    title: {
      text: '各学院人均任务完成数',
      left: '50%',
      top: '18',
      textAlign: 'center',
      textBaseline: 'middle',
      textStyle: {
        color: '#1a1b4e',
        fontStyle: 'normal',
        fontWeight: 'bold',
        fontSize: 30
      }
    },
    tooltip: {
      trigger: 'item'
    },
    grid: {
      containLabel: true
    },
    xAxis: {
      name: '完成数'
    },
    yAxis: {
      type: 'category',
      data: data.map(item => item.name)
    },
    visualMap: {
      orient: 'horizontal',
      left: 'center',
      bottom: 20,
      min: 0,
      max: 20,
      text: ['High', 'Low'],
      dimension: 0,
      inRange: {
        color: ['#FCCF31', '#EA5455', '#F55555']
      }
    },
    series: [
      {
        type: 'bar',
        encode: {
          x: 'num',
          y: 'college'
        },
        data: data
      },
    ]
  })
})

</script>

<style scoped>
    .box {
        display: flex;
        position: relative;
      }

      .border{
        position: absolute;
      }

      .chart{
        width: 100%;
        height: 100%;
        margin-top: 20px;
        scale: 0.9;
      }
</style>
