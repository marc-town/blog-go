import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

/* Bootstrap使用宣言 */
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
Vue.use(BootstrapVue)

/* ElementUI使用宣言 */
import ElementUI from 'element-ui';
import locale from 'element-ui/lib/locale/lang/ja'
import 'element-ui/lib/theme-chalk/index.css';
import 'element-ui/lib/theme-chalk/display.css'
Vue.use(ElementUI, { locale })

/* markdown-editor使用宣言 */
import MavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
Vue.use(MavonEditor)

Vue.config.productionTip = false

// axios setting
const axiosBase = require("axios")
const customAxios = axiosBase.create({
  baseURL: process.env.VUE_APP_API_URL,
  headers: {
    'Content-Type': 'application/json'
  },
  responseType: 'json',
  responseEncoding: 'utf8'
})
Vue.prototype.$axios = customAxios

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
