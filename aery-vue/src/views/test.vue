<template>
  <div>
    <img v-if="captcha" :src="captcha" alt="Captcha" />
    <button @click="test">获取验证码</button>
  </div>
</template>

<script setup lang="ts">
import axios from 'axios';
import { ref } from 'vue';

const captcha = ref(''); // 用于存储验证码的 Base64 数据

async function test() {
  const option = {
    method: "get",
    url: "http://localhost:1234/captcha",
    withCredentials: true, // 确保请求携带 cookie
    headers: {
      'Content-Type': 'application/json',
      // 如果你需要额外的 headers 可以在这里添加，比如 Authorization
    },
  };

  try {
    const response = await axios(option); // 使用 await 让代码更简洁
    console.log(response.data.data['captcha']);
    captcha.value = response.data.data['captcha'];
  } catch (error) {
    console.error("请求失败:", error); // 打印错误信息
    captcha.value = null; // 如果请求失败，清空验证码
  }
}

</script>

<style scoped>
img{
  width: 300px;
  height: 100px;
}
</style>
