import os from 'os';
import fs from 'fs';

// 固定的配置文件路径和适配器名称
const configFilePath = '../config/tsconfig.json';  // 替换为你的路径
const adapterName = 'WLAN';  // 替换为你想要的适配器名称

// 获取指定适配器的 IP 地址
function getIpAddressByAdapter() {
    const interfaces = os.networkInterfaces();
    const adapterInfo = interfaces[adapterName];

    if (!adapterInfo) {
        throw new Error(`未找到适配器 "${adapterName}"`);
    }

    for (const interfaceInfo of adapterInfo) {
        // 过滤掉 IPv6 和内部回环地址，只保留 IPv4
        if (interfaceInfo.family === 'IPv4' && !interfaceInfo.internal) {
            return interfaceInfo.address;
        }
    }

    throw new Error(`适配器 "${adapterName}" 没有找到有效的 IPv4 地址`);
}

// 更新配置文件中的 zurl 的 IP 地址
export function updateConfigIpAddress() {
    const newIpAddress = getIpAddressByAdapter();

    // 读取现有的 config.json 文件
    fs.readFile(configFilePath, 'utf8', (err, data) => {
        if (err) {
            console.error('读取配置文件时出错:', err);
            return;
        }

        // 解析 JSON
        let config = JSON.parse(data);

        // 修改 zurl 的 IP 地址
        const urlParts = new URL(config.zurl);
        urlParts.hostname = newIpAddress; // 更新 IP 地址
        config.zurl = urlParts.toString(); // 更新 zurl

        // 将修改后的 JSON 对象重新写回 config.json
        fs.writeFile(configFilePath, JSON.stringify(config, null, 4), 'utf8', (err) => {
            if (err) {
                console.error('写入配置文件时出错:', err);
                return;
            }
            console.log(`成功更新 zurl 为: ${config.zurl}`);
        });
    });
}
updateConfigIpAddress()