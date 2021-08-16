import {createRouter, createWebHashHistory} from "vue-router"

const routes = [
  {path: "/", redirect: "/home"},
  {
    path: "/login",
    name: "Login",
    component: () => import("@/views/Login.vue")
  },
  {
    path: "/home",
    name: "Home",
    component: () => import('@/views/Home.vue'),
  },
  {
    path: "/theme",
    name: "Theme",
    component: () => import('@/views/Theme.vue'),
  },
  {
    path: "/suite",
    name: "Suite",
    component: () => import('@/views/Suite.vue'),
    props: route => ({themeId: route.query.themeId})
  },
  {
    path: "/suite/:id",
    name: "SuiteDetail",
    component: () => import('@/views/SuiteDetail.vue')
  },
  {
    path: "/tasks",
    name: "Tasks",
    component: () => import('@/views/Task.vue')
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
})

export default router
