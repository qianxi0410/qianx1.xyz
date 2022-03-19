import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

const routes = [{ path: "/", component: () => import("@/components/Hi.vue") }] as RouteRecordRaw[];

export const router = createRouter({
  history: createWebHashHistory(),
  routes,
});
