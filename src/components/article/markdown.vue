<script setup lang="ts">
import { md2HTML } from "../../plugins/markdown";

const light = () => import("@/styles/light.css");
const dark = () => import("@/styles/dark.css");

const $q = useQuasar();

if ($q.dark.isActive) {
  dark();
} else {
  light();
}

const props = defineProps<{ text: string }>();

const html = md2HTML(props.text);
</script>

<template>
  <div class="row justify-center container" style="margin-top: 40px">
    <div class="col-md-6 col-sm-9 col-xs-10 text-grey">
      <div :class="$q.dark.isActive ? 'dark' : 'white'" v-html="html"></div>
    </div>
  </div>
</template>

<style lang="sass">
.title
    font-weight: bold
    transition: all 0.4s ease-in-out
    cursor: pointer
    &::before
        content: "#"
        margin-left: -1em
        position: absolute
        opacity: 0.0
        transition: all 0.4s ease-in-out
    &:hover
        color: $green-3
        transition: all 0.4s ease-in-out
        &::before
            color: $green-3
            opacity: 0.5

.link
  text-decoration: none
  color: $green-3
  background-image: linear-gradient($green-3, $green-3)
  background-position: 0% 100%
  background-repeat: no-repeat
  background-size: 0% 0.7px
  &::before
    content: '「'
    opacity: 1
  &::after
    content: '」'
    opacity: 1
  transition: background-size cubic-bezier(0, 0.5, 0, 1) 0.7s
  &:hover
    text-decoration: none
    background-size: 100% 1.2px
  &:focus
    text-decoration: none
    color: $green-8
    background-size: 100% 1.2px

.strong
  font-weight: bold

.del
  text-decoration: none
  background-color: #000
  color: #000
  padding: 0.1em 0.3em
  border-radius: 0.3em
  transition: all 0.4s ease-in-out
  &:hover
    background-color: rgba(0, 0, 0, 0)
    color: $grey
    text-decoration-style: dotted
    cursor: pointer

.codespan
  color: $green-3
  font-size: 0.8em
  &::before
    content: '`'
    color: $green-3
  &::after
    content: '`'
    color: $green-3

.image
  display: block
  margin: 0 auto
  max-width: 90%
  max-height: 100%
  border-radius: 1em
</style>
