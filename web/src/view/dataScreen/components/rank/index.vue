<template>
  <div class="box">
    <BorderBox4
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
  var maxData = [1343, 1492, 1576, 1676, 1887, 2003, 2376, 2492, 2743]
  var xData = [
    '张三',
    '李四',
    '王五',
    '赵六',
    '灰太狼',
    '喜羊羊',
    '猪猪侠',
    '芙莉莲',
    '芜湖'
  ]

  var tempMaxData = maxData.pop()
  var tempXData = xData.pop()
  var option = {
    title: {
      text: '积分榜单',
      left: 'center',
      top: '15%',
      textStyle: {
        color: '#1a1b4e',
        fontStyle: 'normal',
        fontWeight: 'bold',
        fontSize: 30
      },
    },
    grid: {
      left: 100,
      top: 100,
      height: '60%'
    },
    xAxis: {
      show: false,
      type: 'value'
    },
    yAxis: {
      type: 'category',
      data: xData
    },
    series: [
      {
        name: '最大/最小时延',
        type: 'bar',
        barCategoryGap: '20%',
        label: {
          show: true,
          color: '#feeeed'
        },
        itemStyle: {
          barBorderRadius: 10,
          color: '#fff'
        },
        data: maxData
      }
    ],
    visualMap: {
      orient: 'horizontal',
      show: false,
      min: 0,
      max: 3000,
      text: ['High', 'Low'],
      dimension: 0,
      inRange: {
        color: ['#cde6c7', '#7bbfea', '#f3715c']
      }
    }
  }
  setInterval(() => {
    xData.unshift(tempXData)
    tempXData = xData.pop()

    maxData.unshift(tempMaxData)
    tempMaxData = maxData.pop()

    mychart.setOption(option)
  }, 1000)
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
    }
</style>
