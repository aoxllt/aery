<template>
  <div class="container">
    <div class="register_form">
      <h2>注册账号</h2>
      <div class="el-input-group">
        <label>账 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;号:</label>
        <el-input
            v-model="register_from.user_name"
            class="input-field"
            placeholder="请输入账号"
            clearable
        />
      </div>
      <div class="el-input-group">
        <label>密 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;码:</label>
        <el-input
            v-model="register_from.passwd"
            class="input-field"
            type="password"
            placeholder="请输入密码"
            show-password
        />
      </div>
      <div class="el-input-group">
        <label>确认密码:</label>
        <el-input
            v-model="register_from.repasswd"
            class="input-field"
            type="password"
            placeholder="再次输入密码"
            show-password
        />
      </div>
      <div class="el-input-group">
        <label>邮 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;箱:</label>
        <el-input
            v-model="phoneEmail"
            style="width: 240px"
            placeholder="请输入邮箱"
            clearable
        />
      </div>
      <div class="el-input-group">
        <label>验 证 码:</label>
        <el-input
            v-model="register_from.captcha"
            class="captcha-input"
            placeholder="请输入验证码"
            clearable
        />
        <button class="btn">重新获取</button>
      </div>
      <button class="register_btn" @click="submit">注册</button>
      <div class="register">
        <router-link to="/login">已有账号？点击登录</router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from "vue";

const register_from = reactive({
  user_name: '',
  passwd: '',
  repasswd: '',
  phone_email: '',
  captcha: '',
});

function submit() {
  console.log(register_from);
}
</script>

<style scoped>
.container {
  width: 400px; /* 固定宽度 */
  height: 500px; /* 固定高度 */
  max-width: 100%; /* 最大宽度为100% */
  margin: 0 auto; /* 居中 */
  padding: 20px 30px; /* 上下20px，左右30px的内边距 */
  border-radius: 10px; /* 增加圆角效果 */
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1); /* 加强阴影效果 */
  background-color: #fff; /* 背景颜色 */
  display: flex; /* 使用 flexbox */
  flex-direction: column; /* 垂直排列子元素 */
  justify-content: center; /* 垂直居中 */
  align-items: center; /* 水平居中 */
  box-sizing: border-box; /* 包括内边距和边框在内的宽高计算 */
}

h2 {
  color: #1a1a1a;
  font-family: "华文黑体", sans-serif; /* 指定字体，加入后备字体 */
  margin-bottom: 30px; /* 增加标题底部间距 */
}

.el-input-group {
  margin-bottom: 20px; /* 输入框之间的间距 */
  display: flex; /* 使 label 和输入框在一行 */
  align-items: center; /* 垂直居中对齐 */
  width: 100%; /* 使输入框组宽度为100% */
}

label {
  font-weight: bold; /* 标签加粗 */
  color: #333; /* 标签颜色 */
  flex: 0 0 80px; /* 固定标签宽度为100px */
}

.input-field {
  flex: 1; /* 输入框自适应宽度 */
  min-width: 220px; /* 设置输入框的最小宽度 */
  max-width: 300px; /* 设置输入框的最大宽度 */
}

.captcha-input {
  width: 110px; /* 设置验证码输入框的固定宽度为110px */
  margin-right: 15px; /* 增加验证码输入框与按钮之间的间距 */
}

.btn {
  font-size: 10px; /* 按钮字体大小 */
  cursor: pointer; /* 鼠标指针效果 */
}

.register_btn {
  width: 100%; /* 按钮宽度 */
  padding: 12px; /* 增加内边距 */
  background-color: #409eff; /* 按钮背景色 */
  color: #fff; /* 按钮文字颜色 */
  border: none; /* 去掉边框 */
  border-radius: 5px; /* 圆角效果 */
  cursor: pointer; /* 鼠标指针效果 */
  transition: background-color 0.3s; /* 背景色过渡效果 */
  font-size: 16px; /* 按钮字体大小 */
}

.register_btn:hover {
  background-color: #66b1ff;  /* 鼠标悬停背景色 */
}

.register {
  margin-top: 20px; /* 上间距 */
  text-align: center; /* 文本居中 */
}

.register a {
  color: #409eff; /* 注册链接颜色 */
  text-decoration: none; /* 去掉下划线 */
}

.register a:hover {
  text-decoration: underline; /* 悬停时显示下划线 */
}
</style>
