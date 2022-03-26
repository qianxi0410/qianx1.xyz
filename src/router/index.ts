import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

const routes = [
  { path: "/", component: () => import("@/views/Index.vue") },
  { path: "/blog", component: () => import("@/views/Blog.vue") },
  { path: "/posts/:id", component: () => import("@/views/Post.vue") },
  { path: "/:pathMatch(.*)*", component: () => import("@/views/404.vue") },
] as RouteRecordRaw[];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
