<template>
  <el-footer
      :style="footerStyle"
      @mouseenter="showFooter"
      @mouseleave="hideFooter"
  >
    <h4>detail</h4>Footer
  </el-footer>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';

const footerVisible = ref(false);

// 控制 footer 的显示
const showFooter = () => {
  footerVisible.value = true;
};

const hideFooter = () => {
  footerVisible.value = false;
};

// 动态计算 footer 的样式
const footerStyle = computed(() => ({
  background: 'rgba(255, 255, 255, 0.5)',
  height: footerVisible.value ? '200px' : '50px', // 初始只露出50px
  transition: 'height 0.3s ease',
  overflow: 'hidden',
}));
</script>

<style scoped>
/* 自定义 footer 的基本样式 */
el-footer {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
}
</style>