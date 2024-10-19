<template>
  <el-container class="header">
    <Header/>
    <el-container class="main-container">
      <router-view/>
    </el-container>


    <el-footer><Footer/></el-footer>
  </el-container>
</template>

<script setup lang="ts">

//引入组件
import Header from '../components/index/index_header.vue'

</script>

<style scoped>
.main-container {
  width: 100%;
}
.header {
  height: 60px;
  width: 100%;
  background: #dddddd;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); /* 添加阴影效果 */
  transition: box-shadow 0.3s ease; /* 添加过渡效果 */
}
</style>/