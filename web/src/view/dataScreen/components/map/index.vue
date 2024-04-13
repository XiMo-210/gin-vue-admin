<template>
  <div
    ref="charts"
  />
</template>

<script setup>
import { ref, onMounted } from 'vue'
import * as echarts from 'echarts'
import chinaJSON from './china.json'

const charts = ref()

echarts.registerMap('china', chinaJSON)

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

  const data = [{
    name: '浙江',
    value: 3129
  }, {
    name: '北京',
    value: 11
  }, {
    name: '天津',
    value: 30
  }, {
    name: '河北',
    value: 78
  }, {
    name: '山西',
    value: 83
  }, {
    name: '内蒙古',
    value: 20
  }, {
    name: '辽宁',
    value: 43
  }, {
    name: '吉林',
    value: 20
  }, {
    name: '黑龙江',
    value: 40
  }, {
    name: '上海',
    value: 31
  }, {
    name: '江苏',
    value: 70
  }, {
    name: '安徽',
    value: 106
  }, {
    name: '福建',
    value: 85
  }, {
    name: '江西',
    value: 82
  }, {
    name: '山东',
    value: 93
  }, {
    name: '河南',
    value: 107
  }, {
    name: '湖北',
    value: 50
  }, {
    name: '湖南',
    value: 130
  }, {
    name: '广东',
    value: 127
  }, {
    name: '广西',
    value: 113
  }, {
    name: '海南',
    value: 20
  }, {
    name: '重庆',
    value: 45
  }, {
    name: '四川',
    value: 81
  }, {
    name: '贵州',
    value: 140
  }, {
    name: '云南',
    value: 89
  }, {
    name: '西藏',
    value: 22
  }, {
    name: '陕西',
    value: 40
  }, {
    name: '甘肃',
    value: 80
  }, {
    name: '青海',
    value: 15
  }, {
    name: '宁夏',
    value: 10
  }, {
    name: '新疆',
    value: 24
  }, {
    name: '台湾',
    value: 0
  }, {
    name: '香港',
    value: 0
  }, {
    name: '澳门',
    value: 0
  }]

  mychart.setOption({
    tooltip: {
      trigger: 'item'
    },
    title: {
      text: '全国新生分布',
      itemGap: 30,
      left: 'center',
      top: '8%',
      textStyle: {
        color: '#1a1b4e',
        fontStyle: 'normal',
        fontWeight: 'bold',
        fontSize: 30
      },
    },
    visualMap: {
      min: 0,
      max: 200,
      left: '10%',
      bottom: '5%',
      text: ['高', '低'],
      calculable: true,
      inRange: {
        color: ['#ffffff', '#E0DAFF', '#ADBFFF', '#9CB4FF', '#6A9DFF', '#3889FF']
      }
    },
    geo: {
      map: 'china',
      label: {
        normal: {
          show: true,
          color: '#c1c4c8'
        }
      },
      itemStyle: {
        normal: {
          areaColor: '#fbfbfb',
          borderColor: '#b9b4b7'
        },
        emphasis: {
          areaColor: '#cde6c7',
        }
      }
    },
    series: [
      {
        name: '地区',
        type: 'map',
        mapType: 'china',
        geoIndex: 0,
        label: {
          normal: {
            show: true
          },
          emphasis: {
            show: true
          }
        },
        data: data
      }]
  })
})
</script>

<style scoped>
</style>
