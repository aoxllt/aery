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
          src="../../../src/assets/vue.svg"
          alt="Element logo"
      />
    </el-menu-item>

    <div class="menu-center">
      <el-menu-item index="1">待填</el-menu-item>
      <el-menu-item index="2">待填</el-menu-item>
      <el-menu-item index="3">待填</el-menu-item>
    </div>

    <el-menu-item index="4" v-if="!isLoggedIn" @click="handleSelect('4')">登录</el-menu-item>

    <el-sub-menu v-if="isLoggedIn" index="5" class="expand-menu">
      <template #title>
        <img src="../../../src/assets/img/mushroom-8313142_1280.jpg" alt="User Avatar" style="width: 30px; border-radius: 50%;" />
      </template>
      <el-menu-item index="5-1">个人信息</el-menu-item>
      <el-menu-item index="5-2">设置</el-menu-item>
      <el-menu-item index="5-3">退出</el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import router from "../../router";

const activeIndex = ref('1')
const isLoggedIn = ref(false) // 假设初始为未登录状态

const handleSelect = (key: string) => {
  console.log(key)
  if(key === '1'){
    router.push('/')
  }
  else if(key === '2'){
    router.push('option1')
  }
  else if(key === '3'){
    router.push('/self')
  }
  else if (key === '4') {
    router.push('/login') // 跳转到登录页面
  }
  // 添加处理其他菜单项的逻辑
}
</script>

<style scoped>
.el-menu-demo {
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