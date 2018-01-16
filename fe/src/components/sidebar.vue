<template>
  <v-flex>
      <router-link to="/light" tag="v-btn" active-class="blue lighten-3">
        <v-icon>mdi-lightbulb-on-outline</v-icon>
        Light
     </router-link>
     <router-link to="/climate" tag="v-btn" active-class="blue lighten-3">
        <v-icon>mdi-fan</v-icon>
        Climate
     </router-link>
  </v-flex>
</template>

<style lang="sass">

</style>

<script>
export default {
  
}
</script>
