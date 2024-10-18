<template>
  <el-container>
    <el-header><Header/></el-header>
    <el-aside style="user-select: none;"><Aside/></el-aside>
    <el-container>
      <el-aside class="ChatRoom">
        <ChatRoom/>
      </el-aside>
      <el-main><router-view/></el-main>
    </el-container>


    <el-footer><Footer/></el-footer>
  </el-container>
</template>

<script setup lang="ts">
import URL from '../config/tsconfig.json'

//引入组件
import ChatRoom from '../components/room/chat_room.vue'
import Header from '../components/index/index_header.vue'
import Aside from '../components/index/index_aside.vue'
import Footer from '../components/index/index_footer.vue'
import {onMounted} from "vue";
import axios from "axios";

</script>

<style scoped>
</style>/