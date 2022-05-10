<script lang="ts" setup>
import {
  evaGithubOutline,
  evaGlobe2Outline,
  evaMoonOutline,
  evaPaperPlaneOutline,
  evaSunOutline,
} from '@quasar/extras/eva-icons'

const { locale } = useI18n()

const toggleLang = () => {
  locale.value = locale.value === 'en' ? 'zh-CN' : 'en'
}
</script>

<template>
  <q-banner style="background-color: rgba(0, 0, 0, 0)">
    <q-btn
      flat
      color="grey"
      :ripple="false"
      :icon="evaPaperPlaneOutline"
      to="/"
      label="cd ~"
      no-caps
    />
    <template #action>
      <q-btn flat color="grey" :ripple="false" label="blog" to="/blog" />
      <q-btn
        flat
        fab
        :icon="evaGlobe2Outline"
        color="grey"
        :ripple="false"
        @click="toggleLang"
      />
      <q-btn
        :icon="evaGithubOutline"
        color="grey"
        fab
        :ripple="false"
        flat
        href="https://www.github.com/qianxi0410"
      />
      <q-btn
        fab
        :icon="$q.dark.isActive ? evaSunOutline : evaMoonOutline"
        :ripple="false"
        color="grey"
        flat
        @click="$q.dark.toggle()"
      />
    </template>
  </q-banner>
</template>
