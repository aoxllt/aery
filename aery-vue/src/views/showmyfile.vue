<template>
  <nav style="display: flex; justify-content: space-between; padding: 10px; background-color: #f0f0f0; box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);">
    <button style="background-color: #4CAF50; color: white; border: none; padding: 10px 20px; border-radius: 5px; cursor: pointer; transition: background-color 0.3s;">
      <router-link to="/upload" style="text-decoration: none; color: white;">返回</router-link>
    </button>
  </nav>
  <div class="file-explorer">
    <h2 style="color: #1a1a1a; font-family: 'Microsoft Sans Serif', sans-serif;">
      <strong>个人文件展示</strong>
    </h2>
    <div class="file-tree">
      <ul class="tree-list">
        <li v-for="node in treeData" :key="node.key">
          <TreeItem
              :node="node"
              @updateSelection="updateSelectedFiles"
          />
        </li>
      </ul>
    </div>
    <div class="action-buttons">
      <div class="button-group">
        <button class="download-btn" @click="handleDownloadSelected">下载</button>
        <button class="delete-btn" @click="handleDeleteSelected">删除</button>
      </div>
    </div>
    <div style="color: #1a1a1a" v-if="isDeleting" class="progress-container">
      <span>正在删除：{{ deletefile }} {{ deleteProgress }}%</span>
      <div class="progress-bar" :style="{ width: `${deleteProgress}%` }">
      </div>
    </div>
  </div>
</template>


<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import TreeItem from '../components/upload/TreeItem.vue'; // 确保路径正确
import { initOSSClient, getFileList, deleteFile, downloadFile } from '../utils/oss.ts';
import router from "../router";
import {getSession} from "../utils/getcookie.ts";
import {ElMessage} from "element-plus"; // 假设有这些函数

interface TreeNode {
  title: string;
  key: string;
  size?: number;
  children?: TreeNode[];
  checked?: boolean; // 选中状态
}

const warnMessage=ref('')
const successMessage=ref('')
const isDeleting = ref(false);
const deleteProgress = ref(0);
const deletefile = ref("");
const treeData = ref<TreeNode[]>([]);
const selectedFiles = ref<string[]>([]); // 存储选中的文件

const createTreeData = (files: any[]) => {
  const root: { [key: string]: any } = {};

  files.forEach(file => {
    if (file.name && typeof file.name === 'string') {
      const pathParts = file.name.split('/');
      let currentLevel = root;

      pathParts.forEach((part, index) => {
        if (!currentLevel[part]) {
          currentLevel[part] = {
            title: part,
            key: pathParts.slice(0, index + 1).join('/'),
            children: [],
            isLeaf: index === pathParts.length - 1,
            size: index === pathParts.length - 1 ? file.size : undefined,
            isFile: index === pathParts.length - 1 // 叶子节点标志
          };
        }
        currentLevel = currentLevel[part].children;
      });
    } else {
      console.warn(`Invalid file object: ${JSON.stringify(file)}`);
    }
  });

  const convertToTree = (obj: any) => {
    return Object.values(obj).map((item: any) => ({
      title: item.title,
      key: item.key,
      size: item.size,
      children: item.isLeaf ? undefined : convertToTree(item.children),
      isFile: item.isFile // 传递标志位
    }));
  };

  return convertToTree(root).length > 0 ? convertToTree(root)[0].children : [];
};
const updateSelectedFiles = (node: TreeNode) => {
  if (node.checked) {
    selectedFiles.value.push(node.key);
  } else {
    selectedFiles.value = selectedFiles.value.filter(key => key !== node.key);
  }
};

function findPathInTree(data, path) {
  for (const node of data) {
    // 检查当前节点的 key
    if (node.key === path && node.isFile === true) {
      return node; // 返回找到的文件节点
    }
    // 如果有 children，递归查找
    if (node.children) {
      const found = findPathInTree(node.children, path);
      if (found) {
        return found; // 只返回找到的节点，无论它的 isFile 值
      }
    }
  }
  return null; // 如果没有找到
}

const handleDownloadSelected = async () => {
  for (const key of selectedFiles.value) {
    const filepath = findPathInTree(treeData.value, key);
    if (!filepath) {
      continue;
    }
    try {
      // console.log("Downloading file:", filepath['key']);
      const result = await downloadFile(filepath['key']); // 调用下载逻辑
      // console.log("下载结果:", result);
      if (result && result.content) {
        const blob = new Blob([result.content]); // 创建 Blob 对象
        const url = URL.createObjectURL(blob); // 创建下载链接

        const a = document.createElement('a');
        a.href = url;
        a.download = key.split('/').pop(); // 设定文件名
        document.body.appendChild(a);
        a.click(); // 触发下载
        document.body.removeChild(a);
        URL.revokeObjectURL(url); // 释放链接
      } else {
        warnMessage.value="下载数据无效"
        warning();
        return
      }
    } catch (error) {
      warnMessage.value=`${filepath['key'].split('/')}下载失败`
      warning();
      return
    }
  }
  successMessage.value="全部文件下载完成"
  success()
  return
};



const handleDeleteSelected = async () => {
  isDeleting.value = true; // 开始删除
  deleteProgress.value = 0; // 重置进度

  for (const key of selectedFiles.value) {
    const filepath = findPathInTree(treeData.value, key);
    if (!filepath) {
      continue; // 如果没有找到，跳过当前循环
    }

    try {
      await deleteFile(filepath['key'], (progress) => {
        deleteProgress.value = progress; // 更新进度
        deletefile.value=filepath['key']
      });
      // 更新 UI，例如从 treeData 中移除已删除的节点
    } catch (error) {
      warnMessage.value='删除文件时发生错误'
      warning();
      return
    }
  }

  isDeleting.value = false; // 完成删除
  deleteProgress.value = 0; // 重置进度
  successMessage.value="文件删除完毕"
  success();
  return
};

const success = () => {
  ElMessage({
    message: successMessage.value,
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

onMounted(async () => {
  const res = await getSession();
  if (res.data['value'] === '') {
    warnMessage.value = "非法访问，请登录";
    warning();
    setTimeout(() => {
      router.push('/login');
    }, 1000);
  } else {
    try {
      await initOSSClient();
      const files = await getFileList(res.data['value']); // 获取文件列表
      treeData.value = createTreeData(files); // 处理文件列表为树形结构
      console.log(treeData.value);
    } catch (error) {
      console.error("Error fetching files:", error);
    }
  }
});

</script>

<style scoped>
.file-explorer {
  padding: 20px;
}

.action-buttons {
  margin-bottom: 10px;
  display: flex;
  justify-content: space-between; /* 分散布局 */
}

.button-group {
  display: flex; /* 使下载和删除按钮在一行 */
}

.action-buttons button {
  margin-top: 10px;
  padding: 8px 20px;
  cursor: pointer;
  border: none;
  border-radius: 18px;
  transition: background-color 0.3s ease;
}

.download-btn {
  background-color: #007bff; /* 下载按钮背景色 */
  color: white; /* 字体颜色 */
  margin-right: 10px; /* 增加右侧间距 */
}

.download-btn:hover {
  background-color: #0056b3; /* 悬停效果 */
}

.delete-btn {
  background-color: #dc3545; /* 删除按钮背景色 */
  color: white; /* 字体颜色 */
}

.delete-btn:hover {
  background-color: #c82333; /* 悬停效果 */
}

.file-tree {
  height: 680px;
  overflow-y: auto;
  border: 1px solid #ccc;
  padding: 10px;
  background-color: #f9f9f9;
}

.tree-list {
  list-style: none;
  padding-left: 0;
}

.file-tree li {
  margin: 5px 0;
}

/* 进度条样式 */
.progress-container {
  margin-top: 20px;
  width: 100%;
  background-color: #f0f0f0;
  border-radius: 4px;
  overflow: hidden;
}

.progress-bar {
  height: 20px;
  background-color: #4caf50;
  width: 0;
  transition: width 0.4s ease; /* 添加动画效果 */
  border-radius: 4px 0 0 4px;
  text-align: center;
  color: white;
}
</style>
