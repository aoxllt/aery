import { createRouter, createWebHistory } from 'vue-router';
import index from '../views/index.vue';
import login from '../views/login.vue';
import login_main from '../components/login_main.vue';

const routes = [
    {
        path: '/',
        name: 'Home',
        component: index,
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
        ]
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
