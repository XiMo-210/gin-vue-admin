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

  const data = [
    { value: 200, name: '化学工程学院' },
    { value: 150, name: '材料科学与工程学院' },
    { value: 200, name: '机械工程学院' },
    { value: 150, name: '信息工程学院' },
    { value: 150, name: '管理学院' },
    { value: 100, name: '经济学院' },
    { value: 80, name: '生物工程学院' },
    { value: 100, name: '环境学院' },
    { value: 80, name: '人文学院' },
    { value: 180, name: '药学院、绿色制药协同创新中心' },
    { value: 120, name: '理学院' },
    { value: 80, name: '法学院' },
    { value: 200, name: '外国语学院' },
    { value: 100, name: '设计与建筑学院' },
    { value: 100, name: '土木工程学院' },
    { value: 50, name: '健行学院' },
    { value: 80, name: '公共管理学院' },
    { value: 200, name: '计算机科学与技术学院、软件学院' },
    { value: 300, name: '教育科学与技术学院' }
  ]

  mychart.setOption({
    tooltip: {
      trigger: 'item'
    },
    legend: {
      show: false
    },
    title: [
      {
        text: '各学院新生数',
        left: '50%',
        top: '18',
        textAlign: 'center',
        textBaseline: 'middle',
        textStyle: {
          color: '#1a1b4e',
          fontStyle: 'normal',
          fontWeight: 'bold',
          fontSize: 30
        },
      },
      {
        text: '学生总数',
        left: '50%',
        top: '45%',
        textAlign: 'center',
        textBaseline: 'middle',
        textStyle: {
          color: '#999',
          fontWeight: 'normal',
          fontSize: 20
        }
      },
      {
        text: '3000人',
        left: '50%',
        top: '55%',
        textAlign: 'center',
        textBaseline: 'middle',
        textStyle: {
          color: '#666',
          fontWeight: 'normal',
          fontSize: 35
        }
      }
    ],
    series: [
      {
        name: '学院',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: true,
        top: '20',
        padAngle: 2,
        itemStyle: {
          borderRadius: 10,
        },
        emphasis: {
          label: {
            show: true,
            fontWeight: 'bold',
          }
        },
        data: data
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
