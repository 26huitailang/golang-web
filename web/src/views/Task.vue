<template>
  <n-data-table :columns="columns" :data="data" :pagination="pagination"/>
</template>

<script>
import {defineComponent, h, onMounted, ref} from 'vue'
import {NButton, NProgress, useMessage} from 'naive-ui'
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
      title: 'Progress',
      key: 'progress',
      render(row) {
        const progress = () => {
          return h(
            NProgress,
            {
              type: 'line',
              status: 'info',
              percentage: row.progress
            }
          )
        }
        return progress()
      }
    },
    {
      title: 'CreatedAt',
      key: 'createdAt'
    },
    {
      title: 'FinishedAt',
      key: 'finishedAt'
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
        let data = response.data
        for (let datum of data) {
          datum['progress'] = 0
        }
        console.debug(data)
        rows.value = data
      })
    }
    // const message = (data) => {
    //   console.log(data)
    // }
    const message = useMessage()
    onMounted(createData)
    const handleData = () => {
      rows.value.push({id: 1, name: 'peter', progress: 0})
    }
    return {
      handleData,
      data: rows,
      columns: createColumns({
        sendMail(rowData) {
          message.warning('send mail to ' + rowData.name)
        },
        handleClick(rowData) {
          // TODO: 进度控制
          console.log('rowData: ', rowData)
          rowData.progress += 10
          console.log(rowData)
          console.log(rows)
        }
      }),
      pagination: {
        pageSize: 10
      }
    }
  }
})
</script>
