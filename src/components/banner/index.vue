<script lang="ts" setup>
import {
  evaGithubOutline,
  evaMoonOutline,
  evaPaperPlaneOutline,
  evaSunOutline,
  evaTwitterOutline,
} from "@quasar/extras/eva-icons";
</script>

<template>
  <q-banner :class="$q.dark.isActive ? 'bg-dark' : ''">
    <q-btn flat color="grey" :icon="evaPaperPlaneOutline" to="/" label="cd ~" no-caps />
    <template #action>
      <q-btn flat color="grey" :ripple="false" label="blog" to="blog" />
      <q-btn flat color="grey" :ripple="false" label="next" to="next" />
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
        :icon="evaTwitterOutline"
        :ripple="false"
        color="grey"
        flat
        href="https://twitter.com/qianxi0410"
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
