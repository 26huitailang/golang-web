<template>
  <n-layout style="height: 100%">
    <n-layout-header bordered>颐和园路</n-layout-header>
    <n-layout position="absolute" style="margin-top: 64px" content-style="padding: 24px;" has-sider>
      <Sider></Sider>
      <n-layout content-style="padding: 24px" :native-scrollbar=false>
        <router-view></router-view>
      </n-layout>
    </n-layout>
  </n-layout>
</template>
<script>
import Sider from "../components/Sider.vue";

export default {
  components: {Sider}
}

</script>
<style>
.n-layout-header,
.n-layout-footer {
  background: rgba(128, 128, 128, 0.2);
  padding: 24px;
  height: 64px;
}

.n-layout-sider {
  /*background: rgba(128, 128, 128, 0.3);*/
}

.n-layout-content {
  background: rgba(128, 128, 128, 0.4);
}
</style>
