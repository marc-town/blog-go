import Vue from 'vue'
import Router from 'vue-router'

// import components
import HelloWorld from '@/components/HelloWorld'

Vue.use(Router)

const routes = [
  { path: '/', component: HelloWorld }
]

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: routes
})