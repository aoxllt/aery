<template>
  <div class="container">
    <!-- 鼠标放置时文字右移，线条动态出现 -->
    <div class="fancy-title">
      {{ title }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const title = ref('Hover Over Me');
</script>

<style scoped>

/* 父容器 */
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
  height: 100px;
  position: relative;
  overflow: hidden; /* 隐藏可能出现的滚动条 */
}

/* 文字部分 */
.fancy-title {
  font-size: 2rem;
  padding: 20px 30px;
  border-radius: 15px;
  display: inline-block;
  position: relative;
  cursor: pointer;
  transition: transform 0.5s ease-out; /* 文字移动的过渡效果 */
  user-select: none; /* 禁止文本选择 */
}

/* 文字悬停时右移效果 */
.fancy-title:hover {
  transform: translateX(40px); /* 文字右移 */
}

/* 伪元素：显示动态出现的竖线 */
.fancy-title::after {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  width: 2px; /* 线条宽度 */
  height: 0; /* 初始高度为0 */
  background-color: #007bff; /* 线条颜色 */
  transition: height 0.5s ease-out; /* 线条动态增长 */
}

/* 悬停时线条动态出现 */
.fancy-title:hover::after {
  height: 100%; /* 线条从上到下动态出现 */
}
</style>