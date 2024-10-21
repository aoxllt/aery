<template>
  <div class="tree-node" style="color: #1a1a1a">
    <div class="node-item">
     <div class="content">
       <span class="node-title" @click="toggleChildren">{{ node.title }} ({{ totalSize }} MB)</span>
       <img class="delete" @click="deleteNode" src="../../assets/删除.svg" alt="删除">
     </div>
      <div v-if="node.children && node.children.length && isOpen" class="children-list">
        <show_files
            v-for="child in node.children"
            :key="child.key"
            :node="child"
            @delete="(key) => emit('delete', key)"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { defineProps, defineEmits, ref, computed } from 'vue';
import type { PropType } from 'vue';

interface TreeNode {
  title: string;
  key: string;
  size?: number; // 文件大小，单位：KB
  children?: TreeNode[];
}

const props = defineProps({
  node: Object as PropType<TreeNode>,
});

const isOpen = ref(false);

const totalSize = computed(() => {
  let size = 0;
  const calculateSize = (node: TreeNode) => {
    if (node.size) {
      size += node.size;
    }
    if (node.children) {
      node.children.forEach(calculateSize);
    }
  };
  calculateSize(props.node);
  return (size / 1024/1024).toFixed(2);
});

const toggleChildren = () => {
  isOpen.value = !isOpen.value;
};

const emit = defineEmits(['delete']);

const deleteNode = () => {
  console.log("nowkey:",props.node.key);
  emit('delete', props.node.key); // 直接使用当前节点的 key
};

</script>


<style scoped>
.tree-node {
  padding-left: 20px;
}

.node-item {
  margin: 10px 0;
  padding: 8px;
  border-radius: 4px;
  background-color: #f9f9f9;
  transition: background-color 0.3s;
}

.node-item:hover {
  background-color: #eaeaea;
}

.node-title {
  cursor: pointer;
  font-weight: bold;
  color: #333;
}


.children-list {
  padding-left: 20px;
}
.content {
  display: flex;
  justify-content: space-between; /* 使内容均匀分布 */
  align-items: center; /* 垂直居中对齐 */
}

.delete {
  cursor: pointer; /* 鼠标悬停时显示手指 */
  width: 20px; /* 可以根据需要调整宽度 */
}

</style>
