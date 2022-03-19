<script setup lang="ts">
import { Icon } from "@iconify/vue";
const { t, locale } = useI18n();
const $q = useQuasar();

const en = () => {
  locale.value = "en";
  $q.dark.set(true);
};

const zh = () => {
  locale.value = "zh-CN";
  $q.dark.set(false);
};
</script>

<template>
  <img class="avatar center" src="src/assets/avatar.jpg" />

  <div class="center link">
    <p>{{ t("first") }}</p>
    <p>{{ t("second") }}</p>
    <Icon mr-1 icon="bi:github" :horizontal-flip="true" />
    <a href="https://github.com/qianxi0410" target="_blank">Github</a>
    <div mt-3>
      <input id="lang-1" type="radio" name="lang" checked value="en" @click="en" />
      <label for="lang-1">en</label>

      <input id="lang-2" type="radio" name="lang" value="zh-CN" @click="zh" />
      <label for="lang-2">zh-CN</label>
    </div>
  </div>
</template>

<style scoped></style>
