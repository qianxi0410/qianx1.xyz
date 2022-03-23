<script lang="ts" setup>
const $q = useQuasar();

const current = ref(3);
</script>

<template>
  <div class="row justify-center q-mt-md">
    <div class="col-md-6 col-xs-10 col-sm-9">
      <q-pagination
        v-model="current"
        :max-pages="10"
        :max="10"
        :ripple="false"
        color="grey"
        round
        :active-color="$q.dark.isActive ? '#88888825' : 'green-3'"
      />
    </div>
  </div>
</template>
