import vueI18n from "@intlify/vite-plugin-vue-i18n";
import { quasar, transformAssetUrls } from "@quasar/vite-plugin";
import vue from "@vitejs/plugin-vue";
// eslint-disable-next-line import/order
import { resolve } from "path";
import autoImport from "unplugin-auto-import/vite";
import components from "unplugin-vue-components/vite";
import { defineConfig } from "vite";

export default defineConfig({
  resolve: {
    alias: {
      "@": "/src",
    },
  },
  envDir: "src/env",
  envPrefix: "APP_",
  // mode: "development",
  plugins: [
    vue({
      template: { transformAssetUrls },
    }),
    components({
      dts: "src/components.d.ts",
      dirs: ["src/components", "src/views"],
    }),
    autoImport({
      imports: ["vue", "@vueuse/core", "pinia", "vue-router", "vue-i18n", "quasar"],
      dts: "src/auto-imports.d.ts",
    }),
    vueI18n({
      runtimeOnly: true,
      compositionOnly: true,
      include: [resolve(__dirname, "locales/**")],
    }),
    quasar({
      sassVariables: "src/styles/variables.scss",
    }),
  ],
});
