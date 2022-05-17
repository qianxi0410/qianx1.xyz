<script setup lang="ts">import { handle } from './handle'

const prefix = '<span class="text-green">visitor</span>@<span class="text-blue">qianx1.xyz</span> <span class="text-pink">➞</span> '

const input = ref<HTMLInputElement>()
const content = ref<HTMLDivElement>()

const blocks = [] as string[]
const history = [] as string[]

const handleInput = async (e: KeyboardEvent) => {
  if (e.key !== 'Enter') return
  const value = input.value?.value
  await handle(value!, blocks, history)
  input.value!.value = ''

  // eslint-disable-next-line security/detect-object-injection
  content.value!.innerHTML = blocks.map((block, i) => `<div>${prefix}${history[i]}<br>${block}</div>`).join('')
}

onMounted(() => {
  input.value?.addEventListener('keypress', handleInput)
  input.value?.focus()
})

onUnmounted(() => {
  input.value?.removeEventListener('keypress', handleInput)
})

</script>

<template>
  <div>
    <div class="row justify-center">
      <div class="col-md-7 col-sm-9 col-xs-10">
        <q-card class="shadow-10">
          <q-bar>
            <q-btn
              dense
              flat
              round
              icon="lens"
              size="0.6em"
              color="red"
              to="/"
            />
            <q-btn
              dense
              flat
              round
              icon="lens"
              size="0.6em"
              color="yellow"
            />
            <q-btn
              dense
              flat
              round
              icon="lens"
              size="0.6em"
              color="green"
            />
            <div class="col text-center text-weight-bold">
              @qianxi0410
            </div>
          </q-bar>
          <q-separator />
          <div class="q-pa-md overflow-auto window-height">
            <div ref="content">
              welcome to qianxi0410's blog terminal. <br />
              you can run `help` to see all the commands. <br />
              have fun :D
            </div>
            <label v-html="prefix">
            </label>
            <input ref="input" class="input" :class="$q.dark.isActive ? 'bg-dark text-white ' : 'bg-white text-black'" />
          </div>
        </q-card>
      </div>
    </div>
  </div>
</template>

<style lang="sass" scoped>
.input
  outline: none
  border: none
  &:focus
    border: none
    outline: none
</style>
