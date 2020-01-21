import Vue from 'vue'
import Router from 'vue-router'

// import routers
import ArticleRoute from './article-router'
import DraftRoute from './drafts-router'

// import components
import Index from '@/components/pages/Index'

Vue.use(Router)

const routes = [
  { path: '/', component: Index },
  { ...ArticleRoute },
  { ...DraftRoute }
]

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: routes
})