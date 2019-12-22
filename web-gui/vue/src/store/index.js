import Vue from 'vue'
import Vuex from 'vuex'

// import modules
import articles from './modules/articles/articles'

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    articles,
  }
})
