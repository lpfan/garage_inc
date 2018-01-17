import Vue from 'vue'
import App from './App.vue'
import VueResource from 'vue-resource'
import Vuex from 'vuex'
import VueRouter from 'vue-router'
import Vuetify from 'vuetify'
import { routes } from './routes'
import { store } from './store/store'
import "../node_modules/vuetify/dist/vuetify.min.css"
import "../node_modules/mdi/css/materialdesignicons.css"

Vue.use(Vuetify)
Vue.use(VueRouter)
Vue.use(VueResource)

const router = new VueRouter({
  routes
})

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
