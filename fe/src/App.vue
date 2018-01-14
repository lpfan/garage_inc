<template>
  <div id="app">
    <v-app>
      <v-content>
        <v-toolbar dark color="primary">
          <v-toolbar-side-icon></v-toolbar-side-icon>
          <v-toolbar-title class="white--text">Garage Inc.</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn icon>
            <v-icon>search</v-icon>
          </v-btn>
          <v-btn icon>
            <v-icon>apps</v-icon>
          </v-btn>
          <v-btn icon>
            <v-icon>refresh</v-icon>
          </v-btn>
          <v-btn icon>
            <v-icon>more_vert</v-icon>
          </v-btn>
        </v-toolbar>
        <v-container fluid>
          <v-layout row wrap>
            <v-flex sm4>
              Sidebar
            </v-flex>
            <v-flex sm8>
              Content
            </v-flex>
          </v-layout>
        </v-container>
      </v-content>
    </v-app>
  </div>
</template>

<script>
export default {
  name: 'app',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>

<style lang="scss" scoped>
</style>
