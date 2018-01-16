import LightControls from './components/light.vue'
import Startpage from './components/content.vue'
import ClimateControls from './components/climate.vue'

export const routes = [
    { path: '', component: Startpage},
    { path: '/light', component: LightControls},
    { path: '/climate', component: ClimateControls}
]