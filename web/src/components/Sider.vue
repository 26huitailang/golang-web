<template>
  <n-layout-sider
    bordered
    collapse-mode="width"
    :collapsed-width="64"
    :width="240"
    show-trigger
    :native-scrollbar="false"
  >
    <n-menu
      :collapsed-width="64"
      :collapsed-icon-size="22"
      :options="menuOptions"
      @update:value="handleUpdateValue"
    />
  </n-layout-sider>
</template>

<script>
import {defineComponent, h, ref} from 'vue'
import {NIcon} from 'naive-ui'
import {BookOutline as BookIcon, PersonOutline as PersonIcon, WineOutline as WineIcon} from '@vicons/ionicons5'
import {RouterLink} from "vue-router";
import router from "../router";

function renderIcon(icon) {
  return () => h(NIcon, null, {default: () => h(icon)})
}

const menuOptions = [
  {
    label: 'Home',
    path: '/home',
    key: 'home',
    icon: renderIcon(BookIcon)
  },
  {
    label: 'Theme',
    path: '/theme',
    key: 'theme',
    icon: renderIcon(BookIcon)
  },
  {
    label: 'Suite',
    path: '/suite',
    key: 'suite',
    icon: renderIcon(BookIcon)
  },
  {
    label: 'Tasks',
    path: '/tasks',
    key: 'tasks',
    disabled: false,
    icon: renderIcon(BookIcon)
  },
  {
    label: '舞，舞，舞',
    key: 'dance-dance-dance',
    icon: renderIcon(BookIcon),
    children: [
      {
        type: 'group',
        label: '人物',
        key: 'people',
        children: [
          {
            label: '叙事者',
            key: 'narrator',
            icon: renderIcon(PersonIcon)
          },
          {
            label: '羊男',
            key: 'sheep-man',
            icon: renderIcon(PersonIcon)
          }
        ]
      },
      {
        label: '饮品',
        key: 'beverage',
        icon: renderIcon(WineIcon),
        children: [
          {
            label: '威士忌',
            key: 'whisky'
          }
        ]
      },
      {
        label: '食物',
        key: 'food',
        children: [
          {
            label: '三明治',
            key: 'sandwich'
          }
        ]
      },
      {
        label: '过去增多，未来减少',
        key: 'the-past-increases-the-future-recedes'
      }
    ]
  }
]

export default defineComponent({
  setup() {
    return {
      collapsed: ref(false),
      menuOptions,
      handleUpdateValue(key, item) {
        router.push(item.path)
      }
    }
  }
})
</script>

<style>
.router-link {
  text-decoration: none;
  color: black;
}

.router-link-active {
  text-decoration: none;
  color: black;
}
</style>
