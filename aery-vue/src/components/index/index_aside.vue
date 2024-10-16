<template>
  <div
      class="sidebar"
      :class="{ 'active': sidebarVisible }"
      @click="hideSidebar"
  >
    <p v-if="sidebarVisible">Sidebar Content</p> <!-- 仅在侧边栏可见时显示文字 -->
  </div>
  <div class="content" v-if="!sidebarVisible" @click="showSidebar">
    <div class="fancy-title">{{ tile }}</div> <!-- 仅在侧边栏不可见时显示文本 -->
  </div>
</template>

<script setup lang="ts">
import {ref} from 'vue';

const tile = ref('Main Content');

const sidebarVisible = ref(false);

// 控制侧边栏的显示
const showSidebar = () => {
  sidebarVisible.value = true;
};

const hideSidebar = () => {
  sidebarVisible.value = false;
};
</script>

<style scoped>

.fancy-title {
  position: relative;
  font-size: 2rem;
  padding: 20px 30px;
  border-radius: 15px;
  display: inline-block;
  overflow: hidden; /* 隐藏溢出 */
  user-select: none; /* 禁止文本选择，隐藏闪烁光标 */
}

.sidebar {
  width: 0; /* 初始宽度为0 */
  overflow: hidden;
  transition: width 0.5s ease-out, opacity 0.5s ease-out, box-shadow 0.5s ease; /* 添加减缓效果 */
  background: #403737; /* 背景设置为黑色 */
  color: white; /* 文字颜色为白色 */
  white-space: nowrap; /* 防止换行 */
  opacity: 0; /* 初始透明度为0 */
  display: flex; /* 使用 flexbox 居中内容 */
  align-items: center; /* 垂直居中 */
  justify-content: center; /* 水平居中 */
  border-radius: 15px; /* 圆角效果 */
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.5); /* 添加阴影 */
}

.sidebar.active {
  width: 300px; /* 展开时的宽度 */
  height: 500px; /* 可以根据需要调整 */
  opacity: 1; /* 展开时透明度为1 */
}

.content {
  padding: 20px;
}
</style>