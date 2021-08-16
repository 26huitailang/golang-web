import {createApp} from 'vue'
import App from './App.vue'
import router from "./router";
import {
  create,
  NAnchor,
  NAnchorLink,
  NButton,
  NDataTable,
  NH2,
  NImage,
  NImageGroup,
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
// global css
import '@/styles/index.scss'
import {createStore} from "vuex";

const naive = create({
  components: [
    NButton, NLayout, NLayoutHeader, NLayoutFooter, NLayoutSider, NH2,
    NSpace, NAnchorLink, NSwitch, NAnchor, NMenu, NDataTable, NMessageProvider, NImageGroup, NImage,
  ]
})

const store = createStore({
  state() {
    return {
      count: 0
    }
  },
  mutations: {
    increment(state) {
      state.count++
    }
  }
})

const app = createApp(App)
app.use(store)
app.use(naive)
app.use(router)
app.mount('#app')
