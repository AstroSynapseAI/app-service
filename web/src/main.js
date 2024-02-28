import { createApp } from 'vue';
import { createPinia } from 'pinia';
import router from './router'

import './style.css'
import App from './App.vue'

import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)
app.use(Toast)
app.mount('#app')
