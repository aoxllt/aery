<template>
  <div class="common-layout">
    <el-container style="height: 100vh; width: 100vw; position: relative;">
      <video
          ref="videoElement"
          autoplay
          muted
          style="position: absolute; top: 0; left: 0; width: 100%; height: 100%; object-fit: cover; z-index: -1;"
      >
        <source src="C:\Users\ss\Downloads\11490-230853032_medium (1).mp4" type="video/mp4">
        Your browser does not support the video tag.
      </video>
      <div class="button-container" v-if="showButton">
        <el-button
            class="fade-button"
            @click="navigateToOtherPage"
            :class="{ show: buttonVisible }"
        >
          Go
        </el-button>
      </div>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const showButton = ref(false)
const buttonVisible = ref(false)
const router = useRouter()

onMounted(() => {
  // 在视频开始后延迟3秒显示按钮
  setTimeout(() => {
    showButton.value = true // 显示按钮
    buttonVisible.value = true // 使按钮渐显
  }, 3000); // 延迟3秒
})

const navigateToOtherPage = () => {
  router.push('/index') // 跳转到其他页面
}
</script>

<style scoped>
.common-layout {
  overflow: hidden;
  user-select: none; /* 禁止文本选择，隐藏闪烁光标 */
}

.button-container {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%); /* 居中 */
}

.fade-button {
  background-color: #e19b60; /* 黄色按钮 */
  border-radius: 20px; /* 去棱角 */
  height: 80px;
  width: 160px;
  transition: opacity 1s; /* 渐显效果 */
  opacity: 0; /* 初始透明度 */
  transform: scale(0); /* 初始缩放 */
  animation: scaleUp 1s forwards, pulse 1.5s infinite; /* 添加放大和脉动动画 */
}

.fade-button.show {
  opacity: 1; /* 按钮出现时的透明度 */
  transform: scale(1); /* 还原到正常大小 */
}

@keyframes scaleUp {
  0% {
    transform: scale(0); /* 从小到大 */
    opacity: 0; /* 初始透明度为0 */
  }
  100% {
    transform: scale(1); /* 最终大小 */
    opacity: 1; /* 最终透明度为1 */
  }
}

@keyframes pulse {
  0%, 100% {
    transform: scale(1); /* 正常大小 */
  }
  50% {
    transform: scale(1.05); /* 放大5% */
  }
}
</style>