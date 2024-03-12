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
    dataset: {
      source: [
        ['num', 'college'],
        [3, '教育科学与技术'],
        [5, '化学工程学院'],
        [7, '材料科学与工程学院'],
        [3, '机械工程学院'],
        [5, '信息工程学院'],
        [7, '管理学院'],
        [6, '经济学院'],
        [7, '生物工程学院'],
        [9, '环境学院'],
        [7, '人文学院'],
        [4, '药学院'],
        [7, '理学院'],
        [2, '法学院'],
        [8, '外国语学院'],
        [9, '设计与建筑学院'],
        [4, '土木工程学院'],
        [6, '健行学院'],
        [5, '公共管理学院'],
        [7, '计算机学院']
      ]
    },
    grid: {
      containLabel: true
    },
    xAxis: {
      name: '完成数'
    },
    yAxis: {
      type: 'category'
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
        }
      }
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
