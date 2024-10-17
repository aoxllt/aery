<template>
  <!-- 图标按钮（人形） -->
  <el-button
      v-if="!isChatRoomVisible"
      class="person-icon-btn"
      @click="toggleChatRoom"
  >
  </el-button>

  <!-- ChatRoom 组件 -->
  <div
      v-if="isChatRoomVisible"
      class="custom-component"
      ref="chatRoom"
  >
    <!-- 顶部，包含标题和关闭按钮 -->
    <div class="header" @mousedown="startDrag">
      <h1>Chat Room</h1>
      <el-button
          icon="el-icon-close"
          class="close-btn"
          @click="toggleChatRoom"
      />
    </div>

    <!-- 中间的消息显示区 -->
    <div class="message-area">
      <div v-for="(msg, index) in messages" :key="index" class="message">
        <p>{{ msg }}</p>
      </div>
    </div>

    <!-- 底部的输入框和按钮 -->
    <div class="input-area">
      <el-input v-model="inputText" placeholder="输入你的信息" class="input-box" />
      <el-button type="primary" @click="sendMessage">发送</el-button>
      <el-upload action="#" :show-file-list="false" class="upload-btn">
        <el-button>上传文件</el-button>
      </el-upload>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// 控制 ChatRoom 的显示和隐藏
const isChatRoomVisible = ref(false)

// 用于存储输入框内容
const inputText = ref('')

// 存储聊天消息的列表
const messages = ref([
  '客户：你好，有什么可以帮忙的吗？',
  '你：您好，我有一个问题需要咨询。',
])

// 切换 ChatRoom 显示状态的函数
const toggleChatRoom = () => {
  isChatRoomVisible.value = !isChatRoomVisible.value
}

// 发送消息函数
const sendMessage = () => {
  if (inputText.value.trim() !== '') {
    messages.value.push(`你：${inputText.value}`)
    inputText.value = '' // 清空输入框
  }
}

// 拖动功能实现，仅允许拖动头部
const chatRoom = ref<HTMLElement | null>(null)

const startDrag = (event: MouseEvent) => {
  const element = chatRoom.value
  if (!element) return

  const startX = event.clientX
  const startY = event.clientY
  const initialLeft = element.offsetLeft
  const initialTop = element.offsetTop

  const onMouseMove = (moveEvent: MouseEvent) => {
    const currentX = moveEvent.clientX
    const currentY = moveEvent.clientY
    const deltaX = currentX - startX
    const deltaY = currentY - startY

    element.style.left = `${initialLeft + deltaX}px`
    element.style.top = `${initialTop + deltaY}px`
  }

  const onMouseUp = () => {
    document.removeEventListener('mousemove', onMouseMove)
    document.removeEventListener('mouseup', onMouseUp)
  }

  document.addEventListener('mousemove', onMouseMove)
  document.addEventListener('mouseup', onMouseUp)
}
</script>

<style scoped>
/* 人形图标按钮样式 */
.person-icon-btn {
  position: fixed;
  bottom: 20px;
  right: 20px;
  background-color: #007bff;
  color: white;
  border-radius: 50%;
  width: 50px;
  height: 50px;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 24px;
}

/* ChatRoom 组件整体样式 */
.custom-component {
  height: 500px;
  width: 350px;
  background-color: #f0f0f0;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  border: 1px solid #ccc;
  position: fixed;
  bottom: 20px;
  right: 20px;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

/* 顶部样式，仅允许拖动头部 */
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  background-color: #007bff;
  color: white;
  border-radius: 10px 10px 0 0;
  user-select: none;
  cursor: move; /* 指示可以拖动 */
}

.close-btn {
  color: white;
  font-size: 20px;
  background: transparent;
  border: none;
}

/* 中间消息区样式 */
.message-area {
  flex: 1;
  padding: 10px;
  background-color: white;
  border: 1px solid #ddd;
  overflow-y: auto; /* 添加滚动条 */
  margin-bottom: 10px;
}

.message {
  margin-bottom: 10px;
  padding: 5px;
  border: 1px solid #ddd;
  background-color: #a9e898;
  border-radius: 5px;
}

/* 底部输入区域 */
.input-area {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px;
  border-top: 1px solid #ddd;
}

.input-box {
  flex: 1;
}

.upload-btn {
  margin-left: 10px;
}
</style>