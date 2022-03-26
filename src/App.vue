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

        <Transition name="scale" mode="out-in">
          <router-view />
        </Transition>

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

.scale-enter-active,
.scale-leave-active {
  transition: all 0.5s ease;
}

.scale-enter-from,
.scale-leave-to {
  opacity: 0;
  transform: scale(0.9);
}
</style>
