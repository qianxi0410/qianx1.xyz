import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

const routes = [
  { path: "/", component: () => import("@/views/Index.vue") },
  { path: "/blog", component: () => import("@/views/Posts.vue") },
  { path: "/:pathMatch(.*)*", component: () => import("@/views/404.vue") },
] as RouteRecordRaw[];

export const router = createRouter({
  history: createWebHashHistory(),
  routes,
});
