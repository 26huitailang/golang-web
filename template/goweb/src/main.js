import {createApp} from 'vue'
import App from './App.vue'
import {
    create,
    NAnchor,
    NAnchorLink,
    NButton,
    NH2,
    NLayout,
    NLayoutFooter,
    NLayoutHeader,
    NLayoutSider,
    NMenu,
    NSpace,
    NSwitch,
} from 'naive-ui'

// 通用字体
import 'vfonts/Lato.css'
// 等宽字体
import 'vfonts/FiraCode.css'

const naive = create({
    components: [
        NButton, NLayout, NLayoutHeader, NLayoutFooter, NLayoutSider, NH2,
        NSpace, NAnchorLink, NSwitch, NAnchor, NMenu,
    ]
})
const app = createApp(App)
app.use(naive)
app.mount('#app')
