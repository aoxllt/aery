import OSS from "ali-oss";
import url from '../config/tsconfig.json';

// 获取STS token
export const getSTSToken = async () => {
    const STS_TOKEN_URL = url.zurl + "/sts"; // 获取STS token的接口，后端提供
    const res = await fetch(STS_TOKEN_URL).then(res => res.json());
    const { AccessKeyId, SecurityToken, AccessKeySecret } = res.data.sts.Credentials || {};
    return {
        SecurityToken,
        AccessKeySecret,
        AccessKeyId
    };
};

// 初始化OSS client实例，并挂载为全局变量
export const initOSSClient = async () => {
    const token = await getSTSToken();
    window.globalOSSClient = new OSS({
        region: "oss-cn-fuzhou",
        accessKeyId: token.AccessKeyId,
        accessKeySecret: token.AccessKeySecret,
        stsToken: token.SecurityToken,
        secure: true,
        refreshSTSTokenInterval: 300000,
        bucket: "aery"
    });
    return window.globalOSSClient; // 返回客户端实例
};


// 上传文件并显示进度条的方法
export const OssuploadFile = async (file, objectName, updateProgressCallback) => {
    // 通过已初始化的 OSSClient 实例
    const OSSClient = window.globalOSSClient;

    if (!OSSClient) {
        console.error("OSSClient 未初始化");
        return;
    }

    try {
        const options = {
            // 上传进度回调
            progress: async (percentage) => {
                // 调用外部传入的回调函数更新进度
                if (updateProgressCallback) {
                    updateProgressCallback(Math.floor(percentage * 100));
                }
            },
            // 其他选项，可以设置并发上传数量、断点续传等
            parallel: 5, // 并发上传的分片个数
            partSize: 1024 * 1024, // 每个分片的大小为1MB
        };

        // 调用上传文件方法，file 为文件对象，objectName 为 OSS 完整路径
        return await OSSClient.multipartUpload(objectName, file, options);
    } catch (error) {
        console.error('上传失败：', error);
        throw error;
    }
};
// 获取指定根路径下的文件列表
export const getFileList = async (prefix) => {
    const OSSClient = window.globalOSSClient;

    if (!OSSClient) {
        console.error("OSSClient 未初始化");
        return [];
    }

    try {
        const result = await OSSClient.list({
            prefix,
            'max-keys': 100 // 可根据需要调整每次获取的最大文件数
        });
        return result.objects || [];
    } catch (error) {
        console.error('获取文件列表失败：', error);
        throw error;
    }
};

export const downloadFile = async (objectName) => {
    const OSSClient = window.globalOSSClient;

    if (!OSSClient) {
        console.error("OSSClient 未初始化");
        return null;
    }

    try {
        return await OSSClient.get(objectName); // 返回文件内容
    } catch (error) {
        console.error('下载文件失败：', error);
        throw error;
    }
};

// 删除文件
export const deleteFile = async (objectName, onProgress) => {
    const OSSClient = window.globalOSSClient;

    if (!OSSClient) {
        console.error("OSSClient 未初始化");
        return;
    }

    try {
        // 模拟进度条
        for (let progress = 0; progress <= 100; progress += 20) {
            if (onProgress) {
                onProgress(progress); // 调用进度回调
            }
            await new Promise(resolve => setTimeout(resolve, 100)); // 模拟延时
        }
        return await OSSClient.delete(objectName);
    } catch (error) {
        console.error('删除文件失败：', error);
        throw error;
    }
};
