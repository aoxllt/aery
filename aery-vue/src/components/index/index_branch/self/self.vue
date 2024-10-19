<template>
  <div class="profile-page">
    <!-- 顶部封面图片部分 -->
    <div class="cover-container">
      <img class="cover-image" src="../../../../assets/img/mushroom-8313142_1280.jpg" alt="Cover" />
      <a-button class="upload-cover-btn">上传封面图片</a-button>
      <span class="ip-status">IP 属地未知</span>
    </div>

    <!-- 头像和个人信息部分 -->
    <div class="user-info-container">
      <div class="avatar-container">
        <el-avatar :size="100" src="C:\Users\ss\Desktop\aery\aery-vue\src\assets\vue.svg" class="avatar" />
      </div>
      <div class="user-details">
        <div style="display: flex; align-items: center; margin: 0; line-height: 1.2;">
          <h2 style="color: #1a1a1a; margin: 0;">{{userInfo.nickName}}</h2>
          <div class="gender-icon" :style="{ color: gender === 1 ? '#007bff' : '#ff69b4', marginLeft: '10px' }">
            <el-icon v-if="gender === 1">
              <Male />
            </el-icon>
            <el-icon v-else-if="gender === 0">
              <Female />
            </el-icon>
          </div>
        </div>
        <h1 style="color: #1a1a1a; margin: 0; line-height: 1.2;">账号: {{userInfo.account}}</h1>
        <div style="display: flex; align-items: center; margin: 0; line-height: 1.2;">
          <h1 style="color: #1a1a1a; margin: 0;">个人简介: {{ userInfo.brief }}</h1>
          <el-icon style="color: #1a1a1a" @click="editBrief()">
            <EditPen />
          </el-icon>
        </div>
        <div class="detail-link" @click="show">查看详细资料</div>
      </div>
      <el-button text @click="dialog = true"
      >编辑用户信息</el-button
      >
    </div>

    <!-- 底部信息选项卡 -->
    <el-tabs v-model="activeTab" class="tabs-container">
      <el-tab-pane label="动态" name="activity">我的动态</el-tab-pane>
      <el-tab-pane label="回答" name="answers">我的回答</el-tab-pane>
      <el-tab-pane label="视频" name="videos">我的视频</el-tab-pane>
      <el-tab-pane label="提问" name="questions">我的提问</el-tab-pane>
      <el-tab-pane label="文章" name="articles">我的文章</el-tab-pane>
      <el-tab-pane label="专栏" name="columns">我的专栏</el-tab-pane>
      <el-tab-pane label="想法" name="thoughts">我的想法</el-tab-pane>
      <el-tab-pane label="收藏" name="favorites">我的收藏</el-tab-pane>
      <el-tab-pane label="关注" name="following">我关注的人</el-tab-pane>
    </el-tabs>

    <!-- 右侧附加内容 -->
    <div class="extra-content">
      <el-card>
        <p>开启你的创作之旅</p>
      </el-card>
    </div>

    <!-- 详细资料弹窗 -->
    <el-dialog v-model="showDetails" title="详细资料">
      <p>昵称: {{ userInfo.nickName }}</p>
      <p>id: {{ userInfo.userId }}</p>
      <p>账号: {{ userInfo.account}}</p>
      <div>
        <p v-if="userInfo.sex ===1">性别: 男</p>
        <p v-else>性别: 女</p>
      </div>
      <p>年龄: {{ userInfo.age}}</p>
      <p>生日: {{ userInfo.birth}}</p>
      <p>邮箱: {{ userInfo.email }}</p>
      <p>手机号: {{ userInfo.phone }}</p>
      <p>加入事件: {{ userInfo.createDate }}</p>
      <p>ip: {{ userInfo.currentAddress }}</p>
      <p>其他: {{ userInfo.info }}</p>
      <el-button type="primary" @click="showDetails = false">关闭</el-button>
    </el-dialog>
  </div>

<!--  编辑-->
  <el-drawer
      v-model="dialog"
      title="编辑"
      :before-close="handleClose"
      direction="ltr"
      class="demo-drawer"
  >
    <div class="demo-drawer__content">
      <el-form :model="form">
        <el-form-item label="邮箱" :label-width="formLabelWidth">
          <el-input v-model="form.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="Area" :label-width="formLabelWidth">
          <el-select
              v-model="form.region"
              placeholder="Please select activity area"
          >
            <el-option label="Area1" value="shanghai" />
            <el-option label="Area2" value="beijing" />
          </el-select>
        </el-form-item>
      </el-form>
      <div class="demo-drawer__footer">
        <el-button @click="cancelForm">Cancel</el-button>
        <el-button type="primary" :loading="loading" @click="onClick">
          {{ loading ? 'Submitting ...' : 'Submit' }}
        </el-button>
      </div>
    </div>
  </el-drawer>
</template>

<script setup>
import {onMounted, reactive, ref} from 'vue'
import { ElDrawer, ElMessageBox } from 'element-plus'
import { Female, Male , EditPen } from '@element-plus/icons-vue';
import axios from "axios";
import url from "../../../../config/tsconfig.json"

const formLabelWidth = '80px'
let timer

const gender = ref(); // 默认选择男性，1为男性，0为女性
const activeTab = ref('activity');
const showDetails = ref(false); // 控制详细信息的显示
const dialog = ref(false)
const loading = ref(false)

const form = reactive({
  name: '',
  region: '',
  date1: '',
  date2: '',
  delivery: false,
  type: [],
  resource: '',
  desc: '',
})

const show = () =>{
  showDetails.value = true;
}

// 示例用户信息（可以替换为真实数据）
const userInfo = ref({});

const onClick = () => {
  loading.value = true
  setTimeout(() => {
    loading.value = false
    dialog.value = false
  }, 400)
}

const handleClose = (done) => {
  if (loading.value) {
    return
  }
  ElMessageBox.confirm('Do you want to submit?')
      .then(() => {
        loading.value = true
        timer = setTimeout(() => {
          done()
          // 动画关闭需要一定的时间
          setTimeout(() => {
            loading.value = false
          }, 400)
        }, 2000)
      })
      .catch(() => {
        // catch error
      })
}

const cancelForm = () => {
  loading.value = false
  dialog.value = false
  clearTimeout(timer)
}

const editBrief = () => {
}

//获取用户信息
const fetchUserInfo = async (id) => {
  const response = await axios.get(`http://10.22.72.209:4567/user/${id}`);
  userInfo.value = response.data.data;
  gender.value = userInfo.value.sex;
}

onMounted(()=>{
fetchUserInfo(1);
})

</script>

<style scoped>
/* 外层容器使整个页面内容居中 */
.profile-page {
  display: flex;
  flex-direction: column;
  align-items: center; /* 水平居中 */
  justify-content: flex-start; /* 垂直方向从顶部开始 */
  width: 100%;
  padding-top: 20px; /* 添加顶部空间，让内容不贴着浏览器顶部 */
  user-select: none;
}

/* 封面图片 */
.cover-container {
  position: relative;
  width: 100%;
  height: 200px;
  background-color: #f2f2f2;
  user-select: none;
}

.cover-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  user-select: none;
}

.upload-cover-btn {
  position: absolute;
  top: 20px;
  right: 20px;
  user-select: none;
}

.ip-status {
  position: absolute;
  bottom: 10px;
  right: 20px;
  color: gray;
}

/* 用户信息部分 */
.user-info-container {
  display: flex;
  justify-content: center;
  margin-top: -50px;
  padding: 20px;
  background-color: white;
  width: 80%;
  border: 1px solid #ebebeb;
  border-radius: 8px;
  position: relative;
}

/* 头像突出效果 */
.avatar-container {
  position: absolute;
  top: -50px;
  left: 20px;
}

.avatar {
  border: 3px solid white; /* 增加头像的边框 */
  border-radius: 50%;
}

.user-details {
  margin-right: 850px; /* 可调整此值以控制用户信息的位置 */
}

.detail-link {
  cursor: pointer; /* 鼠标指针变成手形 */
  display: block;
  color: gray;
}

.edit-profile-btn {
  position: absolute;
  right: 20px;
  top: 20px;
  user-select: none;
}

/* 底部选项卡 */
.tabs-container {
  width: 80%;
  margin-top: 20px;
}

.extra-content {
  width: 20%;
  margin-left: 20px;
}

.el-card {
  margin-top: 20px;
}
</style>