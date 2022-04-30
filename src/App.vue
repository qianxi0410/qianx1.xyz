<script setup lang="ts">
useMeta({
  title: `${import.meta.env.APP_TITLE}`,
  meta: {
    description: { name: "description", content: "qianxi's blog" },
    keywords: { name: "keywords", content: "blog qianxi" },
    equiv: { "http-equiv": "Content-Type", content: "text/html; charset=UTF-8" },
    noscript: {
      default: "This is content for browsers with no JS (or disabled JS)",
    },
  },
});
</script>

<template>
  <Blossom />
  <q-layout>
    <q-page-container>
      <q-page>
        <Banner />
        <router-view />
        <Footer />
      </q-page>
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
  scroll-behavior: smooth;
}
</style>
