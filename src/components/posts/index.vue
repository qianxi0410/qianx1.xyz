<script setup lang="ts">
import { getPosts } from "../../api/post";
import { Post } from "../../types";

const posts = ref<Post[]>([]);
const pageQuery = reactive({ page: 1, size: 100 });

onMounted(async () => {
  const { data: postRes } = await getPosts(pageQuery.page, pageQuery.size);
  posts.value = postRes.data;
});
</script>

<template>
  <div>
    <Item v-for="post in posts" :key="post.id" :post="post" />
  </div>
</template>
