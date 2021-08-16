<template>
  <n-image-group>
    <n-space>
      <n-image v-for="image in data.images"
        height="120"
        :src="`/image/${image.path}`"
      />
<!--      <n-image-->
<!--        width="100"-->
<!--        src="https://gw.alipayobjects.com/zos/antfincdn/aPkFc8Sj7n/method-draw-image.svg"-->
<!--      />-->
    </n-space>
  </n-image-group>
</template>

<script>
import {defineComponent, h, onMounted, ref} from 'vue'
import {NButton, NImageGroup, NSpace, NImage, useMessage} from 'naive-ui'
import {getSuiteDetail} from "@/api/suite";
import {useRoute, useRouter} from "vue-router";

const createColumns = ({sendMail, handleClick}) => {
  return [
    {
      title: 'Name',
      key: 'name',
      render(row) {
        const link = () => {
          return h(
            NButton,
            {
              text: true,
              style: {},
              tag: 'a',
              type: 'info',
              target: '_blank',
              onClick: handleClick,
            },
            {
              default: () => row.name
            }
          )
        }
        return link()
      }
    },
    {
      title: 'Id',
      key: 'id'
    },
    {
      title: 'CreatedAt',
      key: 'createdAt'
    },
    {
      title: 'UpdatedAt',
      key: 'updatedAt'
    },
    // {
    //   title: 'Tags',
    //   key: 'tags',
    //   render(row) {
    //     const tags = row.tags.map((tagKey) => {
    //       return h(
    //         NTag,
    //         {
    //           style: {
    //             marginRight: '6px'
    //           },
    //           type: 'info'
    //         },
    //         {
    //           default: () => tagKey
    //         }
    //       )
    //     })
    //     return tags
    //   }
    // },
    // {
    //   title: 'Action',
    //   key: 'actions',
    //   render(row) {
    //     return h(
    //       NButton,
    //       {
    //         size: 'small',
    //         onClick: () => sendMail(row)
    //       },
    //       {default: () => 'Send Email'}
    //     )
    //   }
    // }
  ]
}


export default defineComponent({
  setup() {
    const item = ref([])
    const router = useRouter()
    const route = useRoute()
    console.log(route.params)
    const createData = () => {
      getSuiteDetail(route.params.id).then(response => {
        console.log(response.data)
        item.value = response.data
      })
    }
    // const message = (data) => {
    //   console.log(data)
    // }
    const message = useMessage()
    onMounted(createData)
    console.log(item)
    return {
      data: item,
      columns: createColumns({
        sendMail(rowData) {
          message.warning('send mail to ' + rowData.name)
        },
        handleClick(rowData) {
          router.push('/suite/:id')
        }
      }),
      pagination: {
        pageSize: 10
      }
    }
  }
})
</script>
