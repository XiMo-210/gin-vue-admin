<template>
  <div class="box">
    <BorderBox12 class="border" />
    <div
      ref="leftCharts"
      class="left"
    />
    <div
      ref="charts"
      class="center"
    />
    <div
      ref="rightCharts"
      class="right"
    />
  </div>
</template>

<script setup>
import 'echarts-liquidfill'
import * as echarts from 'echarts'
import { ref, onMounted } from 'vue'

const leftCharts = ref()
onMounted(() => {
  const mycharts = echarts.init(leftCharts.value)
  const data = 4937
  const data1 = data / 5014
  const data2 = data1 * 0.8
  mycharts.setOption({
    graphic: [{
      type: 'group',
      left: 'center',
      top: '60%',
      children: [{
        type: 'text',
        z: 100,
        left: '40',
        top: 'middle',
        style: {
          fill: '#fff',
          text: '注册学生数: ' + data + '人',
          font: '14px Microsoft YaHei'
        }
      }]
    }],
    series: {
      type: 'liquidFill',
      radius: '80%',
      center: ['50%', '50%'],
      data: [data1, data2],
      color: ['rgba(2, 159, 200, 1)', 'rgba(0, 186, 211, .5)'],
      backgroundStyle: {
        borderWidth: 1,
        color: 'RGBA(51, 66, 127, 0.7)',
      },
      label: {
        position: ['50%', '35%'],
        formatter: (data1 * 100).toFixed(1) + '%',
        textStyle: {
          color: '#fff',
          fontSize: 30,
        },
      },
      outline: {
        show: true,
        itemStyle: {
          borderWidth: 2,
          borderColor: '#F1FAFA',
        },
      },
    },
  })
})

const rightCharts = ref()
onMounted(() => {
  const mycharts = echarts.init(rightCharts.value)
  const data = 3482
  const data1 = data / 5014
  const data2 = data1 * 0.8
  mycharts.setOption({
    graphic: [{
      type: 'group',
      left: 'center',
      top: '60%',
      children: [{
        type: 'text',
        z: 100,
        left: '40',
        top: 'middle',
        style: {
          fill: '#fff',
          text: '报到人数: ' + data + '人',
          font: '14px Microsoft YaHei'
        }
      }]
    }],
    series: {
      type: 'liquidFill',
      radius: '80%',
      center: ['50%', '50%'],
      data: [data1,
        {
          value: data2,
          direction: 'left',
        }],
      color: [
        {
          type: 'linear',
          x: 0,
          y: 0,
          x2: 0,
          y2: 1,
          colorStops: [
            {
              offset: 1,
              color: 'rgba(6, 187, 112, 0.3)', // 下
            },
            {
              offset: 0,
              color: 'rgba(11, 201, 199, 0.3)',
            },
          ],
          globalCoord: false,
        },
        {
          type: 'linear',
          x: 0,
          y: 0,
          x2: 0,
          y2: 1,
          colorStops: [
            {
              offset: 1,
              color: 'rgba(6, 187, 112, 1)', // 下
            },
            {
              offset: 0,
              color: 'rgba(11, 201, 199, 1)',
            },
          ],
          globalCoord: false,
        },
      ],
      backgroundStyle: {
        borderWidth: 1,
        color: 'RGBA(51, 66, 127, 0.7)',
      },
      label: {
        position: ['50%', '35%'],
        formatter: (data1 * 100).toFixed(1) + '%',
        textStyle: {
          color: '#fff',
          fontSize: 30,
        },
      },
      outline: {
        show: true,
        itemStyle: {
          borderWidth: 2,
          borderColor: '#F1FAFA',
        },
      },
    },
  })
})

const charts = ref()
onMounted(() => {
  const mycharts = echarts.init(charts.value)
  const data = 4698
  const data1 = data / 5014
  const data2 = data1 * 0.8
  mycharts.setOption({
    graphic: [{
      type: 'group',
      left: 'center',
      top: '60%',
      children: [{
        type: 'text',
        z: 100,
        left: '40',
        top: 'middle',
        style: {
          fill: '#fff',
          text: '认证用户数:  ' + data + '人',
          font: '14px Microsoft YaHei'
        }
      }]
    }],
    series: {
      type: 'liquidFill',
      radius: '80%',
      center: ['50%', '50%'],
      data: [data1, data2],
      color: ['#8190E3', '#71B1FC'],
      backgroundStyle: {
        borderWidth: 1,
        color: 'RGBA(51, 66, 127, 0.7)',
      },
      label: {
        position: ['50%', '35%'],
        formatter: (data1 * 100).toFixed(1) + '%',
        textStyle: {
          color: '#fff',
          fontSize: 30,
        },
      },
      outline: {
        show: true,
        itemStyle: {
          borderWidth: 2,
          borderColor: '#F1FAFA',
        },
      },
    },
  })
})

</script>

<style scoped lang="scss">
  .box {
    margin-top: 30px;
    display: flex;
    position: relative;
  }

  .border{
    position: absolute;
  }

  .left {
    width: 100%;
    height: 100%;
    margin-right: -100px;
  }

  .center {
    width: 100%;
    height: 100%;
  }

  .right {
    width: 100%;
    height: 100%;
    margin-left: -100px;
  }
</style>
