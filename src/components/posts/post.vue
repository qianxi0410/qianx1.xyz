<script lang="ts" setup>
import { Post } from "../../types";
const blog = defineProps<{ post: Post }>();

const create_date_fmt = computed(() => {
  const date = new Date(blog.post.create_date);
  const arr = date.toDateString().substring(4).split(" ");

  let res = `${arr[0]} ${arr[1]},${arr[2]}`;
  if (date.getFullYear() === new Date().getFullYear()) res = res.substring(0, res.length - 5);
  return res;
});
</script>

<template>
  <div class="row justify-center">
    <q-card class="col-md-6 col-xs-10 col-sm-9" style="background-color: rgba(0, 0, 0, 0)" flat>
      <q-card-section>
        <div class="text-h5 q-mt-sm text-grey-7 q-mb-xs">
          <span class="post-title">{{ post.title }}</span>
        </div>
        <div class="text-grey text-body1">
          {{ create_date_fmt }}
          <template v-if="post.tags && post.tags.length > 0">
            <q-chip
              v-for="i in post.tags.length"
              :key="i"
              outline
              size="sm"
              class="text-body2 text-grey"
              style="background-color: rgba(0, 0, 0, 0)"
            >
              {{ post.tags[i - 1] }}
            </q-chip>
          </template>
        </div>
      </q-card-section>
    </q-card>
  </div>
</template>

<style scoped lang="sass">
.post-title
  cursor: pointer
  margin-bottom: 0.5rem
  margin-top: 0.5rem
  transition: color .6s ease-in-out
  &:hover
    color: $green-3
</style>
