<template>
  <n-data-table :columns="columns" :data="data" :pagination="pagination"/>
</template>

<script>
import {defineComponent, onMounted, ref, h} from 'vue'
import {useMessage, NButton} from 'naive-ui'
import {getSuiteList} from "@/api/suite";
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
              onClick: () => handleClick(row),
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
    const rows = ref([])
    const router = useRouter()
    const route = useRoute()

    const createData = () => {
      getSuiteList({themeId: route.params.themeId}).then(response => {
        console.log(response.data)
        rows.value = response.data
      })
    }
    // const message = (data) => {
    //   console.log(data)
    // }
    const message = useMessage()
    onMounted(createData)
    const handleData = () => {
      rows.value.push({id: 1, name: 'peter'})
    }
    return {
      handleData,
      data: rows,
      columns: createColumns({
        sendMail(rowData) {
          message.warning('send mail to ' + rowData.name)
        },
        handleClick(rowData) {
          console.log('rowData: ', rowData)
          router.push(`/suite/${rowData.id}`)
        }
      }),
      pagination: {
        pageSize: 10
      }
    }
  }
})
</script>
