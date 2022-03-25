<script setup lang="ts">
import { getPosts, getTotal } from "../../api/post";
import { Post } from "../../types";

const posts = ref<Post[]>([]);
const pageQuery = reactive({ page: 1, size: 6 });
const total = ref<number>(0);

const max = ref(0);

onMounted(async () => {
  const [{ data: postRes }, { data: totalRes }] = await Promise.all([
    getPosts(pageQuery.page, pageQuery.size),
    getTotal(),
  ]);

  total.value = totalRes.data;
  posts.value = postRes.data;
  max.value = Math.ceil(total.value / pageQuery.size);
});

watch(pageQuery, (newVal) => {
  getPosts(newVal.page, newVal.size).then(({ data }) => {
    posts.value = data.data;
  });
});
</script>

<template>
  <div>
    <Transition name="scale" mode="out-in">
      <template v-if="posts.length > 0">
        <TransitionGroup name="list" tag="div" mod="out-in">
          <Item v-for="post in posts" :key="post.id" :post="post" />
        </TransitionGroup>
      </template>
      <template v-else>
        <Nothing />
      </template>
    </Transition>

    <Transition name="scale" mode="out-in">
      <div v-if="total > pageQuery.size" class="row justify-center q-mt-md q-mb-md">
        <div class="col-md-6 col-xs-10 col-sm-9">
          <q-pagination
            v-model="pageQuery.page"
            :max-pages="3"
            :max="max"
            :ripple="false"
            color="grey"
            :active-color="$q.dark.isActive ? '#88888825' : 'green-3'"
          />
        </div>
      </div>
    </Transition>
  </div>
</template>

<style scoped lang="css">
.list-move,
.list-enter-active,
.list-leave-active {
  opacity: 0.6;
  transition: all 0.5s ease;
}

.list-enter-from,
.list-leave-to {
  transform: translateY(100px);
}

.list-leave-active {
  position: absolute;
}
</style>
