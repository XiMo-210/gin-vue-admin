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

function randomData() {
  return Math.round(Math.random() * 1000)
}

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
    name: '北京',
    value: randomData()
  }, {
    name: '天津',
    value: randomData()
  }, {
    name: '上海',
    value: randomData()
  }, {
    name: '重庆',
    value: randomData()
  }, {
    name: '河北',
    value: randomData()
  }, {
    name: '河南',
    value: randomData()
  }, {
    name: '云南',
    value: randomData()
  }, {
    name: '辽宁',
    value: randomData()
  }, {
    name: '黑龙江',
    value: randomData()
  }, {
    name: '湖南',
    value: randomData()
  }, {
    name: '安徽',
    value: randomData()
  }, {
    name: '山东',
    value: randomData()
  }, {
    name: '新疆',
    value: randomData()
  }, {
    name: '江苏',
    value: randomData()
  }, {
    name: '浙江',
    value: randomData()
  }, {
    name: '江西',
    value: randomData()
  }, {
    name: '湖北',
    value: randomData()
  }, {
    name: '广西',
    value: randomData()
  }, {
    name: '甘肃',
    value: randomData()
  }, {
    name: '山西',
    value: randomData()
  }, {
    name: '内蒙古',
    value: randomData()
  }, {
    name: '陕西',
    value: randomData()
  }, {
    name: '吉林',
    value: randomData()
  }, {
    name: '福建',
    value: randomData()
  }, {
    name: '贵州',
    value: randomData()
  }, {
    name: '广东',
    value: randomData()
  }, {
    name: '青海',
    value: randomData()
  }, {
    name: '西藏',
    value: randomData()
  }, {
    name: '四川',
    value: randomData()
  }, {
    name: '宁夏',
    value: randomData()
  }, {
    name: '海南',
    value: randomData()
  }, {
    name: '台湾',
    value: randomData()
  }, {
    name: '香港',
    value: randomData()
  }, {
    name: '澳门',
    value: randomData()
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
      max: 1000,
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
