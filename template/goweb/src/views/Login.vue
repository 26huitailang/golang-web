<template>
  <h1>Login</h1>
  <button @click="toHome">toHome</button>
</template>

<script>
import {useRoute, useRouter} from "vue-router"

export default {
  name: "Home",
  setup() {
    const route = useRoute()
    const router = useRouter()
    const toHome = (() => {
      router.push({
        name: "home"
      })
    })
    return { toHome }
  }
}
</script>

<style scoped>

</style>