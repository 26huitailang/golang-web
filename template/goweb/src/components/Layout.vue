<template>
  <n-space vertical size="large">
    <n-layout style="height: 700px">
      <n-layout-header>颐和园路</n-layout-header>
      <n-layout position="absolute" style="top: 64px; bottom: 64px;" has-sider>
        <Sider></Sider>
        <n-layout content-style="padding: 24px" :native-scrollbar="false">
          <p>平山道</p>
          <p>平山道</p>
          <p>平山道</p>
        </n-layout>
      </n-layout>
      <n-layout-footer>成府路</n-layout-footer>
    </n-layout>
  </n-space>
</template>
<script>
import Sider from "./Sider.vue";

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