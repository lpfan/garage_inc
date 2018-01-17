import Vue from 'vue'
import Vuex from 'vuex'
import VueResource from 'vue-resource'

Vue.use(Vuex)
Vue.use(VueResource)

export const store = new Vuex.Store({
    state: {
        lamps: {
            1: false,
            2: false,
            3: false,
            4: false,
            5: false,
            6: false
        }
    },
    getters: {
        lampState: (state) => (lampNum) => {
            return state.lamps[lampNum];
        }
    },
    mutations: {
        switchLamp(state, lampNum) {
            state.lamps[lampNum] = ! state.lamps[lampNum];
        }
    },
    actions: {
        switchLamp(context, lampNum) {
            context.commit('switchLamp', lampNum);

            Vue.http.put('http://localhost/api/v1/light/turnlamp')
                .then( response => {
                    console.log(response)
                })
        }
    }
})