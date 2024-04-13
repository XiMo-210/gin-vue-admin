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
  var maxData = []
  var xData = [
    '赵伟返', '钱数芳', '孙江娜', '李秀英', '周   敏', '吴帮静', '郑   丽', '王提强', '冯精磊', '陈   洋',
    '褚   艳', '卫   勇', '蒋平军', '沈规霞', '韩时亮', '杨西刚', '朱嗷艳', '秦课勇', '尤拉军', '许偶霞',
    '何从刚', '吕凑艳', '施   军', '张不勇', '孔破霞', '曹学亮', '严   刚', '华   艳', '金   军', '魏夏霞',
    '陶篇刚', '姜送艳', '戚次军', '谢毕霞', '邹笑亮', '喻   刚', '柏系艳', '吴   军', '窦   霞', '章覆亮'
  ]

  maxData.push(100)
  for (let i = 1; i <= 40; i++) {
    maxData.push(maxData[i - 1] + 10 * (i % 5 + 1))
    xData[i - 1] = (41 - i) + ' ' + xData[i - 1]
  }

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
      max: 2000,
      text: ['High', 'Low'],
      dimension: 0,
      inRange: {
        color: ['#cde6c7', '#7bbfea', '#f3715c']
      }
    }
  }

  var startIdx = 32
  var endIdx = 40

  var slicedMaxData = maxData.slice(startIdx, endIdx)
  var slicedXData = xData.slice(startIdx, endIdx)

  option.yAxis.data = slicedXData
  option.series[0].data = slicedMaxData

  mychart.setOption(option)

  setInterval(() => {
    var slicedMaxData = maxData.slice(startIdx, endIdx)
    var slicedXData = xData.slice(startIdx, endIdx)

    option.yAxis.data = slicedXData
    option.series[0].data = slicedMaxData

    mychart.setOption(option)

    startIdx -= 8
    endIdx -= 8
    if (startIdx < 0) {
      startIdx = 32
      endIdx = 40
    }
  }, 5000)
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
