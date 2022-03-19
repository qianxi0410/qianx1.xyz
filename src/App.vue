<script setup lang="ts">
useHead({
  title: `${import.meta.env.APP_TITLE}`,
  meta: [
    { name: "description", content: "qianxi's blog" },
    { name: "keywords", content: "qianxi, blog" },
    {
      property: "og:locale:alternate",
      content: "zh",
      key: "zh",
    },
    {
      property: "og:locale:alternate",
      content: "en",
      key: "en",
    },
  ],
});
</script>

<template>
  <q-layout>
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<style>
@font-face {
  font-family: "Fira Code";
  src: "https://fonts.googleapis.com/css?family=Fira+Code&display=swap";
}

#app {
  font-family: "Fira Code", monospace;
  font-size: 1.2em;
}
</style>
