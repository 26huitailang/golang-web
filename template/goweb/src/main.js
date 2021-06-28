import {createApp} from 'vue'
import App from './App.vue'
import {router} from "./route";
import {
  create,
  NAnchor,
  NAnchorLink,
  NButton,
  NDataTable,
  NH2,
  NLayout,
  NLayoutFooter,
  NLayoutHeader,
  NLayoutSider,
  NMenu,
  NMessageProvider,
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
    NSpace, NAnchorLink, NSwitch, NAnchor, NMenu, NDataTable, NMessageProvider,
  ]
})
const app = createApp(App)
app.use(naive)
app.use(router)
app.mount('#app')
