<script setup lang="ts">
import {
  evaArrowheadUpOutline,
  evaArrowLeftOutline,
  evaArrowRightOutline,
} from '@quasar/extras/eva-icons'
import { usePost } from '../../composoable/usePost'
import { useToc } from '../../composoable/useToc'
import { Post } from '../../types'

const post = reactive<Post>({} as any)

const create_date_fmt = computed(() => {
  const date = new Date(Number(post.create_time) * 1000)
  const array = date.toDateString().slice(4).split(' ')

  let res = `${array[0]} ${array[1]},${array[2]}`
  if (date.getFullYear() === new Date().getFullYear()) res = res.slice(0, Math.max(0, res.length - 5))
  return res
})

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
    window.history.pushState(undefined, '', `/posts/${post.display_id}`)
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

  const tocs = useToc()
  // eslint-disable-next-line no-console
  console.log(tocs)
  window.history.pushState(undefined, '', `/posts/${post.display_id}`)
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
