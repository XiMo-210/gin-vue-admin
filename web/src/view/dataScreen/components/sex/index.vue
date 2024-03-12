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
            value: 52.7,
            name: '男  覆盖率52.7%'
          },
          {
            name: ' ',
            value: 47.3,
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
            value: 47.3,
            name: '女  覆盖率47.3%',
          },
          {
            name: ' ',
            value: 52.7,
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
