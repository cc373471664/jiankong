import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  // {
  //   path: '/',
  //   name: 'index',
  //   component: () => import(/* webpackChunkName: "about" */ '../views/web/index.vue'),
  //   children:[
  //
  //   ]
  // },
  {
    path: '/urlpei',
    name: 'urlpei',
    component: () => import(/* webpackChunkName: "about" */ '../views/web/urlpei.vue'),
    children:[

    ]
  },


]

const router = new VueRouter({
  routes
})

export default router
