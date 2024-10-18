import { createRouter, createWebHistory } from 'vue-router';
import index from '../views/index.vue';
import login from '../views/login.vue';
import login_main from '../components/login/login_main.vue';
import register from '../components/login/register_main.vue'
import forgot from '../components/login/forgot_password.vue'
import change_password from '../components/login/change_password.vue';
import test from '../views/test.vue'
import index_main from "../components/index/index_branch/index_main.vue";
import option1 from "../components/index/index_branch/option1.vue";

const routes = [
    {
        path: '/',
        name: 'index',
        component: index,
        children: [{
            path: '',
            name: 'main',
            component: index_main,
        },{
            path: 'option1',
            name: 'option1',
            component: option1,
        }]
    },
    {
        path: '/test',
        name: 'test',
        component: test,
    },
    {
        path: '/login',
        name: 'Login',
        component: login,
        children: [
            {
                path: '',  // 默认子路由
                name: 'LoginMain',
                component: login_main,
            },
            {
                path: 'register',
                name: 'RegisterMain',
                component: register,
            },
            {
                path: 'forgot',
                name: 'Forgot',
                component: forgot,
            },
            {
                path: 'change_password',
                name: 'ChangePassword',
                component: change_password,
            }
        ]
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
