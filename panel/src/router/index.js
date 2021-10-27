import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import store from '../store/index';

const routes = [
    {
        path: '/',
        redirect: '/login'
    },
    {
        path: '/login',
        name: 'Login',
        component: Login
    },
    {
        path: '/dashboard',
        name: 'Dashboard',
        component: () => import(/* webpackChunkName: "dashboard" */ '../views/Dashboard.vue'),
        meta: { authenticated: true }
    },
    {
        path: '/moderate',
        name: 'Moderate',
        component: () => import(/* webpackChunkName: "moderate" */ '../views/Moderate.vue'),
        meta: { authenticated: true }
    }
]

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
