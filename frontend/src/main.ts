import { createApp } from 'vue'
import './style.css'
import ElementPlus from 'element-plus'
import App from './App.vue'
import 'element-plus/dist/index.css'
import router from './router'
import store from './store'
import '@/assets/css/iconfont.js';
import "@/assets/css/iconfont.css";
import SvgIcon from '@/components/SvgIcon.vue'
import '@/assets/css/reset.css'
import "virtual:windi.css";
import 'virtual:svg-icons-register'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
const app = createApp(App)

app.use(store)
app.use(router)
app.use(ElementPlus)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }
app.component('SvgIcon', SvgIcon)

app.mount('#app')
