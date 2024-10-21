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
    const OSSClient = new OSS({
        region: "oss-cn-fuzhou",
        accessKeyId: token.AccessKeyId,
        accessKeySecret: token.AccessKeySecret,
        stsToken: token.SecurityToken,
        secure: true,
        refreshSTSTokenInterval: 300000, // 5分钟刷新一次STS token
        bucket: "aery"
    });

    // 将 OSSClient 实例挂载为全局变量
    window.globalOSSClient = OSSClient;

    return OSSClient;
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
        const result = await OSSClient.multipartUpload(objectName, file, options);

        console.log('上传成功：', result);
        return result;
    } catch (error) {
        console.error('上传失败：', error);
        throw error;
    }
};
