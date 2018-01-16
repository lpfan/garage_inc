import Vue from 'vue'
import App from './App.vue'

import VueRouter from 'vue-router'
import Vuetify from 'vuetify'

import { routes } from './routes'

Vue.use(Vuetify)
Vue.use(VueRouter)

const router = new VueRouter({
  routes
})

import "../node_modules/vuetify/dist/vuetify.min.css"
import "../node_modules/mdi/css/materialdesignicons.css"

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
