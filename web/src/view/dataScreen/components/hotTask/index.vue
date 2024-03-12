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
  var chartData = {
    timelineData: ['主线任务', '支线任务', '隐藏任务'],
    seriesData: [
      {
        xdata: ['主线任务1', '主线任务3', '主线任务5'],
        ydata: [275, 175, 146]
      },
      {
        xdata: ['支线任务4', '支线任务2', '支线任务3'],
        ydata: [386, 287, 277]
      },
      {
        xdata: ['隐藏任务1', '隐藏任务3', '隐藏任务6'],
        ydata: [174, 145, 125]
      }
    ]
  }
  var optionVal = chartData.seriesData.map((item, index) => {
    return {
      xAxis: {
        data: item.xdata
      },
      series: [
        {
          name: '',
          type: 'pictorialBar',
          symbolSize: [0, 0],
          symbolOffset: [0, 0],
          label: {
            show: true,
            position: 'top',
            formatter: function(params) {
              return item.ydata[params.dataIndex]
            },
            fontSize: 25,
            fontWeight: 'bold',
            color: '#afb4db'
          },
          data: item.ydata
        },
        {
          type: 'bar',
          data: item.ydata,
          barWidth: 40,
          itemStyle: {
            normal: {
              color: new echarts.graphic.LinearGradient(
                0,
                0,
                0,
                1, // 从上到下变浅
                [
                  { offset: 0, color: '#4e72b8' },
                  { offset: 1, color: '#9b95c9' }
                ],
                false
              ),
              opacity: 1
            }
          }
        }
      ]
    }
  })

  mychart.setOption({
    baseOption: {
      title: {
        text: '今日热点任务',
        left: '50%',
        top: '10',
        textAlign: 'center',
        textBaseline: 'middle',
        textStyle: {
          color: '#1a1b4e',
          fontStyle: 'normal',
          fontWeight: 'bold',
          fontSize: 25
        }
      },
      timeline: {
        data: chartData.timelineData,
        axisType: 'category',
        autoPlay: true,
        playInterval: 5000,
        show: false,
      },
      xAxis: {
        type: 'category'
      },
      yAxis: {
        name: '今日完成（人）',
        type: 'value'
      }
    },
    options: optionVal
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
            scale:0.9
          }
  </style>
