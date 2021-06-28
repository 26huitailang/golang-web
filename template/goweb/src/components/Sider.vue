<template>
  <n-space vertical>
    <n-space> <n-switch v-model:value="collapsed" />collapsed</n-space>
    <n-layout has-sider>
      <n-layout-sider
          bordered
          collapse-mode="width"
          :collapsed-width="64"
          :width="240"
          show-trigger
          :collapsed="collapsed"
      >
        <n-menu
            :collapsed="collapsed"
            :collapsed-width="64"
            :collapsed-icon-size="22"
            :options="menuOptions"
        />
      </n-layout-sider>
    </n-layout>
  </n-space>
</template>

<script>
import { h, defineComponent, ref } from 'vue'
import { NIcon } from 'naive-ui'
import {
  BookOutline as BookIcon,
  PersonOutline as PersonIcon,
  WineOutline as WineIcon
} from '@vicons/ionicons5'
import {RouterLink} from "vue-router";

function renderIcon (icon) {
  return () => h(NIcon, null, { default: () => h(icon) })
}

const menuOptions = [
  {
    label: () => h(
        RouterLink,
        {
          to: "/home",
        },
        'Home'
    ),
    key: 'home',
    icon: renderIcon(BookIcon)
  },
  {
    label: () => h(
        RouterLink,
        {
          to: "/theme",
        },
        'Theme'
    ),
    key: 'theme',
    icon: renderIcon(BookIcon)
  },
  {
    label: () => h(
        RouterLink,
        {
          to: "/suite",
        },
        'Suite'
    ),
    key: 'suite',
    icon: renderIcon(BookIcon)
  },
  {
    label: '1973年的弹珠玩具',
    key: 'pinball-1973',
    icon: renderIcon(BookIcon),
    disabled: true,
    children: [
      {
        label: '鼠',
        key: 'rat'
      }
    ]
  },
  {
    label: '寻羊冒险记',
    key: 'a-wild-sheep-chase',
    disabled: true,
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
  setup () {
    return {
      collapsed: ref(false),
      menuOptions
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
