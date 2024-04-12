<template>
  <div class="top">
    <div class="left">
      <Decoration8 />
    </div>
    <div class="center">
      <Decoration7 class="title">
        <el-text
          tag="b"
          style="font-size: 40px;font-family: Microsoft Yahei;cursor: pointer;"
          @click="goback()"
        >&nbsp;&nbsp;&nbsp;&nbsp;E启新篇 迎新数据实况&nbsp;&nbsp;&nbsp;&nbsp;
        </el-text>
      </Decoration7>
    </div>
    <div class="right">
      <Decoration8 :reverse="true" />
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const goback = () => {
  router.go(-1)
}

router.afterEach(() => {
  location.reload()
})
</script>

<style scoped lang="scss">
  .top {
      display: flex;

      .left {
          flex: 1.5;
      }

      .center {
          flex: 2;

          .title {
              height: 60px;
              text-align: center;
              line-height: 60px;
          }
      }

      .right {
          flex: 1.5;
      }
  }
</style>
