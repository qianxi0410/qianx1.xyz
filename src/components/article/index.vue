<script setup lang="ts">
import {
  evaArrowheadUpOutline,
  evaArrowLeftOutline,
  evaArrowRightOutline,
} from '@quasar/extras/eva-icons'
import { useDate } from '../../composoable/useDate'
import { usePost } from '../../composoable/usePost'
import { Post } from '../../types'
import Comment from './comment.vue'
import Comment1 from './comment.vue'
import Comment2 from './comment.vue'

const post = reactive<Post>({} as any)

const create_date_fmt = computed(() => useDate(post.create_time))

const route = useRoute()
const router = useRouter()

const cancel = watch(
  () => route.params.id,
  // eslint-disable-next-line space-before-function-paren
  async (id) => {
    if (!route.path.startsWith('/posts')) return

    const p = await usePost(id as string)
    if (!p.id) router.push('/404')
    Object.assign(post, p)
  },
)

onUnmounted(() => {
  cancel()
})

const scroll = () => {
  const element = document.querySelector('#app')
  if (element) element.scrollIntoView({ behavior: 'smooth' })
}

const to = (id: string) => {
  router.push({ path: `/posts/${id}` })
  // scroll to top
  scroll()
}

const { y } = useWindowScroll()
const { height } = useWindowSize()

onMounted(async () => {
  if (!route.path.startsWith('/posts')) return

  const id = route.params.id as string

  const p = await usePost(id)
  if (!p.id) router.push('/404')
  Object.assign(post, p)
})

</script>

<template>
  <div v-if="post.id" :key="post.id">
    <Transition mode="out-in" name="fade">
      <q-btn
        v-if="y > height * 1.1"
        round
        flat
        :ripple="false"
        :icon="evaArrowheadUpOutline"
        class="fixed-bottom-right"
        @click="scroll"
      />
    </Transition>
    <div class="row justify-center">
      <div class="col-md-6 col-xs-10 col-sm-9 text-h3">
        {{ post.title }}
      </div>
    </div>
    <div class="row justify-center">
      <div class="col-md-6 col-xs-10 col-sm-9 text-body1 text-grey q-mt-md">
        {{ create_date_fmt }} &nbsp;·&nbsp;
        <span>qianxi0410</span>
      </div>
    </div>
    <Markdown :text="post.content!" />
    <div class="row justify-center">
      <q-toolbar class="col-md-5 col-sm-8 col-xs-9 text-center text-grey">
        <q-btn
          v-show="post.prev"
          :icon="evaArrowLeftOutline"
          rounded
          flat
          :ripple="false"
          @click="to(post.prev)"
        >
        </q-btn>
        <q-space />
        <q-btn
          v-show="post.next"
          :icon="evaArrowRightOutline"
          rounded
          flat
          :ripple="false"
          @click="to(post.next)"
        >
        </q-btn>
      </q-toolbar>
    </div>
    <Comment :issue_numer="post.issue_number" />
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
