<template>
  <div class="container">
    <router-link to="/login"><img class="return" src="../../assets/返回.svg" alt="返回"></router-link>
    <div class="forgot_form">
      <h2>忘记密码</h2>
      <div class="el-input-group">
        <label>账 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;号:</label>
        <el-input
            v-model="forgot.user_name"
            class="input-field"
            placeholder="请输入账号"
            clearable
        />
      </div>
      <div class="el-input-group">
        <label>验 证 器:</label>
        <el-input
            v-model="forgot.email"
            class="input-field"
            placeholder="请输入绑定的邮箱"
            clearable
        />
      </div>
      <div class="el-input-group">
        <label>验 证 码:</label>
        <el-input
            v-model="forgot.captcha"
            class="captcha-input"
            placeholder="请输入验证码"
            clearable
        />
        <button class="captcha_btn" :disabled="isButtonDisabled" @click="get_captcha">{{ buttonLabel }}</button>
      </div>
    </div>
    <button style="align-items: center" class="forgot_btn" @click="submit">提交</button>
  </div>
</template>

<script setup lang="ts">
import {h, reactive, ref} from "vue";
import url from "../../config/tsconfig.json";
import axios from "axios";
import {ElMessage, ElNotification} from "element-plus";
import router from "../../router";

const isButtonDisabled = ref(false); // 按钮禁用状态
const buttonLabel = ref('获取验证码'); // 按钮显示文本
let countdown = 60; // 倒计时
const warnMessage = ref('');
const errorMessage = ref('');
const forgot = reactive({
  user_name: "",
  email: "",
  captcha: ""
});

function submit() {
  if(forgot.user_name===''){
    open1();
    return;
  }
  const options={
    method:"post",
    url:url.zurl+"/forgotpassword",
    headers:{
      'Accept': 'application/json',
      'Content-Type': 'application/json; charset=utf-8',
    },
    withCredentials:true,
    data:forgot
  }
  axios(options).then(res=>{
    console.log(res);
    if (res.data.data['status']===true){
      success();
      setTimeout(() => {
        router.push('/login/change?username='+forgot.user_name);
      }, 1000)
    }else {
      warnMessage.value=res.data.data['message'];
      warning();
      return;
    }
  })
}

function get_captcha() {
  const options = {
    method: "get",
    url: url.zurl+"/registercaptcha?email="+encodeURIComponent(forgot.email),
    withCredentials: true,
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json; charset=utf-8',
    },
  }
  axios(options).then(res=>{
     // console.log(res.data);
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

const get_err = () => {
  ElMessage({
    message: errorMessage.value,
    type: 'error',
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

const open1 = () => {
  ElNotification({
    title: '错误，请重新输入',
    message: h('i', { style: 'color: red' }, '账号不能为空'),
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
    message: '验证成功',
    type: 'success',
    plain: true,
  })
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
  align-items: flex-start; /* 使子元素靠左对齐 */
  box-sizing: border-box; /* 包括内边距和边框在内的宽高计算 */
}

.forgot_form{
  display: flex;
  flex-direction: column;
  align-items: flex-start; /* 使子元素靠左对齐 */
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
  flex: 0 0 100px; /* 固定标签宽度为100px */
}

.input-field {
  flex: 1; /* 输入框自适应宽度 */
  min-width: 220px; /* 设置输入框的最小宽度 */
  max-width: 300px; /* 设置输入框的最大宽度 */
}

.captcha-input {
  width: 130px; /* 设置验证码输入框的固定宽度为110px */
  margin-right: 5px; /* 增加验证码输入框与按钮之间的间距 */
}

/* 提交按钮 */
.forgot_btn {
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

.forgot_btn:hover {
  background-color: #66b1ff; /* 鼠标悬停背景色 */
}

.return{
  width: 30px;
  height: 40px;
  margin-bottom: 40px;

}
.captcha_btn {
  width: 130px;
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
</style>
