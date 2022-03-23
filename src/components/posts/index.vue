<script setup lang="ts">
import { Blog } from "../../types";

const blog: Blog = {
  title: "如何为你的terminal赋能",
  create_date: Date.now(),
  tags: ["terminal", "shell", "linux"],
};
</script>

<template>
  <Post v-for="i in 8" :key="i" :props="blog" />
  <Pagination />
</template>
