<template>
  <div class="container">
    <div class="login_form">
      <h2>欢迎登录</h2>
      <div class="el-input-group">
        <label>账 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;号:</label>
        <el-input
            v-model="login_from.user_name"
            class="input-field"
            placeholder="请输入账号"
            clearable
        />
      </div>
      <div class="el-input-group">
        <label>密 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;码:</label>
        <el-input
            v-model="login_from.password"
            class="input-field"
            type="password"
            placeholder="请输入密码"
            show-password
        />
      </div>
      <div class="el-input-group">
        <label>验 证 码:</label>
        <el-input
            v-model="login_from.captcha"
            class="captcha-input"
            placeholder="验证码"
            clearable
        />
        <img :src="item" class="captcha-image" alt="Captcha Image" @click="receiveCaptcha" />
        <router-link to="/login/forgot" class="forgot-password">忘记密码</router-link>
      </div>
      <div class="button-container">
        <button class="login_btn" @click="submit">登录</button>
        <div class="register">
          <router-link to="/login/register">没有账号？点击注册</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from "vue";

const login_from = reactive({
  user_name: '',
  password: '',
  captcha: '',
});
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
  margin-bottom: 63px; /* 增加标题底部间距 */
  margin-top: 35px;
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
  flex: 0 0 100px; /* 固定标签宽度为100px */
}

.input-field {
  flex: 1; /* 输入框自适应宽度 */
  min-width: 220px; /* 设置输入框的最小宽度 */
  max-width: 300px; /* 设置输入框的最大宽度 */
}

.captcha-input {
  width: 110px; /* 设置验证码输入框的固定宽度为110px */
  margin-right: 15px; /* 增加验证码输入框与图片之间的间距 */
}

.captcha-image {
  cursor: pointer; /* 鼠标指针效果 */
  margin-right: 20px;
}

.login_btn {
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

.login_btn:hover {
  background-color: #66b1ff; /* 鼠标悬停背景色 */
}

.forgot-password {
  font-size: 12px; /* 字体大小 */
  color: #409eff; /* 颜色 */
  text-decoration: none; /* 去掉下划线 */
}

.forgot-password:hover {
  text-decoration: underline; /* 悬停时显示下划线 */
}

.button-container {
  margin-top: 25.5%;
  display: flex; /* 使用 flexbox */
  flex-direction: column; /* 垂直排列子元素 */
  justify-content: flex-end; /* 向下对齐 */
  margin-bottom: 12px; /* 底部间距 */
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
