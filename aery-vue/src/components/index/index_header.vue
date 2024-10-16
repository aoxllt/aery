<template>
  <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
      @select="handleSelect"
  >
    <el-menu-item index="0">
      <img
          style="width: 100px"
          src="C:\Users\ss\Desktop\untitled2\aery-vue\src\assets\vue.svg"
          alt="Element logo"
      />
    </el-menu-item>

    <div class="menu-center">
      <el-menu-item index="1">待填</el-menu-item>
      <el-menu-item index="1">待填</el-menu-item>
      <el-menu-item index="1">待填</el-menu-item>
    </div>

    <el-sub-menu index="2" class="expand-menu">
      <template #title>拓展列</template>
      <el-menu-item index="2-1">登录</el-menu-item>
      <el-menu-item index="2-2">待填</el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import router from "../../router";

const activeIndex = ref('1')
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
  if (key === '2-1') {
    router.push('/login') // 跳转到登录页面
  }
}
</script>

<style scoped>
.el-menu-demo{
  height: 100%;
}

.menu-center {
  display: flex;
  flex: 1;
  justify-content: center; /* 居中 */
  user-select: none; /* 禁止文本选择，隐藏闪烁光标 */
}

.expand-menu {
  margin-left: auto; /* 将拓展列推到右边 */
  user-select: none; /* 禁止文本选择，隐藏闪烁光标 */
}
</style>