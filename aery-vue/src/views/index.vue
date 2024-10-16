<template>
  <div class="common-layout">
    <el-container style="height: 100vh; width: 100vw; position: relative;">
      <video autoplay muted loop style="position: absolute; top: 0; left: 0; width: 100%; height: 100%; object-fit: cover; z-index: -1;">
        <source src="C:\Users\ss\Downloads\11490-230853032_medium (1).mp4" type="video/mp4">
        Your browser does not support the video tag.
      </video>
    </el-container>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.common-layout {
  overflow: hidden;
  user-select: none; /* 禁止文本选择，隐藏闪烁光标 */
}
</style>