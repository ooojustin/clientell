import { createApp } from 'vue';
import App from './App.vue';

import router from './router';
import store from './store';

import VueToast from 'vue-toast-notification';
import 'vue-toast-notification/dist/theme-default.css';

import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap";

createApp(App)
    .use(store)
    .use(router)
    .use(VueToast)
    .mount('#app');
