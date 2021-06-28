<template>
  <h1>Home {{ name }}</h1>
</template>

<script>
import {defineComponent} from "vue";

export default defineComponent({
  setup() {
    const name = 'world'

    return {
      name
    }
  }
})
</script>

<style scoped>

</style>