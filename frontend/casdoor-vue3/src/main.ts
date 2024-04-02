import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import Casdoor from 'casdoor-vue-sdk'

const config = {
    serverUrl: "http://192.168.31.64:8000",
    clientId: "2db6dfeba919f743247a",
    organizationName: "game",
    appName: "game-hosting",
    redirectPath: "/callback",
}

const app = createApp(App)

app.use(router)
app.use(Casdoor, config)

app.mount('#app')
