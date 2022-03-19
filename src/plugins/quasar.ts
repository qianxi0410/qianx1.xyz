import { LoadingBar } from "quasar";
import { router } from "./../router/index";

export const config = {
  plugins: {
    LoadingBar,
  },
  config: {
    loadingBar: {
      color: "green",
      size: "0.24rem",
      position: "top",
    },
  },
};

router.beforeEach(() => {
  LoadingBar.start();
});

router.afterEach(() => {
  LoadingBar.stop();
});
