<template>
  <div class="upload-container">
    <h2 style="color: #1a1a1a;font-family: 'Microsoft Sans Serif',sans-serif"><strong>文件信息预览</strong></h2>
    <div class="file-tree">
      <ul class="tree-list">
        <li v-for="node in treeData" :key="node.key">
          <FileNode :node="node" @delete="handleDelete" />
        </li>
      </ul>
    </div>
    <div class="custom-file">
      <input type="file" webkitdirectory multiple @change="handleFileSelect" id="fileInput" style="display: none;"/>
      <label for="fileInput" class="custom-file-upload">
        选择文件
      </label>
      <button @click="clear" class="delete-button">删除所有文件</button>
      <button @click="uploadFile" class="upload-btn">上传</button>
    </div>
    <div class="upload-jindu" v-if="uploadProgress > 0" style="color: #1a1a1a">
      <div class="progress-text">
        <div><strong>正在上传：{{ currentFileName }} 上传进度：{{ uploadProgress }}% 速度：{{ uploadSpeed }}</strong></div>
      </div>
      <div class="progress-bar">
        <div class="progress" :style="{ width: uploadProgress + '%' }"></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref} from 'vue';
import FileNode from '../components/upload/show_files.vue';
import url from '../config/tsconfig.json'
import axios from "axios";
import {initOSSClient, OssuploadFile} from "../utils/oss.ts";
import {ElMessage} from "element-plus";


interface TreeNode {
  title: string;
  key: string;
  size?: number;
  children?: TreeNode[];
}

const uploadSpeed = ref('0 MB/s');
const uploadProgress=ref(0)
const currentFileName = ref('');
const warnMessage=ref('')
const treeData = ref<TreeNode[]>([]);
const filePaths = ref<string[]>([]);

const createTreeData = (files: FileList) => {
  const root: { [key: string]: any } = {};
  Array.from(files).forEach((file) => {
    const pathParts = file.webkitRelativePath.split('/');
    let currentLevel = root;
    pathParts.forEach((part, index) => {
      if (!currentLevel[part]) {
        currentLevel[part] = {
          title: part,
          key: pathParts.slice(0, index + 1).join('/'),
          children: [],
          isLeaf: index === pathParts.length - 1,
          size: index === pathParts.length - 1 ? file.size : undefined,
        };
      }
      currentLevel = currentLevel[part].children;
    });
  });

  const convertToTree = (obj: any) => {
    return Object.values(obj).map((item: any) => ({
      title: item.title,
      key: item.key,
      size: item.size,
      children: item.isLeaf ? undefined : convertToTree(item.children),
    }));
  };

  return convertToTree(root);
};

const clear = () => {
  treeData.value = []; // 设置为一个空数组
  const inputElement = document.querySelector('input[type="file"]') as HTMLInputElement;
  if (inputElement) {
    inputElement.value = ''; // 清空文件输入框
  }
};

const handleFileSelect = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files) {
    const newFiles = createTreeData(target.files);
    filePaths.value = Array.from(target.files).map(file => file.webkitRelativePath);
    // 合并新文件与已有的 treeData
    treeData.value = mergeTreeData(treeData.value, newFiles);
  }
};

// 处理删除事件
const handleDelete = (key: string) => {
  console.log("Deleting key:", key);

  const deleteNode = (nodes: TreeNode[]): TreeNode[] => {
    return nodes.map(node => {
      if (node.key === key) {
        return null; // 标记为删除
      }
      if (node.children) {
        node.children = deleteNode(node.children); // 递归删除
      }
      return node; // 返回未删除的节点
    }).filter(Boolean) as TreeNode[];
  };

  const cleanEmptyFolders = (nodes: TreeNode[]): TreeNode[] => {
    return nodes
        .map(node => {
          if (node.children) {
            node.children = cleanEmptyFolders(node.children); // 递归处理子节点
          }
          return (node.children?.length || 0) > 0 || (node.size && node.size > 0) ? node : null;
        })
        .filter(Boolean) as TreeNode[]; // 过滤掉 null
  };

  // 删除节点并清理空文件夹
  treeData.value = cleanEmptyFolders(deleteNode(treeData.value)); // 更新树数据

  // 如果 treeData 为空，清空文件输入框并允许用户重新上传
  if (treeData.value.length === 0) {
    // 清空输入框
    const inputElement = document.querySelector('input[type="file"]') as HTMLInputElement;
    if (inputElement) {
      inputElement.value = ''; // 清空文件输入框
    }

    // 提示用户可以重新上传文件
    console.log("树数据为空，请重新上传文件。");
  }
};

const mergeTreeData = (existingNodes: TreeNode[], newNodes: TreeNode[]): TreeNode[] => {
  const existingKeys = new Set(existingNodes.map(node => node.key));

  newNodes.forEach(newNode => {
    if (!existingKeys.has(newNode.key)) {
      existingNodes.push(newNode); // 追加新节点
    }
  });

  return existingNodes; // 返回合并后的节点
};

const uploadFile = async () => {
  try {
    // 获取根路径
    const options = {
      method: 'POST',
      url: url.zurl + "/getsession",
      withCredentials: true,
    };

    const res = await axios(options);
    const rootPath = res.data.data.value;

    // 上传文件逻辑
    for (const filePath of filePaths.value) {
      const file = await fetchFileByPath(filePath);

      if (file) {
        currentFileName.value = file.name; // 设置当前文件名
        const fileUploadPath = `${rootPath}/${filePath}`; // 拼凑完整路径

        // 记录开始时间和已上传字节数
        const startTime = Date.now();
        let uploadedBytes = 0;

        // 上传文件并更新进度
        await OssuploadFile(file, fileUploadPath, (progress) => {
          uploadProgress.value = progress; // 更新进度

          // 计算已上传字节数
          uploadedBytes= (progress / 100) * file.size;
          // 计算网速
          const elapsedTime = (Date.now() - startTime) / 1000; // 转换为秒
          const speed = uploadedBytes / elapsedTime; // 字节每秒
          uploadSpeed.value = (speed / (1024 * 1024)).toFixed(2) + ' MB/s'; // 转换为 MB/s
        });
      } else {
        warnMessage.value = `未找到与路径 ${filePath} 对应的文件`;
        warning();
        clear();
      }
    }
  setTimeout(()=>{
    success();
  },1000)
    clear();
    uploadProgress.value=0
  } catch (error) {
    warnMessage.value = '上传失败,请重试';
    warning();
    uploadProgress.value=0
    clear();
  }
};


// 根据路径获取文件
const fetchFileByPath = async (path: string) => {
  const inputElement = document.querySelector('input[type="file"]') as HTMLInputElement;
  if (!inputElement || !inputElement.files) return null;

  for (const file of inputElement.files) {
    if (file.webkitRelativePath === path) {
      return file; // 返回匹配的文件
    }
  }
  return null;
};
// function handleDownload() {
//   if (!window.globalOSSClient) {
//     console.error("OSS Client is not initialized.");
//     return;
//   }
//
//   window.globalOSSClient.get("test/ss.jpg")
//       .then((result) => {
//         const blob = new Blob([result.content]); // 将文件内容转为 Blob
//         const url = window.URL.createObjectURL(blob); // 创建一个 URL
//         const a = document.createElement('a'); // 创建一个链接
//         a.href = url;
//         a.download = "ss.jpg"; // 设置下载文件的名称
//         document.body.appendChild(a); // 添加链接到 DOM
//         a.click(); // 自动点击下载
//         a.remove(); // 下载后移除链接
//         window.URL.revokeObjectURL(url); // 释放 URL 对象
//       })
//       .catch((error) => {
//         console.error("Error downloading file:", error);
//       });
// }

const success = () => {
  ElMessage({
    message: '上传完成，全部文件已上传',
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

onMounted(() => {
  initOSSClient().then(() => {
    // 这里可以直接调用下载函数，如果需要自动下载
    // handleDownload();
  }).catch((error) => {
    console.error("Error initializing OSS Client:", error);
  });
});
</script>
<style scoped>
.upload-container {
  padding: 20px;
}

input[type='file'] {
  margin-top: 20px; /* 保持输入框与树形结构的距离 */
}

.file-tree {
  height: 750px;
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

.custom-file{
  margin-top: 10px;
}
.custom-file-upload {
  display: inline-block;
  padding: 8px 20px;
  margin-right: 10px;
  cursor: pointer;
  background-color: #007BFF; /* 按钮背景色 */
  color: white; /* 字体颜色 */
  border-radius: 20px; /* 圆角 */
  transition: background-color 0.3s ease;
}

.custom-file-upload:hover {
  background-color: #0056b3; /* 悬停效果 */
}

.delete-button {
  padding: 8px 20px;
  cursor: pointer;
  background-color: #dc3545; /* 删除按钮背景色 */
  color: white; /* 字体颜色 */
  border: none; /* 去掉边框 */
  border-radius: 20px; /* 圆角 */
  transition: background-color 0.3s ease;
}

.delete-button:hover {
  background-color: #c82333; /* 悬停效果 */
}

.upload-btn{
  float: right;
  padding: 8px 20px;
  cursor: pointer;
  background-color: #08c10d; /* 删除按钮背景色 */
  color: white; /* 字体颜色 */
  border: none; /* 去掉边框 */
  border-radius: 18px; /* 圆角 */
  transition: background-color 0.3s ease;
}

.upload-btn:hover {
  background-color: #06a826; /* 悬停效果 */
}

.upload-jindu {
  margin: 20px;
  font-family: Arial, sans-serif;
  text-align: center;
}

.progress-text {
  margin-bottom: 5px; /* 文字与进度条之间的间距 */
}

.progress-bar {
  width: 100%;
  background-color: #e0e0e0;
  border-radius: 10px;
  overflow: hidden;
  height: 25px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}
@keyframes smooth-width {
  from {
    width: 0;
  }
  to {
    width: 100%;
  }
}

.progress {
  height: 100%;
  background-color: #4caf50; /* 进度条颜色 */
  animation: smooth-width 0.8s ease-in-out; /* 动画效果 */
}

</style>
