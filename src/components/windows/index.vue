<script setup lang="ts">import { handle } from './handle'

const prefix = '<span class="text-green">coustom</span>@<span class="text-blue">qianx1.xyz</span> <span class="text-pink">➞</span> '

const input = ref<HTMLInputElement>()

const handleInput = (e: KeyboardEvent) => {
  if (e.key !== 'Enter') return
  const value = input.value?.value
  handle(value!)
}

onMounted(() => {
  input.value?.addEventListener('keypress', handleInput)
})

onUnmounted(() => {
  input.value?.removeEventListener('keypress', handleInput)
})

</script>

<template>
  <div>
    <div class="row justify-center">
      <div class="col-md-7 col-sm-9 col-xs-10">
        <q-card class="shadow-10 window-height">
          <q-bar>
            <q-btn
              dense
              flat
              round
              icon="lens"
              size="0.6em"
              color="red"
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
          <div class="q-pa-md">
            <label v-html="prefix">
            </label>
            <input ref="input" class="bg-dark text-white input" />
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
</style>
