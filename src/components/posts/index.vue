<script setup lang="ts">
import { Blog } from "../../types";

const blogs = ref<Blog[]>();

onMounted(() => {
  const blog = null;
});
</script>

<template>
  <div>
    <template v-if="blogs && blogs.length > 0"></template>
    <template v-else>
      <Nothing />
    </template>
    <Pagination :page="1" :size="8" :total="8" />
  </div>
</template>
