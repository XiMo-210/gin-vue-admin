<template>
  <div class="box">
    <BorderBox11
      title="广告点击数"
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
import 'echarts-wordcloud'

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
    { value: '4012', name: '中国移动' },
    { value: '3275', name: '微精弘' },
    { value: '3054', name: 'KFC' },
    { value: '2985', name: '华为' },
    { value: '2845', name: '校园卡' },
    { value: '2776', name: '校级组织' },
    { value: '2645', name: '小米' },
    { value: '2575', name: '美团' },
    { value: '2475', name: '饿了么' },
    { value: '2424', name: '哔哩哔哩' },
    { value: '2383', name: 'ROG' },
    { value: '2268', name: '明日方舟' },
    { value: '2133', name: '小红书' },
    { value: '2074', name: '网易云' },
    { value: '1727', name: '学习通' },
    { value: '1686', name: '钉钉' },
    { value: '1535', name: '飞书' },
    { value: '1456', name: '校园健身' },
    { value: '1324', name: 'CSDN' },
    { value: '1228', name: '稀土掘金' },
    { value: '1110', name: '联想' },
    { value: '1056', name: '玩家国度' },
  ]
  mychart.setOption({
    tooltip: {
      show: true,
      borderColor: '#fe9a8bb3',
      borderWidth: 1,
      padding: [10, 15, 10, 15],
      confine: true,
      backgroundColor: 'rgba(255, 255, 255, .9)',
      textStyle: {
        color: 'hotpink',
        lineHeight: 22
      },
      extraCssText: 'box-shadow: 0 4px 20px -4px rgba(199, 206, 215, .7);border-radius: 4px;'
    },
    series: [
      {
        type: 'wordCloud',
        shape: 'pentagon',
        left: 'center',
        top: 'center',
        width: '100%',
        height: '100%',
        sizeRange: [14, 50],
        rotationRange: [0, 0],
        rotationStep: 0,
        gridSize: 25,
        drawOutOfBound: false,
        layoutAnimation: true,
        textStyle: {
          fontFamily: 'PingFangSC-Semibold',
          fontWeight: 600,
          color: function(params) {
            const colors = ['#fe9a8bb3', '#fe9a8bb3', '#fe9a8b03', '#9E87FFb3', '#9E87FFb3', '#9E87FFb3', '#fe9a8bb3', '#fe9a8bb3', '#fe9a8bb3', '#73DDFF', '#58D5FF']
            return colors[parseInt(Math.random() * 10)]
          },
        },
        emphasis: {
          focus: 'none',
        },
        data: data,
      },
    ],
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
    scale: 0.8;
  }
</style>
