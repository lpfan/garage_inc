<template>
  <div id="app">
    <v-app>
      <v-content>
        <app-header></app-header>
        <v-container fluid>
          <v-layout row wrap>
            <v-flex sm3>
              <app-sidebar></app-sidebar>
            </v-flex>
            <v-flex sm9>
              <router-view></router-view>
            </v-flex>
          </v-layout>
        </v-container>
      </v-content>
    </v-app>
  </div>
</template>

<script>
  import Header from './components/header.vue';
  import Sidebar from './components/sidebar.vue';
  import Content from './components/content.vue';

  export default {
    name: 'app',
    data () {
      return {
        msg: 'Welcome to Your Vue.js App'
      }
    },
    components: {
      AppHeader: Header,
      AppSidebar: Sidebar,
      AppContent: Content
    }
  }
</script>

<style lang="scss" scoped>
</style>
