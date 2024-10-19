<template>
  <div class="container">
    <div class="register_form">
      <h2>注册账号</h2>
      <div class="message-container">
      <div v-if="usernameError" :class="{'error-message': true, 'success-message': isUsernameAvailable}" class="username-error">
        {{ usernameError || ' ' }}
      </div>
    </div>
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
            v-model="register_from.password"
            class="input-field"
            type="password"
            placeholder="请输入密码"
            show-password
        />
      </div>
      <div class="el-input-group">
        <label>确认密码:</label>
        <el-input
            v-model="register_from.repassword"
            class="input-field"
            type="password"
            placeholder="再次输入密码"
            show-password
        />
      </div>
      <div class="el-input-group">
        <label>邮 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;箱:</label>
        <el-input
            v-model="register_from.email"
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
        <button class="captcha_btn" :disabled="isButtonDisabled" @click="get_captcha">{{ buttonLabel }}</button>
      </div>
      <button class="register_btn" @click="submit">注册</button>
      <div class="register">
        <router-link to="/login">已有账号？点击登录</router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import url from'../../config/tsconfig.json'
import {h, reactive, ref, watch} from "vue";
import {ElMessage, ElNotification} from "element-plus";
import { debounce } from 'lodash-es';
import axios from "axios";
import router from "../../router";

let isCheckingUsername = false;
let usernameError = ref('');
let isUsernameAvailable = ref(false);
const warnMessage = ref('');
const errorMessage = ref('');
const isButtonDisabled = ref(false); // 按钮禁用状态
const buttonLabel = ref('获取验证码'); // 按钮显示文本
let countdown = 60; // 倒计时
let register_from = reactive({
  user_name: '',
  password: '',
  repassword: '',
  email: '',
  captcha: '',
});

function get_captcha() {
  const options = {
    method: "get",
    url: `${url.zurl}/registercaptcha?email=${encodeURIComponent(register_from.email)}`,
    withCredentials: true,
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json; charset=utf-8',
    },
  }
  axios(options).then(res=>{
    // console.log(res.data.data);
    if(res.data.data==null){
      errorMessage.value='错误，请输入邮箱';
      get_err();
      return
    }
    if (res.data.data['status']===false){
      if(res.data.data['message']==='验证码发送失败'){
        warnMessage.value=res.data.data['message'];
        warning();
      }else {
        errorMessage.value='错误，请稍后重新尝试获取';
        get_err();
      }
    }else {
      startCountdown()
    }
  })
}

function username_change(){
  if (register_from.user_name.trim().length === 0) {
    usernameError.value = '';
    isUsernameAvailable.value = false;
    return;  // 如果账号为空，直接返回
  }

  if(register_from.user_name.length<4){
    usernameError.value = '该账号不可用，账号最少四位数。';
    isUsernameAvailable.value = false;
    return;
  }

  if (isCheckingUsername) {
    return;
  }

  isCheckingUsername = true;

  const options = {
    method: 'get',
    url: url.zurl + '/checkusername?username='+encodeURIComponent(register_from.user_name),
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json; charset=utf-8',
    },
  };

  axios(options).then(res => {
    // console.log(res.data.data['status']);
    if (res.data.data['status'] === true) {
      usernameError.value = '该账号已存在，请输入其他账号。';
      isUsernameAvailable.value = false;
      // console.log(usernameError.value); // 打印错误消息
    } else {
      usernameError.value = '该账号可用';
      isUsernameAvailable.value = true;
      // console.log(usernameError.value); // 打印可用消息
    }
  }).catch(() => {
    usernameError.value = '检查账号时出错，请稍后再试。';
    isUsernameAvailable.value = false;
    // console.error('请求失败:', err.response ? err.response.data : err.message);
  }).finally(() => {
    isCheckingUsername = false;
  });
}

function submit() {
  if(register_from.user_name===''){
    open1();
    return;
  }

  if(register_from.password===''){
    open2();
    return;
  }

  if(register_from.password!==register_from.repassword){
    open4();
    return;
  }

  if(register_from.password.length<6){
    open3();
    return;
  }

  if(register_from.user_name.length<3){
    open7();
    return;
  }

  if(register_from.email===''){
    open5();
    return;
  }

  if (register_from.email) {
    const input = register_from.email.trim();
    const emailPattern = /^[\w.-]+@[a-zA-Z\d.-]+\.[a-zA-Z]{2,6}$/;
    if (!emailPattern.test(input)){
      open6();
      return;
    }
  }

  const options={
    method:"POST",
    url:url.zurl+"/register",
    withCredentials:true,
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json; charset=utf-8',
    },
    data:register_from,
  }
  axios(options).then(res=>{
    // console.log(res.data.data);
    if (res.data.data['status']===true){
      success()
      setTimeout(() => {
        router.push('/login');
      }, 1000)
    } else if(res.data.data['status']===false){
      if (res.data.data['message'].substring(0, 2) === '错误'){
        errorMessage.value='服务器出问题，请稍后重试';
        err()
      }else {
        warnMessage.value=res.data.data['message'];
        warning()
      }
    }else {
      errorMessage.value='请稍后重试'
      err()
    }
  })
}

function startCountdown() {
  isButtonDisabled.value = true;
  countdown = 60;
  buttonLabel.value = `${countdown}秒后重新获取`;

  const timer = setInterval(() => {
    countdown--;
    if (countdown > 0) {
      buttonLabel.value = `${countdown}秒后重新获取`;
    } else {
      clearInterval(timer);
      isButtonDisabled.value = false;
      buttonLabel.value = '重新获取';
    }
  }, 1000);
}

const success = () => {
  ElMessage({
    message: '注册成功',
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
    message: '注册失败，'+errorMessage.value,
    type: 'error',
    plain: true,
  })
}
const get_err = () => {
  ElMessage({
    message: errorMessage.value,
    type: 'error',
    plain: true,
  })
}

const open1 = () => {
  ElNotification({
    title: '错误，请重新输入',
    message: h('i', { style: 'color: red' }, '账号不能为空'),
  })
}

const open2 = () => {
  ElNotification({
    title: '错误，请重新输入',
    message: h('i', { style: 'color: red' }, '密码不能为空'),
  })
}

const open3 = () => {
  ElNotification({
    title: '错误，请重新输入',
    message: h('i', { style: 'color: red' }, '密码必须大于6位'),
  })
}

const open4 = () => {
  ElNotification({
    title: '错误，请重新输入',
    message: h('i', { style: 'color: red' }, '两次输入的密码不一致'),
  })
}

const open5 = () => {
  ElNotification({
    title: '错误，请重新输入',
    message: h('i', { style: 'color: red' }, '邮箱不能为空'),
  })
}

const open6 = () => {
  ElNotification({
    title: '错误，请重新输入',
    message: h('i', { style: 'color: red' }, '请输入正确的邮箱'),
  })
}

const open7 = () => {
  ElNotification({
    title: '错误，请重新输入',
    message: h('i', { style: 'color: red' }, '账号最少为4位'),
  });
}

const debouncedUsernameChange = debounce(username_change, 300);
watch(() => register_from.user_name,debouncedUsernameChange)

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
  margin-bottom: 10px; /* 增加标题底部间距 */
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

.captcha_btn {
  background-color: #409eff; /* 按钮默认背景颜色 */
  color: white; /* 按钮文字颜色 */
  border: none; /* 去掉边框 */
  cursor: pointer; /* 鼠标悬停时显示为手型 */
  border-radius: 4px; /* 圆角 */
  padding: 8px;
  font-size: 12px; /* 字体大小 */
  transition: background-color 0.3s ease; /* 添加平滑过渡效果 */
}

.captcha_btn:hover {
  background-color: #66b1ff; /* 悬停时的背景颜色 */
}

.captcha_btn:disabled {
  background-color: #dcdcdc; /* 禁用状态下的背景颜色 */
  color: #ffffff; /* 禁用状态下的文字颜色 */
  cursor: not-allowed; /* 禁用时的鼠标指针 */
}

.captcha_btn:disabled:hover {
  background-color: #dcdcdc; /* 禁用时悬停颜色保持不变 */
}

.message-container {
  top: -20px; /* 调整为负值，使其出现在标题下方 */
  left: 0;
  right: 0;
  height: 20px; /* 固定高度以防止影响布局 */
  overflow: hidden;
}

.error-message {
  font-size: 12px;
  color: red;
  margin-right: 25%;
}

.success-message {
  font-size: 12px;
  color: green;
  margin-right: 49%;
}

</style>
