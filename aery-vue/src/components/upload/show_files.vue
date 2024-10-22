<template>
  <div class="tree-node" style="color: #1a1a1a">
    <div class="node-item">
     <div class="content">
       <span class="icon" v-if="node.children && node.children.length">{{ isOpen ? '📂' : '📁' }}</span>
       <img v-if="!node.children" :src="getIcon(node.title)" alt="file-icon" style="height: 50px;width: 40px" class="file-icon" />
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

const getIcon = (title: string) => {
  const extension = title.split('.').pop()?.toLowerCase();
  switch (extension) {
    case 'pdf':
      return 'src/assets/Pdf.svg'; // 替换为 PDF 图标路径
    case 'doc':
    case 'docx':
      return 'src/assets/doc.svg'; // 替换为 Word 图标路径
    case 'xls':
    case 'xlsx':
      return 'src/assets/xlsx.svg'; // 替换为 Excel 图标路径
    case 'ppt':
    case 'pptx':
      return 'src/assets/pptx.svg'; // 替换为 PowerPoint 图标路径
    case 'jpg':
    case 'jpeg':
      return 'src/assets/Jpg.svg';
    case 'png':
      return 'src/assets/Png.svg';
    default:
      return 'src/assets/sysfile.svg'; // 替换为默认图标路径
  }
};

const toggleChildren = () => {
  isOpen.value = !isOpen.value;
};

const emit = defineEmits(['delete']);

const deleteNode = () => {
  emit('delete', props.node.key); // 直接使用当前节点的 key
};

</script>


<style scoped>
.tree-node {
  padding-left: 20px;
  text-align: left;
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
  align-items: center; /* 垂直居中对齐 */
}

.delete {
  cursor: pointer; /* 鼠标悬停时显示手指 */
  width: 20px; /* 可以根据需要调整宽度 */
}

.icon {
  margin-right: 8px; /* 图标和文本之间的间距 */
  font-size: 1.2em; /* 图标大小 */
}
</style>
