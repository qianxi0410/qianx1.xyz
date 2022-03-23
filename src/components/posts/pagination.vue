<script lang="ts" setup>
const $q = useQuasar();

const props = defineProps<{ page: number; size: number; total: number }>();

const max = Math.ceil(props.total / props.size);

const current = ref(props.page > max ? max : props.page);

const emit = defineEmits<{
  (event: "change", page: number): void;
}>();

watch(current, (newVal) => {
  emit("change", newVal);
});
</script>

<template>
  <div v-if="max !== 1" class="row justify-center q-mt-md q-mb-md">
    <div class="col-md-6 col-xs-10 col-sm-9">
      <q-pagination
        v-model="current"
        :max-pages="3"
        :max="max"
        :ripple="false"
        color="grey"
        :active-color="$q.dark.isActive ? '#88888825' : 'green-3'"
      />
    </div>
  </div>
</template>
