<script setup lang="ts">
const { t } = useI18n();
</script>

<template>
  <div>
    <div class="row justify-center text-grey">
      <div class="col-md-6 col-sm-9 col-xs-10">
        <p class="text-h3">{{ t("404.welcome") }}</p>
        <p class="text-body1">{{ t("404.first") }}</p>
        <p class="text-body1">{{ t("404.second") }}</p>
        <p class="text-body1">{{ t("404.third") }}</p>
        <p class="text-body1">{{ t("404.fourth") }}</p>
      </div>
    </div>
  </div>
</template>
