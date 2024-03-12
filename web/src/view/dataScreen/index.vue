<template>
  <div
    class="container"
    style="max-width: unset;"
  >
    <!-- 数据大屏展示内容区域 -->
    <div
      ref="screen"
      class="screen"
    >
      <!-- 数据大屏顶部 -->
      <div class="top">
        <Top />
      </div>
      <div class="bottom">
        <div class="left">
          <College class="college" />
          <Sex class="sex" />
          <Rank class="rank" />
        </div>
        <div class="center">
          <Number class="number" />
          <Map class="map" />
          <AD class="ad" />
        </div>
        <div class="right">
          <PerTask class="per-task" />
          <DayTask class="day-task" />
          <HotTask class="hot-task" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
// 引入顶部的子组件
import Top from '@/view/dataScreen/components/top/index.vue'
// 引入左侧三个子组件
import College from '@/view/dataScreen/components/college/index.vue'
import Sex from '@/view/dataScreen/components/sex/index.vue'
import Rank from '@/view/dataScreen/components/rank/index.vue'

// 引入中间两个子组件
import Number from '@/view/dataScreen/components/number/index.vue'
import Map from '@/view/dataScreen/components/map/index.vue'
import AD from '@/view/dataScreen/components/ad/index.vue'

// 引入右侧三个子组件
import PerTask from '@/view/dataScreen/components/perTask/index.vue'
import DayTask from '@/view/dataScreen/components/dayTask/index.vue'
import HotTask from '@/view/dataScreen/components/hotTask/index.vue'
// 获取数据大屏展示内容盒子的DOM元素
const screen = ref()
onMounted(() => {
  screen.value.style.transform = `scale(${getScale()}) translate(-50%,-50%)`
})
// 定义大屏缩放比例
function getScale(w = 1920, h = 1100) {
  const ww = window.innerWidth / w
  const wh = window.innerHeight / h
  return ww < wh ? ww : wh
}
// 监听视口变化
window.onresize = () => {
  screen.value.style.transform = `scale(${getScale()}) translate(-50%,-50%)`
}

</script>

<style scoped lang="scss">
.container {
    width: 100vw;
    height: 100vh;

    .screen {
        position: fixed;
        width: 1920px;
        height: 1100px;
        left: 50%;
        top: 50%;
        transform-origin: left top;

        .top {
            width: 100%;
            height: 40px;
        }

        .bottom {
            display: flex;

            .right {
                flex: 1;
                display: flex;
                flex-direction: column;
                margin-left: 40px;

                .per-task {
                    flex: 2;
                }

                .day-task {
                    flex: 1;

                }

                .hot-task {
                    flex: 1;
                }
            }

            .left {
                flex: 1;
                height: 1040px;
                display: flex;
                flex-direction: column;

                .college {
                    flex: 1.2;
                }

                .sex {
                    flex: 1;

                }

                .rank {
                    flex: 1;
                }
            }

            .center {
                flex: 1.5;
                display: flex;
                flex-direction: column;

                .number {
                    flex: 1;
                }

                .map {
                    flex: 3;
                }

                .ad {
                    flex: 1;
                }
            }
        }
    }
}
</style>
