import { LoadingBar } from "quasar";
import { router } from "./../router/index";

export const config = {
  plugins: {
    LoadingBar,
  },
  config: {
    loadingBar: {
      color: "primary",
      size: "0.24rem",
      position: "top",
    },
    brand: {
      primary: "#37944c",
      secondary: "#26A69A",
      accent: "#9C27B0",

      dark: "#24202000",

      positive: "#46ab5e",
      negative: "#cf727c",
      info: "#9db7bd",
      warning: "#c9a542",
    },
  },
};

router.beforeEach(() => {
  LoadingBar.start();
});

router.afterEach(() => {
  LoadingBar.stop();
});
