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

  const data1 = 2494 / 5014
  const data2 = 2520 / 5014

  mychart.setOption({
    title: {
      text: '男女比例',
      left: '50%',
      top: '20%',
      textAlign: 'center',
      textBaseline: 'middle',
      textStyle: {
        color: '#1a1b4e',
        fontStyle: 'normal',
        fontWeight: 'bold',
        fontSize: 30
      },
    },
    series: [
      {
        center: ['25.0%', '50%'],
        radius: ['49%', '50%'],
        clockWise: false,
        hoverAnimation: false,
        top: 20,
        left: 50,
        type: 'pie',
        itemStyle: {
          normal: {
            label: {
              show: true,
              textStyle: {
                fontSize: 15,
                fontWeight: 'bold'
              },
              position: 'center'
            },
            labelLine: {
              'show': false
            },
            color: '#5886f0',
            borderColor: '#5886f0',
            borderWidth: 25
          },
          emphasis: {
            label: {
              textStyle: {
                fontSize: 15,
                fontWeight: 'bold'
              }
            },
            color: '#5886f0',
            borderColor: '#5886f0',
            borderWidth: 25
          }
        },
        data: [
          {
            value: data1,
            name: '男  覆盖率' + (data1 * 100).toFixed(1) + '%'
          },
          {
            name: ' ',
            value: 1 - data1,
            itemStyle: {
              normal: {
                label: {
                  show: false
                },
                labelLine: {
                  show: false
                },
                color: '#5886f0',
                borderColor: '#5886f0',
                borderWidth: 0
              },
              emphasis: {
                color: '#5886f0',
                borderColor: '#5886f0',
                borderWidth: 0
              }
            }
          }
        ]
      },
      {
        center: ['75.0%', '50%'],
        radius: ['49%', '50%'],
        top: 20,
        left: -50,
        clockWise: false,
        hoverAnimation: false,
        type: 'pie',
        itemStyle: {
          normal: {
            label: {
              show: true,
              textStyle: {
                fontSize: 15,
                fontWeight: 'bold'
              },
              position: 'center'
            },
            labelLine: {
              show: false
            },
            color: '#ee3a3a',
            borderColor: '#ee3a3a',
            borderWidth: 25
          },
          emphasis: {
            label: {
              textStyle: {
                fontSize: 15,
                fontWeight: 'bold'
              }
            },
            color: '#ee3a3a',
            borderColor: '#ee3a3a',
            borderWidth: 25
          }
        },
        data: [
          {
            value: data2,
            name: '女  覆盖率' + (data2 * 100).toFixed(1) + '%',
          },
          {
            name: ' ',
            value: 1 - data2,
            itemStyle: {
              normal: {
                label: {
                  show: false
                },
                labelLine: {
                  show: false
                },
                color: '#ee3a3a',
                borderColor: '#ee3a3a',
                borderWidth: 0
              },
              emphasis: {
                color: '#ee3a3a',
                borderColor: '#ee3a3a',
                borderWidth: 0
              }
            }
          }
        ]
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
      }
    </style>
