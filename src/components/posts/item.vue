<script lang="ts" setup>
import { Post } from '../../types'
const blog = defineProps<{ post: Post }>()

const create_date_fmt = computed(() => {
  const date = new Date(Number(blog.post.create_time) * 1000)
  const array = date.toDateString().slice(4).split(' ')

  let res = `${array[0]} ${array[1]},${array[2]}`
  if (date.getFullYear() === new Date().getFullYear()) res = res.slice(0, Math.max(0, res.length - 5))
  return res
})

const tags = computed(() => {
  if (!blog.post.tags) return []
  return blog.post.tags.split('/').map(tag => tag.trim())
})

const router = useRouter()
const to = (id: string) => {
  router.push({ path: `/posts/${id}` })
}
</script>

<template>
  <div class="row justify-center">
    <q-card class="col-md-6 col-xs-10 col-sm-9" style="background-color: rgba(0, 0, 0, 0)" flat>
      <q-card-section>
        <div class="text-h6 text-grey-7">
          <span class="post-title" @click="to(post.id)">{{ post.title }}</span>
        </div>
        <div class="text-grey text-body1">
          {{ create_date_fmt }}
          <template v-if="post.tags && post.tags.length > 0">
            <q-chip
              v-for="i in tags.length"
              :key="i"
              outline
              :ripple="false"
              size="sm"
              class="text-body2 text-grey"
              style="background-color: rgba(0, 0, 0, 0)"
            >
              {{ tags[i - 1] }}
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
