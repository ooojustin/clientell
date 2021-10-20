import { createRouter, createWebHistory } from '@ionic/vue-router';
import { RouteRecordRaw } from 'vue-router';
import AuthTabs from '../views/AuthTabs.vue';
import Login from '../views/Login.vue';
import CreateAccount from '../views/CreateAccount.vue';
import store from '../store/index';

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        redirect: '/auth/login'
    },
    {
        path: '/auth/',
        component: AuthTabs,
        children: [
            {
                path: 'login',
                component: Login
            },
            {
                path: 'createAccount',
                component: CreateAccount
            },
        ]
    },
    {
        path: '/tabs/',
        component: () => import('@/views/Tabs.vue'),
        children: [
            {
                path: '',
                redirect: '/tabs/home' 
            },
            {
                path: 'home',
                component: () => import('@/views/Home.vue'),
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
            next("/auth/login");
        else
            next();
        return;
    } 
    next();
});

export default router;
