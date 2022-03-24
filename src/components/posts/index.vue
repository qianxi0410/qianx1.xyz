<script setup lang="ts">
import { Post } from "../../types";

const posts = ref<Post[]>();

onMounted(() => {
  const blog = null;
});
</script>

<template>
  <div>
    <template v-if="posts && posts.length > 0"></template>
    <template v-else>
      <Nothing />
    </template>
    <Pagination :page="1" :size="8" :total="8" />
  </div>
</template>
