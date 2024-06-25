<template>
  <ACard class="m-20 w-96 mx-auto select-none">
    <div class="flex justify-between items-center mb-1">
      <h1 class="text-2xl font-bold  dark:text-white">生成短链接</h1>
      <a-tooltip content="切换生成或还原链接">
        <ASwitch v-model="isGen" type="round" />
      </a-tooltip>
    </div>
      <div>
        <AInput
            v-model="url"
            :placeholder="isGen ? '请输入长链接' : '请输入短链接'"
            class="w-full mb-5 px-4 py-3 rounded-l border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
        />
      </div>
      <div>
        <AButton
            @click="genHandler"
            type="primary"
            html-type="submit"
            class="py-3 px-6 bg-blue-500 text-white rounded-r-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
        >
          {{ isGen ? '生成' : '还原' }}
        </AButton>
      </div>
  </ACard>
</template>

<script setup lang="ts">

const url = ref('');
const isGen = ref(true)
const genHandler = async () => {
  if (url.value == "") {
    AMessage.warning("暂时不支持空字符串!")
    return
  }
  if (isGen.value) {
    const genData = await $fetch('/api/url/generate', {
      method: "GET",
      params: {
        url: url.value
      }
    })
    const context =  "加密后的字符串:" + genData
    AModal.open({
      title: '生成成功',
      content: () => h('div', {}, context),
      draggable: true,  // 设置弹窗可拖动
      okText: "跳转",
      closable: true,
      onOk: () => {
      }
    });
    return
  }else{
    const getLongData = await $fetch('/api/url/getLong', {
      method: "GET",
      params: {
        url: url.value
      }
    })
    const context =  "还原后的字符串:" + getLongData
    AModal.confirm({
      title: '还原生成',
      content: () => h('div', {}, context),
      draggable: true, // 设置弹窗可拖动
      closable: true,
    });
    return
  }
}
const {data: tdkData} = await useAsyncData<any>(
    'tdk',
    () => $fetch('/api/seo/json/home', {
      method: "GET",
    })
) as any
const title = computed(() => tdkData.value?.title)
const keywords = computed(() => tdkData.value?.keywords)
const description = computed(() => tdkData.value?.description)
useSeoMeta({
  title,
  keywords,
  description,
})
</script>

<style scoped>

</style>
