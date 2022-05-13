<script setup lang="ts">
import { getPosts } from '../../api/post'
import { Post } from '../../types'

const posts = ref<Post[]>([])

onMounted(async () => {
  const { data: postRes } = await getPosts()
  posts.value = postRes.data
})
</script>

<template>
  <div>
    <Item v-for="post in posts" :key="post.id" :post="post" />
  </div>
</template>
