import { createApp } from 'vue';
import './style.css';
import App from './App.vue';
import router from './router';
import Antd from 'ant-design-vue';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import 'ant-design-vue/dist/reset.css';
import { createGlobalConfig } from '@idux/components/config'

const loadIconDynamically = (iconName: string) => {
    return fetch(`/idux-icons/${iconName}.svg`).then(res => res.text())
}

const customConfig = { icon: { loadIconDynamically } }
const globalConfig = createGlobalConfig(customConfig)

const app = createApp(App);
app.use(router);
app.use(ElementPlus);
app.use(Antd);
app.use(globalConfig);
app.mount('#app');
