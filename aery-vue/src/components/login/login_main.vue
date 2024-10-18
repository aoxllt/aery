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
        <img :src="captchaImg" class="captcha-image" alt="Captcha Image" @click="get_captcha" />
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
import url from '../../config/tsconfig.json'
import {onMounted, reactive, ref,h } from "vue";
import axios from "axios";
import {ElMessage, ElNotification} from 'element-plus'
import router from "../../router";

const captchaImg = ref('')
const warnMessage = ref('')
const errorMessage = ref('')
let login_from = reactive({
  user_name: '',
  password: '',
  captcha: '',
})


function get_captcha() {
  const options = {
    method:"get",
    url:url.zurl+"/captcha",
    withCredentials: true,
  }
  axios(options).then(res=>{
    console.log(res);
    captchaImg.value = res.data.data['captcha'];
  })
}

function submit(){
  if (login_from.user_name === '') {
    open1(); // 显示警告弹窗
    return;
  }
  if (login_from.user_name.length<4) {
    open4();
    return;
  }
  if (login_from.password === '') {
    open2();
    return;
  }

  if (login_from.password.length < 6) {
    open3();
    return;
  }
  const options = {
    method: "post",
    url: url.zurl+"/login",
    withCredentials: true,
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json; charset=utf-8',
    },
    data: login_from,
  }
  axios(options).then(res=>{
      // console.log(res.data.data);
      if(res.data.data['status']===true){
        success()
        setTimeout(() => {
          router.push('/');
        }, 1000)
      }else if(res.data.data['status']===false){
        if(res.data.data['message']==="用户名或密码错误"){
          errorMessage.value = res.data.data['message'];
          err()
        }else {
          warnMessage.value = res.data.data['message'];
          warning()
        }
      }
    else {
        errorMessage.value ="错误，请稍后重试"
        err()
      }
  })
}

const success = () => {
  ElMessage({
    message: '登录成功',
    type: 'success',
    plain: true,
  })
}
const warning= () => {
  ElMessage({
    message: warnMessage.value,
    type: 'warning',
    plain: true,
  })
}

const err = () => {
  ElMessage({
    message: '登录失败，'+errorMessage.value,
    type: 'error',
    plain: true,
  })
}

const open1 = () => {
  ElNotification({
    title: '错误，请重新输入',
    message: h('i', { style: 'color: red' }, '账号不能为空'),
  });
}
const open4 = () => {
  ElNotification({
    title: '错误，请重新输入',
    message: h('i', { style: 'color: red' }, '账号最少为4位'),
  });
}
const open2 = () => {
  ElNotification({
    title: '错误，请重新输入',
    message: h('i', { style: 'color: red' }, '密码不能为空'),
  });
}

const open3 = () => {
  ElNotification({
    title: '错误，请重新输入',
    message: h('i', { style: 'color: red' }, '密码必须大于6位'),
  });
}

onMounted(() => {
  get_captcha();
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
  flex: 0 0 80px; /* 固定标签宽度为100px */
}

.input-field {
  flex: 1; /* 输入框自适应宽度 */
  min-width: 220px; /* 设置输入框的最小宽度 */
  max-width: 300px; /* 设置输入框的最大宽度 */
}

.captcha-input {
  width: 180px; /* 设置验证码输入框的固定宽度为110px */
  margin-right: 10px; /* 增加验证码输入框与图片之间的间距 */
}

.captcha-image {
  cursor: pointer; /* 鼠标指针效果 */
  margin-right: 5px;
  width: 120px;
  height: 40px;
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
