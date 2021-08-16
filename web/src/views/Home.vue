<template>
  <h1>Home {{ name }}</h1>
</template>

<script>
import {defineComponent, onMounted} from "vue";

export default defineComponent({
  setup(props) {
    const name = 'world'
    console.log(props)
    return {
      name
    }
  },
  methods: {
    printThis() {
      console.log('this', this.$store.commit('increment'))
      console.log(this.$store.state.count)
    }
  },
  mounted() {
    console.log('mounted')
    this.printThis()
  }
})
</script>

<style scoped>

</style>
