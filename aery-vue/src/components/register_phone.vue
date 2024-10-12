<template>
  <label>电 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;话:</label>
  <el-input
      v-model="phoneEmail"
      style="width: 240px"
      placeholder="请输入电话"
      clearable
  />
  <router-link to="/login/register/register_email" class="email">邮箱登录</router-link>
</template>

<script setup lang="ts">
import { defineProps, defineEmits, computed } from 'vue';

const props = defineProps({
  phone_email: String  // 接收来自父组件的 phone_email 作为 prop
});

const emit = defineEmits(['update:phone_email']);

// 双向绑定，通过 computed 实现 v-model
const phoneEmail = computed({
  get: () => props.phone_email,
  set: (value) => emit('update:phone_email', value),
});
</script>

<style scoped>
label {
  margin-right: 10px; /* 标签和输入框之间的间距 */
  font-weight: bold; /* 标签加粗 */
  color: #333; /* 标签颜色 */
  flex: 0 0 100px; /* 固定标签宽度为100px */
}
.email{
  margin-left: 10px;
}
</style>