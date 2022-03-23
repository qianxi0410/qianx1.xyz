<script setup lang="ts">
const { t } = useI18n();
</script>

<template>
  <!-- avatar -->
  <div class="row justify-center">
    <div class="col col-auto self-center">
      <q-avatar size="11vw" color="grey" class="shadow-8">
        <img src="@/assets/avatar.jpg" />
      </q-avatar>
    </div>
  </div>
  <!-- motto -->
  <div class="row justify-center q-mt-md">
    <p class="col-auto text-body1 text-center self-center">{{ t("profile.motto") }}</p>
  </div>
  <!-- desc -->
  <div class="row justify-center q-mt-md text-grey-7" style="font-size: 1.2em">
    <div class="col-md-6 col-sm-9 col-xs-10">
      <p style="font-weight: bold; font-size: large">{{ t("profile.welcome") }}</p>
      <p>{{ t("profile.intro") }}</p>
      <p v-html="t('profile.desc')"></p>
      <p>
        {{ t("profile.status") }}
      </p>
      <p v-html="t('profile.findme')"></p>
      <p style="word-wrap: break-word">
        {{ `${t("profile.mail")}\`echo bGl1eWloYW4wNDEwQGdtYWlsLmNvbQo= | base64 -d\`` }}
      </p>
    </div>
  </div>
</template>

<style lang="sass">
.a
  text-decoration: none
  position: relative
  color: $grey
  &:after
    content: ""
    position: absolute
    top: 65%
    left: -0.1em
    right: -0.1em
    bottom: 0
    transition: top 200ms cubic-bezier(0, 0.8, 0.13, 1)
    background-color: $green-3
    opacity: 0.4
    margin-right: 0.1rem
    margin-left: 0.1rem
</style>
