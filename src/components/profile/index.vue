<script setup lang="ts"></script>

<template>
  <div>
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
      <p class="col col-auto text-body1 self-center">MAKE THINGS HAPPEN.</p>
    </div>
    <!-- desc -->
    <div class="row justify-center q-mt-md text-grey-7" style="font-size: 1.2em">
      <div class="col-md-6 col-xs-11">
        <p style="font-weight: bold; font-size: large">qianxi / qianxi0410 / 千禧.</p>
        <p>
          My name is qianxi. A graduate student of NWPU and expected to graduate in 2024 spring.
        </p>
        <p>
          Pursue new technologies, hate scientific research, like to write some useful or
          interesting little toys, and put them on my
          <a href="https://github.com/qianxi0410">GitHub</a>.
        </p>
        <p>
          Currently learning rust, distributed systems. Interested in networking related and DevOps.
        </p>
        <p>
          Find me on
          <a href="https://www.zhihu.com/people/qian-xi-8-12-94">ZhiHu</a> or
          <a href="https://twitter.com/qianxi0410">Twitter</a>.
        </p>
        <p>Mail me at: `echo bGl1eWloYW4wNDEwQGdtYWlsLmNvbQo= | base64 -d`.</p>
      </div>
    </div>

    <br />
    <br />
    <!-- copyright -->
    <div class="row justify-center q-mt-auto">
      <div class="col-md-6 text-body2 col-xs-11 text-grey-6">
        <span style="font-family: arial">&copy;</span>qianxi 2022 - {{ new Date().getFullYear() }}
        <a href="https://beian.miit.gov.cn">陕ICP备2021012752号</a>
      </div>
    </div>
    <br />
    <br />
  </div>
</template>

<style lang="sass" scoped>
a
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
