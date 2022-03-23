import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

const routes = [
  { path: "/", component: () => import("@/views/Index.vue") },
  { path: "/blog", component: () => import("@/views/Posts.vue") },
] as RouteRecordRaw[];

export const router = createRouter({
  history: createWebHashHistory(),
  routes,
});
