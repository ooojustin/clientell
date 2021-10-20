import { createRouter, createWebHistory } from '@ionic/vue-router';
import { RouteRecordRaw } from 'vue-router';
import Login from '../views/Login.vue';
import store from '../store/index';

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        redirect: '/login'
    },
    {
        path: '/login',
        component: Login
    },
    {
        path: '/tabs/',
        component: () => import('@/views/Tabs.vue'),
        children: [
            {
                path: '',
                redirect: '/tabs/tab1' 
            },
            {
                path: 'tab1',
                component: () => import('@/views/Tab1.vue'),
                meta: { authenticated: true }
            },
            {
                path: 'tab2',
                component: () => import('@/views/Tab2.vue'),
                meta: { authenticated: true }
            },
            {
                path: 'tab3',
                component: () => import('@/views/Tab3.vue'),
                meta: { authenticated: true }
            }
        ]
    }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

router.beforeEach((to, from, next) => {
    if (to.meta.authenticated) {
        if (!store.getters.isAuthenticated) 
            next("/login");
        else
            next();
        return;
    } 
    next();
});

export default router;
