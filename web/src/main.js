import { createApp } from 'vue';
import { createPinia } from 'pinia';
import router from './router'

import './style.css'
import App from './App.vue'

import Toast, { POSITION } from "vue-toastification";
import "vue-toastification/dist/index.css";

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)
app.use(Toast, {
    position: POSITION.TOP_CENTER,
    timeout: 1500,
    pauseOnHover: false,
    pauseOnLoseFocus: false,
    hideProgressBar: true,
    maxToasts: 1,
    nevestOnTop: false
});

app.mount('#app')
