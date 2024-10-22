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
import Footer from '../components/index/index_footer.vue'
import { updateConfigIpAddress } from "../utils/getLocalip"; // 确保路径正确
import { onMounted } from "vue";

onMounted(async () => {
  try {
    await updateConfigIpAddress(); // 等待更新完成
    console.log("配置文件已更新");
  } catch (error) {
    console.error("更新配置文件时出错:", error);
  }
});
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