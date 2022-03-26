<script setup lang="ts">
import { getPost } from "../../api/post";
import { Post } from "../../types";

const post = reactive<Post>({} as any);

const create_date_fmt = computed(() => {
  const date = new Date(Number(post.create_time));
  const arr = date.toDateString().substring(4).split(" ");

  let res = `${arr[0]} ${arr[1]},${arr[2]}`;
  if (date.getFullYear() === new Date().getFullYear()) res = res.substring(0, res.length - 5);
  return res;
});

const route = useRoute();

onMounted(async () => {
  const id = route.params.id;
  if (id) {
    const { data } = await getPost(id as string);
    Object.assign(post, data.data);
  }
});
</script>

<template>
  <div class="row justify-center">
    <div class="col-md-6 col-xs-10 col-sm-9 text-h3">{{ post.title }}</div>
  </div>
  <div class="row justify-center">
    <div class="col-md-6 col-xs-10 col-sm-9 text-body1 text-grey q-mt-md">
      {{ create_date_fmt }} &nbsp;·&nbsp;
      <span>{{ post.author }}</span>
    </div>
  </div>
  <Markdown v-if="post.id" :text="post.content!" />
</template>
