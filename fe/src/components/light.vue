<template>
  <v-layout>
    <v-flex sm2></v-flex>
    <v-flex sm8>
        <div class="garage-box">
            <div class="lamp left-side-lamp top-side-lamp">
                <button v-on:click="tapLamp(1, $event)" class="red darken-1"></button>
            </div>
            <div class="lamp left-side-lamp bottom-side-lamp">
                <button v-on:click="tapLamp(2, $event)" class="red darken-1"></button>
            </div>
            <div class="lamp right-side-lamp top-side-lamp">
                <button v-on:click="tapLamp(3, $event)" class="red darken-1"></button>
            </div>
            <div class="lamp right-side-lamp bottom-side-lamp">
                <button v-on:click="tapLamp(4, $event)" class="red darken-1"></button>
            </div>
            <div class="lamp middle-side-lamp top-side-lamp">
                <button v-on:click="tapLamp(5, $event)" class="red darken-1"></button>
            </div>
            <div class="lamp middle-side-lamp bottom-side-lamp">
                <button v-on:click="tapLamp(6, $event)" class="red darken-1"></button>
            </div>
        </div>
    </v-flex>
  </v-layout>
</template>

<style>
    .garage-box {
        border: 2px solid #000;
        height: 500px;
        position: relative;
        width: 300px;
    }

    .lamp {
        height: 150px;
        margin: 10px 0;
        position: absolute;
        width: 30px;
    }

    .left-side-lamp {
        left: 10px;
    }

    .right-side-lamp {
        right: 10px;
    }

    .lamp button {
        height: 100%;
        width: 100%;
    }

    .middle-side-lamp {
        height: 30px;
        width: 150px;
        margin: 0 0 0 -75px;
        left: 50%;
    }

    .top-side-lamp {
        top: 70px;
    }

    .bottom-side-lamp {
        bottom: 70px;
    }
</style>


<script>
export default {
  data() {
      return {
          lampState: {
            1: false,
            2: false,
            3: false,
            4: false,
            5: false,
            6: false
          }
      }
  },
  methods: {
      tapLamp(lampNum, event) {
          event.preventDefault();
          this.lampState[lampNum] = !this.lampState[lampNum];
          if ( this.lampState[lampNum] == false ) {
              jquery(event.target).removeClass().addClass('red darken-1');
          } else {
              jquery(event.target).removeClass().addClass('light-green darken-1');
          }
      }
  }
}
</script>
