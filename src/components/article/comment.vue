<script setup lang="ts">

const props = defineProps<{issue_numer: number}>()
const container_dark = ref<HTMLDivElement>()
const container_light = ref<HTMLDivElement>()

const $q = useQuasar()

onMounted(() => {
  const utterances_dark = document.createElement('script')
  utterances_dark.async = true
  utterances_dark.setAttribute('src', 'https://utteranc.es/client.js')
  utterances_dark.setAttribute('repo', 'qianxi0410/qianx1.xyz')
  utterances_dark.setAttribute('issue-number', props.issue_numer.toString())
  utterances_dark.setAttribute('theme', 'github-dark')
  utterances_dark.setAttribute('crossorigin', 'anonymous')

  container_dark.value!.append(utterances_dark)

  const utterances_light = document.createElement('script')
  utterances_light.async = true
  utterances_light.setAttribute('src', 'https://utteranc.es/client.js')
  utterances_light.setAttribute('repo', 'qianxi0410/qianx1.xyz')
  utterances_light.setAttribute('issue-number', props.issue_numer.toString())
  utterances_light.setAttribute('theme', 'github-light')
  utterances_light.setAttribute('crossorigin', 'anonymous')

  container_light.value!.append(utterances_light)
})

</script>

<template>
  <div v-show="$q.dark.isActive" ref="container_dark" class="col-md-6 col-sm-9 col-xs-10"></div>
  <div v-show="!$q.dark.isActive" ref="container_light" class="col-md-6 col-sm-9 col-xs-10"></div>
</template>
