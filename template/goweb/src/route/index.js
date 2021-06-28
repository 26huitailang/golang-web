import {createRouter, createWebHashHistory} from "vue-router"

const home = () => import("../views/Home.vue")
const login = () => import("../views/Login.vue")
const theme = () => import("../views/Theme.vue")

const routes = [
  {path: "/", redirect: "/home"},
  {
    path: "/home",
    name: "home",
    component: home
  },
  {
    path: "/theme",
    name: "theme",
    component: theme
  },
  {
    path: "/login",
    name: "login",
    component: login
  }
]

export const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
})
