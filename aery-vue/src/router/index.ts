import { createRouter, createWebHistory } from 'vue-router';
import index from '../views/index.vue';
import login from '../views/login.vue';
import login_main from '../components/login_main.vue';
import register from '../components/register_main.vue'
import register_phone from '../components/register_phone.vue';
import register_email from '../components/register_email.vue';
import forgot from '../components/forgot_password.vue'
import change_password from '../components/change_password.vue';
import test from '../views/test.vue'
const routes = [
    {
        path: '/',
        name: 'Home',
        component: index,
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
                children: [
                    {
                        path: '',
                        name: 'Register_phone',
                        component: register_phone,
                    },
                    {
                        path: 'register_email',
                        name: 'Register_email',
                        component: register_email,
                    }
                ]
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
