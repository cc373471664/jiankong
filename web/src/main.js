import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
//axios配置
import {fetch,post} from './api/http.js'
Vue.prototype.$post =post;
Vue.prototype.$fetch = fetch;
//ant配置
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
Vue.use(Antd);
//引入iconfont图标
import './assets/css/iconfont/iconfont.css'

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
