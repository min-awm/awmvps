import { createApp } from 'vue'
import App from './App.vue'
import router from '@/router'
import "@/assets/css/index.scss";
import { createPinia } from 'pinia'
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";
// import VueSimpleContextMenu from 'vue-simple-context-menu';
// import 'vue-simple-context-menu/dist/vue-simple-context-menu.css';

const app = createApp(App)
// app.component('vue-simple-context-menu', VueSimpleContextMenu);
app.use(router)
app.use(createPinia())
app.use(Toast);
app.mount('#app')
