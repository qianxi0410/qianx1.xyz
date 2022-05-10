<script setup lang="ts">
const { t } = useI18n()
</script>

<template>
  <br />
  <div class="row justify-center">
    <div class="col-md-6 col-sm-9 text-body2 col-xs-10 text-grey-6">
      <span style="font-family: arial">&copy;</span>{{ t("profile.name") }} 2022 -
      {{ new Date().getFullYear() }}
      <a class="a" href="https://beian.miit.gov.cn">陕ICP备2021012752号</a>
    </div>
  </div>
  <br />
</template>

<style lang="sass" scoped>
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
