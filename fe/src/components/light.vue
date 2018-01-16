<template>
  <v-layout>
    <v-flex sm2></v-flex>
    <v-flex sm8>
        <div class="garage-box">
            <div class="lamp left-side-lamp top-side-lamp">
            </div>
            <div class="lamp left-side-lamp bottom-side-lamp">
            </div>
            <div class="lamp right-side-lamp top-side-lamp">
            </div>
            <div class="lamp right-side-lamp bottom-side-lamp">
            </div>
            <div class="lamp middle-side-lamp top-side-lamp">
            </div>
            <div class="lamp middle-side-lamp bottom-side-lamp">
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
        border: 1px solid #000;
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
  
}
</script>
