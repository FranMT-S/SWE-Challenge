import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './style.css'

import { createVfm } from 'vue-final-modal'

import PrimeVue from "primevue/config";
import Button from 'primevue/button';
import Message from 'primevue/message';
import Tooltip from 'primevue/tooltip';
import OverlayPanel from 'primevue/overlaypanel';
import Paginator from 'primevue/paginator';
import StyleClass from 'primevue/styleclass';

import "primevue/resources/themes/lara-light-indigo/theme.css";
import "primeicons/primeicons.css";
import "primevue/resources/primevue.min.css"

import 'vue-final-modal/style.css'



const vfm = createVfm()
const app = createApp(App)

app.use(store).use(router)

app.use(vfm)

app.use(PrimeVue)
.directive('tooltip', Tooltip)
.directive('styleclass', StyleClass)
.component('Message',Message)
.component('Button',Button)
.component('OverlayPanel',OverlayPanel)
.component('Paginator',Paginator)

app.mount('#app')
