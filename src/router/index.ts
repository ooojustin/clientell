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
                path: 'settings',
                component: () => import('@/views/Settings.vue'),
                meta: { authenticated: true }
            },
            {
                path: 'search',
                component: () => import('@/views/Search.vue'),
                meta: { authenticated: true },
                children: [
                    {
                        path: '',
                        redirect: '/tabs/search/name'
                    },
                    {
                        path: 'name',
                        component: () => import('@/views/SearchByName.vue')
                    },
                    {
                        path: 'address',
                        component: () => import('@/views/SearchByAddress.vue')
                    }
                ]
            }
        ]
    },
    {
        path: '/createRating/:id',
        component: () => import('@/views/CreateRating.vue'),
        meta: { authenticated: true }
    },
    {
        path: '/createPerson',
        component: () => import('@/views/CreatePerson.vue'),
        meta: { authenticated: true }
    },
    {
        path: '/person/:id',
        component: () => import('@/views/Person.vue'),
        meta: { authenticated: true }
    },
    {
        path: '/ratings',
        component: () => import('@/views/UserRatings.vue'),
        meta: { authenticated: true }
    },
    {
        path: '/faq',
        component: () => import('@/views/FAQ.vue'),
        meta: { authenticated: true }
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
