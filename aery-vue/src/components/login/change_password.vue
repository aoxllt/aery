<template>
  <div class="container">
    <router-link to="/login/forgot"><img class="return" src="../../assets/返回.svg" alt="返回"></router-link>
    <div class="change_form">
      <h2>修改密码</h2>
      <div class="el-input-group">
        <label>密 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;码:</label>
        <el-input
            v-model="change.password"
            class="input-field"
            type="password"
            placeholder="请输入密码"
            show-password
        />
      </div>
      <div class="el-input-group">
        <label>确认密码:</label>
        <el-input
            v-model="change.repassword"
            class="input-field"
            type="password"
            placeholder="再次输入密码"
            show-password
        />
      </div>
    </div>
    <button class="change_btn" @click="submit">提交</button>
  </div>
</template>

<script setup lang="ts">
import {h, onMounted, reactive, ref} from "vue";
import axios from "axios";
import url from '../../config/tsconfig.json'
import router from "../../router";
import {ElMessage, ElNotification} from "element-plus";
import { useRoute } from 'vue-router';

const errorMessage=ref('')
const warnMessage=ref('')
const route = useRoute();  // 获取当前路由
const change = reactive({
  username: ref(route.query.username),
  password: "",
  repassword: ""
});

function checkCookie(){
  const options={
    method:"get",
    url:url.zurl+"/interceptor/changepassword",
    withCredentials: true,
  }
  axios(options).then(res=>{
    if(!res.data.data['status']){
      router.push("/login")
    }
  })
}
function submit() {
  if(change.password===''){
    open2();
    return;
  }

  if(change.password!==change.repassword){
    open4();
    return;
  }

  if(change.password.length<6){
    open3();
    return;
  }
  const options = {
    method:"post",
    url:url.zurl+"/changepassword",
    withCredentials: true,
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json; charset=utf-8',
    },
    data:change,
  }
  axios(options).then(res=>{
    if (res.data.data['status']===false){
      if(res.data.data['message']==="非法访问"){
        errorMessage.value=res.data.data['message'];
        err();
        setTimeout(() => {
          router.push('/login');
        }, 1000)
      }else{
        warnMessage.value=res.data.data['message'];
        warning();
        return;
      }
    }else if(res.data.data['status']===true){
      success();
      setTimeout(() => {
        router.push('/login');
      }, 1000)
    }
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

const success = () => {
  ElMessage({
    message: '修改成功',
    type: 'success',
    plain: true,
  })
}

const err = () => {
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

onMounted(()=>{
  checkCookie();
})
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

.change_form{
  display: flex;
  flex-direction: column;
  align-items: flex-start; /* 使子元素靠左对齐 */
}

h2 {
  color: #1a1a1a;
  font-family: "华文黑体", sans-serif; /* 指定字体，加入后备字体 */
  margin-bottom: 60px; /* 增加标题底部间距 */
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

/* 提交按钮 */
.change_btn {
  margin-top: 50px;
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

.change_btn:hover {
  background-color: #66b1ff; /* 鼠标悬停背景色 */
}

.return{
  width: 30px;
  height: 40px;
  margin-bottom: 40px;

}
</style>
